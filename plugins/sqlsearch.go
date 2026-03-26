package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	models "github.com/Mr-Destructive/meetgor.com/models"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

// PostForSQL represents the structure of a post for the SQL.js database
type PostForSQL struct {
	Title       string   `json:"title"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
	Content     string   `json:"content"` // Plain text content
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Series      string   `json:"series"`
	Status      string   `json:"status"`
	Slug        string   `json:"slug"`
}

// SQLSearchData holds all the data for the SQL.js database
type SQLSearchData struct {
	Posts      []PostForSQL `json:"posts"`
	Tils       []PostForSQL `json:"tils"`
	Newsletter []PostForSQL `json:"newsletter"`
	Work       []PostForSQL `json:"work"`
	Projects   []PostForSQL `json:"projects"`
	AllPosts   []PostForSQL `json:"all_posts"`
}

type SQLSearchPlugin struct {
	PluginName string
}

func (p *SQLSearchPlugin) Name() string {
	return p.PluginName
}

func (p *SQLSearchPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *SQLSearchPlugin) Requires() []string {
	return []string{"readPosts"}
}

func (p *SQLSearchPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

// stripHTML removes HTML tags and decodes HTML entities from a string.
func stripHTML(htmlContent string) string {
	// Remove HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	s := re.ReplaceAllString(htmlContent, "")
	// Replace common HTML entities (a more robust solution might use a library)
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&#39;", "'")
	// Replace multiple spaces with a single space
	re = regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// readMarkdownFile reads a markdown file, extracts front matter, and converts content to plain text.
func readMarkdownFile(filePath string) (*PostForSQL, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var frontmatterObj models.FrontMatter
	var contentBytes []byte
	success := false

	// Attempt to detect JSON front matter
	jsonSeparator := []byte("}\n\n")
	jsonIndex := bytes.Index(fileBytes, jsonSeparator)

	if jsonIndex != -1 {
		frontmatterBytes := fileBytes[:jsonIndex+1] // Keep closing brace
		contentBytes = fileBytes[jsonIndex+2:]      // Skip the separator

		if err := json.Unmarshal(frontmatterBytes, &frontmatterObj); err == nil {
			success = true
		}
	}

	// Attempt to detect YAML front matter if JSON failed or not found
	if !success {
		yamlSeparator := []byte("---\n\n")
		yamlIndex := bytes.Index(fileBytes, yamlSeparator)

		if yamlIndex != -1 {
			frontmatterBytes := fileBytes[:yamlIndex]
			contentBytes = fileBytes[yamlIndex+len(yamlSeparator):]

			if err := yaml.Unmarshal(frontmatterBytes, &frontmatterObj); err == nil {
				success = true
			}
		}
	}

	if !success {
		return nil, nil // No valid front matter found
	}

	// Convert Markdown to HTML, then strip HTML for plain content
	var htmlContentBuffer bytes.Buffer
	if err := goldmark.Convert(contentBytes, &htmlContentBuffer); err != nil {
		return nil, err
	}
	plainContent := stripHTML(htmlContentBuffer.String())

	// Ensure tags is not nil
	if frontmatterObj.Tags == nil {
		frontmatterObj.Tags = []string{}
	}

	// Extract series from Extras if present
	series := ""
	if s, ok := frontmatterObj.Extras["series"].(string); ok {
		series = s
	}

	// Determine slug
	slug := frontmatterObj.Slug
	if slug == "" {
		slug = Slugify(frontmatterObj.Title)
	}
	frontmatterObj.Type = NormalizePostType(frontmatterObj.Type)

	return &PostForSQL{
		Title:       frontmatterObj.Title,
		Date:        frontmatterObj.Date,
		Tags:        frontmatterObj.Tags,
		Content:     plainContent,
		Description: frontmatterObj.Description,
		Type:        frontmatterObj.Type,
		Series:      series,
		Status:      frontmatterObj.Status,
		Slug:        slug,
	}, nil
}

func (p *SQLSearchPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config

	sqlData := SQLSearchData{
		Posts:      []PostForSQL{},
		Tils:       []PostForSQL{},
		Newsletter: []PostForSQL{},
		Work:       []PostForSQL{},
		Projects:   []PostForSQL{},
		AllPosts:   []PostForSQL{},
	}

	for _, post := range ssg.Posts {
		if strings.EqualFold(post.Frontmatter.Status, "draft") {
			continue
		}
		typed := NormalizePostType(post.Frontmatter.Type)

		slug := strings.Trim(post.Frontmatter.Slug, "/")
		content := stripHTML(string(post.Content))
		series := ""
		if post.Frontmatter.Extras != nil {
			if s, ok := post.Frontmatter.Extras["series"].(string); ok {
				series = s
			}
		}

		tags := append([]string{}, post.Frontmatter.Tags...)
		if tags == nil {
			tags = []string{}
		}

		sqlPost := PostForSQL{
			Title:       post.Frontmatter.Title,
			Date:        post.Frontmatter.Date,
			Tags:        tags,
			Content:     content,
			Description: post.Frontmatter.Description,
			Type:        typed,
			Series:      series,
			Status:      post.Frontmatter.Status,
			Slug:        slug,
		}

		switch typed {
		case "posts":
			sqlData.Posts = append(sqlData.Posts, sqlPost)
		case "til":
			sqlData.Tils = append(sqlData.Tils, sqlPost)
		case "newsletter":
			sqlData.Newsletter = append(sqlData.Newsletter, sqlPost)
		case "work":
			sqlData.Work = append(sqlData.Work, sqlPost)
		case "project", "projects":
			sqlData.Projects = append(sqlData.Projects, sqlPost)
		default:
			sqlData.Posts = append(sqlData.Posts, sqlPost)
		}

		sqlData.AllPosts = append(sqlData.AllPosts, sqlPost)
	}

	jsonData, err := json.MarshalIndent(sqlData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling SQL data to JSON: %w", err)
	}

	outputDir := filepath.Join(".", config.Blog.OutputDir)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %w", outputDir, err)
	}

	jsonFilePath := filepath.Join(outputDir, "posts.json")
	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing posts.json: %w", err)
	}

	log.Printf("Generated %s with SQL search data.", jsonFilePath)
	return nil
}

func init() {
	RegisterPlugin("SQLSearch", func() Plugin {
		return &SQLSearchPlugin{PluginName: "SQLSearch"}
	})
}
