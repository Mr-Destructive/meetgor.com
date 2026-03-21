package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
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

	// Get all posts
	rows, err := db.QueryContext(ctx,
		`SELECT id, title, slug, type_id, tags, metadata FROM posts ORDER BY created_at DESC`)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	fmt.Println("\nPost Tags Sync Check:")
	fmt.Println("=====================\n")

	totalPosts := 0
	for rows.Next() {
		totalPosts++
		var id, title, slug, typeID string
		var tags sql.NullString
		var metadata sql.NullString

		if err := rows.Scan(&id, &title, &slug, &typeID, &tags, &metadata); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}

		if !metadata.Valid {
			continue
		}

		var meta map[string]interface{}
		if err := json.Unmarshal([]byte(metadata.String), &meta); err != nil {
			continue
		}

		// Check if tags exist in metadata
		metaTags, hasMetaTags := meta["tags"]
		dbTags := ""
		if tags.Valid {
			dbTags = tags.String
		}

		fmt.Printf("[%s] %s (%s)\n", typeID, slug, id)
		fmt.Printf("   Title: %s\n", title)
		if hasMetaTags && metaTags != nil {
			fmt.Printf("   Metadata tags: %v\n", metaTags)
		} else {
			fmt.Printf("   Metadata tags: (none)\n")
		}
		if len(dbTags) > 0 {
			fmt.Printf("   DB tags: %s\n", dbTags)
		} else {
			fmt.Printf("   DB tags: (empty)\n")
		}
		fmt.Println()
	}

	fmt.Printf("\nSummary: %d total posts in database\n", totalPosts)
}
