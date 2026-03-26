package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	ctx := context.Background()

	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	if dbName == "" || dbToken == "" {
		fmt.Printf("❌ Missing Turso credentials\n")
		os.Exit(1)
	}

	// Build connection string
	conn := dbName
	if !strings.HasPrefix(conn, "libsql://") && !strings.HasPrefix(conn, "https://") {
		conn = "libsql://" + conn
	}
	if dbToken != "" {
		sep := "?"
		if strings.Contains(conn, "?") {
			sep = "&"
		}
		conn = fmt.Sprintf("%s%sauthToken=%s", conn, sep, dbToken)
	}

	db, err := sql.Open("libsql", conn)
	if err != nil {
		fmt.Printf("❌ Database connection failed: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Get all newsletter posts
	rows, err := db.QueryContext(ctx, "SELECT slug FROM posts WHERE slug LIKE 'newsletter/%' ORDER BY slug")
	if err != nil {
		fmt.Printf("❌ Query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var slugs []string
	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			fmt.Printf("⚠️  Scan error: %v\n", err)
			continue
		}
		slugs = append(slugs, slug)
	}

	fmt.Printf("📋 Found %d newsletter posts in DB\n", len(slugs))
	for _, slug := range slugs {
		fmt.Printf("  - %s\n", slug)
	}

	if len(slugs) == 0 {
		fmt.Printf("✅ No newsletter posts to delete\n")
		os.Exit(0)
	}

	// Confirm deletion
	fmt.Printf("\n⚠️  About to DELETE all %d newsletter posts. Type 'yes' to confirm: ", len(slugs))
	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "yes" {
		fmt.Printf("❌ Deletion cancelled\n")
		os.Exit(1)
	}

	// Delete all newsletter posts
	q := libsqlssg.New(db)
	deleted := 0

	for _, slug := range slugs {
		err := q.DeletePost(ctx, slug)
		if err != nil {
			fmt.Printf("❌ Error deleting %s: %v\n", slug, err)
			continue
		}
		fmt.Printf("🗑️  Deleted: %s\n", slug)
		deleted++
	}

	fmt.Printf("\n✅ Deleted %d newsletter posts\n", deleted)
}
