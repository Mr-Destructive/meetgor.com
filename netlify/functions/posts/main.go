package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	libsqlssg "github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	log.SetOutput(io.Discard)
}

const (
	githubUsername   = "mr-destructive"
	githubRepoName   = "meetgor.com"
	githubEventType  = "content-update"
	githubApiUrlBase = "https://api.github.com/repos/%s/%s/dispatches"
	githubTimeout    = 20 * time.Second
)

var (
	openDB              = sql.Open
	closeDB             = func(db *sql.DB) error { return db.Close() }
	httpDo              = func(req *http.Request) (*http.Response, error) { return http.DefaultClient.Do(req) }
	triggerGitHubAction = defaultTriggerGitHubAction
)

type postPayload struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Title    string                 `json:"title"`
	Body     string                 `json:"body"`
	Slug     string                 `json:"slug,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC] %v\n%s", r, debug.Stack())
		}
	}()

	log.Printf("[REQUEST] method=%s path=%s", request.HTTPMethod, request.Path)

	if request.HTTPMethod == http.MethodOptions {
		return optionsResponse(), nil
	}

	path := normalizePath(request.Path)
	if path == "" {
		if request.HTTPMethod != http.MethodPost {
			log.Printf("[ERROR] slug required for method=%s", request.HTTPMethod)
			return jsonError(http.StatusBadRequest, "slug required"), nil
		}
	}

	db, err := openDatabase()
	if err != nil {
		log.Printf("[ERROR] database connection failed")
		return jsonError(http.StatusInternalServerError, "database connection failed"), nil
	}
	defer closeDB(db)

	ctx := context.Background()

	switch request.HTTPMethod {
	case http.MethodPost:
		if path != "" {
			log.Printf("[ERROR] invalid path for POST: %s", path)
			return jsonError(http.StatusBadRequest, "invalid path for POST"), nil
		}
		return handleCreate(ctx, db, request)
	case http.MethodGet:
		return handleGet(ctx, db, path)
	case http.MethodPut:
		return handleUpdate(ctx, db, path, request)
	case http.MethodDelete:
		return handleDelete(ctx, db, path, request)
	default:
		return jsonError(http.StatusMethodNotAllowed, "method not allowed"), nil
	}
}

func handleCreate(ctx context.Context, db *sql.DB, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload, err := readPayload(request.Body)
	if err != nil {
		log.Printf("[ERROR] invalid payload: %v", err)
		return jsonError(http.StatusBadRequest, "invalid payload"), nil
	}

	if err := validateCreate(payload); err != nil {
		log.Printf("[ERROR] validation failed: %v", err)
		return jsonError(http.StatusBadRequest, err.Error()), nil
	}

	log.Printf("[AUTH] attempting login for user: %s", payload.Username)
	user, err := authenticate(ctx, db, payload.Username, payload.Password)
	if err != nil {
		log.Printf("[ERROR] authentication failed for user: %s", payload.Username)
		return jsonError(http.StatusUnauthorized, "authentication failed"), nil
	}
	log.Printf("[AUTH] user authenticated: %s (id=%d)", user.Name, user.ID)

	slug := pickSlug(payload)
	if slug == "" {
		log.Printf("[ERROR] empty slug generated")
		return jsonError(http.StatusBadRequest, "slug or title required"), nil
	}
	log.Printf("[CREATE] slug=%s title=%s type=%v", slug, payload.Title, payload.Metadata["type"])

	payload.Metadata = ensureMetadata(payload.Metadata, slug, payload.Title)
	meta, err := json.Marshal(payload.Metadata)
	if err != nil {
		log.Printf("[ERROR] metadata serialization failed")
		return jsonError(http.StatusInternalServerError, "failed to serialize metadata"), nil
	}

	q := libsqlssg.New(db)
	post, err := q.CreatePost(ctx, libsqlssg.CreatePostParams{
		Title:    payload.Title,
		Slug:     slug,
		Body:     payload.Body,
		Metadata: string(meta),
		AuthorID: user.ID,
	})
	if err != nil {
		log.Printf("[ERROR] database insert failed: %v", err)
		return jsonError(http.StatusInternalServerError, "failed to insert post"), nil
	}

	log.Printf("[BUILD] triggering GitHub Action")
	if err := triggerBuild(); err != nil {
		log.Printf("[ERROR] build trigger failed: %v", err)
		return jsonError(http.StatusBadGateway, "build trigger failed"), nil
	}

	log.Printf("[SUCCESS] post created: %s", post.Slug)
	return jsonResponse(http.StatusOK, map[string]interface{}{
		"slug": post.Slug,
	}), nil
}

func handleGet(ctx context.Context, db *sql.DB, slug string) (events.APIGatewayProxyResponse, error) {
	if slug == "" {
		return jsonError(http.StatusBadRequest, "slug required"), nil
	}

	q := libsqlssg.New(db)
	posts, err := q.GetPostsBySlugType(ctx, slug)
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to query post"), nil
	}
	if len(posts) == 0 || posts[0].Deleted.Bool {
		return jsonError(http.StatusNotFound, "post not found"), nil
	}

	return jsonResponse(http.StatusOK, map[string]interface{}{
		"title":    posts[0].Title,
		"body":     posts[0].Body,
		"metadata": posts[0].Metadata,
		"slug":     posts[0].Slug,
	}), nil
}

func handleUpdate(ctx context.Context, db *sql.DB, slug string, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if slug == "" {
		return jsonError(http.StatusBadRequest, "slug required"), nil
	}

	payload, err := readPayload(request.Body)
	if err != nil {
		return jsonError(http.StatusBadRequest, err.Error()), nil
	}

	user, err := authenticate(ctx, db, payload.Username, payload.Password)
	if err != nil {
		return jsonError(http.StatusUnauthorized, "authentication failed"), nil
	}

	q := libsqlssg.New(db)
	posts, err := q.GetPostsBySlugType(ctx, slug)
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to query post"), nil
	}
	if len(posts) == 0 {
		return jsonError(http.StatusNotFound, "post not found"), nil
	}

	post := posts[0]
	metadata, _ := parseMetadata(post.Metadata)
	metadata = mergeMetadata(metadata, payload.Metadata, slug, payload.Title)

	newTitle := post.Title
	newBody := post.Body
	if payload.Title != "" {
		newTitle = payload.Title
	}
	if payload.Body != "" {
		newBody = payload.Body
	}

	metaBytes, err := json.Marshal(metadata)
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to serialize metadata"), nil
	}

	if _, err := db.ExecContext(ctx,
		"UPDATE posts SET title = ?, body = ?, metadata = ?, updated_at = CURRENT_TIMESTAMP, author_id = ? WHERE id = ?",
		newTitle, newBody, string(metaBytes), user.ID, post.ID); err != nil {
		return jsonError(http.StatusInternalServerError, "failed to update post"), nil
	}

	if err := triggerBuild(); err != nil {
		return jsonError(http.StatusBadGateway, err.Error()), nil
	}

	return jsonResponse(http.StatusOK, map[string]interface{}{
		"slug": slug,
	}), nil
}

func handleDelete(ctx context.Context, db *sql.DB, slug string, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if slug == "" {
		return jsonError(http.StatusBadRequest, "slug required"), nil
	}

	payload, err := readPayload(request.Body)
	if err != nil {
		return jsonError(http.StatusBadRequest, err.Error()), nil
	}

	if _, err := authenticate(ctx, db, payload.Username, payload.Password); err != nil {
		return jsonError(http.StatusUnauthorized, "authentication failed"), nil
	}

	if _, err := db.ExecContext(ctx, "UPDATE posts SET deleted = 1 WHERE slug = ?", slug); err != nil {
		return jsonError(http.StatusInternalServerError, "failed to delete post"), nil
	}

	if err := triggerBuild(); err != nil {
		return jsonError(http.StatusBadGateway, err.Error()), nil
	}

	return jsonResponse(http.StatusOK, map[string]string{"slug": slug}), nil
}

func awsHandler() {
	lambda.Start(handler)
}

func main() {
	awsHandler()
}

func normalizePath(path string) string {
	for _, prefix := range []string{"/.netlify/functions/posts", "/posts"} {
		if strings.HasPrefix(path, prefix) {
			path = strings.TrimPrefix(path, prefix)
			break
		}
	}
	path = strings.Trim(path, "/")
	return strings.TrimSpace(path)
}

func readPayload(body string) (postPayload, error) {
	var payload postPayload
	if body == "" {
		return payload, fmt.Errorf("body required")
	}
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		return payload, fmt.Errorf("invalid json: %w", err)
	}
	return payload, nil
}

func validateCreate(payload postPayload) error {
	if payload.Username == "" || payload.Password == "" {
		return fmt.Errorf("username and password required")
	}
	if payload.Title == "" {
		return fmt.Errorf("title required")
	}
	if payload.Body == "" {
		return fmt.Errorf("body required")
	}
	return nil
}

func pickSlug(payload postPayload) string {
	if payload.Slug != "" {
		return slugify(payload.Slug)
	}
	return slugify(payload.Title)
}

func ensureMetadata(meta map[string]interface{}, slug, title string) map[string]interface{} {
	if meta == nil {
		meta = make(map[string]interface{})
	}
	meta["slug"] = slug
	if title != "" {
		meta["title"] = title
	}
	if _, ok := meta["date"]; !ok {
		meta["date"] = time.Now().Format(time.RFC3339)
	}
	return meta
}

func mergeMetadata(existing map[string]interface{}, override map[string]interface{}, slug, title string) map[string]interface{} {
	if existing == nil {
		existing = make(map[string]interface{})
	}
	for k, v := range override {
		existing[k] = v
	}
	if slug != "" {
		existing["slug"] = slug
	}
	if title != "" {
		existing["title"] = title
	}
	if _, ok := existing["date"]; !ok {
		existing["date"] = time.Now().Format(time.RFC3339)
	}
	return existing
}

func parseMetadata(raw string) (map[string]interface{}, error) {
	if strings.TrimSpace(raw) == "" {
		return map[string]interface{}{}, nil
	}
	var meta map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &meta); err != nil {
		return nil, err
	}
	return meta, nil
}

func slugify(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	re := regexp.MustCompile(`[^a-z0-9]+`)
	input = re.ReplaceAllString(input, "-")
	input = strings.ReplaceAll(input, "--", "-")
	return strings.Trim(input, "-")
}

func authenticate(ctx context.Context, db *sql.DB, username, password string) (libsqlssg.GetUserRow, error) {
	q := libsqlssg.New(db)
	user, err := q.GetUser(ctx, username)
	if err != nil {
		return user, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}
	return user, nil
}

func jsonResponse(statusCode int, payload interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, PUT, DELETE, OPTIONS",
		},
		Body: string(body),
	}
}

func jsonError(statusCode int, message string) events.APIGatewayProxyResponse {
	return jsonResponse(statusCode, map[string]string{"error": message})
}

func optionsResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
		},
		Body: "",
	}
}

func openDatabase() (*sql.DB, error) {
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	if dbName == "" || dbToken == "" {
		return nil, fmt.Errorf("missing Turso credentials")
	}
	conn := fmt.Sprintf("libsql://%s?authToken=%s", dbName, dbToken)
	return openDB("libsql", conn)
}

func triggerBuild() error {
	token := os.Getenv("GITHUB_PAT_FOR_TRIGGER")
	if token == "" {
		return fmt.Errorf("missing GitHub token")
	}
	return triggerGitHubAction(token)
}

func defaultTriggerGitHubAction(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	url := fmt.Sprintf(githubApiUrlBase, githubUsername, githubRepoName)
	payload := map[string]string{"event_type": githubEventType}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := httpDo(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("github error: %d %s", resp.StatusCode, string(data))
	}
	return nil
}

// test rebuild
