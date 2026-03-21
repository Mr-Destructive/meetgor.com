package main

import (
	"bufio"
	"context"
	"database/sql"
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

	fmt.Println("Post types in use:")
	fmt.Println("==================\n")

	// Get all post types with count of posts
	rows, err := db.QueryContext(ctx,
		`SELECT pt.id, pt.name, COUNT(p.id) as count 
		 FROM post_types pt 
		 LEFT JOIN posts p ON pt.id = p.type_id AND p.status != 'deleted'
		 GROUP BY pt.id
		 ORDER BY count DESC, pt.id`)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	totalTypes := 0
	totalPosts := 0
	for rows.Next() {
		var id, name string
		var count int
		if err := rows.Scan(&id, &name, &count); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}
		totalTypes++
		totalPosts += count
		if count > 0 {
			fmt.Printf("✓ %s (%s): %d posts\n", id, name, count)
		} else {
			fmt.Printf("  %s (%s): 0 posts (unused)\n", id, name, count)
		}
	}

	fmt.Printf("\nTotal: %d types, %d posts\n", totalTypes, totalPosts)
}
