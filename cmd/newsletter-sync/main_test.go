package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSortNewsletterFeedPostsNewestFirst(t *testing.T) {
	posts := []newsletterFeedPost{
		{Slug: "older", Date: "2026-01-01"},
		{Slug: "newest", Date: "2026-03-01"},
		{Slug: "middle", Date: "2026-02-01"},
	}

	sortNewsletterFeedPosts(posts)

	if posts[0].Slug != "newest" || posts[1].Slug != "middle" || posts[2].Slug != "older" {
		t.Fatalf("unexpected sort order: %+v", posts)
	}
}

func TestPlanNewsletterMaterializationStopsAtExisting(t *testing.T) {
	posts := []newsletterFeedPost{
		{Slug: "newest", Date: "2026-03-01"},
		{Slug: "existing", Date: "2026-02-01"},
		{Slug: "older", Date: "2026-01-01"},
	}
	existing := map[string]bool{
		"existing": true,
	}

	planned := planNewsletterMaterialization(posts, existing)

	if len(planned) != 1 {
		t.Fatalf("expected 1 planned post, got %d", len(planned))
	}
	if planned[0].Slug != "newest" {
		t.Fatalf("expected newest post first, got %q", planned[0].Slug)
	}
}

func TestMaterializeNewsletterResponseWritesOnlyNewerPosts(t *testing.T) {
	dir := t.TempDir()
	outputDir := filepath.Join(dir, "posts", "newsletter")
	createdFiles := filepath.Join(dir, "created_files.txt")
	responsePath := filepath.Join(dir, "response.json")

	response := `{
		"posts": [
			{"title":"Newest","slug":"newest","date":"2026-03-01","link":"https://example.com/newest","body":"<p>newest</p>","description":"Newest"},
			{"title":"Existing","slug":"existing","date":"2026-02-01","link":"https://example.com/existing","body":"<p>existing</p>","description":"Existing"},
			{"title":"Older","slug":"older","date":"2026-01-01","link":"https://example.com/older","body":"<p>older</p>","description":"Older"}
		]
	}`
	if err := os.WriteFile(responsePath, []byte(response), 0o644); err != nil {
		t.Fatalf("write response: %v", err)
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		t.Fatalf("mkdir output: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outputDir, "existing.md"), []byte("---\ntitle: Existing\nslug: existing\n---\nexisting"), 0o644); err != nil {
		t.Fatalf("write existing file: %v", err)
	}

	count, err := materializeNewsletterResponse(responsePath, outputDir, createdFiles)
	if err != nil {
		t.Fatalf("materialize response: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 created post, got %d", count)
	}

	newFile := filepath.Join(outputDir, "newest.md")
	if _, err := os.Stat(newFile); err != nil {
		t.Fatalf("expected newest file to exist: %v", err)
	}

	olderFile := filepath.Join(outputDir, "older.md")
	if _, err := os.Stat(olderFile); !os.IsNotExist(err) {
		t.Fatalf("expected older file to be skipped, got err=%v", err)
	}

	createdBytes, err := os.ReadFile(createdFiles)
	if err != nil {
		t.Fatalf("read created files: %v", err)
	}
	created := strings.TrimSpace(string(createdBytes))
	if created != newFile {
		t.Fatalf("expected created file list to contain only %q, got %q", newFile, created)
	}

	content, err := os.ReadFile(newFile)
	if err != nil {
		t.Fatalf("read newest file: %v", err)
	}
	text := string(content)
	if !strings.Contains(text, "title: \"Newest\"") {
		t.Fatalf("missing title in frontmatter: %s", text)
	}
	if !strings.Contains(text, "slug: newest") {
		t.Fatalf("missing slug in frontmatter: %s", text)
	}
	if !strings.Contains(text, "canonical_url: https://example.com/newest") {
		t.Fatalf("missing canonical url in frontmatter: %s", text)
	}
}
