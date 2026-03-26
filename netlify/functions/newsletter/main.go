package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

type RSSPost struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Date        string `json:"date"`
	Link        string `json:"link"`
	Body        string `json:"body"`
	Source      string `json:"source"`
	Description string `json:"description"`
}

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	// Get newsletter feed URL from environment or use default
	feedURL := os.Getenv("NEWSLETTER_FEED_URL")
	if feedURL == "" {
		feedURL = "https://techstructively.substack.com/feed"
	}

	// Fetch and parse RSS feed
	posts, err := fetchNewsletter(ctx, feedURL)
	if err != nil {
		log.Printf("Error fetching newsletter: %v", err)
		return errorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to fetch newsletter: %v", err)), nil
	}

	// Create markdown files in posts/newsletter
	postsDir := "/tmp/posts/newsletter"
	if err := os.MkdirAll(postsDir, 0o755); err != nil {
		log.Printf("Error creating posts directory: %v", err)
		return errorResponse(http.StatusInternalServerError, "Failed to create directory"), nil
	}

	saved := 0
	skipped := 0

	for _, post := range posts {
		// Generate frontmatter
		frontmatter := generateFrontmatter(post)

		// Create markdown file
		filename := post.Slug + ".md"
		filepath := postsDir + "/" + filename

		content := frontmatter + "\n\n" + post.Body

		if err := os.WriteFile(filepath, []byte(content), 0o644); err != nil {
			log.Printf("Error writing file %s: %v", filepath, err)
			skipped++
			continue
		}

		log.Printf("Saved newsletter post: %s", post.Title)
		saved++
	}

	return jsonResponse(http.StatusOK, map[string]interface{}{
		"message": "Newsletter sync completed",
		"saved":   saved,
		"skipped": skipped,
		"total":   len(posts),
	}), nil
}

func fetchNewsletter(ctx context.Context, feedURL string) ([]RSSPost, error) {
	client := &http.Client{Timeout: 20 * time.Second}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set user agent to avoid 403 from Substack
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/rss+xml, application/atom+xml, application/xml;q=0.9, */*;q=0.1")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Referer", "https://www.google.com/")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("feed returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse feed
	parser := gofeed.NewParser()
	feed, err := parser.ParseString(string(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed: %w", err)
	}

	return parseRSSItems(feed), nil
}

func parseRSSItems(feed *gofeed.Feed) []RSSPost {
	posts := []RSSPost{}

	for _, item := range feed.Items {
		if item == nil || item.Title == "" {
			continue
		}

		date := ""
		if item.PublishedParsed != nil {
			date = item.PublishedParsed.Format("2006-01-02")
		}

		// Get body content
		body := item.Content
		if body == "" {
			body = item.Description
		}

		slug := slugFromTitle(item.Title)

		post := RSSPost{
			Title:       item.Title,
			Slug:        "newsletter/" + date + "-" + slug,
			Date:        date,
			Link:        item.Link,
			Body:        body,
			Source:      "substack",
			Description: item.Description,
		}

		posts = append(posts, post)
	}

	return posts
}

func generateFrontmatter(post RSSPost) string {
	frontmatter := fmt.Sprintf(`---
type: newsletter
title: "%s"
description: "%s"
status: published
slug: %s
date: %s
tags: ["newsletter", "substack"]
canonical_url: %s
---`, 
		escapeFrontmatterValue(post.Title),
		escapeFrontmatterValue(truncate(post.Description, 100)),
		post.Slug,
		post.Date,
		post.Link,
	)
	return frontmatter
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

func escapeFrontmatterValue(s string) string {
	// Escape quotes in YAML frontmatter
	return strings.ReplaceAll(s, `"`, `\"`)
}

func truncate(s string, length int) string {
	if len(s) > length {
		return s[:length] + "..."
	}
	return s
}

func jsonResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
		},
		Body: string(body),
	}
}

func errorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return jsonResponse(statusCode, map[string]string{"error": message})
}
