package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// SyncOptions defines the sync configuration
type SyncOptions struct {
	Root     string
	AuthorID int64
	DryRun   bool
}

// SyncStats holds statistics about the sync operation
type SyncStats struct {
	Added   int
	Updated int
	Skipped int
}

// Post metadata structure
type PostMetadata struct {
	Title       string
	Slug        string
	ContentHash string
	Type        string
}

func setupHashDB(t interface{ Fatalf(string, ...interface{}) }) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	ddl := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		is_admin BOOLEAN DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS post_types (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		slug TEXT NOT NULL UNIQUE,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS posts (
		id TEXT PRIMARY KEY,
		type_id TEXT NOT NULL,
		title TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE,
		content TEXT NOT NULL,
		excerpt TEXT,
		status TEXT DEFAULT 'draft',
		is_featured INTEGER DEFAULT 0,
		tags TEXT,
		metadata TEXT,
		created_at REAL,
		updated_at REAL,
		published_at REAL,
		CONSTRAINT fk_posts_type_id FOREIGN KEY (type_id) REFERENCES post_types(id)
	);
	`

	if _, err := db.Exec(ddl); err != nil {
		t.Fatalf("exec ddl: %v", err)
	}

	return db
}

func syncMarkdownFiles(ctx context.Context, db *sql.DB, opts SyncOptions) (SyncStats, error) {
	stats := SyncStats{}

	// Walk through markdown files in root directory
	err := filepath.Walk(opts.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .md files
		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		// Read markdown file
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Parse front matter and content
		meta, body := parseFrontMatter(string(content))

		// Calculate content hash based on title+body only (not metadata)
		hash := calculateHash(meta["title"], body)

		// Check if post exists and hash matches
		slug := meta["slug"]
		if slug == "" {
			slug = slugify(meta["title"])
		}

		q := libsqlssg.New(db)
		posts, err := q.GetPostsBySlugType(ctx, slug)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		// Extract hash from existing post metadata to compare
		existingHash := ""
		if len(posts) > 0 && posts[0].Metadata.Valid {
			existingMeta := parseMetadata(posts[0].Metadata.String)
			existingHash = existingMeta["hash"]
		}

		if existingHash == hash {
			// Post exists and hash matches - skip
			stats.Skipped++
			return nil
		}

		// Update metadata with hash for comparison on next sync
		meta["hash"] = hash
		meta["slug"] = slug

		// Prepare post data
		postType := meta["type"]
		if postType == "" {
			postType = "post"
		}

		// Build metadata JSON
		metadataJSON := buildMetadataJSON(meta)

		// Only insert/update if not dry-run
		if !opts.DryRun {
			if len(posts) == 0 {
				// Insert new post
				_, err = q.CreatePost(ctx, libsqlssg.CreatePostParams{
					ID:       generateID(),
					TypeID:   postType,
					Title:    meta["title"],
					Slug:     slug,
					Content:  body,
					Metadata: sql.NullString{String: metadataJSON, Valid: true},
				})
				if err != nil {
					return err
				}
				stats.Added++
			} else {
				// Update existing post
				err = q.UpdatePost(ctx, libsqlssg.UpdatePostParams{
					Title:    meta["title"],
					Content:  body,
					Metadata: sql.NullString{String: metadataJSON, Valid: true},
					Slug:     slug,
				})
				if err != nil {
					return err
				}
				stats.Updated++
			}

			// Rewrite file with updated metadata
			updatedContent := reconstructFrontMatter(meta, body)
			if err := ioutil.WriteFile(path, []byte(updatedContent), 0644); err != nil {
				return err
			}
		} else {
			// Dry-run: just count what would be added
			if len(posts) == 0 {
				stats.Added++
			} else {
				stats.Updated++
			}
		}

		return nil
	})

	return stats, err
}

func parseMetadata(jsonStr string) map[string]string {
	meta := make(map[string]string)
	if jsonStr == "" {
		return meta
	}
	
	// Simple JSON parser for metadata
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return meta
	}
	
	for k, v := range obj {
		if str, ok := v.(string); ok {
			meta[k] = str
		}
	}
	return meta
}

func parseFrontMatter(content string) (map[string]string, string) {
	meta := make(map[string]string)

	lines := strings.Split(content, "\n")
	if len(lines) < 2 || lines[0] != "---" {
		return meta, content
	}

	// Find closing ---
	endIdx := -1
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			endIdx = i
			break
		}
	}

	if endIdx == -1 {
		return meta, content
	}

	// Parse front matter
	for i := 1; i < endIdx; i++ {
		parts := strings.SplitN(lines[i], ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			meta[key] = value
		}
	}

	// Body is everything after closing ---
	body := strings.Join(lines[endIdx+1:], "\n")
	body = strings.TrimSpace(body)

	return meta, body
}

func reconstructFrontMatter(meta map[string]string, body string) string {
	lines := []string{"---"}

	for key, value := range meta {
		lines = append(lines, fmt.Sprintf("%s: %s", key, value))
	}

	lines = append(lines, "---")
	lines = append(lines, body)

	return strings.Join(lines, "\n")
}

func calculateHash(title, body string) string {
	data := title + "|" + body
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func slugify(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	input = strings.ReplaceAll(input, " ", "-")
	// Remove non-alphanumeric except dash
	result := []rune{}
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result = append(result, r)
		}
	}
	return strings.Trim(string(result), "-")
}

func buildMetadataJSON(meta map[string]string) string {
	// Simple JSON builder for metadata
	pairs := []string{}
	for key, value := range meta {
		// Escape quotes in value
		value = strings.ReplaceAll(value, `"`, `\"`)
		pairs = append(pairs, fmt.Sprintf(`"%s": "%s"`, key, value))
	}
	return "{" + strings.Join(pairs, ", ") + "}"
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func main() {
	var (
		root     = flag.String("root", "posts", "Root directory for markdown files")
		authorID = flag.Int64("author", 1, "Default author ID")
		dryRun   = flag.Bool("dry-run", false, "Dry run - show what would happen without writing")
	)

	flag.Parse()

	ctx := context.Background()

	// Open database
	dbURL := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	if dbURL == "" || dbToken == "" {
		log.Fatalf("TURSO_DATABASE_NAME and TURSO_DATABASE_AUTH_TOKEN environment variables required")
	}

	connStr := fmt.Sprintf("libsql://%s?authToken=%s", dbURL, dbToken)
	db, err := sql.Open("libsql", connStr)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer db.Close()

	// Run sync
	stats, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     *root,
		AuthorID: *authorID,
		DryRun:   *dryRun,
	})
	if err != nil {
		log.Fatalf("sync failed: %v", err)
	}

	// Log results
	if *dryRun {
		log.Printf("[DRY-RUN] Would add=%d, update=%d, skip=%d", stats.Added, stats.Updated, stats.Skipped)
	} else {
		log.Printf("[SYNC] added=%d updated=%d skipped=%d", stats.Added, stats.Updated, stats.Skipped)
	}
}
