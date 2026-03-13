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

func TestParseHours(t *testing.T) {
	if parseHours("2") != 2 {
		t.Fatalf("expected 2")
	}
	if parseHours("bad") != 1 {
		t.Fatalf("expected default 1")
	}
}

func TestParseConfig(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--sync-all", "--dry-run", "--verbose", "--output-dir", "out", "--hours", "2"}
	oldName := os.Getenv("TURSO_DATABASE_NAME")
	oldTok := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	os.Setenv("TURSO_DATABASE_NAME", "libsql://x")
	os.Setenv("TURSO_DATABASE_AUTH_TOKEN", "t")
	defer func() {
		os.Setenv("TURSO_DATABASE_NAME", oldName)
		os.Setenv("TURSO_DATABASE_AUTH_TOKEN", oldTok)
	}()
	cfg, err := parseConfig()
	if err != nil || !cfg.SyncAll || !cfg.DryRun || !cfg.Verbose || cfg.OutputDir != "out" {
		t.Fatalf("parseConfig failed")
	}
}

func TestConnectDB(t *testing.T) {
	old := dbOpen
	dbOpen = func(string, string) (*sql.DB, error) { return sql.Open("sqlite3", ":memory:") }
	defer func() { dbOpen = old }()
	if _, err := connectDB("x"); err != nil {
		t.Fatalf("connectDB failed")
	}
}

func TestSlugify(t *testing.T) {
	if slugify("Hello World!") != "hello-world" {
		t.Fatalf("slugify mismatch")
	}
}

func TestParseTime(t *testing.T) {
	if _, err := parseTime("2024-01-01"); err != nil {
		t.Fatalf("parseTime failed")
	}
	if _, err := parseTime("bad"); err == nil {
		t.Fatalf("expected error")
	}
}

func TestParseMetadata(t *testing.T) {
	m, err := parseMetadata(`{"title":"t"}`)
	if err != nil || m.Title != "t" {
		t.Fatalf("parse metadata failed")
	}
	m, err = parseMetadata("{bad}")
	if err != nil {
		t.Fatalf("parse metadata should not error")
	}
}

func TestDetermineFilePath(t *testing.T) {
	meta := &PostMetadata{Type: "posts"}
	path, err := determineFilePath(meta, "slug", "out")
	if err != nil || !strings.Contains(path, "out") {
		t.Fatalf("determineFilePath failed")
	}
}

func TestCreatePostContentAndWrite(t *testing.T) {
	post := Post{Title: "T", Body: "B", Created: "2024-01-01", Updated: "2024-01-01"}
	meta := &PostMetadata{Type: "posts"}
	content, err := createPostContent(post, meta, "slug", "body", true, log.Default())
	if err != nil || !strings.Contains(content, "title") {
		t.Fatalf("createPostContent failed")
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "a.md")
	if err := writeFile(path, content); err != nil {
		t.Fatalf("writeFile failed")
	}
}

func TestCreatePostContentDraft(t *testing.T) {
	post := Post{Title: "T", Body: "B", Created: "2024-01-01", Updated: "2024-01-01"}
	meta := &PostMetadata{Type: "posts", Draft: true}
	content, err := createPostContent(post, meta, "slug", "body", false, log.Default())
	if err != nil {
		t.Fatalf("createPostContent failed for draft: %v", err)
	}
	if !strings.Contains(content, `"draft": true`) {
		t.Fatalf("expected draft flag in frontmatter: %s", content)
	}
}

func TestShouldWriteFileUpToDate(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a.md")
	if err := os.WriteFile(path, []byte("body"), 0644); err != nil {
		t.Fatal(err)
	}
	now := time.Now().Add(-time.Hour)
	if err := os.Chtimes(path, now, now); err != nil {
		t.Fatal(err)
	}
	ok, reason, err := shouldWriteFile(path, now.Add(-time.Minute))
	if err != nil || ok {
		t.Fatalf("expected false for up-to-date file, got ok=%v reason=%s err=%v", ok, reason, err)
	}
}

func TestProcessPostMissingSlug(t *testing.T) {
	dir := t.TempDir()
	post := Post{
		ID:       3,
		Title:    "",
		Body:     "body",
		Created:  "2024-01-01T00:00:00Z",
		Updated:  "2024-01-01T01:00:00Z",
		Metadata: `{}`,
	}
	cfg := &Config{OutputDir: dir, DryRun: false, SyncAll: true}
	logger := log.New(io.Discard, "", 0)
	if err := processPost(post, cfg, logger); err == nil {
		t.Fatalf("expected error when slug cannot be determined")
	}
}

func TestEffectiveSlugTitle(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a.md")
	content := "---\ntitle: Hello\nslug: hi\n---\nBody"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	slug, ok := getEffectiveSlugFromFile(path, log.Default(), false)
	if !ok || slug != "hi" {
		t.Fatalf("slug not read")
	}
	title, ok := getEffectiveTitleFromFile(path, log.Default(), false)
	if !ok || title != "Hello" {
		t.Fatalf("title not read")
	}
}

func TestGetExistingFileSlugs(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "posts"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "posts", "a.md"), []byte("---\ntitle: Hello\nslug: hi\n---\n"), 0644); err != nil {
		t.Fatal(err)
	}
	slugs, titles, err := getExistingFileSlugs(dir, log.Default(), false)
	if err != nil {
		t.Fatal(err)
	}
	if !slugs["hi"] || !titles["Hello"] {
		t.Fatalf("expected slug and title")
	}
}

// TestShouldWriteFile is covered in main_test.go
