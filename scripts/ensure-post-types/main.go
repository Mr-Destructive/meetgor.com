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

	// Ensure all post types exist
	postTypes := []struct {
		id   string
		name string
	}{
		{"link", "Link"},
		{"posts", "Posts"},
		{"projects", "Projects"},
		{"work", "Work"},
		{"til", "TIL"},
		{"newsletter", "Newsletter"},
		{"sqlog", "SQLog"},
		{"article", "Article"},
		{"quote", "Quote"},
		{"tutorial", "Tutorial"},
		{"snippet", "Snippet"},
		{"post", "Post"},
	}

	var created []string

	for _, pt := range postTypes {
		var exists int
		err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM post_types WHERE id = ?", pt.id).Scan(&exists)
		if err != nil || exists == 0 {
			if !*dryRun {
				_, err = db.ExecContext(ctx,
					"INSERT OR IGNORE INTO post_types (id, name, slug) VALUES (?, ?, ?)",
					pt.id, pt.name, pt.id)
				if err != nil {
					fmt.Printf("ERROR creating post type %s: %v\n", pt.id, err)
					continue
				}
			}
			created = append(created, pt.id)
		}
	}

	if len(created) == 0 {
		fmt.Println("All post types already exist")
		return
	}

	fmt.Printf("Found %d post types to create:\n", len(created))
	for _, id := range created {
		fmt.Printf("  - %s\n", id)
	}

	if *dryRun {
		fmt.Println("\n[DRY-RUN] To apply changes, run: go run scripts/ensure-post-types/main.go -dry-run=false")
	} else {
		fmt.Printf("\nCreated %d post types\n", len(created))
	}
}
