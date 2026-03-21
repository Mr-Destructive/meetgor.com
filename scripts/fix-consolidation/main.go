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

	fmt.Println("Fixing consolidation: 'article' → 'post'")
	fmt.Println("=========================================\n")

	// Check posts with 'article' type that came from 'posts'
	var articleCount int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE type_id = 'article'").Scan(&articleCount)
	fmt.Printf("Posts with type 'article': %d\n\n", articleCount)

	if articleCount == 0 {
		fmt.Println("No posts with 'article' type found")
		return
	}

	if *dryRun {
		fmt.Println("[DRY-RUN] Changes that would be made:")
		fmt.Printf("  - %d posts: 'article' → 'post'\n", articleCount)
		fmt.Println("\nTo apply changes, run: go run scripts/fix-consolidation/main.go -dry-run=false")
		return
	}

	fmt.Println("Applying fix...")

	// Change 'article' back to 'post'
	_, err = db.ExecContext(ctx, "UPDATE posts SET type_id = 'post' WHERE type_id = 'article'")
	if err != nil {
		fmt.Printf("ERROR updating 'article' to 'post': %v\n", err)
		return
	}
	fmt.Printf("✓ Updated %d posts: 'article' → 'post'\n", articleCount)

	// Delete 'article' post type and recreate 'post'
	_, err = db.ExecContext(ctx, "DELETE FROM post_types WHERE id = 'article'")
	if err != nil {
		fmt.Printf("ERROR deleting 'article' post type: %v\n", err)
		return
	}

	_, err = db.ExecContext(ctx, "INSERT OR IGNORE INTO post_types (id, name, slug) VALUES (?, ?, ?)", "post", "Post", "post")
	if err != nil {
		fmt.Printf("ERROR creating 'post' post type: %v\n", err)
		return
	}

	fmt.Println("✓ Deleted post type: 'article'")
	fmt.Println("✓ Created post type: 'post'")
	fmt.Println("\nFix complete")
}
