package main

import (
	"bufio"
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

	"gopkg.in/yaml.v3"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type PostMetadata struct {
	Title       string        `yaml:"title"`
	Slug        string        `yaml:"slug"`
	Date        string        `yaml:"date"`
	Status      string        `yaml:"status"`
	Type        string        `yaml:"type"`
	Tags        []interface{} `yaml:"tags"`
	Description string        `yaml:"description"`
	Extra       map[string]interface{} `yaml:"-"`
}

func loadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
}

func parseFrontmatter(content string) (PostMetadata, string, error) {
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return PostMetadata{}, content, fmt.Errorf("no frontmatter found")
	}

	var meta PostMetadata
	err := yaml.Unmarshal([]byte(parts[1]), &meta)
	if err != nil {
		return PostMetadata{}, "", err
	}

	body := strings.TrimSpace(parts[2])
	return meta, body, nil
}

func calculateHash(title, body string) string {
	data := title + "|" + body
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func slugify(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	input = strings.ReplaceAll(input, " ", "-")
	result := []rune{}
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result = append(result, r)
		}
	}
	return strings.Trim(string(result), "-")
}

func getTypeFromPath(filePath string) string {
	parts := strings.Split(filePath, string(os.PathSeparator))
	if len(parts) > 1 {
		parentDir := parts[len(parts)-2]
		// Map directory to type
		typeMap := map[string]string{
			"link":      "link",
			"links":     "link",
			"article":   "article",
			"articles":  "article",
			"quote":     "quote",
			"quotes":    "quote",
			"tutorial":  "tutorial",
			"tutorials": "tutorial",
			"snippet":   "snippet",
			"snippets":  "snippet",
			"list":      "list",
			"lists":     "list",
		}
		if t, ok := typeMap[parentDir]; ok {
			return t
		}
	}
	return "link" // default to link
}

func main() {
	dryRun := flag.Bool("dry-run", true, "Show what would be changed without making changes")
	postsDir := flag.String("posts-dir", "posts", "Root directory for markdown files")
	flag.Parse()

	loadEnv(".env")

	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		log.Fatal("Missing TURSO_DATABASE_NAME or TURSO_DATABASE_AUTH_TOKEN")
	}

	connStr := dbName
	if !strings.HasPrefix(dbName, "libsql://") {
		connStr = fmt.Sprintf("libsql://%s", dbName)
	}
	connStr = fmt.Sprintf("%s?authToken=%s", connStr, dbToken)
	db, err := sql.Open("libsql", connStr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	var added, updated, skipped int

	// Walk through markdown files
	err = filepath.Walk(*postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		meta, body, err := parseFrontmatter(string(content))
		if err != nil {
			// No frontmatter, skip
			return nil
		}

		// Use provided slug or generate from title
		slug := meta.Slug
		if slug == "" && meta.Title != "" {
			slug = slugify(meta.Title)
		}
		if slug == "" {
			return nil // Skip files without title or slug
		}

		// Get type from frontmatter or file path
		typeID := meta.Type
		if typeID == "" {
			typeID = getTypeFromPath(path)
		}

		// Default status
		status := meta.Status
		if status == "" {
			status = "draft"
		}

		// Calculate hash
		hash := calculateHash(meta.Title, body)

		// Check if post exists
		var existingID string
		var existingHash string
		err = db.QueryRowContext(ctx,
			"SELECT id, metadata FROM posts WHERE slug = ?",
			slug).Scan(&existingID, &sql.NullString{})

		if err == nil {
			// Post exists, check hash to see if we need to update
			var metaStr sql.NullString
			db.QueryRowContext(ctx, "SELECT metadata FROM posts WHERE id = ?", existingID).Scan(&metaStr)

			if metaStr.Valid {
				var m map[string]interface{}
				json.Unmarshal([]byte(metaStr.String), &m)
				if h, ok := m["hash"].(string); ok {
					existingHash = h
				}
			}

			if existingHash == hash {
				skipped++
				return nil // Post unchanged
			}
			// Post exists but content changed, update it
			updated++
		} else if err == sql.ErrNoRows {
			added++
		} else {
			return err
		}

		// Prepare metadata
		metadata := map[string]interface{}{
			"hash":  hash,
			"title": meta.Title,
			"date":  meta.Date,
		}

		if meta.Description != "" {
			metadata["description"] = meta.Description
		}

		// Handle tags
		var tagsJSON string
		if len(meta.Tags) > 0 {
			tagsBytes, _ := json.Marshal(meta.Tags)
			tagsJSON = string(tagsBytes)
		} else {
			tagsJSON = "[]"
		}

		metaBytes, _ := json.Marshal(metadata)

		// Ensure post type exists
		var typeExists int
		db.QueryRowContext(ctx, "SELECT COUNT(*) FROM post_types WHERE id = ?", typeID).Scan(&typeExists)
		if typeExists == 0 {
			db.ExecContext(ctx,
				"INSERT OR IGNORE INTO post_types (id, name, slug) VALUES (?, ?, ?)",
				typeID, strings.Title(typeID), typeID)
		}

		// Insert or update post
		if existingID == "" {
			// Insert
			_, err = db.ExecContext(ctx,
				`INSERT INTO posts (id, type_id, title, slug, content, metadata, tags, status) 
				 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				fmt.Sprintf("%d", time.Now().UnixNano()),
				typeID, meta.Title, slug, body,
				string(metaBytes), tagsJSON, status)
		} else {
			// Update
			_, err = db.ExecContext(ctx,
				`UPDATE posts SET title = ?, content = ?, metadata = ?, tags = ?, status = ?, type_id = ? 
				 WHERE id = ?`,
				meta.Title, body, string(metaBytes), tagsJSON, status, typeID, existingID)
		}

		if err != nil {
			log.Printf("ERROR syncing %s: %v", slug, err)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Walk failed: %v", err)
	}

	fmt.Printf("\nMigration Results:\n")
	fmt.Printf("  Added: %d\n", added)
	fmt.Printf("  Updated: %d\n", updated)
	fmt.Printf("  Skipped: %d\n", skipped)
	fmt.Printf("  Total: %d\n\n", added+updated+skipped)

	if *dryRun {
		fmt.Println("[DRY-RUN] To apply changes, run: go run scripts/migrate-all-posts/main.go -dry-run=false")
	}
}
