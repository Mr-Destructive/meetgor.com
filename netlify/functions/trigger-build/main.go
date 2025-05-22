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
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
	githubUsername   = "mr-destructive"
	githubRepoName   = "mr-destructive.github.io"
	githubBranch     = "main"
	timestampFileURL = "https://raw.githubusercontent.com/%s/%s/%s/.last_build_timestamp"
	githubApiUrlBase = "https://api.github.com/repos/%s/%s/dispatches"
	githubEventType  = "content-update"

	// Time constants
	githubTimeout = 20 * time.Second

	// Default old timestamp for fallback
	defaultOldTime = "2000-01-01T00:00:00Z"
)

// Post represents a blog post/content item from the database
type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	HTMLContent string `json:"html_content"`
	Slug        string `json:"slug"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Type        string `json:"type"`
	PublishedAt string `json:"published_at,omitempty"`
	Status      string `json:"status,omitempty"`
}

// WorkflowPost represents content to be sent to GitHub Action
type WorkflowPost struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// GitHubActionPayload represents the complete payload for GitHub dispatch
type GitHubActionPayload struct {
	Timestamp string         `json:"timestamp"`
	Posts     []WorkflowPost `json:"posts"`
}

// Config holds all configuration values
type Config struct {
	GitHubPAT            string
	NetlifyTriggerSecret string
	TursoDBURL           string
	TursoAuthToken       string
}

// ProcessPost converts a post to markdown with frontmatter
func ProcessPost(post Post) (string, error) {
	// Convert HTML to Markdown
	markdownContent, err := htmltomarkdown.ConvertString(post.HTMLContent)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	// Generate frontmatter
	frontmatter := generateFrontmatter(post)

	return frontmatter + markdownContent, nil
}

// generateFrontmatter creates YAML frontmatter for the post
func generateFrontmatter(post Post) string {
	// Ensure we have valid dates
	date := post.CreatedAt
	if !isValidRFC3339(date) {
		log.Printf("Warning: Invalid CreatedAt for post %s, using current time", post.ID)
		date = time.Now().UTC().Format(time.RFC3339)
	}

	lastmod := post.UpdatedAt
	if !isValidRFC3339(lastmod) {
		lastmod = date // Fall back to creation date
	}

	// Escape quotes in title
	title := strings.ReplaceAll(post.Title, "\"", "\\\"")

	// Build frontmatter
	var frontmatter strings.Builder
	frontmatter.WriteString("---\n")
	frontmatter.WriteString(fmt.Sprintf("title: \"%s\"\n", title))
	frontmatter.WriteString(fmt.Sprintf("date: %s\n", date))
	frontmatter.WriteString(fmt.Sprintf("lastmod: %s\n", lastmod))
	frontmatter.WriteString(fmt.Sprintf("slug: %s\n", post.Slug))
	frontmatter.WriteString(fmt.Sprintf("type: %s\n", post.Type))

	// Add published_at if available
	if post.PublishedAt != "" && isValidRFC3339(post.PublishedAt) {
		frontmatter.WriteString(fmt.Sprintf("publishedAt: %s\n", post.PublishedAt))
	}

	// Add status if available
	if post.Status != "" {
		frontmatter.WriteString(fmt.Sprintf("status: %s\n", post.Status))
	}

	frontmatter.WriteString("---\n\n")

	return frontmatter.String()
}

// loadConfig loads and validates environment variables
func loadConfig() (*Config, error) {
	config := &Config{
		GitHubPAT:            os.Getenv("GITHUB_PAT_FOR_TRIGGER"),
		NetlifyTriggerSecret: os.Getenv("NETLIFY_TRIGGER_SECRET"),
		TursoDBURL:           os.Getenv("TURSO_DATABASE_URL"),
		TursoAuthToken:       os.Getenv("TURSO_AUTH_TOKEN"),
	}

	// Validate required fields
	if config.GitHubPAT == "" {
		return nil, fmt.Errorf("GITHUB_PAT_FOR_TRIGGER is required")
	}
	if config.TursoDBURL == "" {
		return nil, fmt.Errorf("TURSO_DATABASE_URL is required")
	}
	if config.TursoAuthToken == "" {
		return nil, fmt.Errorf("TURSO_AUTH_TOKEN is required")
	}

	return config, nil
}

// createResponse creates standardized API Gateway response with CORS headers
func createResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
			"Content-Type":                 "application/json",
		},
		Body: body,
	}
}

// createErrorResponse creates a JSON error response
func createErrorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	errorBody := map[string]string{"error": message}
	body, _ := json.Marshal(errorBody)
	return createResponse(statusCode, string(body))
}

// createSuccessResponse creates a JSON success response
func createSuccessResponse(message string, data interface{}) events.APIGatewayProxyResponse {
	response := map[string]interface{}{
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}
	body, _ := json.Marshal(response)
	return createResponse(http.StatusOK, string(body))
}

// validateTriggerSecret checks the X-Trigger-Secret header if configured
func validateTriggerSecret(request events.APIGatewayProxyRequest, expectedSecret string) bool {
	if expectedSecret == "" {
		return true // No validation required
	}

	// Check for the header (case-insensitive)
	requestSecret := getHeaderValue(request.Headers, "X-Trigger-Secret")
	return requestSecret == expectedSecret
}

// getHeaderValue gets header value case-insensitively
func getHeaderValue(headers map[string]string, key string) string {
	// Try exact match first
	if val, ok := headers[key]; ok {
		return val
	}

	// Try lowercase
	if val, ok := headers[strings.ToLower(key)]; ok {
		return val
	}

	return ""
}

// generateSlug creates a URL-friendly slug from title
func generateSlug(title string) string {
	if title == "" {
		return ""
	}

	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

// determineFilePath determines the file path based on post type and slug
func determineFilePath(post Post) string {
	slug := post.Slug
	if slug == "" {
		slug = generateSlug(post.Title)
		if slug == "" {
			slug = fmt.Sprintf("post-%s", post.ID) // Ultimate fallback
		}
	}

	var baseDir string = "posts/"
	folderPost := strings.ToLower(post.Type)
	return fmt.Sprintf("%s%s/%s.md", baseDir, folderPost, slug)
}

// isValidRFC3339 checks if a string is a valid RFC3339 timestamp
func isValidRFC3339(timeStr string) bool {
	if timeStr == "" {
		return false
	}
	_, err := time.Parse(time.RFC3339, timeStr)
	return err == nil
}

// parseTimestamp attempts to parse various timestamp formats to RFC3339
func parseTimestamp(timeStr string) (string, error) {
	if timeStr == "" {
		return "", fmt.Errorf("empty timestamp")
	}

	// Try RFC3339 first
	if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
		return t.UTC().Format(time.RFC3339), nil
	}

	// Try common database formats
	formats := []string{
		"2006-01-02 15:04:05",         // MySQL/PostgreSQL
		"2006-01-02T15:04:05",         // ISO without timezone
		"2006-01-02 15:04:05.000000",  // With microseconds
		"2006-01-02T15:04:05.000000Z", // ISO with microseconds
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t.UTC().Format(time.RFC3339), nil
		}
	}

	return "", fmt.Errorf("unparseable timestamp format: %s", timeStr)
}

// getPostModificationTime returns the effective modification time for comparison
func getPostModificationTime(post Post) (time.Time, error) {
	// Prefer UpdatedAt, fall back to CreatedAt
	timeStr := post.UpdatedAt
	if timeStr == "" {
		timeStr = post.CreatedAt
	}

	if timeStr == "" {
		return time.Time{}, fmt.Errorf("no valid timestamp found for post %s", post.ID)
	}

	// Parse the timestamp
	normalizedTime, err := parseTimestamp(timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse timestamp for post %s: %w", post.ID, err)
	}

	return time.Parse(time.RFC3339, normalizedTime)
}

// fetchLastBuildTimestamp retrieves the last build timestamp from GitHub
func fetchLastBuildTimestamp(githubPAT string) (time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	url := fmt.Sprintf(timestampFileURL, githubUsername, githubRepoName, githubBranch)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubPAT))
	req.Header.Set("Accept", "application/vnd.github.v3.raw")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to fetch timestamp: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		log.Println("Timestamp file not found, using default old timestamp")
		return time.Parse(time.RFC3339, defaultOldTime)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return time.Time{}, fmt.Errorf("GitHub API error: %d - %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to read response: %w", err)
	}

	timestampStr := strings.TrimSpace(string(body))
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp format '%s': %w", timestampStr, err)
	}

	return timestamp, nil
}

// fetchPostsFromDB retrieves posts from Turso database
func fetchPostsFromDB(dbURL, authToken string) ([]Post, error) {

	// Connect to Turso
	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	defer db.Close()

	// Query for posts - only fetch published posts or posts without status
	query := `
		SELECT 
			id, 
			title, 
			html_content, 
			COALESCE(slug, '') as slug,
			type, 
			published_at,
			updated_at, 
			created_at,
			COALESCE(status, 'published') as status
		FROM posts 
		WHERE COALESCE(status, 'published') IN ('published', 'public')
		ORDER BY COALESCE(updated_at, created_at) DESC
		LIMIT 1000`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		var publishedAt sql.NullString

		err := rows.Scan(
			&p.ID, &p.Title, &p.HTMLContent, &p.Slug, &p.Type,
			&publishedAt, &p.UpdatedAt, &p.CreatedAt, &p.Status,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		// Handle nullable published_at
		if publishedAt.Valid {
			if normalized, err := parseTimestamp(publishedAt.String); err == nil {
				p.PublishedAt = normalized
			}
		}

		// Normalize timestamps
		if normalized, err := parseTimestamp(p.CreatedAt); err == nil {
			p.CreatedAt = normalized
		} else {
			log.Printf("Warning: Invalid CreatedAt for post %s: %s", p.ID, p.CreatedAt)
			continue // Skip posts with invalid creation dates
		}

		if normalized, err := parseTimestamp(p.UpdatedAt); err == nil {
			p.UpdatedAt = normalized
		} else {
			p.UpdatedAt = p.CreatedAt // Fall back to creation date
		}

		// Validate required fields
		p.Title = strings.TrimSpace(p.Title)
		p.HTMLContent = strings.TrimSpace(p.HTMLContent)

		if p.Title == "" || p.HTMLContent == "" {
			log.Printf("Skipping post %s: missing title or content", p.ID)
			continue
		}

		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	log.Printf("Fetched %d valid posts from database", len(posts))
	return posts, nil
}

// triggerGitHubAction sends a repository dispatch event to GitHub
func triggerGitHubAction(githubPAT string, payload GitHubActionPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	dispatchURL := fmt.Sprintf(githubApiUrlBase, githubUsername, githubRepoName)

	// Create dispatch payload
	dispatchPayload := map[string]interface{}{
		"event_type":     githubEventType,
		"client_payload": payload,
	}

	body, err := json.Marshal(dispatchPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", dispatchURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubPAT))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// handler is the main Netlify function handler
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Content sync function called: %s %s", request.HTTPMethod, request.Path)

	// Handle CORS preflight
	if request.HTTPMethod == "OPTIONS" {
		return createResponse(http.StatusOK, ""), nil
	}

	// Validate HTTP method
	if request.HTTPMethod != "POST" {
		return createErrorResponse(http.StatusMethodNotAllowed, "Only POST method is allowed"), nil
	}

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Printf("Configuration error: %v", err)
		return createErrorResponse(http.StatusInternalServerError, "Server configuration error"), nil
	}

	// Validate trigger secret
	if !validateTriggerSecret(request, config.NetlifyTriggerSecret) {
		log.Println("Unauthorized: Invalid or missing trigger secret")
		return createErrorResponse(http.StatusUnauthorized, "Invalid or missing trigger secret"), nil
	}

	// Fetch last build timestamp
	lastBuildTime, err := fetchLastBuildTimestamp(config.GitHubPAT)
	if err != nil {
		log.Printf("Error fetching last build timestamp: %v", err)
		// Use default old time to ensure we process content
		lastBuildTime, _ = time.Parse(time.RFC3339, defaultOldTime)
	}
	log.Printf("Last build timestamp: %s", lastBuildTime.Format(time.RFC3339))

	// Fetch posts from database
	posts, err := fetchPostsFromDB(config.TursoDBURL, config.TursoAuthToken)
	if err != nil {
		log.Printf("Database error: %v", err)
		return createErrorResponse(http.StatusInternalServerError, "Failed to fetch content from database"), nil
	}

	// Process posts that have been modified since last build
	var workflowPosts []WorkflowPost
	var processedCount int

	for _, post := range posts {
		postTime, err := getPostModificationTime(post)
		if err != nil {
			log.Printf("Skipping post %s: %v", post.ID, err)
			continue
		}

		// Only process posts modified after last build
		if !postTime.After(lastBuildTime) {
			continue
		}

		log.Printf("Processing updated post: %s (%s)", post.ID, post.Title)

		// Convert to markdown with frontmatter
		content, err := htmltomarkdown.ConvertString(post.HTMLContent)
		if err != nil {
			log.Printf("Error processing post %s: %v", post.ID, err)
			continue
		}

		// Determine file path
		filePath := determineFilePath(post)

		workflowPosts = append(workflowPosts, WorkflowPost{
			Path:    filePath,
			Content: content,
		})
		processedCount++
	}

	// Check if we have any content to update
	if len(workflowPosts) == 0 {
		log.Println("No new or updated content found")
		return createSuccessResponse("No new content to update", map[string]interface{}{
			"last_build_time": lastBuildTime.Format(time.RFC3339),
			"posts_checked":   len(posts),
		}), nil
	}

	// Create GitHub Action payload
	currentTimestamp := time.Now().UTC().Format(time.RFC3339)
	payload := GitHubActionPayload{
		Timestamp: currentTimestamp,
		Posts:     workflowPosts,
	}

	// Trigger GitHub Action
	if err := triggerGitHubAction(config.GitHubPAT, payload); err != nil {
		log.Printf("Failed to trigger GitHub Action: %v", err)
		return createErrorResponse(http.StatusBadGateway, "Failed to trigger content build"), nil
	}

	log.Printf("Successfully triggered GitHub Action with %d posts", len(workflowPosts))

	return createSuccessResponse("Content update triggered successfully", map[string]interface{}{
		"timestamp":       currentTimestamp,
		"posts_updated":   len(workflowPosts),
		"posts_checked":   len(posts),
		"last_build_time": lastBuildTime.Format(time.RFC3339),
	}), nil
}

func main() {
	lambda.Start(handler)
}
