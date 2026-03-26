package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"

	models "github.com/Mr-Destructive/meetgor.com/models"
	"github.com/Mr-Destructive/meetgor.com/plugins"
	"gopkg.in/yaml.v3"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func WalkAndListFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return files, err
	}
	return files, nil
}

func ReadFiles(files []string) ([][]byte, error) {
	var filesBytes = [][]byte{}
	for _, file := range files {
		fileBytes, err := os.ReadFile(file)
		if err != nil {
			return filesBytes, err
		}
		filesBytes = append(filesBytes, fileBytes)
	}
	return filesBytes, nil
}

func deriveDescription(content []byte, limit int) string {
	if limit <= 0 {
		limit = 160
	}
	text := string(content)
	replacements := []struct {
		pattern *regexp.Regexp
		replace string
	}{
		{regexp.MustCompile("(?s)```.*?```+"), " "},
		{regexp.MustCompile(`(?s)<[^>]+>`), " "},
		{regexp.MustCompile(`!\[[^\]]*\]\([^)]+\)`), " "},
		{regexp.MustCompile(`\[[^\]]+\]\([^)]+\)`), " "},
		{regexp.MustCompile("`{1,3}[^`]+`{1,3}"), " "},
		{regexp.MustCompile(`(?m)^#{1,6}\s+`), " "},
		{regexp.MustCompile(`(?m)^>\s?`), " "},
		{regexp.MustCompile(`(?m)^[-*+]\s+`), " "},
		{regexp.MustCompile(`(?m)^\d+\.\s+`), " "},
	}
	for _, r := range replacements {
		text = r.pattern.ReplaceAllString(text, r.replace)
	}
	text = strings.Join(strings.Fields(text), " ")
	if len(text) <= limit {
		return text
	}
	return strings.TrimSpace(text[:limit])
}

func normalizeContentLinks(content string, frontmatter *models.FrontMatter) string {
	content = normalizeSeriesLinks(content)
	content = normalizeBareLinks(content)
	if frontmatter != nil && frontmatter.Type == "links" {
		if frontmatter.Extras != nil {
			if rawLink, ok := frontmatter.Extras["link"].(string); ok && rawLink != "" {
				content = normalizeRootRelativeLinks(content, rawLink)
			}
		}
	}
	return content
}

func normalizeSeriesLinks(content string) string {
	re := regexp.MustCompile(`(/series/)([^)\s"'<>]+)`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		series := parts[2]
		if decoded, err := url.PathUnescape(series); err == nil {
			series = decoded
		}
		series = plugins.Slugify(series)
		return parts[1] + series
	})
}

func normalizeBareLinks(content string) string {
	re := regexp.MustCompile(`\]\(([^)]+)\)`)
	updated := re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		urlPart := parts[1]
		if strings.HasPrefix(urlPart, "/login") || strings.HasPrefix(urlPart, "/logout") {
			return "](#)"
		}
		if strings.HasPrefix(urlPart, "/") || strings.HasPrefix(urlPart, "#") {
			return match
		}
		if strings.Contains(urlPart, "://") {
			return match
		}
		if strings.HasPrefix(urlPart, "@") {
			return "](" + "https://twitter.com/" + strings.TrimPrefix(urlPart, "@") + ")"
		}
		if strings.HasPrefix(urlPart, "www.") || strings.Contains(urlPart, ".") {
			return "](" + "https://" + urlPart + ")"
		}
		return match
	})
	updated = normalizeBareHTMLAttrs(updated)
	return updated
}

func normalizeBareHTMLAttrs(content string) string {
	re := regexp.MustCompile(`\b(href|src)=["']([^"']+)["']`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		attr := parts[1]
		urlPart := parts[2]
		if strings.HasPrefix(urlPart, "/login") || strings.HasPrefix(urlPart, "/logout") {
			return attr + "=\"#\""
		}
		if strings.HasPrefix(urlPart, "/") || strings.HasPrefix(urlPart, "#") {
			return match
		}
		if strings.Contains(urlPart, "://") {
			return match
		}
		if strings.Contains(urlPart, "{{") || strings.Contains(urlPart, "{%") {
			return attr + "=\"#\""
		}
		if strings.HasPrefix(urlPart, "@") {
			return attr + "=\"https://twitter.com/" + strings.TrimPrefix(urlPart, "@") + "\""
		}
		if strings.HasPrefix(urlPart, "www.") || strings.Contains(urlPart, ".") {
			return attr + "=\"https://" + urlPart + "\""
		}
		return match
	})
}

