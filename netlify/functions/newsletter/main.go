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

type SyncResponse struct {
	Message string     `json:"message"`
	Posts   []RSSPost  `json:"posts"`
	Total   int        `json:"total"`
	New     int        `json:"new"`
}

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	feedURL := os.Getenv("NEWSLETTER_FEED_URL")
	if feedURL == "" {
		feedURL = "https://techstructively.substack.com/feed"
	}

	log.Printf("📰 Fetching from: %s", feedURL)

	// Fetch RSS feed
	posts, err := fetchNewsletter(ctx, feedURL)
	if err != nil {
		log.Printf("❌ Error: %v", err)
		return errorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to fetch newsletter: %v", err)), nil
	}

	log.Printf("✅ Found %d posts in feed", len(posts))

	response := SyncResponse{
		Message: "Newsletter sync completed",
		Posts:   posts,
		Total:   len(posts),
		New:     len(posts), // GitHub Action will filter duplicates
	}

	return jsonResponse(http.StatusOK, response), nil
}

func fetchNewsletter(ctx context.Context, feedURL string) ([]RSSPost, error) {
	client := &http.Client{Timeout: 20 * time.Second}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/rss+xml, application/atom+xml, application/xml;q=0.9, */*;q=0.1")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("feed returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

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

func jsonResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(body),
	}
}

func errorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return jsonResponse(statusCode, map[string]string{"error": message})
}
