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

type SeriesPlugin struct {
	PluginName string
}

func (p *SeriesPlugin) Name() string {
	return p.PluginName
}

func (p *SeriesPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	seriesPost := make(map[string][]models.Post)
	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		CleanPostFrontmatter(&post, ssg)
		if post.Frontmatter.Extras == nil {
			continue
		}
		if post.Frontmatter.Extras["series"] == nil {
			continue
		}
		seriesList := post.Frontmatter.Extras["series"].([]interface{})
		for _, series := range seriesList {
			series := Slugify(series.(string))
			seriesPost[series] = append(seriesPost[series], post)
		}
	}
	fmt.Println("seriesPost", len(seriesPost))

	for series, posts := range seriesPost {
		buffer := bytes.Buffer{}
		templatePath := config.Blog.PagesConfig["series"].TemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}
		feed := models.Feed{
			Title: series,
			Type:  series,
			Slug:  ssg.Config.Blog.PrefixURL + "series/" + series,
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
		feedPath := filepath.Join(".", config.Blog.OutputDir, "series", feed.Type)
		err = os.MkdirAll(feedPath, os.ModePerm)
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
	}
}

func init() {
	RegisterPlugin("Series", reflect.TypeOf(SeriesPlugin{
		PluginName: "Series",
	}))
}
