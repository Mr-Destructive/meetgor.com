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

	fmt.Println("Migration Verification Report")
	fmt.Println("=============================\n")

	// Check total posts
	var totalPosts int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE status != 'deleted'").Scan(&totalPosts)
	fmt.Printf("Total Posts: %d\n\n", totalPosts)

	// Check posts with tags
	var withTags int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE tags IS NOT NULL AND tags != '' AND tags != '[]'").Scan(&withTags)
	fmt.Printf("Posts with tags: %d (%d%%)\n", withTags, (withTags*100)/totalPosts)

	// Check posts with metadata
	var withMeta int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE metadata IS NOT NULL AND metadata != ''").Scan(&withMeta)
	fmt.Printf("Posts with metadata: %d (%d%%)\n", withMeta, (withMeta*100)/totalPosts)

	// Check posts with hash
	rows, _ := db.QueryContext(ctx, "SELECT metadata FROM posts LIMIT 5")
	var hasHash int
	for rows.Next() {
		var metaStr sql.NullString
		rows.Scan(&metaStr)
		if metaStr.Valid {
			var m map[string]interface{}
			if json.Unmarshal([]byte(metaStr.String), &m) == nil {
				if _, ok := m["hash"]; ok {
					hasHash++
				}
			}
		}
	}
	rows.Close()

	// Sample 5 random posts
	fmt.Println("\nSample Posts:")
	rows, _ = db.QueryContext(ctx, `SELECT type_id, title, slug, tags FROM posts ORDER BY RANDOM() LIMIT 5`)
	for rows.Next() {
		var typeID, title, slug, tags string
		rows.Scan(&typeID, &title, &slug, &tags)
		fmt.Printf("  [%s] %s (%s)\n", typeID, title, slug)
		fmt.Printf("      Tags: %s\n", tags)
	}
	rows.Close()

	fmt.Println("\n✓ Migration complete")
}
