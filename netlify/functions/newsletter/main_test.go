package main

import (
	"testing"
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

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
