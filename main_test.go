package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Mr-Destructive/meetgor.com/models"
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

func TestCreateFeedsKeepsPostsArchiveAuthoritative(t *testing.T) {
	dir := t.TempDir()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll("templates", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll("public", 0755); err != nil {
		t.Fatal(err)
	}

	tmpl := `{{define "feed.html"}}ARCHIVE {{.FeedInfo.Title}}|{{range .FeedInfo.Posts}}{{.Frontmatter.Title}};{{end}}{{end}}{{define "post.html"}}POST {{.Post.Frontmatter.Title}}{{end}}`
	if err := os.WriteFile(filepath.Join("templates", "feed.html"), []byte(tmpl), 0644); err != nil {
		t.Fatal(err)
	}

	ssg := &models.SSG{
		Config: models.SSG_CONFIG{
			Blog: models.BlogConfig{
				OutputDir:           "public",
				BaseUrl:             "https://example.com",
				TemplatesDir:        "templates",
				DefaultFeedTemplate: "feed.html",
				DefaultPostTemplate: "post.html",
				Themes: map[string]models.Theme{
					"default":   {},
					"secondary": {},
				},
				PagesConfig: map[string]models.PageConfig{
					"posts": {
						TemplatePath:     "post.html",
						FeedTemplatePath: "feed.html",
					},
				},
			},
		},
		TemplateFS: template.Must(template.New("").Parse(tmpl)),
		Posts: []models.Post{
			{
				Frontmatter: models.FrontMatter{
					Title:  "Archive Root",
					Date:   "2024-01-01",
					Status: "published",
					Type:   "posts",
					Slug:   "posts",
				},
			},
			{
				Frontmatter: models.FrontMatter{
					Title:  "First",
					Date:   "2024-01-02",
					Status: "published",
					Type:   "posts",
					Slug:   "posts/first",
				},
			},
		},
		FeedPosts: []models.Feed{
			{
				Title: "Posts",
				Type:  "posts",
				Slug:  "posts",
				Posts: []models.Post{
					{
						Frontmatter: models.FrontMatter{
							Title:  "First",
							Date:   "2024-01-02",
							Status: "published",
							Type:   "posts",
							Slug:   "posts/first",
						},
					},
					{
						Frontmatter: models.FrontMatter{
							Title:  "Archive Root",
							Date:   "2024-01-01",
							Status: "published",
							Type:   "posts",
							Slug:   "posts",
						},
					},
				},
			},
		},
	}

	if err := (&CreateFeedsPlugin{PluginName: "createFeeds"}).Execute(ssg); err != nil {
		t.Fatalf("create feeds failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join("public", "posts", "index.html"))
	if err != nil {
		t.Fatalf("archive page missing: %v", err)
	}
	got := string(data)
	if !strings.Contains(got, "ARCHIVE Posts") {
		t.Fatalf("expected archive page to win, got %q", got)
	}
	if !strings.Contains(got, "First;") {
		t.Fatalf("expected archive listing to include first post, got %q", got)
	}
	if _, err := os.Stat(filepath.Join("public", "posts", "first", "index.html")); err != nil {
		t.Fatalf("detail page missing: %v", err)
	}
}

func TestReadPostsStripsEmbeddedJSONFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "broken.md")
	content := `---
title: Broken JSON Body
date: 2024-01-03
type: til
status: published
slug: broken-json-body
---
{
"title": "Broken JSON Body",
"type": "til",
"date": "2024-01-03",
"slug": "broken-json-body"
}

## Body
`
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
	if strings.HasPrefix(strings.TrimSpace(posts[0].Markdown), "{") {
		t.Fatalf("expected embedded json to be stripped, got %q", posts[0].Markdown)
	}
	if !strings.Contains(posts[0].Markdown, "## Body") {
		t.Fatalf("expected markdown body to remain, got %q", posts[0].Markdown)
	}
}

func TestIndexPluginListsAllFeedTypes(t *testing.T) {
	dir := t.TempDir()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(wd)
	})
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll("static", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll("public", 0755); err != nil {
		t.Fatal(err)
	}
	templateContent := `{{define "index.html"}}HOME|{{range .FeedPosts}}{{if ne .Type "all-content"}}{{.Title}}:{{len .Posts}};{{end}}{{end}}{{end}}`
	if err := os.WriteFile(filepath.Join("static", "index.html"), []byte(templateContent), 0644); err != nil {
		t.Fatal(err)
	}

	ssg := &models.SSG{
		Config: models.SSG_CONFIG{
			Blog: models.BlogConfig{
				OutputDir: "public",
				StaticDir: "static",
				BaseUrl:   "https://example.com",
				Themes: map[string]models.Theme{
					"default":   {},
					"secondary": {},
				},
			},
		},
		FeedPosts: []models.Feed{
			{Title: "Posts", Type: "posts", Slug: "posts", Posts: []models.Post{{}, {}}},
			{Title: "TIL", Type: "til", Slug: "til", Posts: []models.Post{{}}},
			{Title: "Newsletter", Type: "newsletter", Slug: "newsletter", Posts: []models.Post{{}, {}, {}}},
			{Title: "All Content", Type: "all-content", Slug: "all-content", Posts: []models.Post{{}}},
		},
	}

	if err := (&IndexPlugin{PluginName: "index"}).Execute(ssg); err != nil {
		t.Fatalf("index plugin failed: %v", err)
	}
	data, err := os.ReadFile(filepath.Join("public", "index.html"))
	if err != nil {
		t.Fatalf("home page missing: %v", err)
	}
	got := string(data)
	if !strings.Contains(got, "Posts:2;") || !strings.Contains(got, "TIL:1;") || !strings.Contains(got, "Newsletter:3;") {
		t.Fatalf("expected all feed types in homepage, got %q", got)
	}
	if strings.Contains(got, "All Content:") {
		t.Fatalf("expected all-content to be hidden from the type summary, got %q", got)
	}
}
