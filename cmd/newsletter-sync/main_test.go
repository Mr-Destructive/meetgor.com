package main

import (
	"encoding/json"
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

func TestMaterializeNewsletterResponseCreatesLinkPostsFromReadAndWatchedSections(t *testing.T) {
	originalFetcher := linkMetadataFetcher
	linkMetadataFetcher = func(rawURL string) (linkMetadata, error) {
		switch rawURL {
		case "https://example.com/article-one":
			return linkMetadata{
				Title:       "Example Article One",
				Description: "Example article description",
				ImageURL:    "https://example.com/article-one.png",
			}, nil
		default:
			return linkMetadata{}, nil
		}
	}
	t.Cleanup(func() {
		linkMetadataFetcher = originalFetcher
	})

	dir := t.TempDir()
	outputDir := filepath.Join(dir, "posts", "newsletter")
	createdFiles := filepath.Join(dir, "created_files.txt")
	responsePath := filepath.Join(dir, "response.json")

	response := `{
		"posts": [
			{
				"title": "Techstructive Weekly #99",
				"slug": "techstructive-weekly-99",
				"date": "2026-03-01",
				"link": "https://example.com/newsletter-99",
				"body": "## Read\n\n- [Article One](https://example.com/article-one)\n  - First thought\n  - Second thought\n\n## Watched\n\n- [Video One](https://www.youtube.com/watch?v=abc123)\n  - Nice video\n",
				"description": "Weekly newsletter"
			}
		]
	}`
	if err := os.WriteFile(responsePath, []byte(response), 0o644); err != nil {
		t.Fatalf("write response: %v", err)
	}

	count, err := materializeNewsletterResponse(responsePath, outputDir, createdFiles)
	if err != nil {
		t.Fatalf("materialize response: %v", err)
	}
	if count != 3 {
		t.Fatalf("expected 3 created posts, got %d", count)
	}

	newsletterFile := filepath.Join(outputDir, "techstructive-weekly-99.md")
	if _, err := os.Stat(newsletterFile); err != nil {
		t.Fatalf("expected newsletter file to exist: %v", err)
	}

	linkDir := filepath.Join(dir, "posts", "links")
	articleFile := filepath.Join(linkDir, "example-article-one.md")
	videoFile := filepath.Join(linkDir, "video-one.md")
	if _, err := os.Stat(articleFile); err != nil {
		t.Fatalf("expected article link file to exist: %v", err)
	}
	if _, err := os.Stat(videoFile); err != nil {
		t.Fatalf("expected video link file to exist: %v", err)
	}

	articleContent, err := os.ReadFile(articleFile)
	if err != nil {
		t.Fatalf("read article link file: %v", err)
	}
	articleText := string(articleContent)
	if !strings.Contains(articleText, "title: \"Thoughts: Example Article One\"") {
		t.Fatalf("missing prefixed title in article link file: %s", articleText)
	}
	if !strings.Contains(articleText, "newsletter: techstructive-weekly-99") {
		t.Fatalf("missing newsletter source in article link file: %s", articleText)
	}
	if !strings.Contains(articleText, "link: \"https://example.com/article-one\"") {
		t.Fatalf("missing link in article link file: %s", articleText)
	}
	if !strings.Contains(articleText, "image_url: \"https://example.com/article-one.png\"") {
		t.Fatalf("missing image url in article link file: %s", articleText)
	}

	videoContent, err := os.ReadFile(videoFile)
	if err != nil {
		t.Fatalf("read video link file: %v", err)
	}
	videoText := string(videoContent)
	if !strings.Contains(videoText, "https://i.ytimg.com/vi/abc123/hqdefault.jpg") {
		t.Fatalf("missing youtube thumbnail in video link file: %s", videoText)
	}
}

func TestMaterializeLinkPostsSkipsExistingContentHash(t *testing.T) {
	originalFetcher := linkMetadataFetcher
	linkMetadataFetcher = func(rawURL string) (linkMetadata, error) {
		return linkMetadata{
			Title:       "Example Article One",
			Description: "Example article description",
			ImageURL:    "https://example.com/article-one.png",
		}, nil
	}
	t.Cleanup(func() {
		linkMetadataFetcher = originalFetcher
	})

	dir := t.TempDir()
	outputDir := filepath.Join(dir, "posts", "newsletter")
	createdFiles := filepath.Join(dir, "created_files.txt")
	linkDir := filepath.Join(dir, "posts", "links")

	if err := os.MkdirAll(linkDir, 0o755); err != nil {
		t.Fatalf("mkdir links: %v", err)
	}

	commentary := buildCommentaryBody([]string{"First thought", "Second thought"})
	hash := computeLinkHash("Thoughts: Example Article One", commentary, "https://example.com/article-one")
	existing := `---
title: "Legacy Copy"
slug: legacy-copy
type: links
hash: ` + hash + `
---
## Commentary

- First thought
- Second thought
`
	if err := os.WriteFile(filepath.Join(linkDir, "legacy-copy.md"), []byte(existing), 0o644); err != nil {
		t.Fatalf("write existing link file: %v", err)
	}

	response := `{
		"posts": [
			{
				"title": "Techstructive Weekly #100",
				"slug": "techstructive-weekly-100",
				"date": "2026-03-08",
				"link": "https://example.com/newsletter-100",
				"body": "## Read\n\n- [Article One](https://example.com/article-one)\n  - First thought\n  - Second thought\n",
				"description": "Weekly newsletter"
			}
		]
	}`
	responsePath := filepath.Join(dir, "response.json")
	if err := os.WriteFile(responsePath, []byte(response), 0o644); err != nil {
		t.Fatalf("write response: %v", err)
	}

	count, err := materializeNewsletterResponse(responsePath, outputDir, createdFiles)
	if err != nil {
		t.Fatalf("materialize response: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected only newsletter file to be created, got %d", count)
	}

	if _, err := os.Stat(filepath.Join(linkDir, "example-article-one.md")); !os.IsNotExist(err) {
		t.Fatalf("expected duplicate link post to be skipped, err=%v", err)
	}
}

func TestMaterializeLinkPostsSkipsExistingLinkURL(t *testing.T) {
	originalFetcher := linkMetadataFetcher
	linkMetadataFetcher = func(rawURL string) (linkMetadata, error) {
		return linkMetadata{
			Title:       "Updated Example Article One",
			Description: "Example article description",
			ImageURL:    "https://example.com/article-one.png",
		}, nil
	}
	t.Cleanup(func() {
		linkMetadataFetcher = originalFetcher
	})

	dir := t.TempDir()
	outputDir := filepath.Join(dir, "posts", "newsletter")
	createdFiles := filepath.Join(dir, "created_files.txt")
	linkDir := filepath.Join(dir, "posts", "links")

	if err := os.MkdirAll(linkDir, 0o755); err != nil {
		t.Fatalf("mkdir links: %v", err)
	}

	existing := `---
title: "Thoughts: Example Article One"
slug: example-article-one
type: links
link: "https://example.com/article-one"
hash: deadbeef
---
## Commentary

- Existing comment
`
	if err := os.WriteFile(filepath.Join(linkDir, "example-article-one.md"), []byte(existing), 0o644); err != nil {
		t.Fatalf("write existing link file: %v", err)
	}

	response := `{
		"posts": [
			{
				"title": "Techstructive Weekly #101",
				"slug": "techstructive-weekly-101",
				"date": "2026-03-09",
				"link": "https://example.com/newsletter-101",
				"body": "## Read\n\n- [Updated Example Article One](https://example.com/article-one)\n  - First thought\n  - Second thought\n",
				"description": "Weekly newsletter"
			}
		]
	}`
	responsePath := filepath.Join(dir, "response.json")
	if err := os.WriteFile(responsePath, []byte(response), 0o644); err != nil {
		t.Fatalf("write response: %v", err)
	}

	count, err := materializeNewsletterResponse(responsePath, outputDir, createdFiles)
	if err != nil {
		t.Fatalf("materialize response: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected only newsletter file to be created, got %d", count)
	}

	if _, err := os.Stat(filepath.Join(linkDir, "updated-example-article-one.md")); !os.IsNotExist(err) {
		t.Fatalf("expected duplicate link URL to be skipped, err=%v", err)
	}
}

func TestYouTubeThumbnailURL(t *testing.T) {
	tests := []struct {
		rawURL   string
		expected string
	}{
		{"https://youtu.be/abc123", "https://i.ytimg.com/vi/abc123/hqdefault.jpg"},
		{"https://www.youtube.com/watch?v=xyz789", "https://i.ytimg.com/vi/xyz789/hqdefault.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.rawURL, func(t *testing.T) {
			if got := youtubeThumbnailURL(tt.rawURL); got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestBuildMetadataPayloadForLinks(t *testing.T) {
	metadata := map[string]interface{}{
		"type":        "links",
		"date":        "2026-03-01",
		"newsletter":  "techstructive-weekly-99",
		"source":      "newsletter",
		"description": "Example article description",
		"image_url":   "https://example.com/article-one.png",
	}

	payload := buildMetadataPayload("links", metadata, "https://example.com/article-one", "hash123")

	if payload["type"] != "links" {
		t.Fatalf("expected links type, got %#v", payload["type"])
	}
	if payload["link"] != "https://example.com/article-one" {
		t.Fatalf("expected link in payload, got %#v", payload["link"])
	}
	if payload["canonical_url"] != "https://example.com/article-one" {
		t.Fatalf("expected canonical url in payload, got %#v", payload["canonical_url"])
	}
	if payload["newsletter"] != "techstructive-weekly-99" {
		t.Fatalf("expected newsletter in payload, got %#v", payload["newsletter"])
	}
	if payload["content_hash"] != "hash123" {
		t.Fatalf("expected content hash in payload, got %#v", payload["content_hash"])
	}

	if tagsJSON := buildTagsJSON(nil, "links"); len(tagsJSON) != 0 {
		t.Fatalf("expected no default tags for links, got %s", string(tagsJSON))
	}

	if tagsJSON := buildTagsJSON(nil, "newsletter"); len(tagsJSON) == 0 {
		t.Fatal("expected default newsletter tags")
	} else {
		var tags []string
		if err := json.Unmarshal(tagsJSON, &tags); err != nil {
			t.Fatalf("unmarshal tags: %v", err)
		}
		if len(tags) != 2 || tags[0] != "newsletter" || tags[1] != "substack" {
			t.Fatalf("unexpected newsletter tags payload: %#v", tags)
		}
	}
}
