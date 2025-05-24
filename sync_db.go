package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
	timestampFile = ".last_build_timestamp"
	timeFormat    = "2006-01-02 15:04:05"
	rfc3339Format = time.RFC3339
)

// Post represents a database post record
type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Body     string `json:"body"`
	Metadata string `json:"metadata"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	AuthorID int64  `json:"author_id"`
	Deleted  bool   `json:"deleted"`
}

// PostMetadata represents the metadata structure
type PostMetadata struct {
	PostDir     string   `json:"post_dir,omitempty"`
	Type        string   `json:"type,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Draft       bool     `json:"draft,omitempty"`
	Description string   `json:"description,omitempty"`
	Slug        string   `json:"slug,omitempty"`
}

// Config holds application configuration
type Config struct {
	DBUrl      string
	SyncAll    bool
	TimeCutoff time.Time
	OutputDir  string
	DryRun     bool
	Verbose    bool
}

func main() {
	logger := log.New(os.Stdout, "[SYNC] ", log.LstdFlags|log.Lshortfile)

	config, err := parseConfig()
	if err != nil {
		logger.Fatalf("Configuration error: %v", err)
	}

	if config.Verbose {
		logger.Printf("Starting content sync with config: SyncAll=%t, DryRun=%t", config.SyncAll, config.DryRun)
	}

	// Connect to database
	db, err := connectDB(config.DBUrl)
	if err != nil {
		logger.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Fetch posts
	posts, err := fetchPosts(db, config, logger)
	if err != nil {
		logger.Fatalf("Failed to fetch posts: %v", err)
	}

	if len(posts) == 0 {
		logger.Println("No posts to process")
		return
	}

	logger.Printf("Found %d posts to process", len(posts))

	// Process posts
	processedCount := 0
	errorCount := 0

	for _, post := range posts {
		if config.Verbose {
			logger.Printf("Processing post: ID=%d, Title='%s'", post.ID, post.Title)
		}

		if err := processPost(post, config, logger); err != nil {
			logger.Printf("Error processing post ID %d: %v", post.ID, err)
			errorCount++
			continue
		}
		processedCount++
	}

	logger.Printf("Processing complete: %d processed, %d errors", processedCount, errorCount)

	// Update timestamp if not dry run and we processed posts
	if !config.DryRun && processedCount > 0 {
		if err := updateTimestamp(); err != nil {
			logger.Printf("Warning: Failed to update timestamp: %v", err)
		} else if config.Verbose {
			logger.Println("Updated build timestamp")
		}
	}
}

// parseConfig parses command line arguments and environment variables
func parseConfig() (*Config, error) {
	config := &Config{
		OutputDir: "content",
		DryRun:    false,
		Verbose:   false,
	}

	// Parse environment variables
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		return nil, fmt.Errorf("missing required environment variables: TURSO_DATABASE_NAME and TURSO_DATABASE_AUTH_TOKEN")
	}

	config.DBUrl = fmt.Sprintf("%s?authToken=%s", dbName, dbToken)

	// Parse command line arguments
	args := os.Args[1:]
	for i, arg := range args {
		switch arg {
		case "sync_all", "--sync-all", "-a":
			config.SyncAll = true
		case "sync_post", "--sync-posts", "-p":
			config.SyncAll = false
		case "--dry-run", "-d":
			config.DryRun = true
		case "--verbose", "-v":
			config.Verbose = true
		case "--output-dir", "-o":
			if i+1 < len(args) {
				config.OutputDir = args[i+1]
			}
		case "--hours":
			if i+1 < len(args) {
				hours := parseHours(args[i+1])
				config.TimeCutoff = time.Now().Add(time.Duration(-hours) * time.Hour)
			}
		}
	}

	// Set default time cutoff if not sync all
	if !config.SyncAll && config.TimeCutoff.IsZero() {
		config.TimeCutoff = getLastBuildTime()
	}

	return config, nil
}

// parseHours safely parses hours from string
func parseHours(s string) int {
	var hours int
	fmt.Sscanf(s, "%d", &hours)
	if hours <= 0 {
		hours = 1 // Default to 1 hour
	}
	return hours
}

