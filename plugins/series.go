package plugins

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mr-Destructive/meetgor.com/models"
)

type SeriesPlugin struct {
	PluginName string
}

func (p *SeriesPlugin) Name() string {
	return p.PluginName
}

func (p *SeriesPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *SeriesPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (p *SeriesPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (p *SeriesPlugin) Execute(ssg *models.SSG) error {
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

		// Handle both string and slice cases for series
		var seriesList []interface{}
		switch v := post.Frontmatter.Extras["series"].(type) {
		case []interface{}:
			seriesList = v
		case string:
			seriesList = []interface{}{v}
		default:
			// Skip if it's neither a string nor a slice
			continue
		}

		for _, series := range seriesList {
			series := Slugify(series.(string))
			fmt.Println("series", series)
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
			Years: YearsFromPosts(ssg.Posts),
		}
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			return err
		}
		feedPath := filepath.Join(".", config.Blog.OutputDir, "series", feed.Type)
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
	return nil
}

func init() {
	RegisterPlugin("Series", func() Plugin {
		return &SeriesPlugin{PluginName: "Series"}
	})
}
