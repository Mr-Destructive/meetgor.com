package plugins

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Mr-Destructive/meetgor.com/models"
)

type URL struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type SitemapPlugin struct {
	PluginName string
}

var sitemapNow = time.Now

func (s *SitemapPlugin) Name() string {
	return s.PluginName
}

func (s *SitemapPlugin) Phase() Phase {
	return PhasePostProcess
}

func (s *SitemapPlugin) Requires() []string {
	return []string{"readPosts"}
}

func (s *SitemapPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (s *SitemapPlugin) Execute(ssg *models.SSG) error {
	config := &ssg.Config
	baseURL := config.Blog.BaseUrl
	prefixURL := config.Blog.PrefixURL
	outputDir := config.Blog.OutputDir

	urlSet := map[string]bool{}
	var urls []URL
	addURL := func(loc string, lastMod string, changeFreq string, priority string) {
		if loc == "" {
			return
		}
		if urlSet[loc] {
			return
		}
		urlSet[loc] = true
		urls = append(urls, URL{
			Loc:        loc,
			LastMod:    lastMod,
			ChangeFreq: changeFreq,
			Priority:   priority,
		})
	}

	homeURL := fmt.Sprintf("%s/%s", baseURL, prefixURL)
	addURL(homeURL, sitemapNow().Format("2006-01-02"), "weekly", "0.9")

	// Add all configured pages (feeds + static pages)
	for pageName := range config.Blog.PagesConfig {
		pageURL := fmt.Sprintf("%s/%s%s", baseURL, prefixURL, pageName)
		addURL(pageURL, sitemapNow().Format("2006-01-02"), "weekly", "0.7")
	}

	// Add feed index pages from generated feeds
	for _, feed := range ssg.FeedPosts {
		feedURL := fmt.Sprintf("%s/%s%s", baseURL, prefixURL, feed.Slug)
		addURL(feedURL, sitemapNow().Format("2006-01-02"), "weekly", "0.7")
	}

	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		pageURL := fmt.Sprintf("%s/%s%s/%s", baseURL, prefixURL, post.Frontmatter.Type, post.Frontmatter.Slug)
		lastMod := post.Frontmatter.Date
		if lastMod == "" {
			lastMod = sitemapNow().Format("2006-01-02")
		}
		addURL(pageURL, lastMod, "weekly", "0.8")
	}

	sitemap := Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}
	xmlData, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return fmt.Errorf("error generating XML: %w", err)
	}

	sitemapPath := filepath.Join(outputDir, "sitemap.xml")
	err = os.WriteFile(sitemapPath, []byte(xml.Header+string(xmlData)), 0644)
	if err != nil {
		return fmt.Errorf("error writing sitemap.xml: %w", err)
	}
	fmt.Println("Sitemap generated at", sitemapPath)
	return nil
}

func init() {
	RegisterPlugin("Sitemap", func() Plugin {
		return &SitemapPlugin{PluginName: "Sitemap"}
	})
}
