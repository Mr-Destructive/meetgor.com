package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
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
