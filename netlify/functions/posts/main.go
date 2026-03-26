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

	"github.com/Mr-Destructive/meetgor.com/plugins"
	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	// Logs are enabled for debugging
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
	Content  string                 `json:"content"`
	Body     string                 `json:"body"` // backward compatibility
	Slug     string                 `json:"slug,omitempty"`
	TypeID   string                 `json:"type_id,omitempty"`
	Excerpt  string                 `json:"excerpt,omitempty"`
	Tags     string                 `json:"tags,omitempty"`
	Status   string                 `json:"status,omitempty"`
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

	// Get type_id from payload or metadata
	typeID := resolvePostType(payload)

	payload.Metadata = ensureMetadata(payload.Metadata, slug, payload.Title)
	payload.Metadata["type"] = typeID
	meta, err := json.Marshal(payload.Metadata)
	if err != nil {
		log.Printf("[ERROR] metadata serialization failed")
		return jsonError(http.StatusInternalServerError, "failed to serialize metadata"), nil
	}

	q := libsqlssg.New(db)

	// Accept either Content or Body field (backward compatibility)
	content := payload.Content
	if content == "" {
		content = payload.Body
	}

	// For "links" type posts with empty content, create content from metadata
	if typeID == "links" && content == "" {
		if commentary, ok := payload.Metadata["commentary"].(string); ok && commentary != "" {
			content = commentary
		}
	}

	// Clean up metadata - for links type, keep only url/link, remove title/commentary/type
	cleanMeta := make(map[string]interface{})
	if typeID == "links" {
		// For links, keep essential fields
		if link, ok := payload.Metadata["link"]; ok {
			cleanMeta["url"] = link
		}
		if date, ok := payload.Metadata["date"]; ok {
			cleanMeta["date"] = date
		}
		if tags, ok := payload.Metadata["tags"]; ok {
			cleanMeta["tags"] = tags
		}
	} else {
		// For other types, copy all metadata except type
		for k, v := range payload.Metadata {
			if k != "type" {
				cleanMeta[k] = v
			}
		}
	}

	metaBytes, _ := json.Marshal(cleanMeta)
	meta = metaBytes

	// Extract tags from metadata or payload
	var tagsJSON sql.NullString
	if tagsVal, ok := payload.Metadata["tags"]; ok {
		if tagsStr, ok := tagsVal.(string); ok {
			// Tags as string (comma-separated)
			tagsArr := strings.Split(tagsStr, ",")
			for i := range tagsArr {
				tagsArr[i] = strings.TrimSpace(tagsArr[i])
			}
			tagsBytes, _ := json.Marshal(tagsArr)
			tagsJSON = sql.NullString{String: string(tagsBytes), Valid: true}
		} else if tagsArr, ok := tagsVal.([]interface{}); ok {
			// Tags as array
			tagsBytes, _ := json.Marshal(tagsArr)
			tagsJSON = sql.NullString{String: string(tagsBytes), Valid: true}
		}
	}

	log.Printf("[INSERT] preparing post insert: title=%s slug=%s type_id=%s", payload.Title, slug, typeID)

	// Ensure post type exists - if not, create it
	var typeExists int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM post_types WHERE id = ?", typeID).Scan(&typeExists)
	if err != nil || typeExists == 0 {
		// Create the post type if it doesn't exist
		log.Printf("[INFO] creating post type: %s", typeID)
		_, err = db.ExecContext(ctx, "INSERT INTO post_types (id, name, slug) VALUES (?, ?, ?)", typeID, strings.Title(typeID), typeID)
		if err != nil {
			log.Printf("[WARN] failed to create post type %s: %v", typeID, err)
		}
	}
	status := payload.Status
	if status == "" {
		status = "draft"
	}

	_, err = q.CreatePost(ctx, libsqlssg.CreatePostParams{
		ID:       generateID(),
		TypeID:   typeID,
		Title:    payload.Title,
		Slug:     slug,
		Content:  content,
		Metadata: sql.NullString{String: string(meta), Valid: true},
		Tags:     tagsJSON,
		Status:   sql.NullString{String: status, Valid: true},
	})

	if err != nil {
		log.Printf("[ERROR] database insert failed: %v", err)
		log.Printf("[ERROR] post params: title=%s slug=%s content_len=%d type_id=%s",
			payload.Title, slug, len(content), typeID)
		return jsonError(http.StatusInternalServerError, fmt.Sprintf("failed to insert post: %v", err)), nil
	}

	log.Printf("[BUILD] triggering GitHub Action")
	if err := triggerBuild(); err != nil {
		log.Printf("[ERROR] build trigger failed: %v", err)
		return jsonError(http.StatusBadGateway, "build trigger failed"), nil
	}

	log.Printf("[SUCCESS] post created: %s", slug)
	return jsonResponse(http.StatusOK, map[string]interface{}{
		"slug": slug,
	}), nil
}

