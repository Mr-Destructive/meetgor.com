package plugins

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
	
	models "github.com/mr-destructive/mr-destructive.github.io/models"
)

type StaticPageRendererPlugin struct {
	PluginName string
}

func (p *StaticPageRendererPlugin) Name() string {
	return p.PluginName
}

func (p *StaticPageRendererPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	templateFS := os.DirFS(config.Blog.TemplatesDir)

	outputPath := filepath.Join(".", config.Blog.OutputDir)
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Check if public directory exists and process markdown files
	publicPath := filepath.Join(".", "public")
	if _, err := os.Stat(publicPath); err == nil {
		// Public directory exists, process markdown files
		files, err := os.ReadDir(publicPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
				// Process markdown files like about.md, contact.md
				pageName := strings.TrimSuffix(file.Name(), ".md")
				
				// Check if this page is configured in pages config
				if pageConfig, exists := config.Blog.PagesConfig[pageName]; exists && pageConfig.TemplatePath != "" && pageName != "posts" {
					content, err := os.ReadFile(filepath.Join(publicPath, file.Name()))
					if err != nil {
						log.Printf("Error reading file %s: %v", file.Name(), err)
						continue
					}

					// Parse frontmatter and content
					frontmatter, markdownContent := parseFrontMatter(content)
					
					// Convert markdown to HTML
					var htmlContent bytes.Buffer
					if err := goldmark.Convert([]byte(markdownContent), &htmlContent); err != nil {
						log.Printf("Error converting markdown for %s: %v", file.Name(), err)
						continue
					}

					pageOutputPath := filepath.Join(outputPath, pageName)
					err = os.MkdirAll(pageOutputPath, os.ModePerm)
					if err != nil {
						log.Fatal(err)
					}
					outputPagePath := filepath.Join(pageOutputPath, "index.html")

					buffer := bytes.Buffer{}
					context := models.TemplateContext{
						Themes: models.ThemeCombo{
							Default:   config.Blog.Themes["default"],
							Secondary: config.Blog.Themes["secondary"],
						},
						Config: models.SSG_CONFIG{
							Blog:      config.Blog,
							AdminMode: config.AdminMode,
						},
						Post: models.Post{
							Frontmatter: frontmatter,
							Content:     template.HTML(htmlContent.String()),
						},
						FeedInfo: models.Feed{
							Title: strings.ToTitle(pageName),
							Type:  pageName,
							Slug:  config.Blog.PrefixURL + pageName,
						},
					}
					fmt.Println("Rendering static page:", pageName, "using template:", pageConfig.TemplatePath)

					// Parse all templates that might be needed
			patterns := []string{pageConfig.TemplatePath, "*.html"}
			t, err := template.ParseFS(templateFS, patterns...)
			if err != nil {
				log.Printf("Error parsing templates for %s: %v", pageConfig.TemplatePath, err)
				// Fallback to just the main template
				t, err = template.ParseFS(templateFS, pageConfig.TemplatePath)
				if err != nil {
					log.Fatal(err)
				}
			}
			err = t.ExecuteTemplate(&buffer, pageConfig.TemplatePath, context)
					if err != nil {
						log.Fatal(err)
					}
					err = os.WriteFile(outputPagePath, buffer.Bytes(), 0660)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	// Also process pages from config that don't have markdown files
	for pageName, pageConfig := range config.Blog.PagesConfig {
		// Only process pages that have a template path and are not the 'posts' type
		if pageConfig.TemplatePath != "" && pageName != "posts" {
			// Skip if we already processed a markdown file for this page
			markdownFile := filepath.Join(publicPath, pageName+".md")
			if _, err := os.Stat(markdownFile); err == nil {
				continue // Already processed above
			}

			pageOutputPath := filepath.Join(outputPath, pageName)
			err = os.MkdirAll(pageOutputPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			outputPagePath := filepath.Join(pageOutputPath, "index.html")

			buffer := bytes.Buffer{}
			context := models.TemplateContext{
				Themes: models.ThemeCombo{
					Default:   config.Blog.Themes["default"],
					Secondary: config.Blog.Themes["secondary"],
				},
				Config: models.SSG_CONFIG{
					Blog:      config.Blog,
					AdminMode: config.AdminMode,
				},
				FeedInfo: models.Feed{
					Title: strings.ToTitle(pageName),
					Type:  pageName,
					Slug:  config.Blog.PrefixURL + pageName,
				},
			}
			fmt.Println("Rendering static page:", pageName, "using template:", pageConfig.TemplatePath)

			// Parse all templates that might be needed
			patterns := []string{pageConfig.TemplatePath, "*.html"}
			t, err := template.ParseFS(templateFS, patterns...)
			if err != nil {
				log.Printf("Error parsing templates for %s: %v", pageConfig.TemplatePath, err)
				// Fallback to just the main template
				t, err = template.ParseFS(templateFS, pageConfig.TemplatePath)
				if err != nil {
					log.Fatal(err)
				}
			}
			err = t.ExecuteTemplate(&buffer, pageConfig.TemplatePath, context)
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(outputPagePath, buffer.Bytes(), 0660)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func parseFrontMatter(content []byte) (models.FrontMatter, string) {
	// Simple frontmatter parser - assumes YAML frontmatter between ---
	lines := strings.Split(string(content), "\n")
	frontmatter := models.FrontMatter{}
	contentStart := len(lines) // Default to whole content if no frontmatter
	
	if len(lines) > 0 && strings.TrimSpace(lines[0]) == "---" {
		// Found start of frontmatter
		for i := 1; i < len(lines); i++ {
			if strings.TrimSpace(lines[i]) == "---" {
				// Found end of frontmatter
				contentStart = i + 1
				break
			}
		}
		
		// Parse YAML frontmatter if we found both --- markers
		if contentStart < len(lines) {
			frontmatterContent := strings.Join(lines[1:contentStart-1], "\n")
			yaml.Unmarshal([]byte(frontmatterContent), &frontmatter)
		}
	}
	
	// Return frontmatter and content
	contentLines := lines[contentStart:]
	return frontmatter, strings.Join(contentLines, "\n")
}
