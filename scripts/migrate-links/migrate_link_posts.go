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

func loadEnvMigration(filename string) {
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

func migrationMain() {
	dryRun := flag.Bool("dry-run", true, "Show what would be changed without making changes")
	flag.Parse()

	loadEnvMigration(".env")

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

	// Find all link-type posts that need migration
	// Look for posts where metadata contains type='links' OR type_id is link/links
	rows, err := db.QueryContext(ctx,
		`SELECT id, title, slug, content, metadata 
		 FROM posts 
		 WHERE (type_id IN ('link', 'links') OR metadata LIKE '%"type":"links"%') 
		 AND status != 'deleted'`)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	var updates []struct {
		id       string
		title    string
		slug     string
		oldBody  string
		newBody  string
		oldMeta  string
		newMeta  string
	}

	for rows.Next() {
		var id, title, slug, content string
		var metadata sql.NullString
		if err := rows.Scan(&id, &title, &slug, &content, &metadata); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}

		if !metadata.Valid {
			// No metadata, skip
			continue
		}

		var meta map[string]interface{}
		if err := json.Unmarshal([]byte(metadata.String), &meta); err != nil {
			fmt.Printf("SKIP %s: invalid JSON metadata\n", slug)
			continue
		}

		// Check if needs migration
		hasCommentary := meta["commentary"] != nil
		hasTitle := meta["title"] != nil

		if !hasCommentary && !hasTitle {
			// Already clean, skip
			continue
		}

		// Extract new content from commentary
		newContent := content
		if commentary, ok := meta["commentary"].(string); ok && commentary != "" && content == "" {
			newContent = commentary
		}

		// Build clean metadata
		cleanMeta := make(map[string]interface{})
		if link, ok := meta["link"]; ok {
			cleanMeta["url"] = link
		}
		if date, ok := meta["date"]; ok {
			cleanMeta["date"] = date
		}
		if tags, ok := meta["tags"]; ok {
			cleanMeta["tags"] = tags
		}

		newMetaBytes, _ := json.Marshal(cleanMeta)

		updates = append(updates, struct {
			id      string
			title   string
			slug    string
			oldBody string
			newBody string
			oldMeta string
			newMeta string
		}{
			id:      id,
			title:   title,
			slug:    slug,
			oldBody: content,
			newBody: newContent,
			oldMeta: metadata.String,
			newMeta: string(newMetaBytes),
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
		if u.oldBody != u.newBody {
			fmt.Printf("    Content: %d → %d bytes\n", len(u.oldBody), len(u.newBody))
		}
		fmt.Printf("    Metadata: %d → %d bytes\n", len(u.oldMeta), len(u.newMeta))
	}

	if *dryRun {
		fmt.Println("\n[DRY-RUN] To apply changes, run: go run scripts/migrate_link_posts.go -dry-run=false")
		return
	}

	fmt.Println("\nApplying migrations...")
	for _, u := range updates {
		_, err := db.ExecContext(ctx,
			"UPDATE posts SET content = ?, metadata = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
			u.newBody, u.newMeta, u.id)

		if err != nil {
			fmt.Printf("ERROR updating %s: %v\n", u.slug, err)
			continue
		}
		fmt.Printf("✓ Migrated %s\n", u.slug)
	}

	fmt.Printf("\nMigration complete: %d posts updated\n", len(updates))
}

func main() {
	migrationMain()
}