// getLastBuildTime gets the last build timestamp or defaults to 1 hour ago
func getLastBuildTime() time.Time {
	if data, err := os.ReadFile(timestampFile); err == nil {
		if t, err := time.Parse(rfc3339Format, strings.TrimSpace(string(data))); err == nil {
			return t
		}
	}
	return time.Now().Add(-1 * time.Hour)
}

// updateTimestamp updates the build timestamp file
func updateTimestamp() error {
	timestamp := time.Now().UTC().Format(rfc3339Format)
	return os.WriteFile(timestampFile, []byte(timestamp), 0644)
}

// connectDB establishes database connection with validation
func connectDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// fetchPosts retrieves posts from database based on configuration
func fetchPosts(db *sql.DB, config *Config, logger *log.Logger) ([]Post, error) {
	var query string
	var args []interface{}

	if config.SyncAll {
		query = "SELECT id, title, slug, body, metadata, created_at, updated_at, author_id, deleted FROM posts WHERE deleted = false ORDER BY updated_at DESC"
	} else {
		query = "SELECT id, title, slug, body, metadata, created_at, updated_at, author_id, deleted FROM posts WHERE deleted = false AND (created_at > ? OR updated_at > ?) ORDER BY updated_at DESC"
		timeStr := config.TimeCutoff.Format(timeFormat)
		args = append(args, timeStr, timeStr)
	}

	if config.Verbose {
		logger.Printf("Executing query: %s", query)
		if len(args) > 0 {
			logger.Printf("Query args: %v", args)
		}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var metadata sql.NullString

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Body,
			&metadata,
			&post.Created,
			&post.Updated,
			&post.AuthorID,
			&post.Deleted,
		)
		if err != nil {
			logger.Printf("Error scanning row: %v", err)
			continue
		}

		if metadata.Valid {
			post.Metadata = metadata.String
		} else {
			post.Metadata = "{}"
		}

		// Skip deleted posts (extra safety)
		if post.Deleted {
			continue
		}

		// Skip posts with empty body or title
		if strings.TrimSpace(post.Body) == "" || strings.TrimSpace(post.Title) == "" {
			logger.Printf("Skipping post ID %d: empty title or body", post.ID)
			continue
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return posts, nil
}

// processPost handles the conversion and writing of a single post
func processPost(post Post, config *Config, logger *log.Logger) error {
	// Parse metadata
	metadata, err := parseMetadata(post.Metadata)
	if err != nil {
		return fmt.Errorf("failed to parse metadata: %w", err)
	}

	// Convert HTML to Markdown

	bodyMd, err := htmltomarkdown.ConvertString(post.Body)
	if err != nil {
		return fmt.Errorf("failed to convert HTML to Markdown: %w", err)
	}

	// Generate slug if needed
	slug := post.Slug
	if slug == "" && metadata.Slug != "" {
		slug = metadata.Slug
	}
	if slug == "" {
		slug = slugify(post.Title)
	}
	if slug == "" {
		return fmt.Errorf("could not generate slug for post")
	}

	// Determine file path
	filePath, err := determineFilePath(metadata, slug, config.OutputDir)
	if err != nil {
		return fmt.Errorf("failed to determine file path: %w", err)
	}

	// Create frontmatter and content
	content, err := createPostContent(post, metadata, slug, bodyMd)
	if err != nil {
		return fmt.Errorf("failed to create post content: %w", err)
	}

	// Write file
	if config.DryRun {
		logger.Printf("DRY RUN: Would write to %s (%d bytes)", filePath, len(content))
		return nil
	}

	if err := writeFile(filePath, content); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if config.Verbose {
		logger.Printf("Written post to: %s", filePath)
	}

	return nil
}

// parseMetadata safely parses JSON metadata
func parseMetadata(metadataStr string) (*PostMetadata, error) {
	metadata := &PostMetadata{}

	if metadataStr == "" {
		return metadata, nil
	}

	// Try to parse as JSON
	if err := json.Unmarshal([]byte(metadataStr), metadata); err != nil {
		// If JSON parsing fails, create empty metadata and log warning
		log.Printf("Warning: Failed to parse metadata as JSON, using defaults: %v", err)
		return &PostMetadata{}, nil
	}

	return metadata, nil
}

// determineFilePath determines the output file path based on metadata
func determineFilePath(metadata *PostMetadata, slug, outputDir string) (string, error) {
	var dir string

	// Determine directory from metadata or post type
	if metadata.PostDir != "" {
		dir = filepath.Join(outputDir, metadata.PostDir)
	} else {
		postDir := strings.ToLower(metadata.Type)
		if postDir == "" {
			postDir = "posts"
		}
	}

	filename := fmt.Sprintf("%s.md", slug)
	return filepath.Join(dir, filename), nil
}

// createPostContent creates the final markdown content with frontmatter
func createPostContent(post Post, metadata *PostMetadata, slug, bodyMd string) (string, error) {
	// Parse dates
	createdTime, err := parseTime(post.Created)
	if err != nil {
		log.Printf("Warning: Invalid created time for post %d, using current time", post.ID)
		createdTime = time.Now()
	}

	updatedTime, err := parseTime(post.Updated)
	if err != nil {
		updatedTime = createdTime
	}

	// Build frontmatter
	frontmatter := map[string]interface{}{
		"title":   post.Title,
		"slug":    slug,
		"date":    createdTime.Format(rfc3339Format),
		"lastmod": updatedTime.Format(rfc3339Format),
	}

	// Add metadata fields to frontmatter
	if metadata.Type != "" {
		frontmatter["type"] = metadata.Type
	}
	if len(metadata.Tags) > 0 {
		frontmatter["tags"] = metadata.Tags
	}
	if len(metadata.Categories) > 0 {
		frontmatter["categories"] = metadata.Categories
	}
	if metadata.Description != "" {
		frontmatter["description"] = metadata.Description
	}
	if metadata.Draft {
		frontmatter["draft"] = true
	}

	// Convert frontmatter to YAML
	var frontmatterStr strings.Builder
	frontmatterStr.WriteString("---\n")

	for key, value := range frontmatter {
		switch v := value.(type) {
		case string:
			frontmatterStr.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, strings.ReplaceAll(v, "\"", "\\\"")))
		case bool:
			frontmatterStr.WriteString(fmt.Sprintf("%s: %t\n", key, v))
		case []string:
			if len(v) > 0 {
				frontmatterStr.WriteString(fmt.Sprintf("%s:\n", key))
				for _, item := range v {
					frontmatterStr.WriteString(fmt.Sprintf("  - \"%s\"\n", strings.ReplaceAll(item, "\"", "\\\"")))
				}
			}
		default:
			frontmatterStr.WriteString(fmt.Sprintf("%s: %v\n", key, v))
		}
	}

	frontmatterStr.WriteString("---\n\n")

	return frontmatterStr.String() + bodyMd, nil
}

