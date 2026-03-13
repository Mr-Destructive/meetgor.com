package plugins

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type YearPlugin struct {
	PluginName string
}

func (p *YearPlugin) Name() string {
	return p.PluginName
}

func (p *YearPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *YearPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (p *YearPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (p *YearPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	yearWisePosts := make(map[string][]models.Post)
	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		CleanPostFrontmatter(&post, ssg)
		date := post.Frontmatter.Date
		parsedDate, err := time.Parse("2006-01-02", string(date))
		if err != nil {
			continue
		}
		if parsedDate.Year() < 2000 {
			continue
		}
		yearWisePosts[parsedDate.Format("2006")] = append(yearWisePosts[parsedDate.Format("2006")], post)
	}
	for year, posts := range yearWisePosts {
		buffer := bytes.Buffer{}
		templatePath := config.Blog.PagesConfig["tags"].TemplatePath
		if templatePath == "" {
			templatePath = config.Blog.DefaultFeedTemplate
		}
		feed := models.Feed{
			Title: year,
			Type:  year,
			Slug:  ssg.Config.Blog.PrefixURL + year,
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
	return nil
}

func init() {
	RegisterPlugin("YearWise", func() Plugin {
		return &YearPlugin{PluginName: "YearWise"}
	})
}
