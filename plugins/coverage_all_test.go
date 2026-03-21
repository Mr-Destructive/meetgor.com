package plugins

import (
	"bytes"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mr-destructive/meetgor.com/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

func TestNewPlugin(t *testing.T) {
	if p, ok := NewPlugin("LinkPreview"); !ok || p == nil {
		t.Fatalf("expected LinkPreview plugin")
	}
	if p, ok := NewPlugin("NonExistent"); ok || p != nil {
		t.Fatalf("expected nil for unknown plugin")
	}
}

func TestMissingLinksFillerExecute(t *testing.T) {
	root := t.TempDir()
	oldwd, _ := os.Getwd()
	defer func() { _ = os.Chdir(oldwd) }()
	if err := os.Chdir(root); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	newsDir := filepath.Join(root, "posts", "newsletter")
	if err := os.MkdirAll(newsDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	newsletter := "## Read\n- [Test Post](https://example.com)\n"
	if err := os.WriteFile(filepath.Join(newsDir, "2024-07-27-techstructive-weekly-0.md"), []byte(newsletter), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	post := models.Post{
		Frontmatter: models.FrontMatter{
			Title:  "Test Post",
			Slug:   "test-post",
			Type:   "links",
			Extras: map[string]interface{}{},
		},
		Content: template.HTML("**Source:** 2024-07-27-techstructive-weekly-0"),
	}
	ssg := &models.SSG{Posts: []models.Post{post}}

	p := &MissingLinksFiller{PluginName: "MissingLinksFiller"}
	if err := p.Execute(ssg); err != nil {
		t.Fatalf("execute: %v", err)
	}
	link, ok := ssg.Posts[0].Frontmatter.Extras["link"]
	if !ok || link.(string) != "https://example.com" {
		t.Fatalf("expected link to be filled")
	}
}

func TestSyncDBPostsPluginEarlyExit(t *testing.T) {
	p := &SyncDBPostsPlugin{PluginName: "SyncDB"}
	ssg := &models.SSG{Config: models.SSG_CONFIG{AdminMode: true}}
	if err := p.Execute(ssg); err != nil {
		t.Fatalf("expected nil: %v", err)
	}
}

func TestSyncDBPostsPluginRunSuccess(t *testing.T) {
	old := syncDBExecCommand
	defer func() { syncDBExecCommand = old }()
	syncDBExecCommand = func(name string, args ...string) *exec.Cmd {
		return exec.Command("true")
	}

	_ = os.Setenv("SYNC_DB_ON_BUILD", "1")
	defer os.Unsetenv("SYNC_DB_ON_BUILD")

	p := &SyncDBPostsPlugin{PluginName: "SyncDB"}
	ssg := &models.SSG{Config: models.SSG_CONFIG{Blog: models.BlogConfig{PostsDir: "posts"}}}
	if err := p.Execute(ssg); err != nil {
		t.Fatalf("expected nil: %v", err)
	}
}

func TestSeasonalEffectsPlugin(t *testing.T) {
	root := t.TempDir()
	oldwd, _ := os.Getwd()
	defer func() { _ = os.Chdir(oldwd) }()
	if err := os.Chdir(root); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	outputDir := "public"
	if err := os.MkdirAll(filepath.Join(root, outputDir), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	ssg := &models.SSG{Config: models.SSG_CONFIG{Blog: models.BlogConfig{OutputDir: outputDir}}}
	p := &SeasonalEffectsPlugin{PluginName: "SeasonalEffects"}
	if err := p.Execute(ssg); err != nil {
		t.Fatalf("execute: %v", err)
	}
	path := filepath.Join(root, outputDir, "seasonal-effects.js")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected seasonal-effects.js")
	}
}

func TestDbPluginExecute(t *testing.T) {
	root := t.TempDir()
	outDir := filepath.Join(root, "public")
	ssg := &models.SSG{
		Config:     models.SSG_CONFIG{Blog: models.BlogConfig{OutputDir: outDir}},
		TemplateFS: template.Must(template.New("editor_template.html").Parse(`{{define "editor_template.html"}}ok{{end}}`)),
	}
	p := &DbPlugin{PluginName: "Db"}
	if err := p.Execute(ssg); err != nil {
		t.Fatalf("execute: %v", err)
	}
	path := filepath.Join(".", outDir, "editor", "index.html")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected editor index: %v", err)
	}
}

func TestGoAndSQLPlaygroundExtenders(t *testing.T) {
	// Ensure Extend and RegisterFuncs are callable without panic.
	md := goldmark.New(goldmark.WithExtensions(&GoPlaygroundExtender{}, &SQLPlaygroundExtender{}))
	var buf bytes.Buffer
	_ = md.Convert([]byte("```go\npackage main\n```"), &buf)
	gr := newGoPlaygroundRenderer().(*goPlaygroundRenderer)
	fr := &fakeRegister{}
	gr.RegisterFuncs(fr)
	if !fr.called {
		t.Fatalf("expected RegisterFuncs to register")
	}
}

type fakeRegister struct{ called bool }

func (f *fakeRegister) Register(kind ast.NodeKind, fn renderer.NodeRendererFunc) {
	f.called = true
}
