package main

import (
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
)

func TestSlugFromTitle(t *testing.T) {
	tests := []struct {
		title    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"How to Build APIs", "how-to-build-apis"},
		{"Test---Multiple-Dashes", "test-multiple-dashes"},
		{"123 Numbers Start", "123-numbers-start"},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			result := slugFromTitle(tt.title)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestEscapeFrontmatterValue(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`Hello "World"`, `Hello \"World\"`},
		{`Simple text`, `Simple text`},
		{`Multiple "quotes" here`, `Multiple \"quotes\" here`},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := escapeFrontmatterValue(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"Hello World", 5, "Hello..."},
		{"Short", 10, "Short"},
		{"Exactly ten", 11, "Exactly ten"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := truncate(tt.input, tt.length)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGenerateFrontmatter(t *testing.T) {
	post := RSSPost{
		Title:       "Test Article",
		Slug:        "newsletter/2024-01-01-test-article",
		Date:        "2024-01-01",
		Link:        "https://example.com/article",
		Description: "A test article description",
	}

	frontmatter := generateFrontmatter(post)

	// Check required fields are present
	if !contains(frontmatter, "type: newsletter") {
		t.Error("frontmatter missing type: newsletter")
	}
	if !contains(frontmatter, "title: \"Test Article\"") {
		t.Error("frontmatter missing title")
	}
	if !contains(frontmatter, "slug: newsletter/2024-01-01-test-article") {
		t.Error("frontmatter missing slug")
	}
	if !contains(frontmatter, "date: 2024-01-01") {
		t.Error("frontmatter missing date")
	}
	if !contains(frontmatter, "canonical_url: https://example.com/article") {
		t.Error("frontmatter missing canonical_url")
	}
	if !contains(frontmatter, "tags: [\"newsletter\", \"substack\"]") {
		t.Error("frontmatter missing tags")
	}
}

func TestParseRSSItemsSortsNewestFirst(t *testing.T) {
	older := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	newer := time.Date(2025, time.February, 1, 0, 0, 0, 0, time.UTC)

	feed := &feedStub{
		Items: []*rssItemStub{
			{
				Title:           "Older",
				Link:            "https://example.com/older",
				Content:         "older body",
				PublishedParsed: &older,
			},
			{
				Title:           "Newer",
				Link:            "https://example.com/newer",
				Content:         "newer body",
				PublishedParsed: &newer,
			},
		},
	}

	posts := parseRSSItems(feed.toFeed())
	if len(posts) != 2 {
		t.Fatalf("expected 2 posts, got %d", len(posts))
	}
	if posts[0].Title != "Newer" || posts[1].Title != "Older" {
		t.Fatalf("unexpected order: %+v", posts)
	}
}

type rssItemStub struct {
	Title           string
	Link            string
	Content         string
	Description     string
	PublishedParsed *time.Time
}

type feedStub struct {
	Items []*rssItemStub
}

func (f *feedStub) toFeed() *gofeed.Feed {
	items := make([]*gofeed.Item, 0, len(f.Items))
	for _, item := range f.Items {
		items = append(items, &gofeed.Item{
			Title:           item.Title,
			Link:            item.Link,
			Content:         item.Content,
			Description:     item.Description,
			PublishedParsed: item.PublishedParsed,
		})
	}
	return &gofeed.Feed{Items: items}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