func normalizeRootRelativeLinks(content string, rawLink string) string {
	base := rawLink
	if !strings.HasPrefix(base, "http://") && !strings.HasPrefix(base, "https://") {
		base = "https://" + base
	}
	parsed, err := url.Parse(base)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return content
	}
	baseURL := parsed.Scheme + "://" + parsed.Host

	replacements := []struct {
		re *regexp.Regexp
	}{
		{regexp.MustCompile(`\]\((/[^)]+)\)`)},
		{regexp.MustCompile(`href="/([^"]+)"`)},
		{regexp.MustCompile(`src="/([^"]+)"`)},
		{regexp.MustCompile(`\]\(([^)]+)\)`)},
	}
	updated := content
	updated = replacements[0].re.ReplaceAllStringFunc(updated, func(match string) string {
		parts := replacements[0].re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		return "](" + baseURL + parts[1] + ")"
	})
	updated = replacements[1].re.ReplaceAllStringFunc(updated, func(match string) string {
		parts := replacements[1].re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		return `href="` + baseURL + "/" + parts[1] + `"`
	})
	updated = replacements[2].re.ReplaceAllStringFunc(updated, func(match string) string {
		parts := replacements[2].re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		return `src="` + baseURL + "/" + parts[1] + `"`
	})
	updated = replacements[3].re.ReplaceAllStringFunc(updated, func(match string) string {
		parts := replacements[3].re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		urlPart := parts[1]
		if strings.HasPrefix(urlPart, "/") || strings.Contains(urlPart, "://") || strings.HasPrefix(urlPart, "#") {
			return match
		}
		if strings.Contains(urlPart, "{{") || strings.Contains(urlPart, "{%") {
			return "](" + "#" + ")"
		}
		if strings.HasPrefix(urlPart, "www.") || strings.Contains(urlPart, ".") {
			return "](" + "https://" + urlPart + ")"
		}
		return match
	})
	return updated
}

func normalizeFrontmatterImageURL(frontmatter *models.FrontMatter) {
	if frontmatter == nil || frontmatter.ImageUrl == "" {
		return
	}
	img := strings.TrimSpace(frontmatter.ImageUrl)
	if strings.HasPrefix(img, "http://") || strings.HasPrefix(img, "https://") {
		return
	}
	link := ""
	if frontmatter.Extras != nil {
		if rawLink, ok := frontmatter.Extras["link"].(string); ok {
			link = rawLink
		}
	}
	if link == "" {
		return
	}
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}
	parsed, err := url.Parse(link)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return
	}
	baseURL := parsed.Scheme + "://" + parsed.Host
	if strings.HasPrefix(img, "/") {
		frontmatter.ImageUrl = baseURL + img
		return
	}
	frontmatter.ImageUrl = baseURL + "/" + img
}

func escapePathSegment(value string) string {
	escaped := url.PathEscape(value)
	return strings.ReplaceAll(escaped, "+", "%2B")
}

func ReadPosts(files []string) ([]models.Post, error) {
	var posts []models.Post

	// Iterate through files
	for _, filePath := range files {
		if strings.EqualFold(filepath.Base(filePath), "_index.md") {
			continue
		}
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		var success bool
		var frontmatterObj models.FrontMatter
		var contentBytes []byte
		var requiredFields []string = []string{"title", "description", "status", "type", "date", "slug", "tags", "image_url"}

		// Attempt to detect JSON front matter
		jsonSeparator := []byte("}\n\n")
		jsonIndex := strings.Index(string(fileBytes), string(jsonSeparator))

		if jsonIndex != -1 {
			frontmatterBytes := fileBytes[:jsonIndex+1] // Keep closing brace
			contentBytes = fileBytes[jsonIndex+2:]      // Skip the separator

			// Unmarshal into a temporary map to capture extra fields
			tempMap := make(map[string]interface{})
			if err := json.Unmarshal(frontmatterBytes, &tempMap); err == nil {
				success = true

				// Extract known fields into the struct
				if err := json.Unmarshal(frontmatterBytes, &frontmatterObj); err != nil {
					log.Printf("Error parsing JSON front matter: %v", err)
					continue
				}

				// Remove known keys and store the rest in Extras
				for _, key := range requiredFields {
					delete(tempMap, key)
				}
				frontmatterObj.Extras = tempMap
			}
		}

		// Attempt to detect YAML front matter
		if !success {
			frontmatterBytes, content := parseYAMLFrontmatter(fileBytes)
			if frontmatterBytes == nil {
				log.Printf("No valid front matter found in file")
				continue
			}
			contentBytes = content

			// Unmarshal into a temporary map to capture extra fields
			tempMap := make(map[string]interface{})
			if err := yaml.Unmarshal(frontmatterBytes, &tempMap); err == nil {
				// Extract known fields into the struct
				if err := yaml.Unmarshal(frontmatterBytes, &frontmatterObj); err != nil {
					log.Printf("Error parsing YAML front matter: %v", err)
					continue
				}

				// Remove known keys and store the rest in Extras
				for _, key := range requiredFields {
					delete(tempMap, key)
				}
				frontmatterObj.Extras = tempMap
			} else {
				log.Printf("Error parsing YAML front matter: %v", err)
				continue
			}
		}
		if strings.TrimSpace(frontmatterObj.Description) == "" {
			frontmatterObj.Description = deriveDescription(contentBytes, 160)
		}
		normalizeFrontmatterImageURL(&frontmatterObj)
		contentBytes = []byte(normalizeContentLinks(string(contentBytes), &frontmatterObj))
		// Convert Markdown to HTML
		var contentBuffer bytes.Buffer
		mdOptions := []goldmark.Option{
			goldmark.WithExtensions(
				&plugins.SQLPlaygroundExtender{},
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
				parser.WithAttribute(),
			),
		}
		if frontmatterObj.Type == "newsletter" {
			mdOptions = append(mdOptions, goldmark.WithRendererOptions(html.WithUnsafe()))
		}
		md := goldmark.New(mdOptions...)
		// Create a new parser context
		ctx := parser.NewContext()
		if err := md.Convert(contentBytes, &contentBuffer, parser.WithContext(ctx)); err != nil {
			log.Printf("Error processing Markdown: %v", err)
			continue
		}
		frontmatterObj.Type = plugins.NormalizePostType(frontmatterObj.Type)
		// Append post
		posts = append(posts, models.Post{
			Frontmatter: frontmatterObj,
			Content:     template.HTML(contentBuffer.String()),
			Markdown:    string(contentBytes),
			SourcePath:  filePath,
		})
	}

	return posts, nil
}

