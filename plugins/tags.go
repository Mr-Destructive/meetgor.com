package plugins

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Mr-Destructive/meetgor.com/models"
)

type TagsPlugin struct {
	PluginName string
}

func (p *TagsPlugin) Name() string {
	return p.PluginName
}

func (p *TagsPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *TagsPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (p *TagsPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (p *TagsPlugin) Execute(ssg *models.SSG) error {
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
	type TagLink struct {
		Title string
		Slug  string
		Count int
	}
	tagLinks := make([]TagLink, 0, len(tagPosts))

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
			Years: YearsFromPosts(ssg.Posts),
		}
		err := ssg.TemplateFS.ExecuteTemplate(&buffer, templatePath, context)
		if err != nil {
			return err
		}
		feedPath := filepath.Join(".", config.Blog.OutputDir, "tags", feed.Type)
		err = os.MkdirAll(feedPath, os.ModePerm)
		if err != nil {
			return err
		}
		feedPath2 := filepath.Join(".", config.Blog.OutputDir, "tag", feed.Type)
		err = os.MkdirAll(feedPath2, os.ModePerm)
		if err != nil {
			return err
		}
		outputFeedPath := fmt.Sprintf("%s/index.html", feedPath)
		err = os.WriteFile(outputFeedPath, buffer.Bytes(), 0660)
		if err != nil {
			return err
		}
		outputFeedPath2 := fmt.Sprintf("%s/index.html", feedPath2)
		err = os.WriteFile(outputFeedPath2, buffer.Bytes(), 0660)
		if err != nil {
			return err
		}

		tagLinks = append(tagLinks, TagLink{
			Title: tag,
			Slug:  tag,
			Count: len(posts),
		})
	}

	if ssg.TemplateFS.Lookup("tags_index.html") == nil {
		return nil
	}

	sort.Slice(tagLinks, func(i, j int) bool {
		if tagLinks[i].Count == tagLinks[j].Count {
			return tagLinks[i].Title < tagLinks[j].Title
		}
		return tagLinks[i].Count > tagLinks[j].Count
	})

	indexBuffer := bytes.Buffer{}
	indexContext := struct {
		Config       *models.SSG_CONFIG
		TagFeedLinks []TagLink
		Years        []string
	}{
		Config:       config,
		TagFeedLinks: tagLinks,
		Years:        YearsFromPosts(ssg.Posts),
	}
	if err := ssg.TemplateFS.ExecuteTemplate(&indexBuffer, "tags_index.html", indexContext); err != nil {
		return err
	}
	tagsIndexPath := filepath.Join(".", config.Blog.OutputDir, "tags", "index.html")
	if err := os.MkdirAll(filepath.Dir(tagsIndexPath), os.ModePerm); err != nil {
		return err
	}
	if err := os.WriteFile(tagsIndexPath, indexBuffer.Bytes(), 0660); err != nil {
		return err
	}
	return nil
}

func init() {
	RegisterPlugin("Tags", func() Plugin {
		return &TagsPlugin{PluginName: "Tags"}
	})
}
