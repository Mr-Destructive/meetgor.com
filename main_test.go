package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadPostsNormalizesBlogType(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "legacy.md")
	content := "---\ntitle: Legacy\ndate: 2024-01-01\ntype: blog\nstatus: published\n---\nHello"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	posts, err := ReadPosts([]string{path})
	if err != nil {
		t.Fatalf("ReadPosts failed: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
	if got := posts[0].Frontmatter.Type; got != "posts" {
		t.Fatalf("expected blog to normalize to posts, got %q", got)
	}
}
