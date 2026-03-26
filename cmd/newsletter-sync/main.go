package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	dbInsert := flag.Bool("db-insert", false, "Insert posts to Turso database instead of creating markdown files")
	flag.Parse()

	if *dbInsert {
		insertToDB()
		return
	}

	feedURL := os.Getenv("NEWSLETTER_FEED_URL")
	if feedURL == "" {
		feedURL = "https://techstructively.substack.com/feed"
	}

	postsDir := "posts/newsletter"

	// Fetch RSS feed
	fmt.Printf("📰 Fetching from: %s\n", feedURL)
	feed, err := fetchFeed(feedURL)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Feed: %s\n", feed.Title)
	fmt.Printf("📄 Total items: %d\n", len(feed.Items))

	// Ensure directory exists
	os.MkdirAll(postsDir, 0o755)

	// Check existing posts
	existing := make(map[string]bool)
	entries, _ := os.ReadDir(postsDir)
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".md") {
			slug := strings.TrimSuffix(e.Name(), ".md")
			existing[slug] = true
		}
	}

	fmt.Printf("📦 Existing posts: %d\n", len(existing))

	// Process feed items
	newCount := 0
	for _, item := range feed.Items {
		if item == nil || item.Title == "" {
			continue
		}

		// Generate slug from date + title
		date := ""
		if item.PublishedParsed != nil {
			date = item.PublishedParsed.Format("2006-01-02")
		}

		slug := date + "-" + slugFromTitle(item.Title)

		// Skip if already exists
		if existing[slug] {
			continue
		}

		// Create markdown file
		body := item.Content
		if body == "" {
			body = item.Description
		}

		frontmatter := fmt.Sprintf(`---
type: newsletter
title: "%s"
date: %s
slug: newsletter/%s
tags: ["newsletter", "substack"]
canonical_url: %s
status: published
---

`, escapeQuotes(item.Title), date, slug, item.Link)

		filename := filepath.Join(postsDir, slug+".md")
		content := frontmatter + body

		if err := os.WriteFile(filename, []byte(content), 0o644); err != nil {
			fmt.Printf("❌ Error writing %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✨ Added: %s\n", slug)
		newCount++
	}

	fmt.Printf("\n✅ Summary: %d new posts\n", newCount)
	os.Exit(0)
}

func insertToDB() {
	ctx := context.Background()

	// Read list of newly created files from workflow
	createdFiles := make(map[string]bool)
	if data, err := os.ReadFile("/tmp/created_files.txt"); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				createdFiles[strings.TrimSpace(line)] = true
			}
		}
	}

	// If no workflow files, sync ALL local files
	if len(createdFiles) == 0 {
		postsDir := "posts/newsletter"
		entries, _ := os.ReadDir(postsDir)
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".md") {
				createdFiles[filepath.Join(postsDir, e.Name())] = true
			}
		}
	}

	if len(createdFiles) == 0 {
		fmt.Printf("⏭️  No files to sync\n")
		os.Exit(0)
	}

	fmt.Printf("📝 Found %d files to sync\n", len(createdFiles))

	// Open database
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	if dbName == "" || dbToken == "" {
		fmt.Printf("❌ Missing Turso credentials\n")
		os.Exit(1)
	}

	// Build connection string - handle both formats
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

	q := libsqlssg.New(db)
	inserted := 0
	seenHashes := make(map[string]bool) // Track hashes we've already processed in this run

	// Process only newly created files
	for filePath := range createdFiles {
		// Read markdown file
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("⚠️  Error reading %s: %v\n", filePath, err)
			continue
		}

		// Parse frontmatter and body
		parts := strings.Split(string(content), "---")
		if len(parts) < 3 {
			fmt.Printf("⚠️  Invalid markdown format: %s\n", filePath)
			continue
		}

		// Parse YAML frontmatter (simple parsing)
		frontmatter := parts[1]
		body := strings.TrimSpace(parts[2])

		title := extractYAML(frontmatter, "title")
		slug := strings.TrimSuffix(filepath.Base(filePath), ".md")
		link := extractYAML(frontmatter, "canonical_url")

		// Compute content hash
		contentHash := computeHash(title + body)

		// Skip if we've seen this hash in this same run (prevents duplicates within same feed fetch)
		if seenHashes[contentHash] {
			fmt.Printf("⏭️  Duplicate content (same hash): %s\n", slug)
			continue
		}
		seenHashes[contentHash] = true

		// Check if title already exists in database - if yes, skip (duplicate)
		var existingTitle string
		db.QueryRowContext(ctx, "SELECT title FROM posts WHERE title = ? LIMIT 1", title).Scan(&existingTitle)
		if existingTitle != "" {
			fmt.Printf("⏭️  Skip: title already exists\n")
			continue
		}

		// Insert only (hash is unique key)
		metadata := map[string]interface{}{
			"canonical_url": link,
			"type":          "newsletter",
			"source":        "substack",
			"content_hash":  contentHash,
		}
		metaJSON, _ := json.Marshal(metadata)
		tagsJSON, _ := json.Marshal([]string{"newsletter", "substack"})

		_, err = q.CreatePost(ctx, libsqlssg.CreatePostParams{
			ID:       fmt.Sprintf("%d", time.Now().UnixNano()),
			TypeID:   "newsletter",
			Title:    title,
			Slug:     slug,
			Content:  body,
			Metadata: sql.NullString{String: string(metaJSON), Valid: true},
			Tags:     sql.NullString{String: string(tagsJSON), Valid: true},
			Status:   sql.NullString{String: "published", Valid: true},
		})

		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		fmt.Printf("📤 Inserted: %s\n", title)
		inserted++
	}

	fmt.Printf("\n✅ Synced %d posts to Turso\n", inserted)
	os.Exit(0)
}

func extractYAML(yaml, key string) string {
	for _, line := range strings.Split(yaml, "\n") {
		if strings.HasPrefix(line, key+":") {
			val := strings.TrimSpace(strings.TrimPrefix(line, key+":"))
			val = strings.Trim(val, `"`)
			return val
		}
	}
	return ""
}

func computeHash(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:])
}

func parseMetadataJSON(metaStr string) map[string]interface{} {
	if metaStr == "" {
		return map[string]interface{}{}
	}
	var meta map[string]interface{}
	json.Unmarshal([]byte(metaStr), &meta)
	return meta
}

func postWithHashExists(ctx context.Context, db *sql.DB, contentHash string) (bool, error) {
	// Query posts table where metadata JSON contains the content_hash
	query := `SELECT COUNT(*) FROM posts WHERE metadata LIKE ?`
	var count int
	err := db.QueryRowContext(ctx, query, "%"+contentHash+"%").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func fetchFeed(url string) (*gofeed.Feed, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/rss+xml, application/atom+xml")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("feed returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	parser := gofeed.NewParser()
	feed, err := parser.ParseString(string(body))
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func slugFromTitle(title string) string {
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s := strings.ToLower(title)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		s = "post"
	}
	return s
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
