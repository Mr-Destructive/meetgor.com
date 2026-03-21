package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

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

func main() {
	dryRun := flag.Bool("dry-run", true, "Show what would be changed without making changes")
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

	// Find all posts where metadata contains tags but tags field is empty/null
	rows, err := db.QueryContext(ctx,
		`SELECT id, title, slug, metadata, tags FROM posts WHERE status != 'deleted'`)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	var updates []struct {
		id       string
		title    string
		slug     string
		oldTags  string
		newTags  string
	}

	for rows.Next() {
		var id, title, slug string
		var metadata sql.NullString
		var tags sql.NullString

		if err := rows.Scan(&id, &title, &slug, &metadata, &tags); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}

		if !metadata.Valid {
			continue
		}

		var meta map[string]interface{}
		if err := json.Unmarshal([]byte(metadata.String), &meta); err != nil {
			continue
		}

		// Check if metadata has tags but db tags field is empty
		metaTags, hasMetaTags := meta["tags"]
		if !hasMetaTags || metaTags == nil {
			continue
		}

		oldTagsStr := ""
		if tags.Valid {
			oldTagsStr = tags.String
		}

		// Already has tags, skip
		if oldTagsStr != "" && oldTagsStr != "[]" && oldTagsStr != "null" {
			continue
		}

		// Extract tags from metadata
		var newTagsJSON string
		switch v := metaTags.(type) {
		case string:
			// Tags as comma-separated string
			tagArr := strings.Split(v, ",")
			for i := range tagArr {
				tagArr[i] = strings.TrimSpace(tagArr[i])
			}
			tagsBytes, _ := json.Marshal(tagArr)
			newTagsJSON = string(tagsBytes)
		case []interface{}:
			// Tags as array
			tagsBytes, _ := json.Marshal(v)
			newTagsJSON = string(tagsBytes)
		default:
			continue
		}

		updates = append(updates, struct {
			id      string
			title   string
			slug    string
			oldTags string
			newTags string
		}{
			id:      id,
			title:   title,
			slug:    slug,
			oldTags: oldTagsStr,
			newTags: newTagsJSON,
		})
	}

	if len(updates) == 0 {
		fmt.Println("No posts need migration")
		return
	}

	fmt.Printf("\nFound %d posts to migrate\n\n", len(updates))

	for i, u := range updates {
		fmt.Printf("[%d] %s (%s)\n", i+1, u.slug, u.id)
		fmt.Printf("    Title: %s\n", u.title)
		fmt.Printf("    Tags: %s → %s\n\n", u.oldTags, u.newTags)
	}

	if *dryRun {
		fmt.Println("[DRY-RUN] To apply changes, run: go run scripts/migrate-metadata-tags/main.go -dry-run=false")
		return
	}

	fmt.Println("Applying migrations...")
	for _, u := range updates {
		_, err := db.ExecContext(ctx,
			"UPDATE posts SET tags = ? WHERE id = ?",
			u.newTags, u.id)

		if err != nil {
			fmt.Printf("ERROR updating %s: %v\n", u.slug, err)
			continue
		}
		fmt.Printf("✓ Migrated %s\n", u.slug)
	}

	fmt.Printf("\nMigration complete: %d posts updated\n", len(updates))
}
