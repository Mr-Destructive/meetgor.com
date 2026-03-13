package plugins

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func withTempCwd(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	wd, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
	return dir
}

func minimalSSG(t *testing.T, outputDir string) *models.SSG {
	t.Helper()
	cfg := models.SSG_CONFIG{
		Blog: models.BlogConfig{
			OutputDir: outputDir,
			BaseUrl:   "https://example.com",
			PrefixURL: "",
			Themes: map[string]models.Theme{
				"default":   {},
				"secondary": {},
			},
			PagesConfig: map[string]models.PageConfig{
				"tags":   {TemplatePath: "feed.html"},
				"series": {TemplatePath: "feed.html"},
			},
			DefaultFeedTemplate: "feed.html",
		},
	}
	tmpl := template.New("")
	tmpl = template.Must(tmpl.Parse(`{{define "feed.html"}}ok{{end}}{{define "editor_template.html"}}ok{{end}}{{define "feeds_template.html"}}ok{{end}}`))
	return &models.SSG{
		Config:     cfg,
		Posts:      []models.Post{},
		TemplateFS: tmpl,
	}
}

func TestSlugify(t *testing.T) {
	if Slugify("Hello World!") != "hello-world" {
		t.Fatalf("slugify mismatch")
	}
}

func TestGenerateID(t *testing.T) {
	if generateID("Hello World!") != "hello-world" {
		t.Fatalf("generateID mismatch")
	}
}

func TestGenerateTOC(t *testing.T) {
	md := goldmark.New(goldmark.WithParserOptions(parser.WithAutoHeadingID()))
	src := []byte("# Title\n## Sub\n")
	reader := md.Parser()
	ctx := parser.NewContext()
	node := reader.Parse(text.NewReader(src), parser.WithContext(ctx))
	toc := GenerateTOC(node, src)
	if len(toc) != 1 || toc[0].Text != "Title" {
		t.Fatalf("unexpected toc")
	}
}

