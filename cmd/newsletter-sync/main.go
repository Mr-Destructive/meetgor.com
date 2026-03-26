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
	"sort"
	"strings"
	"time"

	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"
	"github.com/mmcdole/gofeed"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gopkg.in/yaml.v2"
)

type newsletterFeedPost struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Date        string `json:"date"`
	Link        string `json:"link"`
	Body        string `json:"body"`
	Description string `json:"description"`
}

type newsletterFeedResponse struct {
	Posts []newsletterFeedPost `json:"posts"`
}

func main() {
	dbInsert := flag.Bool("db-insert", false, "Insert posts to Turso database instead of creating markdown files")
	materializeJSON := flag.String("materialize-json", "", "Materialize newsletter markdown files from a Netlify response JSON file")
	materializeOutputDir := flag.String("materialize-output-dir", "posts/newsletter", "Directory for newsletter markdown files")
	materializeCreatedFiles := flag.String("materialize-created-files", "/tmp/created_files.txt", "File that records newly created newsletter markdown files")
	flag.Parse()

	if *materializeJSON != "" {
		count, err := materializeNewsletterResponse(*materializeJSON, *materializeOutputDir, *materializeCreatedFiles)
		if err != nil {
			fmt.Printf("❌ Materialize failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Materialized %d newsletter posts\n", count)
		return
	}

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

		date := ""
		if item.PublishedParsed != nil {
			date = item.PublishedParsed.Format("2006-01-02")
		}

		slug := slugFromTitle(item.Title)

		// Skip if already exists
		if existing[slug] {
			continue
		}

		// Create markdown file
		body := item.Content
		if body == "" {
			body = item.Description
		}

		frontmatter := generateFrontmatter(item.Title, slug, date, item.Link, item.Description)

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

func materializeNewsletterResponse(responsePath, outputDir, createdFilesPath string) (int, error) {
	data, err := os.ReadFile(responsePath)
	if err != nil {
		return 0, err
	}

	var response newsletterFeedResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return 0, err
	}

	sortNewsletterFeedPosts(response.Posts)

	existing, err := readExistingNewsletterSlugs(outputDir)
	if err != nil {
		return 0, err
	}

	planned := planNewsletterMaterialization(response.Posts, existing)
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return 0, err
	}
	if err := os.WriteFile(createdFilesPath, []byte(""), 0o644); err != nil {
		return 0, err
	}

	created := 0
	for _, post := range planned {
		if err := writeNewsletterPost(outputDir, createdFilesPath, post); err != nil {
			return created, err
		}
		created++
	}

	return created, nil
}

func sortNewsletterFeedPosts(posts []newsletterFeedPost) {
	sort.SliceStable(posts, func(i, j int) bool {
		return newsletterPostSortKey(posts[i]).After(newsletterPostSortKey(posts[j]))
	})
}

func planNewsletterMaterialization(posts []newsletterFeedPost, existing map[string]bool) []newsletterFeedPost {
	planned := make([]newsletterFeedPost, 0, len(posts))
	for _, post := range posts {
		if post.Slug == "" {
			continue
		}
		if existing[post.Slug] {
			break
		}
		planned = append(planned, post)
	}
	return planned
}

func readExistingNewsletterSlugs(outputDir string) (map[string]bool, error) {
	existing := make(map[string]bool)
	entries, err := os.ReadDir(outputDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		existing[strings.TrimSuffix(entry.Name(), ".md")] = true
	}
	return existing, nil
}

func writeNewsletterPost(outputDir, createdFilesPath string, post newsletterFeedPost) error {
	body := post.Body
	if body == "" {
		body = post.Description
	}
	content := generateFrontmatter(post.Title, post.Slug, post.Date, post.Link, post.Description) + body
	filePath := filepath.Join(outputDir, post.Slug+".md")
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return err
	}

	file, err := os.OpenFile(createdFilesPath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, filePath); err != nil {
		return err
	}

	fmt.Printf("✨ Created: %s\n", filePath)
	return nil
}