// parseTime attempts to parse time from various formats
func parseTime(timeStr string) (time.Time, error) {
	timeStr = strings.TrimSpace(timeStr)
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	// Try RFC3339 first
	if t, err := time.Parse(rfc3339Format, timeStr); err == nil {
		return t, nil
	}

	// Try common SQLite format
	if t, err := time.Parse(timeFormat, timeStr); err == nil {
		return t, nil
	}

	// Try without seconds
	if t, err := time.Parse("2006-01-02 15:04", timeStr); err == nil {
		return t, nil
	}

	// Try date only
	if t, err := time.Parse("2006-01-02", timeStr); err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("could not parse time: %s", timeStr)
}

// writeFile safely writes content to file, creating directories as needed
func writeFile(filePath, content string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// slugify creates a URL-friendly slug from a string
func slugify(input string) string {
	if input == "" {
		return ""
	}

	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace common characters
	replacements := map[string]string{
		"&":  "and",
		"@":  "at",
		"+":  "plus",
		"=":  "equals",
		"%":  "percent",
		"#":  "hash",
		"'":  "",
		"\"": "",
	}

	for old, new := range replacements {
		slug = strings.ReplaceAll(slug, old, new)
	}

	// Remove non-alphanumeric characters except spaces and hyphens
	reg := regexp.MustCompile(`[^a-z0-9\s\-]+`)
	slug = reg.ReplaceAllString(slug, "")

	// Replace multiple spaces/hyphens with single hyphen
	reg = regexp.MustCompile(`[\s\-]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	// Limit length
	if len(slug) > 60 {
		slug = slug[:60]
		// Try to cut at word boundary
		if lastHyphen := strings.LastIndex(slug, "-"); lastHyphen > 30 {
			slug = slug[:lastHyphen]
		}
	}

	return slug
}