func TestTagsSeriesYearWise(t *testing.T) {
	withTempCwd(t)
	ssg := minimalSSG(t, "public")
	ssg.Config.Blog.PagesConfig["tags"] = models.PageConfig{TemplatePath: "feed.html"}
	ssg.Config.Blog.PagesConfig["series"] = models.PageConfig{TemplatePath: "feed.html"}
	ssg.Posts = []models.Post{
		{
			Frontmatter: models.FrontMatter{
				Title: "Post A", Date: "2024-01-02", Status: "published", Type: "posts",
				Tags:   []string{"go"},
				Extras: map[string]interface{}{"series": "Alpha"},
			},
		},
	}

	if err := (&TagsPlugin{PluginName: "Tags"}).Execute(ssg); err != nil {
		t.Fatalf("tags plugin failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "tags", "go", "index.html")); err != nil {
		t.Fatalf("tags output missing: %v", err)
	}

	if err := (&SeriesPlugin{PluginName: "Series"}).Execute(ssg); err != nil {
		t.Fatalf("series plugin failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "series", "alpha", "index.html")); err != nil {
		t.Fatalf("series output missing: %v", err)
	}

	if err := (&YearPlugin{PluginName: "YearWise"}).Execute(ssg); err != nil {
		t.Fatalf("yearwise plugin failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "2024", "index.html")); err != nil {
		t.Fatalf("yearwise output missing: %v", err)
	}
}

func TestBeforeAfterPosts(t *testing.T) {
	ssg := minimalSSG(t, "public")
	ssg.Posts = []models.Post{
		{Frontmatter: models.FrontMatter{Title: "A", Date: "2024-01-01", Status: "published", Type: "posts", Tags: []string{"go"}, Slug: "a"}},
		{Frontmatter: models.FrontMatter{Title: "B", Date: "2024-02-01", Status: "published", Type: "posts", Tags: []string{"go"}, Slug: "b"}},
		{Frontmatter: models.FrontMatter{Title: "C", Date: "2024-03-01", Status: "published", Type: "posts", Tags: []string{"go"}, Slug: "c"}},
	}
	if err := (&BeforeAfterPostsPlugin{PluginName: "BeforeAfterPosts"}).Execute(ssg); err != nil {
		t.Fatalf("before/after plugin failed: %v", err)
	}
	extras := ssg.Posts[1].Frontmatter.Extras
	if extras == nil {
		t.Fatalf("expected extras")
	}
	if extras["nav_previous"] == nil || extras["nav_next"] == nil {
		t.Fatalf("expected nav prev/next")
	}
}

func TestSeasonalEffects(t *testing.T) {
	js := generateSeasonalEffectsJS()
	if !strings.Contains(js, "Seasonal Effects") {
		t.Fatalf("unexpected js content")
	}
}

func TestRSSAndSitemap(t *testing.T) {
	withTempCwd(t)
	ssg := minimalSSG(t, "public")
	if err := os.MkdirAll("public", 0755); err != nil {
		t.Fatal(err)
	}
	ssg.Posts = []models.Post{
		{Frontmatter: models.FrontMatter{Title: "A", Date: "2024-01-01", Status: "published", Type: "posts", Slug: "posts/a"}},
	}
	rssNow = func() time.Time { return time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC) }
	sitemapNow = func() time.Time { return time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC) }
	if err := (&RSSPlugin{PluginName: "RSS"}).Execute(ssg); err != nil {
		t.Fatalf("rss plugin failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "rss.xml")); err != nil {
		t.Fatalf("rss missing: %v", err)
	}
	if err := (&SitemapPlugin{PluginName: "Sitemap"}).Execute(ssg); err != nil {
		t.Fatalf("sitemap plugin failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "sitemap.xml")); err != nil {
		t.Fatalf("sitemap missing: %v", err)
	}
}

func TestStaticPageRenderer(t *testing.T) {
	dir := withTempCwd(t)
	templatesDir := filepath.Join(dir, "templates")
	publicDir := filepath.Join(dir, "public")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(publicDir, "about.md"), []byte("---\ntitle: About\n---\nHi"), 0644); err != nil {
		t.Fatal(err)
	}
	templateContent := `{{define "page.html"}}<h1>{{.Post.Frontmatter.Title}}</h1>{{end}}`
	if err := os.WriteFile(filepath.Join(templatesDir, "page.html"), []byte(templateContent), 0644); err != nil {
		t.Fatal(err)
	}

	ssg := minimalSSG(t, "public")
	ssg.Config.Blog.TemplatesDir = "templates"
	ssg.Config.Blog.PagesConfig = map[string]models.PageConfig{
		"about": {TemplatePath: "page.html"},
	}
	if err := (&StaticPageRendererPlugin{PluginName: "StaticPageRenderer"}).Execute(ssg); err != nil {
		t.Fatalf("static page renderer failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "about", "index.html")); err != nil {
		t.Fatalf("static page missing: %v", err)
	}
}

func TestLinkPreviewPlugin(t *testing.T) {
	dir := withTempCwd(t)
	path := filepath.Join(dir, "link.md")
	markdown := `---
type: links
title: "Link Post"
link: https://example.com
status: published
---

Body`
	if err := os.WriteFile(path, []byte(markdown), 0644); err != nil {
		t.Fatal(err)
	}

	originalFactory := linkPreviewClientFactory
	linkPreviewClientFactory = func() *http.Client {
		return &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			body := `<meta property="og:image" content="https://example.com/og.png">`
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
		})}
	}
	defer func() { linkPreviewClientFactory = originalFactory }()

	ssg := &models.SSG{
		Posts: []models.Post{
			{
				Frontmatter: models.FrontMatter{
					Type:   "links",
					Status: "published",
					Extras: map[string]interface{}{"link": "https://example.com"},
				},
				SourcePath: path,
			},
		},
	}

	if err := (&LinkPreviewPlugin{PluginName: "LinkPreview"}).Execute(ssg); err != nil {
		t.Fatalf("link preview failed: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	expected := "image_url: https://example.com/og.png"
	if !strings.Contains(string(data), expected) {
		t.Fatalf("expected image_url in front matter")
	}
	if ssg.Posts[0].Frontmatter.ImageUrl != "https://example.com/og.png" {
		t.Fatalf("expected post frontmatter to update image_url")
	}
}

func TestSQLSearch(t *testing.T) {
	dir := withTempCwd(t)
	postsDir := filepath.Join(dir, "posts")
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := "---\ntitle: T\ndate: 2024-01-01\ntype: posts\nstatus: published\n---\nHello"
	if err := os.WriteFile(filepath.Join(postsDir, "a.md"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	ssg := minimalSSG(t, "public")
	ssg.Config.Blog.PostsDir = "posts"
	if err := (&SQLSearchPlugin{PluginName: "SQLSearch"}).Execute(ssg); err != nil {
		t.Fatalf("sql search failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join("public", "posts.json")); err != nil {
		t.Fatalf("sqlsearch output missing: %v", err)
	}
}

func TestReadMarkdownFileVariants(t *testing.T) {
	dir := withTempCwd(t)
	path1 := filepath.Join(dir, "a.md")
	content1 := "{\n  \"title\": \"J\",\n  \"date\": \"2024-01-01\",\n  \"type\": \"posts\",\n  \"status\": \"published\"\n}\n\nHello"
	if err := os.WriteFile(path1, []byte(content1), 0644); err != nil {
		t.Fatal(err)
	}
	post, err := readMarkdownFile(path1)
	if err != nil || post == nil || post.Title != "J" {
		t.Fatalf("json frontmatter read failed")
	}

	path2 := filepath.Join(dir, "b.md")
	content2 := "---\ntitle: Y\ndate: 2024-01-01\ntype: posts\nstatus: published\n---\n\nHi"
	if err := os.WriteFile(path2, []byte(content2), 0644); err != nil {
		t.Fatal(err)
	}
	post, err = readMarkdownFile(path2)
	if err != nil || post == nil || post.Title != "Y" {
		t.Fatalf("yaml frontmatter read failed")
	}

	path3 := filepath.Join(dir, "c.md")
	if err := os.WriteFile(path3, []byte("no fm"), 0644); err != nil {
		t.Fatal(err)
	}
	post, err = readMarkdownFile(path3)
	if err != nil || post != nil {
		t.Fatalf("expected nil for missing frontmatter")
	}
}

func TestDBCreatePostPayload(t *testing.T) {
	payload := Payload{
		Title: "Hello",
		Post:  "Body",
		Metadata: map[string]interface{}{
			"type": "posts",
		},
	}
	post, err := CreatePostPayload(payload, 1, "Meet")
	if err != nil {
		t.Fatalf("CreatePostPayload error: %v", err)
	}
	if post.Title != "Hello" || post.AuthorID != 1 {
		t.Fatalf("unexpected post payload")
	}
}

func TestGoSQLPlaygroundRender(t *testing.T) {
	md := goldmark.New(goldmark.WithExtensions(extension.GFM))
	src := []byte("```go\npackage main\n```\n```sql\nselect 1;\n```")
	ctx := parser.NewContext()
	node := md.Parser().Parse(text.NewReader(src), parser.WithContext(ctx))

	var goBuf bytes.Buffer
	var sqlBuf bytes.Buffer
	gr := newGoPlaygroundRenderer().(*goPlaygroundRenderer)
	sr := newSQLPlaygroundRenderer().(*sqlPlaygroundRenderer)

	var goFound, sqlFound bool
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() == ast.KindFencedCodeBlock {
			f := n.(*ast.FencedCodeBlock)
			lang := string(f.Language(src))
			if lang == "go" {
				goFound = true
				bw := bufio.NewWriter(&goBuf)
				_, _ = gr.renderFencedCodeBlock(bw, src, n, true)
				_, _ = gr.renderFencedCodeBlock(bw, src, n, false)
				_ = bw.Flush()
			}
			if lang == "sql" {
				sqlFound = true
				bw := bufio.NewWriter(&sqlBuf)
				_, _ = sr.renderFencedCodeBlock(bw, src, n, true)
				_, _ = sr.renderFencedCodeBlock(bw, src, n, false)
				_ = bw.Flush()
			}
		}
		return ast.WalkContinue, nil
	})
	if !goFound || !sqlFound {
		t.Fatalf("expected fenced code blocks")
	}
	if !strings.Contains(goBuf.String(), "go-playground") || !strings.Contains(sqlBuf.String(), "sql-playground") {
		t.Fatalf("playground render missing")
	}
}

func TestPluginErrorBranches(t *testing.T) {
	ssg := minimalSSG(t, "missing/output")
	ssg.Config.Blog.PostsDir = "missing"
	if err := (&SQLSearchPlugin{PluginName: "SQLSearch"}).Execute(ssg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat(filepath.Join("missing/output", "posts.json")); err != nil {
		t.Fatalf("expected posts.json to exist: %v", err)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
