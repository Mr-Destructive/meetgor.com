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

// RSSFeed represents the structure of the RSS feed
type RSSFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel RSSChannel `xml:"channel"`
}

// RSSChannel represents an RSS channel
type RSSChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	PubDate     string    `xml:"pubDate"`
	Items       []RSSItem `xml:"item"`
}

// RSSItem represents an item in the RSS feed
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Content     string `xml:"content"`
}

// RSSPlugin generates an RSS feed for the blog posts
type RSSPlugin struct {
	PluginName string
}

func (p *RSSPlugin) Name() string {
	return p.PluginName
}

func (p *RSSPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	baseURL := config.Blog.BaseUrl
	rssFilePath := filepath.Join(config.Blog.OutputDir, "rss.xml")

	// Collect all published posts
	var rssItems []RSSItem
	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}

		pubDate, err := time.Parse("2006-01-02", post.Frontmatter.Date)
		if err != nil {
			log.Printf("Error parsing post date: %v", err)
			continue
		}

		rssItems = append(rssItems, RSSItem{
			Title:       post.Frontmatter.Title,
			Link:        fmt.Sprintf("%s/%s", baseURL, post.Frontmatter.Slug),
			Description: post.Frontmatter.Description,
			PubDate:     pubDate.Format(time.RFC1123),
			Content:     string(post.Markdown),
		})
	}

	// Create the RSS feed
	rssFeed := RSSFeed{
		Version: "2.0",
		Channel: RSSChannel{
			Title:       config.Blog.Name,
			Link:        baseURL,
			Description: config.Blog.Description,
			Language:    "en-us",
			PubDate:     time.Now().Format(time.RFC1123),
			Items:       rssItems,
		},
	}

	// Generate XML
	xmlBytes, err := xml.MarshalIndent(rssFeed, "", "  ")
	if err != nil {
		log.Fatalf("Error generating RSS XML: %v", err)
	}

	// Write RSS XML to file
	err = os.WriteFile(rssFilePath, xmlBytes, 0666)
	if err != nil {
		log.Fatalf("Error writing RSS file: %v", err)
	}

	fmt.Println("RSS feed generated:", rssFilePath)
}

func init() {
	RegisterPlugin("RSS", reflect.TypeOf(RSSPlugin{
		PluginName: "RSS",
	}))
}