func newsletterPostSortKey(post newsletterFeedPost) time.Time {
	if post.Date != "" {
		if parsed, err := time.Parse("2006-01-02", post.Date); err == nil {
			return parsed
		}
	}
	return time.Time{}
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

		frontmatter, body, err := splitNewsletterFrontmatter(string(content))
		if err != nil {
			fmt.Printf("⚠️  Invalid markdown format: %s\n", filePath)
			continue
		}

		// Try to parse as JSON first
		var metadata map[string]interface{}
		err = json.Unmarshal([]byte(frontmatter), &metadata)

		// If JSON fails, try YAML
		if err != nil {
			err = yaml.Unmarshal([]byte(frontmatter), &metadata)
			if err != nil {
				fmt.Printf("⚠️  Cannot parse metadata (JSON or YAML): %s\n", filePath)
				continue
			}
		}

		title, ok := metadata["title"].(string)
		if !ok {
			fmt.Printf("⚠️  Missing title in: %s\n", filePath)
			continue
		}

		slug := extractMetadataString(metadata, "slug")
		slug = normalizeNewsletterSlug(slug)
		if slug == "" {
			slug = strings.TrimSuffix(filepath.Base(filePath), ".md")
		}

		link := extractMetadataString(metadata, "canonical_url")
		if link == "" {
			link = extractMetadataString(metadata, "link")
		}

		// Compute content hash
		contentHash := computeHash(title + body)

		// Skip if we've seen this hash in this same run (prevents duplicates within same feed fetch)
		if seenHashes[contentHash] {
			fmt.Printf("⏭️  Duplicate content (same hash): %s\n", slug)
			continue
		}
		seenHashes[contentHash] = true

		// Skip if this slug already exists in the database.
		if _, err := q.GetPostBySlug(ctx, slug); err == nil {
			fmt.Printf("⏭️  Skip: slug already exists\n")
			continue
		}

		// Skip if we already inserted this exact content hash in the past.
		if exists, err := postWithHashExists(ctx, db, contentHash); err == nil && exists {
			fmt.Printf("⏭️  Skip: content hash already exists\n")
			continue
		}

		// Insert only (hash is unique key)
		metaPayload := map[string]interface{}{
			"canonical_url": link,
			"type":          "newsletter",
			"source":        "substack",
			"content_hash":  contentHash,
		}
		metaJSON, _ := json.Marshal(metaPayload)
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

func computeHash(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:])
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

func generateFrontmatter(title, slug, date, link, description string) string {
	var builder strings.Builder
	builder.WriteString("---\n")
	builder.WriteString(fmt.Sprintf("title: \"%s\"\n", escapeQuotes(title)))
	if date != "" {
		builder.WriteString(fmt.Sprintf("date: %s\n", date))
	}
	builder.WriteString(fmt.Sprintf("slug: %s\n", slug))
	builder.WriteString("type: newsletter\n")
	builder.WriteString("status: published\n")
	builder.WriteString("source: newsletter\n")
	builder.WriteString(fmt.Sprintf("canonical_url: %s\n", link))
	if description != "" {
		builder.WriteString(fmt.Sprintf("description: \"%s\"\n", escapeQuotes(description)))
	}
	builder.WriteString("tags: [\"newsletter\", \"substack\"]\n")
	builder.WriteString("---\n\n")
	return builder.String()
}

func splitNewsletterFrontmatter(content string) (string, string, error) {
	trimmed := strings.TrimSpace(content)
	if trimmed == "" {
		return "", "", fmt.Errorf("empty content")
	}

	if strings.HasPrefix(trimmed, "---") {
		lines := strings.Split(trimmed, "\n")
		end := -1
		for i := 1; i < len(lines); i++ {
			if strings.TrimSpace(lines[i]) == "---" {
				end = i
				break
			}
		}
		if end == -1 {
			return "", "", fmt.Errorf("missing closing frontmatter delimiter")
		}
		frontmatter := strings.Join(lines[1:end], "\n")
		body := strings.Join(lines[end+1:], "\n")
		return frontmatter, strings.TrimSpace(body), nil
	}

	if strings.HasPrefix(trimmed, "{") {
		if end := strings.Index(trimmed, "\n\n"); end != -1 {
			frontmatter := strings.TrimSpace(trimmed[:end])
			body := strings.TrimSpace(trimmed[end+2:])
			return frontmatter, body, nil
		}
	}

	return "", "", fmt.Errorf("unsupported markdown format")
}

func extractMetadataString(metadata map[string]interface{}, key string) string {
	if metadata == nil {
		return ""
	}
	if value, ok := metadata[key]; ok {
		if s, ok := value.(string); ok {
			return s
		}
	}
	return ""
}

func normalizeNewsletterSlug(slug string) string {
	slug = strings.TrimSpace(slug)
	slug = strings.TrimPrefix(slug, "newsletter/")
	slug = strings.TrimPrefix(slug, "/")
	return slug
}