func parseYAMLFrontmatter(fileBytes []byte) ([]byte, []byte) {
	lines := strings.Split(string(fileBytes), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return nil, nil
	}
	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return nil, nil
	}
	frontmatterContent := strings.Join(lines[1:end], "\n")
	content := strings.Join(lines[end+1:], "\n")
	return []byte(frontmatterContent), []byte(content)
}

func ReadTemplates(files []string) ([]string, error) {
	templateStrs := []string{}
	filesBytes, err := ReadFiles(files)
	if err != nil {
		return templateStrs, err
	}
	for _, fileBytes := range filesBytes {
		templateStrs = append(templateStrs, string(fileBytes))
	}
	return templateStrs, nil
}

func Copy(src string, dst string) error {
	srcFiles, err := WalkAndListFiles(src)
	if err != nil {
		return err
	}
	for _, srcFileName := range srcFiles {
		relPath, err := filepath.Rel(src, srcFileName)
		if err != nil {
			return err
		}

		srcFile, err := os.Open(srcFileName)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstPath := filepath.Join(dst, relPath)
		err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		if err != nil {
			return err
		}
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyCustom(src string, dst string, skipDirs []string) error {
	srcFiles, err := WalkAndListFiles(src)
	if err != nil {
		return err
	}
	fmt.Println("srcFiles:", len(srcFiles))

	for _, srcFileName := range srcFiles {
		relPath, err := filepath.Rel(src, srcFileName)
		if err != nil {
			return err
		}

		// Skip copying anything inside "admin"
		if strings.HasPrefix(relPath, "admin") {
			continue
		}
		skip := false
		for _, dir := range skipDirs {
			if dir == "" {
				continue
			}
			if strings.HasPrefix(relPath, dir+string(os.PathSeparator)) || relPath == dir {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		srcFile, err := os.Open(srcFileName)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstPath := filepath.Join(dst, relPath)
		err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		if err != nil {
			return err
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func GeneratePages(config models.SSG_CONFIG) error {
	src := config.Blog.StaticDir
	templateFS := os.DirFS(config.Blog.TemplatesDir)
	mdFiles := []string{}
	filterMDFile := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			mdFiles = append(mdFiles, path)
		}
		return nil
	}
	err := filepath.Walk(src, filterMDFile)
	if err != nil {
		return err
	}
	for _, mdFile := range mdFiles {
		mdFileName := filepath.Base(mdFile)
		content, err := ReadPosts([]string{string(mdFile)})
		if err != nil {
			return err
		}
		if len(content) == 0 {
			return fmt.Errorf("no content found in file: %s", mdFile)
		}
		post := content[0]

		mdFileName = strings.TrimSuffix(mdFileName, filepath.Ext(mdFileName))
		feed := models.Feed{
			Title: post.Frontmatter.Title,
			Type:  mdFileName,
			Slug:  mdFileName,
			Posts: []models.Post{post},
		}
		context := models.TemplateContext{
			FeedPosts: []models.Feed{feed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: feed,
			Post:     post,
			Config: models.SSG_CONFIG{
				Blog: config.Blog,
			},
		}
		buffer := bytes.Buffer{}
		t := template.New("").Funcs(template.FuncMap{
			"dateOnly": func(dateStr string) string {
				if len(dateStr) >= 10 {
					return dateStr[:10]
				}
				return dateStr
			},
			"slugify": func(value string) string {
				return plugins.Slugify(value)
			},
			"pathEscape": func(value string) template.URL {
				return template.URL(escapePathSegment(value))
			},
		})
		t, err = t.ParseFS(templateFS, "*.html")
		if err != nil {
			return err
		}
		err = t.ExecuteTemplate(&buffer, "default_page_template.html", context)
		if err != nil {
			return err
		}
		//create a folder with mdFileName
		err = os.MkdirAll(filepath.Join(".", config.Blog.OutputDir, mdFileName), os.ModePerm)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(".", config.Blog.OutputDir, mdFileName, "index.html"), buffer.Bytes(), 0660)
		if err != nil {
			return err
		}
	}
	return nil
}

type PluginManager struct {
	Plugins []plugins.Plugin
}

func (p *PluginManager) Register(plugin plugins.Plugin) {
	p.Plugins = append(p.Plugins, plugin)
}

func (p *PluginManager) ExecuteAll(ssg *models.SSG) error {
	fmt.Println("Running plugins")
	for _, plugin := range p.Plugins {
		fmt.Println("Running plugin:", plugin.Name())
		if err := plugin.Execute(ssg); err != nil {
			return fmt.Errorf("plugin %s: %w", plugin.Name(), err)
		}
	}
	return nil
}

type PostReaderPlugin struct {
	PluginName string
}

func (c *PostReaderPlugin) Name() string {
	return c.PluginName
}

func (c *PostReaderPlugin) Phase() plugins.Phase {
	return plugins.PhaseRead
}

func (c *PostReaderPlugin) Requires() []string {
	return nil
}

func (c *PostReaderPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func (c *PostReaderPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	postFolder := config.Blog.PostsDir
	var postFiles []string
	postFiles, err := WalkAndListFiles(postFolder)
	if err != nil {
		return err
	}
	postsList, err := ReadPosts(postFiles)
	if err != nil {
		return err
	}
	for i := range postsList {
		if postsList[i].Frontmatter.Status == "draft" {
			continue
		}
		plugins.CleanPostFrontmatter(&postsList[i], ssg)
	}

	ssg.Posts = dedupePosts(postsList, config.Blog.PrefixURL)
	fmt.Println("Posts:", len(ssg.Posts))
	return nil
}

func dedupePosts(posts []models.Post, prefixURL string) []models.Post {
	byKey := map[string]models.Post{}
	for _, post := range posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		key := normalizeSlugKey(post.Frontmatter.Slug, prefixURL)
		if key == "" {
			continue
		}
		existing, ok := byKey[key]
		if !ok || isPreferredPost(post, existing) {
			byKey[key] = post
		}
	}

	deduped := make([]models.Post, 0, len(byKey))
	for _, post := range byKey {
		deduped = append(deduped, post)
	}
	return deduped
}

func normalizeSlugKey(slug string, prefixURL string) string {
	if slug == "" {
		return ""
	}
	if prefixURL != "" {
		slug = strings.TrimPrefix(slug, prefixURL)
	}
	slug = strings.TrimPrefix(slug, "/")
	return slug
}

func isPreferredPost(candidate models.Post, existing models.Post) bool {
	cScore := canonicalScore(candidate)
	eScore := canonicalScore(existing)
	if cScore != eScore {
		return cScore > eScore
	}
	cDate, cOk := parseDate(candidate.Frontmatter.Date)
	eDate, eOk := parseDate(existing.Frontmatter.Date)
	if cOk && eOk {
		if cDate.After(eDate) {
			return true
		}
		if eDate.After(cDate) {
			return false
		}
	} else if cOk && !eOk {
		return true
	} else if !cOk && eOk {
		return false
	}
	return len(candidate.SourcePath) < len(existing.SourcePath)
}

func canonicalScore(post models.Post) int {
	path := filepath.ToSlash(post.SourcePath)
	postType := plugins.NormalizePostType(post.Frontmatter.Type)
	canonicalDir := "/posts/" + postType + "/"
	if strings.Contains(path, canonicalDir) {
		return 2
	}
	if strings.Contains(path, "/posts/") {
		return 1
	}
	return 0
}

func parseDate(dateStr string) (time.Time, bool) {
	if len(dateStr) < 10 {
		return time.Time{}, false
	}
	parsed, err := time.Parse("2006-01-02", dateStr[:10])
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

type RenderTemplatesPlugin struct {
	PluginName string
}

func (c *RenderTemplatesPlugin) Name() string {
	return c.PluginName
}

func (c *RenderTemplatesPlugin) Phase() plugins.Phase {
	return plugins.PhaseRender
}

func (c *RenderTemplatesPlugin) Requires() []string {
	return []string{"readPosts"}
}

func (c *RenderTemplatesPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func SortPosts(posts []models.Post) []models.Post {
	sort.Slice(posts, func(i, j int) bool {
		if posts[i].Frontmatter.Date == "" {
			return false
		}
		if posts[j].Frontmatter.Date == "" {
			return true
		}
		date1, err1 := time.Parse("2006-01-02", posts[i].Frontmatter.Date[:10])
		if err1 != nil {
			return false
		}
		date2, err2 := time.Parse("2006-01-02", posts[j].Frontmatter.Date[:10])
		if err2 != nil {
			return false
		}
		return date1.After(date2)
	})
	return posts
}

type oEmbedResponse struct {
	HTML string `json:"html"`
}

var twitterOEmbedBaseURL = "https://publish.twitter.com/oembed?url="
var httpClient = &http.Client{Timeout: 5 * time.Second}
var httpGet = func(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "meetgor-ssg/1.0")
	return httpClient.Do(req)
}
var listenAndServe = http.ListenAndServe
var fatalf = log.Fatal

func getTwitterEmbed(url string) string {
	apiURL := twitterOEmbedBaseURL + url
	resp, err := httpGet(apiURL)
	if err != nil {
		fmt.Println("Error fetching Twitter embed:", err)
		return fmt.Sprintf(`<div class="embed-container"><a href="%s" target="_blank">%s</a></div>`, url, url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching Twitter embed:", resp.Status)
		return fmt.Sprintf(`<div class="embed-container"><a href="%s" target="_blank">%s</a></div>`, url, url)
	}

	body, _ := io.ReadAll(resp.Body)
	var oembed oEmbedResponse
	if err := json.Unmarshal(body, &oembed); err != nil {
		fmt.Println("Error parsing Twitter embed response:", err)
		return fmt.Sprintf(`<div class="embed-container"><a href="%s" target="_blank">%s</a></div>`, url, url)
	}

	return fmt.Sprintf(`<div class="embed-container">%s</div>`, oembed.HTML)
}

func generateEmbed(postContent string) string {
	re := regexp.MustCompile(`{%\s*embed\s*(https?://[^\s]+)\s*%}`)
	replacedContent := re.ReplaceAllStringFunc(postContent, func(match string) string {
		url := re.FindStringSubmatch(match)[1]

		if strings.Contains(url, "x.com") || strings.Contains(url, "twitter.com") {
			return getTwitterEmbed(url)
		}
		if strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be") {
			videoID := extractYouTubeID(url)
			return fmt.Sprintf(`<iframe width="560" height="315" src="https://www.youtube.com/embed/%s" frameborder="0" allowfullscreen></iframe>`, videoID)
		}
		return fmt.Sprintf(`<iframe src="%s" width="600" height="400" frameborder="0"></iframe>`, url)
	})
	return replacedContent
}

func extractYouTubeID(url string) string {
	if strings.Contains(url, "youtube.com/watch?v=") {
		return strings.Split(strings.TrimPrefix(url, "https://www.youtube.com/watch?v="), "&")[0]
	}
	if strings.Contains(url, "youtu.be/") {
		return strings.Split(strings.TrimPrefix(url, "https://youtu.be/"), "?")[0]
	}
	return ""
}

func (c *RenderTemplatesPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	templateFS := os.DirFS(config.Blog.TemplatesDir)
	ssg.FS = templateFS

	// Create template with custom functions first
	t := template.New("").Funcs(template.FuncMap{
		"dateOnly": func(dateStr string) string {
			// Extract just the date part (YYYY-MM-DD) from the beginning
			if len(dateStr) >= 10 {
				return dateStr[:10]
			}
			return dateStr
		},
		"slugify": func(value string) string {
			return plugins.Slugify(value)
		},
		"pathEscape": func(value string) template.URL {
			return template.URL(escapePathSegment(value))
		},
	})

	t, err := t.ParseFS(templateFS, "*.html")

	ssg.TemplateFS = t
	if err != nil {
		return err
	}
	var prefixURL string = ""
	if config.Blog.PrefixURL != "" {
		prefixURL = config.Blog.PrefixURL
	}

	// render the templates with the content
	outputPath := filepath.Join(".", config.Blog.OutputDir)
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return err
	}
	feedPosts := make(map[string][]models.Post)
	postsCopy := slices.Clone(ssg.Posts)
	SortPosts(postsCopy)
	ssg.Posts = postsCopy

	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		if post.Frontmatter.Date == "" {
			continue
		}
		postType := plugins.NormalizePostType(post.Frontmatter.Type)
		templatePath := config.Blog.PagesConfig[postType].TemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultPostTemplate
		}
		buffer := bytes.Buffer{}
		postSlug := post.Frontmatter.Slug
		if postSlug == "" {
			postSlug = plugins.Slugify(post.Frontmatter.Title)
		}
		if strings.Contains(postSlug, "/") {
			slugPath := postSlug
			if prefixURL != "" {
				slugPath = strings.TrimPrefix(slugPath, prefixURL)
			}
			slugPath = strings.TrimPrefix(slugPath, "/")
			typePrefix := postType + "/"
			if strings.HasPrefix(slugPath, typePrefix) {
				postSlug = strings.TrimPrefix(slugPath, typePrefix)
			} else {
				postSlug = slugPath
			}
		}
		post.Frontmatter.Type = postType
		post.Frontmatter.Slug = prefixURL + postType + "/" + postSlug
		postPath := filepath.Join(outputPath, postType, postSlug)
		//outputDirPath := filepath.Join(postPath, postSlug)
		err = os.MkdirAll(postPath, os.ModePerm)
		if err != nil {
			return err
		}
		outputPostPath := filepath.Join(postPath, "index.html")
		post.Content = template.HTML(generateEmbed(string(post.Content)))
		context := models.TemplateContext{
			Post: post,
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			Config: models.SSG_CONFIG{
				Blog:      config.Blog,
				AdminMode: config.AdminMode,
			},
			Years: plugins.YearsFromPosts(ssg.Posts),
		}
		err := t.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			return err
		}
		err = os.WriteFile(outputPostPath, buffer.Bytes(), 0660)
		if err != nil {
			return err
		}
		feedPosts[postType] = append(feedPosts[postType], post)
	}
	// we no longer log individual feed counts
	feedPostLists := []models.Feed{}
	for postType, posts := range feedPosts {
		postsCopy = slices.Clone(posts)
		SortPosts(postsCopy)
		posts = postsCopy
		feed := models.Feed{
			Title: strings.ToTitle(postType),
			Type:  postType,
			Slug:  prefixURL + postType,
			Posts: posts,
		}
		fmt.Println("Feed", feed.Title, feed.Slug)
		feedPostLists = append(feedPostLists, feed)
	}

	// Render a dedicated archive page for each feed type.
	// This keeps /posts, /til, /newsletter, etc. aligned with their type feeds.
	for _, feed := range feedPostLists {
		templatePath := config.Blog.PagesConfig[feed.Type].FeedTemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}

		buffer := bytes.Buffer{}
		templateFS := os.DirFS(config.Blog.TemplatesDir)
		t := template.New("").Funcs(template.FuncMap{
			"dateOnly": func(dateStr string) string {
				if len(dateStr) >= 10 {
					return dateStr[:10]
				}
				return dateStr
			},
			"slugify": func(value string) string {
				return plugins.Slugify(value)
			},
			"pathEscape": func(value string) template.URL {
				return template.URL(escapePathSegment(value))
			},
		})
		t, err = t.ParseFS(templateFS, "*.html")
		if err != nil {
			return err
		}

		context := models.TemplateContext{
			FeedPosts: []models.Feed{feed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: feed,
			Config: models.SSG_CONFIG{
				Blog:      config.Blog,
				AdminMode: config.AdminMode,
			},
			Years: plugins.YearsFromPosts(ssg.Posts),
		}
		if err := t.ExecuteTemplate(&buffer, templatePath, context); err != nil {
			return err
		}

		feedPath := filepath.Join(outputPath, feed.Type)
		if err := os.MkdirAll(feedPath, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(feedPath, "index.html"), buffer.Bytes(), 0660); err != nil {
			return err
		}
	}
	ssg.FeedPosts = feedPostLists
	return nil
}

// "createFeeds",
type CreateFeedsPlugin struct {
	PluginName string
}

func (c *CreateFeedsPlugin) Name() string {
	return c.PluginName
}

func (c *CreateFeedsPlugin) Phase() plugins.Phase {
	return plugins.PhaseRender
}

func (c *CreateFeedsPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (c *CreateFeedsPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func (c *CreateFeedsPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config

	// Generate individual feed pages
	for _, feed := range ssg.FeedPosts {
		buffer := bytes.Buffer{}
		templatePath := config.Blog.PagesConfig[feed.Type].FeedTemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}

		context := models.TemplateContext{
			FeedPosts: []models.Feed{feed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: feed,
			Config: models.SSG_CONFIG{
				Blog:      config.Blog,
				AdminMode: config.AdminMode,
			},
			Years: plugins.YearsFromPosts(ssg.Posts),
		}
		fmt.Println("Post:", feed.Title, len(feed.Posts))
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			return err
		}
		feedPath := filepath.Join(".", config.Blog.OutputDir, feed.Type)
		err = os.MkdirAll(feedPath, os.ModePerm)
		if err != nil {
			return err
		}
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
		if err != nil {
			return err
		}
	}

	// Generate all-content feed with all posts
	allPosts := []models.Post{}
	for _, feed := range ssg.FeedPosts {
		allPosts = append(allPosts, feed.Posts...)
	}
	// Sort all posts by date
	SortPosts(allPosts)

	if len(allPosts) > 0 {
		allContentFeed := models.Feed{
			Title: "All Content",
			Type:  "all-content",
			Slug:  "all-content",
			Posts: allPosts,
		}

		buffer := bytes.Buffer{}
		templatePath := config.Blog.PagesConfig["all-content"].FeedTemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}

		context := models.TemplateContext{
			FeedPosts: []models.Feed{allContentFeed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: allContentFeed,
			Config: models.SSG_CONFIG{
				Blog:      config.Blog,
				AdminMode: config.AdminMode,
			},
			Years: plugins.YearsFromPosts(ssg.Posts),
		}

		fmt.Println("Creating all-content feed with", len(allPosts), "posts")
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			return err
		}

		feedPath := filepath.Join(".", config.Blog.OutputDir, "all-content")
		err = os.MkdirAll(feedPath, os.ModePerm)
		if err != nil {
			return err
		}
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
		if err != nil {
			return err
		}
	}

	// create a folder for post if the post type is "post"
	// as this should be the /posts/<slug> as well as /<slug>
	for _, post := range ssg.Posts {
		if plugins.NormalizePostType(post.Frontmatter.Type) == "posts" {
			postPath := filepath.Join(".", config.Blog.OutputDir, post.Frontmatter.Slug)
			err := os.MkdirAll(postPath, os.ModePerm)
			if err != nil {
				return err
			}
			buffer := bytes.Buffer{}
			context := models.TemplateContext{
				Post: post,
				Themes: models.ThemeCombo{
					Default:   config.Blog.Themes["default"],
					Secondary: config.Blog.Themes["secondary"],
				},
				Config: models.SSG_CONFIG{
					Blog:      config.Blog,
					AdminMode: config.AdminMode,
				},
				Years: plugins.YearsFromPosts(ssg.Posts),
			}
			err = ssg.TemplateFS.ExecuteTemplate(&buffer, config.Blog.DefaultPostTemplate, context)
			if err != nil {
				return err
			}
			outputPostPath := fmt.Sprintf("%s/index.html", postPath)
			err = os.WriteFile(outputPostPath, buffer.Bytes(), 0660)
			if err != nil {
				return err
			}
		}
	}

	// Re-render the archive pages after post detail pages so a bad slug cannot
	// overwrite the type archive path (for example /posts/).
	for _, feed := range ssg.FeedPosts {
		templatePath := config.Blog.PagesConfig[feed.Type].FeedTemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}

		buffer := bytes.Buffer{}
		templateFS := os.DirFS(config.Blog.TemplatesDir)
		t := template.New("").Funcs(template.FuncMap{
			"dateOnly": func(dateStr string) string {
				if len(dateStr) >= 10 {
					return dateStr[:10]
				}
				return dateStr
			},
			"slugify": func(value string) string {
				return plugins.Slugify(value)
			},
			"pathEscape": func(value string) template.URL {
				return template.URL(escapePathSegment(value))
			},
		})
		t, err := t.ParseFS(templateFS, "*.html")
		if err != nil {
			return err
		}

		context := models.TemplateContext{
			FeedPosts: []models.Feed{feed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: feed,
			Config: models.SSG_CONFIG{
				Blog:      config.Blog,
				AdminMode: config.AdminMode,
			},
			Years: plugins.YearsFromPosts(ssg.Posts),
		}
		if err := t.ExecuteTemplate(&buffer, templatePath, context); err != nil {
			return err
		}

		feedPath := filepath.Join(".", config.Blog.OutputDir, feed.Type)
		if err := os.MkdirAll(feedPath, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(feedPath, "index.html"), buffer.Bytes(), 0660); err != nil {
			return err
		}
	}
	return nil
}

// "createFeeds",
// "copyStaticFiles",
type CopyStaticFilesPlugin struct {
	PluginName string
}

func (c *CopyStaticFilesPlugin) Name() string {
	return c.PluginName
}

func (c *CopyStaticFilesPlugin) Phase() plugins.Phase {
	return plugins.PhasePostProcess
}

func (c *CopyStaticFilesPlugin) Requires() []string {
	return nil
}

func (c *CopyStaticFilesPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func (c *CopyStaticFilesPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	err := Copy(config.Blog.StaticDir, config.Blog.OutputDir)
	if err != nil {
		return err
	}
	err = GeneratePages(*config)
	if err != nil {
		return err
	}
	return nil
}

type IndexPlugin struct {
	PluginName string
}

func (c *IndexPlugin) Name() string {
	return c.PluginName
}

func (c *IndexPlugin) Phase() plugins.Phase {
	return plugins.PhasePostProcess
}

func (c *IndexPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (c *IndexPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func (c *IndexPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config

	buffer := bytes.Buffer{}
	templateFS := os.DirFS(config.Blog.StaticDir)
	t, err := template.ParseFS(templateFS, "index.html")
	if err != nil {
		return err
	}
	context := models.TemplateContext{
		Themes: models.ThemeCombo{
			Default:   config.Blog.Themes["default"],
			Secondary: config.Blog.Themes["secondary"],
		},
		Config: models.SSG_CONFIG{
			Blog:      config.Blog,
			AdminMode: config.AdminMode,
		},
		FeedPosts: ssg.FeedPosts,
		Years:     plugins.YearsFromPosts(ssg.Posts),
	}
	err = t.ExecuteTemplate(&buffer, "index.html", context)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(".", config.Blog.OutputDir, "index.html"), buffer.Bytes(), 0660)
	if err != nil {
		return err
	}
	return nil
}

type CopyDirsPlugin struct {
	PluginName string
}

func (c *CopyDirsPlugin) Name() string {
	return c.PluginName
}

func (c *CopyDirsPlugin) Phase() plugins.Phase {
	return plugins.PhasePostProcess
}

func (c *CopyDirsPlugin) Requires() []string {
	return []string{"index"}
}

func (c *CopyDirsPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminRun
}

func (c *CopyDirsPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	for _, dest := range config.Blog.DuplicateOutputTo {
		destPath := filepath.Join(".", config.Blog.OutputDir, dest)
		err := CopyCustom(config.Blog.OutputDir, destPath, config.Blog.DuplicateOutputTo)
		if err != nil {
			return err
		}
	}
	return nil
}

// "server"
type ServerPlugin struct {
	PluginName string
}

func (c *ServerPlugin) Name() string {
	return c.PluginName
}

func (c *ServerPlugin) Phase() plugins.Phase {
	return plugins.PhaseServe
}

func (c *ServerPlugin) Requires() []string {
	return nil
}

func (c *ServerPlugin) AdminPolicy() plugins.AdminPolicy {
	return plugins.AdminOnly
}

func (c *ServerPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	http.Handle("/", http.FileServer(http.Dir(config.Blog.OutputDir)))
	fmt.Println("Listening on port 3030")
	if err := listenAndServe(":3030", nil); err != nil {
		return err
	}
	return nil
}

func main() {

	args := os.Args
	devEnv := false
	if len(args) > 1 {
		if args[1] == "dev" {
			devEnv = true
		}
	}
	// read the config
	// read all the plugins from config file
	ssg := models.SSG{}
	configbytes, err := os.ReadFile(models.SSG_CONFIG_FILE_NAME)
	if err != nil {
		fatalf(err)
	}
	var config models.SSG_CONFIG
	err = json.Unmarshal(configbytes, &config)
	if err != nil {
		fatalf(err)
	}
	ssg.Config = config
	if devEnv {
		ssg.Config.Blog.PrefixURL = ""
	}

	pluginsList, err := buildPlugins(config)
	if err != nil {
		fatalf(err)
	}

	normalPlugins := filterPluginsForMode(pluginsList, RunModeNormal)
	if err := validatePluginOrder(normalPlugins); err != nil {
		fatalf(err)
	}
	if err := (&PluginManager{Plugins: normalPlugins}).ExecuteAll(&ssg); err != nil {
		fatalf(err)
	}

	originalOutputDir := ssg.Config.Blog.OutputDir
	ssg.Config.AdminMode = true
	ssg.Config.Blog.OutputDir = path.Join(ssg.Config.Blog.OutputDir, ssg.Config.Blog.AdminDir)
	adminPlugins := filterPluginsForMode(pluginsList, RunModeAdmin)
	if err := validatePluginOrder(adminPlugins); err != nil {
		fatalf(err)
	}
	if err := (&PluginManager{Plugins: adminPlugins}).ExecuteAll(&ssg); err != nil {
		fatalf(err)
	}
	ssg.Config.Blog.OutputDir = originalOutputDir
	ssg.Config.AdminMode = false

	if devEnv {
		servePlugins := filterPluginsForMode(pluginsList, RunModeServe)
		if err := validatePluginOrder(servePlugins); err != nil {
			fatalf(err)
		}
		if err := (&PluginManager{Plugins: servePlugins}).ExecuteAll(&ssg); err != nil {
			fatalf(err)
		}
	}
}

func LoadPlugin(pluginName string) (plugins.Plugin, error) {
	pluginInstance, exists := plugins.NewPlugin(pluginName)
	if !exists {
		return nil, fmt.Errorf("plugin %s not found", pluginName)
	}
	return pluginInstance, nil
}

var corePluginFactories = map[string]func() plugins.Plugin{
	"readPosts":       func() plugins.Plugin { return &PostReaderPlugin{PluginName: "readPosts"} },
	"renderTemplates": func() plugins.Plugin { return &RenderTemplatesPlugin{PluginName: "renderTemplates"} },
	"createFeeds":     func() plugins.Plugin { return &CreateFeedsPlugin{PluginName: "createFeeds"} },
	"copyStaticFiles": func() plugins.Plugin { return &CopyStaticFilesPlugin{PluginName: "copyStaticFiles"} },
	"index":           func() plugins.Plugin { return &IndexPlugin{PluginName: "index"} },
	"copyDirs":        func() plugins.Plugin { return &CopyDirsPlugin{PluginName: "copyDirs"} },
	"server":          func() plugins.Plugin { return &ServerPlugin{PluginName: "server"} },
}

type RunMode int

const (
	RunModeNormal RunMode = iota
	RunModeAdmin
	RunModeServe
)

func resolvePlugin(name string) (plugins.Plugin, error) {
	if factory, ok := corePluginFactories[name]; ok {
		return factory(), nil
	}
	return LoadPlugin(name)
}

func buildPlugins(config models.SSG_CONFIG) ([]plugins.Plugin, error) {
	pluginsList := make([]plugins.Plugin, 0, len(config.Plugins))
	for _, name := range config.Plugins {
		plugin, err := resolvePlugin(name)
		if err != nil {
			return nil, err
		}
		pluginsList = append(pluginsList, plugin)
	}
	return pluginsList, nil
}

func pluginPhase(plugin plugins.Plugin) plugins.Phase {
	if meta, ok := plugin.(plugins.PluginMetadata); ok {
		return meta.Phase()
	}
	return plugins.PhaseOther
}

func pluginRequires(plugin plugins.Plugin) []string {
	if meta, ok := plugin.(plugins.PluginMetadata); ok {
		return meta.Requires()
	}
	return nil
}

func pluginAdminPolicy(plugin plugins.Plugin) plugins.AdminPolicy {
	if meta, ok := plugin.(plugins.PluginMetadata); ok {
		return meta.AdminPolicy()
	}
	return plugins.AdminRun
}

func filterPluginsForMode(all []plugins.Plugin, mode RunMode) []plugins.Plugin {
	filtered := make([]plugins.Plugin, 0, len(all))
	for _, plugin := range all {
		phase := pluginPhase(plugin)
		if mode == RunModeServe {
			if phase != plugins.PhaseServe {
				continue
			}
			filtered = append(filtered, plugin)
			continue
		}
		if phase == plugins.PhaseServe {
			continue
		}
		policy := pluginAdminPolicy(plugin)
		if mode == RunModeAdmin && policy == plugins.AdminSkip {
			continue
		}
		if mode == RunModeNormal && policy == plugins.AdminOnly {
			continue
		}
		filtered = append(filtered, plugin)
	}
	return filtered
}

func phaseWeight(phase plugins.Phase) int {
	switch phase {
	case plugins.PhaseRead:
		return 10
	case plugins.PhaseTransform:
		return 20
	case plugins.PhaseRender:
		return 30
	case plugins.PhasePostProcess:
		return 40
	case plugins.PhaseServe:
		return 50
	default:
		return 60
	}
}

func validatePluginOrder(pluginsList []plugins.Plugin) error {
	seen := make(map[string]int, len(pluginsList))
	lastPhaseWeight := -1

	for i, plugin := range pluginsList {
		name := plugin.Name()
		if _, ok := seen[name]; ok {
			return fmt.Errorf("duplicate plugin name: %s", name)
		}
		seen[name] = i

		weight := phaseWeight(pluginPhase(plugin))
		if weight < lastPhaseWeight {
			return fmt.Errorf("plugin %s is out of phase order", name)
		}
		lastPhaseWeight = weight
	}

	for i, plugin := range pluginsList {
		for _, req := range pluginRequires(plugin) {
			reqIndex, ok := seen[req]
			if !ok {
				return fmt.Errorf("plugin %s requires missing plugin %s", plugin.Name(), req)
			}
			if reqIndex >= i {
				return fmt.Errorf("plugin %s requires %s to run before it", plugin.Name(), req)
			}
		}
	}
	return nil
}
