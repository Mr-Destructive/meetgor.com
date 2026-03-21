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

	// Find all post types that are not used
	rows, err := db.QueryContext(ctx,
		`SELECT pt.id, pt.name FROM post_types pt 
		 LEFT JOIN posts p ON pt.id = p.type_id 
		 WHERE p.id IS NULL`)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	var unused []struct {
		id   string
		name string
	}

	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatalf("Scan failed: %v", err)
		}
		unused = append(unused, struct {
			id   string
			name string
		}{id, name})
	}

	if len(unused) == 0 {
		fmt.Println("No unused post types found")
		return
	}

	fmt.Printf("Found %d unused post types:\n\n", len(unused))
	for i, ut := range unused {
		fmt.Printf("[%d] %s (%s)\n", i+1, ut.id, ut.name)
	}

	if *dryRun {
		fmt.Println("\n[DRY-RUN] To apply changes, run: go run scripts/cleanup-post-types/main.go -dry-run=false")
		return
	}

	fmt.Println("\nDeleting unused post types...")
	for _, ut := range unused {
		_, err := db.ExecContext(ctx, "DELETE FROM post_types WHERE id = ?", ut.id)
		if err != nil {
			fmt.Printf("ERROR deleting %s: %v\n", ut.id, err)
			continue
		}
		fmt.Printf("✓ Deleted %s\n", ut.id)
	}

	fmt.Printf("\nCleanup complete: %d post types removed\n", len(unused))
}
