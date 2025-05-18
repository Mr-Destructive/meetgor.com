package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"

	models "github.com/mr-destructive/mr-destructive.github.io/models"
	"github.com/mr-destructive/mr-destructive.github.io/plugins"
	"gopkg.in/yaml.v3"

	"github.com/yuin/goldmark"
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

func ReadPosts(files []string) ([]models.Post, error) {
	var posts []models.Post

	// Read file contents
	filesBytes, err := ReadFiles(files)
	if err != nil {
		return nil, err
	}

	// Iterate through files
	for _, fileBytes := range filesBytes {
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
			yamlSeparator := []byte("---\n\n")
			yamlIndex := strings.Index(string(fileBytes), string(yamlSeparator))

			if yamlIndex != -1 {
				frontmatterBytes := fileBytes[:yamlIndex]
				contentBytes = fileBytes[yamlIndex+len(yamlSeparator):]

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
			} else {
				log.Printf("No valid front matter found in file", frontmatterObj)
				log.Println(string(fileBytes))
				continue
			}
		}
		// Convert Markdown to HTML
		var contentBuffer bytes.Buffer
		if err := goldmark.Convert(contentBytes, &contentBuffer); err != nil {
			log.Printf("Error processing Markdown: %v", err)
			continue
		}

		// Append post
		posts = append(posts, models.Post{
			Frontmatter: frontmatterObj,
			Content:     template.HTML(contentBuffer.String()),
			Markdown:    string(contentBytes),
		})
	}

	return posts, nil
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
		srcFile, err := os.Open(srcFileName)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		srcFileName = filepath.Base(srcFileName)
		dstPath := filepath.Join(dst, srcFileName)
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func CopyCustom(src string, dst string) error {
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
			log.Fatal(err)
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
			log.Fatal(err)
		}
		postContent := string(content[0].Content)
		if err != nil {
			log.Fatal(err)
		}

		mdFileName = strings.TrimSuffix(mdFileName, filepath.Ext(mdFileName))
		feed := models.Feed{
			Title: strings.ToTitle(mdFileName),
			Type:  mdFileName,
			Slug:  mdFileName,
			Posts: []models.Post{
				{
					Content: template.HTML(string(postContent)),
					Frontmatter: models.FrontMatter{
						Title: mdFileName,
						Slug:  mdFileName,
					},
				},
			},
		}
		context := models.TemplateContext{
			FeedPosts: []models.Feed{feed},
			Themes: models.ThemeCombo{
				Default:   config.Blog.Themes["default"],
				Secondary: config.Blog.Themes["secondary"],
			},
			FeedInfo: feed,
			Config: models.SSG_CONFIG{
				Blog: config.Blog,
			},
		}
		buffer := bytes.Buffer{}
		t, err := template.ParseFS(templateFS, "default_page_template.html")
		err = t.ExecuteTemplate(&buffer, "default_page_template.html", context)
		if err != nil {
			log.Fatal(err)
		}
		//create a folder with mdFileName
		err = os.MkdirAll(filepath.Join(".", config.Blog.OutputDir, mdFileName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(filepath.Join(".", config.Blog.OutputDir, mdFileName, "index.html"), buffer.Bytes(), 0660)
		if err != nil {
			log.Fatal(err)
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

func (p *PluginManager) ExecuteAll(ssg *models.SSG) {
	fmt.Println("Running plugins")
	for _, plugin := range p.Plugins {
		fmt.Println("Running plugin:", plugin.Name())
		plugin.Execute(ssg)
	}
}

// Config Plugin
type ConfigPlugin struct {
	PluginName string
}

func (c *ConfigPlugin) Name() string {
	return c.PluginName
}

type PostReaderPlugin struct {
	PluginName string
}

func (c *PostReaderPlugin) Name() string {
	return c.PluginName
}

func (c *PostReaderPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	postFolder := config.Blog.PostsDir
	var postFiles []string
	postFiles, err := WalkAndListFiles(postFolder)
	if err != nil {
		log.Fatal(err)
	}
	postsList, err := ReadPosts(postFiles)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range postsList {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		plugins.CleanPostFrontmatter(&post, ssg)
	}
	ssg.Posts = postsList
	fmt.Println("Posts:", len(ssg.Posts))
}

type RenderTemplatesPlugin struct {
	PluginName string
}

func (c *RenderTemplatesPlugin) Name() string {
	return c.PluginName
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

func getTwitterEmbed(url string) string {
	apiURL := "https://publish.twitter.com/oembed?url=" + url
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching Twitter embed:", err)
		return fmt.Sprintf(`<div class="embed-container"><a href="%s" target="_blank">%s</a></div>`, url, url)
	}
	defer resp.Body.Close()

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

func (c *RenderTemplatesPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	templateFS := os.DirFS(config.Blog.TemplatesDir)
	ssg.FS = templateFS
	t, err := template.ParseFS(templateFS, "*.html")
	ssg.TemplateFS = t
	if err != nil {
		log.Fatal(err)
	}
	var prefixURL string = ""
	if config.Blog.PrefixURL != "" {
		prefixURL = config.Blog.PrefixURL
	}

	// render the templates with the content
	outputPath := filepath.Join(".", config.Blog.OutputDir)
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	feedPosts := make(map[string][]models.Post)
	postsCopy := slices.Clone(ssg.Posts)
	SortPosts(postsCopy)
	ssg.Posts = postsCopy

	for _, post := range ssg.Posts {
		fmt.Println("Post:", post.Frontmatter.Title, post.Frontmatter.Type)
		if post.Frontmatter.Status == "draft" {
			continue
		}
		if post.Frontmatter.Date == "" {
			continue
		}
		postType := post.Frontmatter.Type
		if postType == "" {
			postType = "posts"
		}
		templatePath := config.Blog.PagesConfig[postType].TemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultPostTemplate
		}
		buffer := bytes.Buffer{}
		postSlug := post.Frontmatter.Slug
		if postSlug == "" {
			postSlug = plugins.Slugify(post.Frontmatter.Title)
		}
		post.Frontmatter.Slug = prefixURL + postType + "/" + postSlug
		postPath := filepath.Join(outputPath, postType, postSlug)
		//outputDirPath := filepath.Join(postPath, postSlug)
		err = os.MkdirAll(postPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		outputPostPath := filepath.Join(postPath, "index.html")
		for i := range ssg.Posts {
			if ssg.Posts[i].Frontmatter.Date == "" {
				continue
			}
			ssg.Posts[i].Frontmatter.Date = ssg.Posts[i].Frontmatter.Date[:10] // Truncate in case of time component
		}
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
		}
		err := t.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(outputPostPath, buffer.Bytes(), 0660)
		feedPosts[postType] = append(feedPosts[postType], post)
	}
	for postType, posts := range feedPosts {
		fmt.Println(postType)
		fmt.Println(len(posts))
	}
	for key, value := range feedPosts {
		fmt.Println(key, len(value))
	}
	feedPostLists := []models.Feed{}
	for postType, posts := range feedPosts {
		postsCopy = slices.Clone(posts)
		SortPosts(postsCopy)
		posts = postsCopy
		fmt.Println(config.Blog.PagesConfig[postType].Emoji)
		feed := models.Feed{
			Title: strings.ToTitle(postType) + " " + config.Blog.PagesConfig[postType].Emoji,
			Type:  postType,
			Slug:  prefixURL + postType,
			Posts: posts,
		}
		fmt.Println("Feed", feed.Title, feed.Slug)
		feedPostLists = append(feedPostLists, feed)
	}
	ssg.FeedPosts = feedPostLists
}

// "createFeeds",
type CreateFeedsPlugin struct {
	PluginName string
}

func (c *CreateFeedsPlugin) Name() string {
	return c.PluginName
}

func (c *CreateFeedsPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
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
		}
		fmt.Println("Post:", feed.Title, len(feed.Posts))
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			log.Fatal(err)
		}
		feedPath := filepath.Join(".", config.Blog.OutputDir, feed.Type)
		err = os.MkdirAll(feedPath, os.ModePerm)
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
	}
	// create a folder for post if the post type is "post"
	// as this should be the /posts/<slug> as well as /<slug>
	for _, post := range ssg.Posts {
		if post.Frontmatter.Type == "posts" {
			postPath := filepath.Join(".", config.Blog.OutputDir, post.Frontmatter.Slug)
			err := os.MkdirAll(postPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
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
			}
			err = ssg.TemplateFS.ExecuteTemplate(&buffer, config.Blog.DefaultPostTemplate, context)
			if err != nil {
				log.Fatal(err)
			}
			outputPostPath := fmt.Sprintf("%s/index.html", postPath)
			err = os.WriteFile(outputPostPath, buffer.Bytes(), 0660)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

// "createFeeds",
// "copyStaticFiles",
type CopyStaticFilesPlugin struct {
	PluginName string
}

func (c *CopyStaticFilesPlugin) Name() string {
	return c.PluginName
}

func (c *CopyStaticFilesPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	err := Copy(config.Blog.StaticDir, config.Blog.OutputDir)
	err = GeneratePages(*config)
	if err != nil {
		log.Fatal(err)
	}
}

type IndexPlugin struct {
	PluginName string
}

func (c *IndexPlugin) Name() string {
	return c.PluginName
}

func (c *IndexPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config

	buffer := bytes.Buffer{}
	templateFS := os.DirFS(config.Blog.StaticDir)
	t, err := template.ParseFS(templateFS, "index.html")
	if err != nil {
		log.Fatal(err)
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
	}
	err = t.ExecuteTemplate(&buffer, "index.html", context)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(".", config.Blog.OutputDir, "index.html"), buffer.Bytes(), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

// admin
type AdminPlugin struct {
	PluginName string
}

func (c *AdminPlugin) Name() string {
	return c.PluginName
}

func (c *AdminPlugin) Execute(ssg *models.SSG) {
}

// "server"
type ServerPlugin struct {
	PluginName string
}

func (c *ServerPlugin) Name() string {
	return c.PluginName
}

func (c *ServerPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	http.Handle("/", http.FileServer(http.Dir(config.Blog.OutputDir)))
	fmt.Println("Listening on port 3030")
	http.ListenAndServe(":3030", nil)
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
		log.Fatal(err)
	}
	var config models.SSG_CONFIG
	err = json.Unmarshal(configbytes, &config)
	if err != nil {
		log.Fatal(err)
	}
	ssg.Config = config
	if devEnv {
		ssg.Config.Blog.PrefixURL = ""
	}

	// loading in the posts -> post folder
	// load in the templates
	// create feeds
	// load in the static files
	// pack the html pages and static files in a folder
	// serve the folder
	pluginManager := PluginManager{}
	for _, plugin := range config.Plugins {
		switch plugin {
		case "readPosts":
			pluginManager.Register(&PostReaderPlugin{PluginName: "readPosts"})
		case "renderTemplates":
			pluginManager.Register(&RenderTemplatesPlugin{PluginName: "renderTemplates"})
		case "createFeeds":
			pluginManager.Register(&CreateFeedsPlugin{PluginName: "createFeeds"})
		case "copyStaticFiles":
			pluginManager.Register(&CopyStaticFilesPlugin{PluginName: "copyStaticFiles"})
		case "index":
			pluginManager.Register(&IndexPlugin{PluginName: "index"})
		default:

			//userPlugin := plugins.UserPlugin{PluginName: plugin}
			//pluginManager.Register(&userPlugin)

			if plugin != "server" {
				pluginStruct, err := LoadPlugin(plugin)
				if err != nil {
					log.Printf("Error loading plugin %s: %v", plugin, err)
					continue
				}
				fmt.Println("Load", plugin, pluginStruct)
				pluginManager.Register(pluginStruct)
			} else {
				continue
			}
		}
	}
	pluginManager.ExecuteAll(&ssg)
	pluginManager = PluginManager{}
	originalOutputDir := ssg.Config.Blog.OutputDir
	ssg.Config.AdminMode = true
	ssg.Config.Blog.OutputDir = path.Join(ssg.Config.Blog.OutputDir, ssg.Config.Blog.AdminDir)
	for _, plugin := range config.Plugins {
		switch plugin {
		case "readPosts":
			pluginManager.Register(&PostReaderPlugin{PluginName: "readPosts"})
		case "renderTemplates":
			pluginManager.Register(&RenderTemplatesPlugin{PluginName: "renderTemplates"})
		case "createFeeds":
			pluginManager.Register(&CreateFeedsPlugin{PluginName: "createFeeds"})
		case "copyStaticFiles":
			pluginManager.Register(&CopyStaticFilesPlugin{PluginName: "copyStaticFiles"})
		case "index":
			pluginManager.Register(&IndexPlugin{PluginName: "index"})
		default:

			//userPlugin := plugins.UserPlugin{PluginName: plugin}
			//pluginManager.Register(&userPlugin)

			if plugin != "server" {
				pluginStruct, err := LoadPlugin(plugin)
				if err != nil {
					log.Printf("Error loading plugin %s: %v", plugin, err)
					continue
				}
				fmt.Println("Load", plugin, pluginStruct)
				pluginManager.Register(pluginStruct)
			} else {
				continue
			}
		}
	}
	pluginManager.ExecuteAll(&ssg)
	pluginManager = PluginManager{}
	pluginManager.Register(&AdminPlugin{PluginName: "admin"})
	pluginManager.ExecuteAll(&ssg)
	ssg.Config.Blog.OutputDir = originalOutputDir
	pluginManager = PluginManager{}
	if devEnv {
		pluginManager.Register(&ServerPlugin{PluginName: "server"})
	}
	pluginManager.ExecuteAll(&ssg)
}

func LoadPlugin(pluginName string) (plugins.Plugin, error) {
	pluginType, exists := plugins.GetPluginType(pluginName)
	if !exists {
		return nil, fmt.Errorf("plugin %s not found", pluginName)
	}

	pluginValue := reflect.New(pluginType).Interface()
	pluginInstance, ok := pluginValue.(plugins.Plugin)
	if !ok {
		return nil, fmt.Errorf("type %s does not implement Plugin interface", pluginName)
	}
	val := reflect.ValueOf(pluginInstance).Elem()
	if field := val.FieldByName("PluginName"); field.IsValid() && field.CanSet() {
		field.SetString(pluginName)
	}
	return pluginInstance, nil
}
