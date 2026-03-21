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

	// Fix specific posts that should be link
	fixes := []struct {
		slug      string
		newTypeID string
	}{
		{"stitch-design-with-ai", "link"},
		{"give-django-your-time-and-money-not-your-tokens", "link"},
	}

	var updates []struct {
		slug    string
		id      string
		title   string
		oldType string
		newType string
	}

	for _, fix := range fixes {
		var id, title, typeID string
		err := db.QueryRowContext(ctx,
			"SELECT id, title, type_id FROM posts WHERE slug = ? AND status != 'deleted'",
			fix.slug).Scan(&id, &title, &typeID)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("Post not found: %s\n", fix.slug)
			} else {
				fmt.Printf("Error querying %s: %v\n", fix.slug, err)
			}
			continue
		}

		if typeID != fix.newTypeID {
			updates = append(updates, struct {
				slug    string
				id      string
				title   string
				oldType string
				newType string
			}{
				slug:    fix.slug,
				id:      id,
				title:   title,
				oldType: typeID,
				newType: fix.newTypeID,
			})
		}
	}

	if len(updates) == 0 {
		fmt.Println("No posts need type migration")
		return
	}

	fmt.Printf("\nFound %d posts to fix\n\n", len(updates))

	for i, u := range updates {
		fmt.Printf("[%d] %s (%s)\n", i+1, u.slug, u.id)
		fmt.Printf("    Title: %s\n", u.title)
		fmt.Printf("    Type: %s → %s\n\n", u.oldType, u.newType)
	}

	if *dryRun {
		fmt.Println("[DRY-RUN] To apply changes, run: go run scripts/fix-post-types/main.go -dry-run=false")
		return
	}

	fmt.Println("Applying migrations...")
	for _, u := range updates {
		_, err := db.ExecContext(ctx,
			"UPDATE posts SET type_id = ? WHERE id = ?",
			u.newType, u.id)

		if err != nil {
			fmt.Printf("ERROR updating %s: %v\n", u.slug, err)
			continue
		}
		fmt.Printf("✓ Fixed type for %s: %s → %s\n", u.slug, u.oldType, u.newType)
	}

	fmt.Printf("\nMigration complete: %d posts updated\n", len(updates))
}
