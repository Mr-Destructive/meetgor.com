package plugins

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	models "github.com/mr-destructive/mr-destructive.github.io/models"
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

func (p *SQLSearchPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	postsDir := config.Blog.PostsDir // Typically "content/posts"

	sqlData := SQLSearchData{
		Posts:      []PostForSQL{},
		Tils:       []PostForSQL{},
		Newsletter: []PostForSQL{},
		Work:       []PostForSQL{},
		Projects:   []PostForSQL{},
		AllPosts:   []PostForSQL{},
	}

	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}
		if filepath.Ext(path) != ".md" {
			return nil // Process only markdown files
		}

		post, err := readMarkdownFile(path)
		if err != nil {
			log.Printf("Error reading markdown file %s: %v", path, err)
			return nil // Continue processing other files
		}
		if post == nil || post.Status == "draft" {
			return nil // Skip if no valid front matter or is a draft
		}

		// Add to specific type array
		switch post.Type {
		case "posts":
			sqlData.Posts = append(sqlData.Posts, *post)
		case "til":
			sqlData.Tils = append(sqlData.Tils, *post)
		case "newsletter":
			sqlData.Newsletter = append(sqlData.Newsletter, *post)
		case "work":
			sqlData.Work = append(sqlData.Work, *post)
		case "project": // Assuming "project" type maps to "projects" table
			sqlData.Projects = append(sqlData.Projects, *post)
		default:
			// If type is not explicitly handled, add to general posts or log
			sqlData.Posts = append(sqlData.Posts, *post)
		}

		// Add to combined all_posts array
		sqlData.AllPosts = append(sqlData.AllPosts, *post)

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking posts directory: %v", err)
	}

	jsonData, err := json.MarshalIndent(sqlData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling SQL data to JSON: %v", err)
	}

	outputDir := filepath.Join(".", config.Blog.OutputDir)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory %s: %v", outputDir, err)
	}

	jsonFilePath := filepath.Join(outputDir, "posts.json")
	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing posts.json: %v", err)
	}

	log.Printf("Generated %s with SQL search data.", jsonFilePath)
}
