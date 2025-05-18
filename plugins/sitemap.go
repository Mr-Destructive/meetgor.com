package plugins

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
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

func (s *SitemapPlugin) Name() string {
	return s.PluginName
}

func (s *SitemapPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	baseURL := config.Blog.BaseUrl
	prefixURL := config.Blog.PrefixURL
	outputDir := config.Blog.OutputDir

	var urls []URL
	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		pageURL := fmt.Sprintf("%s/%s%s/%s", baseURL, prefixURL, post.Frontmatter.Type, post.Frontmatter.Slug)
		lastMod := post.Frontmatter.Date
		if lastMod == "" {
			lastMod = time.Now().Format("2006-01-02")
		}
		urls = append(urls, URL{
			Loc:        pageURL,
			LastMod:    lastMod,
			ChangeFreq: "weekly",
			Priority:   "0.8",
		})
	}

	sitemap := Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}
	xmlData, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		log.Fatal("Error generating XML:", err)
	}

	sitemapPath := filepath.Join(outputDir, "sitemap.xml")
	err = os.WriteFile(sitemapPath, []byte(xml.Header+string(xmlData)), 0644)
	if err != nil {
		log.Fatal("Error writing sitemap.xml:", err)
	}
	fmt.Println("Sitemap generated at", sitemapPath)
}

func init() {
	RegisterPlugin("Sitemap", reflect.TypeOf(SitemapPlugin{
		PluginName: "Sitemap",
	}))
}