func handleGet(ctx context.Context, db *sql.DB, slug string) (events.APIGatewayProxyResponse, error) {
	if slug == "" {
		return jsonError(http.StatusBadRequest, "slug required"), nil
	}

	q := libsqlssg.New(db)
	post, err := q.GetPostBySlug(ctx, slug)
	if err == sql.ErrNoRows {
		return jsonError(http.StatusNotFound, "post not found"), nil
	}
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to query post"), nil
	}

	return jsonResponse(http.StatusOK, map[string]interface{}{
		"title":    post.Title,
		"content":  post.Content,
		"metadata": post.Metadata.String,
		"slug":     post.Slug,
		"status":   post.Status.String,
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

	_, err = authenticate(ctx, db, payload.Username, payload.Password)
	if err != nil {
		return jsonError(http.StatusUnauthorized, "authentication failed"), nil
	}

	q := libsqlssg.New(db)
	post, err := q.GetPostBySlug(ctx, slug)
	if err == sql.ErrNoRows {
		return jsonError(http.StatusNotFound, "post not found"), nil
	}
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to query post"), nil
	}

	parsedMeta, _ := parseMetadata(post.Metadata.String)
	parsedMeta = mergeMetadata(parsedMeta, payload.Metadata, slug, payload.Title)
	parsedMeta["type"] = resolvePostType(postPayload{
		TypeID:   fmt.Sprint(parsedMeta["type"]),
		Metadata: parsedMeta,
	})

	newTitle := post.Title
	newContent := post.Content
	if payload.Title != "" {
		newTitle = payload.Title
	}
	// Accept either Content or Body field
	updateContent := payload.Content
	if updateContent == "" {
		updateContent = payload.Body
	}
	if updateContent != "" {
		newContent = updateContent
	}

	metaBytes, err := json.Marshal(parsedMeta)
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to serialize metadata"), nil
	}

	// Extract tags from metadata for update
	var tagsJSON sql.NullString
	if tagsVal, ok := parsedMeta["tags"]; ok {
		if tagsStr, ok := tagsVal.(string); ok {
			tagsArr := strings.Split(tagsStr, ",")
			for i := range tagsArr {
				tagsArr[i] = strings.TrimSpace(tagsArr[i])
			}
			tagsBytes, _ := json.Marshal(tagsArr)
			tagsJSON = sql.NullString{String: string(tagsBytes), Valid: true}
		} else if tagsArr, ok := tagsVal.([]interface{}); ok {
			tagsBytes, _ := json.Marshal(tagsArr)
			tagsJSON = sql.NullString{String: string(tagsBytes), Valid: true}
		}
	}

	if err := q.UpdatePost(ctx, libsqlssg.UpdatePostParams{
		Title:    newTitle,
		Content:  newContent,
		Metadata: sql.NullString{String: string(metaBytes), Valid: true},
		Tags:     tagsJSON,
		Slug:     slug,
	}); err != nil {
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

	q := libsqlssg.New(db)
	if err := q.DeletePost(ctx, slug); err != nil {
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
	// Accept either Content or Body field (backward compatibility)
	content := payload.Content
	if content == "" {
		content = payload.Body
	}
	// Links type doesn't require content
	postType := resolvePostType(payload)
	if postType != "links" && content == "" {
		return fmt.Errorf("content required")
	}
	return nil
}

func resolvePostType(payload postPayload) string {
	typeID := strings.TrimSpace(payload.TypeID)
	if typeID == "" {
		if postType, ok := payload.Metadata["type"].(string); ok && postType != "" {
			typeID = postType
		}
	}
	typeID = plugins.NormalizePostType(typeID)
	if typeID == "" {
		typeID = "posts"
	}
	return typeID
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

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
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
