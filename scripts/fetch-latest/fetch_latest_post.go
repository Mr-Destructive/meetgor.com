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

func latestPostMain() {
	loadEnv(".env")

	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		log.Fatal("Missing TURSO_DATABASE_NAME or TURSO_DATABASE_AUTH_TOKEN")
	}

	// dbName should already include the protocol
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

	var id, typeID, title, slug, content, status string
	var metadata sql.NullString
	var createdAt, updatedAt sql.NullTime

	err = db.QueryRowContext(ctx,
		`SELECT id, type_id, title, slug, content, status, metadata, created_at, updated_at 
		 FROM posts 
		 WHERE status != 'deleted' 
		 ORDER BY created_at DESC 
		 LIMIT 1`).
		Scan(&id, &typeID, &title, &slug, &content, &status, &metadata, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println("No posts found")
		return
	}
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	post := map[string]interface{}{
		"id":         id,
		"type_id":    typeID,
		"title":      title,
		"slug":       slug,
		"content":    content,
		"status":     status,
	}
	
	if createdAt.Valid {
		post["created_at"] = createdAt.Time.Unix()
	}
	if updatedAt.Valid {
		post["updated_at"] = updatedAt.Time.Unix()
	}

	if metadata.Valid {
		var meta interface{}
		if err := json.Unmarshal([]byte(metadata.String), &meta); err == nil {
			post["metadata"] = meta
		} else {
			post["metadata"] = metadata.String
		}
	}

	out, _ := json.MarshalIndent(post, "", "  ")
	fmt.Println(string(out))
}

func main() {
	latestPostMain()
}
