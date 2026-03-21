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

	descriptions := map[string]string{
		"link":       "Links to external articles, resources, and interesting web content",
		"post":       "Regular blog posts covering various topics and thoughts",
		"article":    "In-depth article content with detailed analysis and insights",
		"newsletter": "Content from weekly newsletters and email communications",
		"sqlog":      "SQL-related posts and database tutorials from the SQLog series",
		"til":        "Today I Learned - short posts about quick tips and discoveries",
		"tutorial":   "Step-by-step tutorials for learning new concepts and technologies",
		"project":    "Project showcases and portfolio pieces",
		"work":       "Work-related posts and professional experiences",
		"quote":      "Interesting quotes and thoughts worth sharing",
		"snippet":    "Code snippets and useful code examples",
		"list":       "List-based content - recommendations, favorites, and collections",
	}

	fmt.Println("Adding descriptions to post types")
	fmt.Println("=================================\n")

	if *dryRun {
		fmt.Println("[DRY-RUN] Descriptions to be added:\n")
		for typeID, desc := range descriptions {
			fmt.Printf("  %s: %s\n", typeID, desc)
		}
		fmt.Println("\nTo apply changes, run: go run scripts/add-descriptions/main.go -dry-run=false")
		return
	}

	fmt.Println("Applying descriptions...\n")
	updated := 0

	for typeID, desc := range descriptions {
		_, err = db.ExecContext(ctx,
			"UPDATE post_types SET description = ? WHERE id = ?",
			desc, typeID)

		if err != nil {
			fmt.Printf("ERROR updating %s: %v\n", typeID, err)
			continue
		}

		fmt.Printf("✓ %s: %s\n", typeID, desc)
		updated++
	}

	fmt.Printf("\nUpdated %d post types with descriptions\n", updated)
}
