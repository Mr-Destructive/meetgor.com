package main

import (
	"bufio"
	"context"
	"database/sql"
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

	fmt.Println("Consolidating post types: 'links' → 'link', 'posts' → 'article'")
	fmt.Println("================================================================\n")

	// Check posts with 'links' type
	var linksCount int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE type_id = 'links'").Scan(&linksCount)
	fmt.Printf("Posts with type 'links': %d\n", linksCount)

	// Check posts with 'posts' type
	var postsCount int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE type_id = 'posts'").Scan(&postsCount)
	fmt.Printf("Posts with type 'posts': %d\n\n", postsCount)

	if linksCount == 0 && postsCount == 0 {
		fmt.Println("No posts with 'links' or 'posts' type found")
		return
	}

	if *dryRun {
		fmt.Println("[DRY-RUN] Changes that would be made:")
		fmt.Printf("  - %d posts: 'links' → 'link'\n", linksCount)
		fmt.Printf("  - %d posts: 'posts' → 'article'\n", postsCount)
		fmt.Println("\nTo apply changes, run: go run scripts/consolidate-types/main.go -dry-run=false")
		return
	}

	fmt.Println("Applying consolidation...")

	// Consolidate 'links' to 'link'
	if linksCount > 0 {
		_, err = db.ExecContext(ctx, "UPDATE posts SET type_id = 'link' WHERE type_id = 'links'")
		if err != nil {
			fmt.Printf("ERROR updating 'links' to 'link': %v\n", err)
		} else {
			fmt.Printf("✓ Updated %d posts: 'links' → 'link'\n", linksCount)
		}
	}

	// Consolidate 'posts' to 'article'
	if postsCount > 0 {
		_, err = db.ExecContext(ctx, "UPDATE posts SET type_id = 'article' WHERE type_id = 'posts'")
		if err != nil {
			fmt.Printf("ERROR updating 'posts' to 'article': %v\n", err)
		} else {
			fmt.Printf("✓ Updated %d posts: 'posts' → 'article'\n", postsCount)
		}
	}

	// Now delete the old post types
	_, err = db.ExecContext(ctx, "DELETE FROM post_types WHERE id IN ('links', 'posts')")
	if err != nil {
		fmt.Printf("ERROR deleting post types: %v\n", err)
		return
	}

	fmt.Println("✓ Deleted post types: 'links', 'posts'")
	fmt.Println("\nConsolidation complete")
}
