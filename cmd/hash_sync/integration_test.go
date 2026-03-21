package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	libsqlssg "github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
)

// TestGitHubPushScenario simulates a direct GitHub push to posts/ directory
// This is the TDD test for Option 1: GitHub Action sync
func TestGitHubPushScenario(t *testing.T) {
	// RED: Test that hash_sync properly handles a GitHub push with new markdown files
	db := setupHashDB(t)
	defer db.Close()
	ctx := context.Background()

	// Setup author
	if _, err := db.ExecContext(ctx, `INSERT INTO authors (username, name, password) VALUES ('author','Author','pwd')`); err != nil {
		t.Fatalf("insert author: %v", err)
	}

	// Create temp posts directory simulating GitHub repo
	tempDir := t.TempDir()

	// Scenario: User edits markdown directly on GitHub and pushes
	post1 := filepath.Join(tempDir, "new-post-1.md")
	post1Content := `---
title: New Post from GitHub
slug: new-post-github
type: posts
---
This was edited directly on GitHub without using the web editor.
`
	if err := os.WriteFile(post1, []byte(post1Content), 0644); err != nil {
		t.Fatalf("write post1: %v", err)
	}

	// Run hash_sync (simulating GitHub Action running cmd/hash_sync)
	stats, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     tempDir,
		AuthorID: 1,
		DryRun:   false,
	})
	if err != nil {
		t.Fatalf("sync error: %v", err)
	}

	// GREEN: Verify post was added to Turso
	if stats.Added != 1 {
		t.Fatalf("expected 1 post added, got %d", stats.Added)
	}

	// Verify the file was rewritten with hash metadata
	rewrittenBytes, err := os.ReadFile(post1)
	if err != nil {
		t.Fatalf("read rewritten file: %v", err)
	}
	rewrittenContent := string(rewrittenBytes)
	if !contains(rewrittenContent, "content_hash") {
		t.Fatalf("expected content_hash in rewritten file")
	}

	// Verify Turso has the post
	q := libsqlssg.New(db)
	posts, err := q.GetPostsBySlugType(ctx, "new-post-github")
	if err != nil {
		t.Fatalf("query posts: %v", err)
	}
	if len(posts) == 0 {
		t.Fatalf("expected post in Turso, got none")
	}
	if posts[0].Title != "New Post from GitHub" {
		t.Fatalf("title mismatch: got %s", posts[0].Title)
	}
}

// TestGitHubPushWithoutEditor verifies GitHub-only changes bypass the editor
func TestGitHubPushWithoutEditor(t *testing.T) {
	db := setupHashDB(t)
	defer db.Close()
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, `INSERT INTO authors (username, name, password) VALUES ('author','Author','pwd')`); err != nil {
		t.Fatalf("insert author: %v", err)
	}

	tempDir := t.TempDir()

	// Create a post directly in markdown
	post := filepath.Join(tempDir, "github-only.md")
	content := `---
title: GitHub Only Post
slug: github-only
---
Never touched the editor UI.
`
	if err := os.WriteFile(post, []byte(content), 0644); err != nil {
		t.Fatalf("write post: %v", err)
	}

	// Run sync
	stats, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     tempDir,
		AuthorID: 1,
		DryRun:   false,
	})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	if stats.Added != 1 {
		t.Fatalf("expected 1 added, got %d", stats.Added)
	}

	// Verify Turso has it
	q := libsqlssg.New(db)
	posts, err := q.GetPostsBySlugType(ctx, "github-only")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
}

// TestGitHubActionWithDryRun verifies --dry-run flag prevents writes
func TestGitHubActionWithDryRun(t *testing.T) {
	db := setupHashDB(t)
	defer db.Close()
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, `INSERT INTO authors (username, name, password) VALUES ('author','Author','pwd')`); err != nil {
		t.Fatalf("insert author: %v", err)
	}

	tempDir := t.TempDir()
	post := filepath.Join(tempDir, "test.md")
	if err := os.WriteFile(post, []byte("---\ntitle: Test\nslug: test\n---\nContent"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	// Dry run should not write to DB or update file
	stats, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     tempDir,
		AuthorID: 1,
		DryRun:   true,
	})
	if err != nil {
		t.Fatalf("sync: %v", err)
	}

	// Should report what would happen
	if stats.Added != 1 {
		t.Fatalf("dry-run: expected 1 added (reported), got %d", stats.Added)
	}

	// But DB should be empty (no actual insert in dry-run)
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE slug = 'test'").Scan(&count); err != nil {
		t.Fatalf("query count: %v", err)
	}
	if count != 0 {
		t.Fatalf("dry-run: expected no post in DB, got %d", count)
	}
}

// TestMultipleGitHubPushes verifies repeated runs don't duplicate
func TestMultipleGitHubPushes(t *testing.T) {
	db := setupHashDB(t)
	defer db.Close()
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, `INSERT INTO authors (username, name, password) VALUES ('author','Author','pwd')`); err != nil {
		t.Fatalf("insert author: %v", err)
	}

	tempDir := t.TempDir()
	post := filepath.Join(tempDir, "post.md")
	if err := os.WriteFile(post, []byte("---\ntitle: Post\nslug: post\n---\nContent v1"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	// First push
	stats1, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     tempDir,
		AuthorID: 1,
		DryRun:   false,
	})
	if err != nil {
		t.Fatalf("sync1: %v", err)
	}
	if stats1.Added != 1 {
		t.Fatalf("expected 1 added, got %d", stats1.Added)
	}

	// Second push (no changes)
	stats2, err := syncMarkdownFiles(ctx, db, SyncOptions{
		Root:     tempDir,
		AuthorID: 1,
		DryRun:   false,
	})
	if err != nil {
		t.Fatalf("sync2: %v", err)
	}
	if stats2.Skipped != 1 {
		t.Fatalf("expected 1 skipped (unchanged), got added=%d updated=%d skipped=%d", stats2.Added, stats2.Updated, stats2.Skipped)
	}

	// Verify only one post in DB
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE slug = 'post'").Scan(&count); err != nil {
		t.Fatalf("query: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 post in DB, got %d", count)
	}
}

func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
