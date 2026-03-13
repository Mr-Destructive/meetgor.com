package main

import (
	"database/sql"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetEffectiveSlugFromFile(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "post.md")
	content := `---
slug: custom-slug
title: Title
---
body`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	logger := log.New(io.Discard, "", 0)
	slug, ok := getEffectiveSlugFromFile(file, logger, true)
	if !ok || slug != "custom-slug" {
		t.Fatalf("expected custom-slug, got %s (ok=%t)", slug, ok)
	}
}

func TestGetEffectiveTitleFromFile(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "post.md")
	content := `---
title: YAML Title
---
body`
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	logger := log.New(io.Discard, "", 0)
	title, ok := getEffectiveTitleFromFile(file, logger, true)
	if !ok || title != "YAML Title" {
		t.Fatalf("expected YAML Title, got %s", title)
	}
}

func TestParseConfigRequiresEnv(t *testing.T) {
	t.Setenv("TURSO_DATABASE_NAME", "db")
	t.Setenv("TURSO_DATABASE_AUTH_TOKEN", "token")
	os.Args = []string{"sync_db", "--output-dir", "out", "--hours", "2"}
	cfg, err := parseConfig()
	if err != nil {
		t.Fatalf("parseConfig failed: %v", err)
	}
	if cfg.OutputDir != "out" {
		t.Fatalf("expected output dir out, got %s", cfg.OutputDir)
	}
}

func TestTimestampHelpers(t *testing.T) {
	withTempDir(t, func() {
		first := getLastBuildTime()
		if time.Since(first) > time.Hour*2 {
			t.Fatalf("expected recent time, got %v", first)
		}
		if err := updateTimestamp(); err != nil {
			t.Fatalf("updateTimestamp failed: %v", err)
		}
		if _, err := os.Stat(timestampFile); err != nil {
			t.Fatalf("timestamp file missing: %v", err)
		}
	})
}

func TestParseMetadataInvalid(t *testing.T) {
	meta, err := parseMetadata("notjson")
	if err != nil {
		t.Fatalf("parseMetadata should not error: %v", err)
	}
	if meta == nil || meta.Title != "" {
		t.Fatalf("expected empty metadata, got %+v", meta)
	}
}

func TestCreatePostContent(t *testing.T) {
	post := Post{
		Title:   "Title",
		Created: "2024-01-01T00:00:00Z",
		Updated: "2024-01-01T01:00:00Z",
	}
	cfg := &Config{DryRun: true}
	logger := log.New(io.Discard, "", 0)
	meta := &PostMetadata{Tags: []string{"a"}}
	content, err := createPostContent(post, meta, "slug", "body", cfg.Verbose, logger)
	if err != nil {
		t.Fatalf("createPostContent failed: %v", err)
	}
	if !strings.Contains(content, `"title":`) {
		t.Fatalf("expected JSON front matter, got %s", content)
	}
}

func TestParseTimeFormats(t *testing.T) {
	cases := []string{"2024-01-01T00:00:00Z", "2006-01-02 15:04:05"}
	for _, c := range cases {
		if _, err := parseTime(c); err != nil {
			t.Fatalf("parseTime failed for %s: %v", c, err)
		}
	}
	if _, err := parseTime(""); err == nil {
		t.Fatalf("expected error for empty time")
	}
}

func TestWriteFileAndSlugify(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "demo.md")
	if err := writeFile(path, "content"); err != nil {
		t.Fatalf("writeFile failed: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "content" {
		t.Fatalf("unexpected file content: %v %v", data, err)
	}
	if slugify("Hello World!") != "hello-world" {
		t.Fatalf("slugify mismatch")
	}
}

func TestProcessPostDryRun(t *testing.T) {
	dir := t.TempDir()
	post := Post{
		ID:       1,
		Title:    "Title",
		Body:     "body",
		Created:  "2024-01-01T00:00:00Z",
		Updated:  "2024-01-01T01:00:00Z",
		Metadata: "{}",
	}
	cfg := &Config{OutputDir: dir, DryRun: true, SyncAll: true}
	logger := log.New(io.Discard, "", 0)
	if err := processPost(post, cfg, logger); err != nil {
		t.Fatalf("processPost failed: %v", err)
	}
}

func TestFetchPosts(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}
	defer db.Close()
	ddl := `
CREATE TABLE posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	slug TEXT,
	body TEXT,
	metadata TEXT,
	created_at TEXT,
	updated_at TEXT,
	author_id INTEGER,
	deleted BOOLEAN
)`
	if _, err := db.Exec(ddl); err != nil {
		t.Fatalf("create table: %v", err)
	}
	now := time.Now().Format(timeFormat)
	_, err = db.Exec(`INSERT INTO posts(title, slug, body, metadata, created_at, updated_at, author_id, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, 0)`,
		"Title", "slug", "body", "{}", now, now, 1)
	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}
	cfg := &Config{SyncAll: true, OutputDir: t.TempDir()}
	logger := log.New(io.Discard, "", 0)
	posts, err := fetchPosts(db, cfg, nil, nil, logger)
	if err != nil {
		t.Fatalf("fetchPosts failed: %v", err)
	}
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
}

func TestFetchPostsDifferential(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}
	defer db.Close()
	ddl := `
CREATE TABLE posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	slug TEXT,
	body TEXT,
	metadata TEXT,
	created_at TEXT,
	updated_at TEXT,
	author_id INTEGER,
	deleted BOOLEAN
)`
	if _, err := db.Exec(ddl); err != nil {
		t.Fatalf("create table: %v", err)
	}
	now := time.Now().Format(timeFormat)
	_, err = db.Exec(`INSERT INTO posts(title, slug, body, metadata, created_at, updated_at, author_id, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, 0)`,
		"Title", "slug", "body", "{}", now, now, 1)
	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}
	cfg := &Config{SyncAll: false, OutputDir: t.TempDir(), TimeCutoff: time.Now().Add(time.Hour)}
	logger := log.New(io.Discard, "", 0)
	diffSlug := map[string]bool{"slug": true}
	posts, err := fetchPosts(db, cfg, diffSlug, nil, logger)
	if err != nil {
		t.Fatalf("fetchPosts differential failed: %v", err)
	}
	if len(posts) != 0 {
		t.Fatalf("expected 0 posts when slug already exists, got %d", len(posts))
	}
}

func TestProcessPostWritesFile(t *testing.T) {
	dir := t.TempDir()
	post := Post{
		ID:       2,
		Title:    "Write Test",
		Body:     "content",
		Created:  "2024-01-01T00:00:00Z",
		Updated:  "2024-01-01T01:00:00Z",
		Metadata: `{"slug":"write-test","type":"posts"}`,
	}
	cfg := &Config{OutputDir: dir, DryRun: false, SyncAll: true}
	logger := log.New(io.Discard, "", 0)
	if err := processPost(post, cfg, logger); err != nil {
		t.Fatalf("processPost failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "posts", "write-test.md")); err != nil {
		t.Fatalf("expected file to be written, got %v", err)
	}
}

func withTempDir(t *testing.T, fn func()) {
	dir := t.TempDir()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(old)
	})
	fn()
}
