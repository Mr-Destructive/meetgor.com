package plugins

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type TagsPlugin struct {
	PluginName string
}

func (p *TagsPlugin) Name() string {
	return p.PluginName
}

func (p *TagsPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	tagPosts := make(map[string][]models.Post)
	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		CleanPostFrontmatter(&post, ssg)
		for _, tag := range post.Frontmatter.Tags {
			tagPosts[tag] = append(tagPosts[tag], post)
		}
	}
	for tag, posts := range tagPosts {
		buffer := bytes.Buffer{}
		templatePath := config.Blog.PagesConfig["tags"].TemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}
		feed := models.Feed{
			Title: tag,
			Type:  tag,
			Slug:  ssg.Config.Blog.PrefixURL + "tags/" + tag,
			Posts: posts,
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
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			log.Fatal(err)
		}
		feedPath := filepath.Join(".", config.Blog.OutputDir, "tags", feed.Type)
		err = os.MkdirAll(feedPath, os.ModePerm)
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
	}
}

func init() {
	RegisterPlugin("Tags", reflect.TypeOf(TagsPlugin{
		PluginName: "Tags",
	}))
}
