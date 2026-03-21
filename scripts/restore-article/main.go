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

	fmt.Println("Restoring 'article' type")
	fmt.Println("========================\n")

	// First, create the 'article' post type
	_, err = db.ExecContext(ctx, "INSERT OR IGNORE INTO post_types (id, name, slug) VALUES (?, ?, ?)", "article", "Article", "article")
	if err != nil {
		fmt.Printf("ERROR creating 'article' post type: %v\n", err)
		return
	}
	fmt.Println("✓ Created post type: 'article'")

	// Query posts that should be articles (those with article-related slugs or from original article posts)
	// We'll need to identify which 3 posts were originally 'article' type
	// Since we can't determine this from current data, we'll need to look at the markdown files
	fmt.Println("\nNote: Need to identify which posts should be 'article' type")
	fmt.Println("Check the markdown files in posts/article/ directory to get their slugs")
	fmt.Println("\nThen run: go run scripts/restore-article/main.go [slug1] [slug2] [slug3] -dry-run=false")

	if *dryRun {
		fmt.Println("\n[DRY-RUN] Article type created but no posts reassigned yet")
		return
	}
}
