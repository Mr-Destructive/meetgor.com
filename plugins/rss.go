package plugins

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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
    Type        string `xml:"type"`
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

		if post.Frontmatter.Type == "post" || post.Frontmatter.Type == "posts" {
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

	// Create separate RSS feeds for each tag and type
	postsByTag := make(map[string][]RSSItem)
	postsByType := make(map[string][]RSSItem)

	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}

		pubDate, err := time.Parse("2006-01-02", post.Frontmatter.Date)
		if err != nil {
			log.Printf("Error parsing post date: %v", err)
			continue
		}

		rssItem := RSSItem{
			Title:       post.Frontmatter.Title,
			Link:        fmt.Sprintf("%s/%s", baseURL, post.Frontmatter.Slug),
			Description: post.Frontmatter.Description,
			PubDate:     pubDate.Format(time.RFC1123),
			Content:     string(post.Markdown),
			Type:        post.Frontmatter.Type,
		}

		for _, tag := range post.Frontmatter.Tags {
			postsByTag[tag] = append(postsByTag[tag], rssItem)
		}
		postsByType[post.Frontmatter.Type] = append(postsByType[post.Frontmatter.Type], rssItem)
	}

	type FeedLink struct {
		URL   string
		Title string
		Count int
	}

	var tagFeedLinks []FeedLink
	var typeFeedLinks []FeedLink

	for tag, items := range postsByTag {
		tagFeed := RSSFeed{
			Version: "2.0",
			Channel: RSSChannel{
				Title:       fmt.Sprintf("%s - Tag: %s", config.Blog.Name, tag),
				Link:        baseURL,
				Description: fmt.Sprintf("Posts tagged with %s", tag),
				Language:    "en-us",
				PubDate:     time.Now().Format(time.RFC1123),
				Items:       items,
			},
		}
		tagRssFilePath := filepath.Join(config.Blog.OutputDir, "tags", tag, "rss.xml")
		if err := os.MkdirAll(filepath.Dir(tagRssFilePath), 0755); err != nil {
			log.Fatalf("Error creating tag directory: %v", err)
		}
		writeRSSFeed(tagFeed, tagRssFilePath)
		tagFeedLinks = append(tagFeedLinks, FeedLink{URL: fmt.Sprintf("%s/tags/%s/rss.xml", config.Blog.PrefixURL, tag), Title: tag, Count: len(items)})
	}

	sort.Slice(tagFeedLinks, func(i, j int) bool {
		return tagFeedLinks[i].Count > tagFeedLinks[j].Count
	})

	for postType, items := range postsByType {
		typeFeed := RSSFeed{
			Version: "2.0",
			Channel: RSSChannel{
				Title:       fmt.Sprintf("%s - Type: %s", config.Blog.Name, postType),
				Link:        baseURL,
				Description: fmt.Sprintf("Posts of type %s", postType),
				Language:    "en-us",
				PubDate:     time.Now().Format(time.RFC1123),
				Items:       items,
			},
		}
		typeRssFilePath := filepath.Join(config.Blog.OutputDir, "type", postType, "rss.xml")
		if err := os.MkdirAll(filepath.Dir(typeRssFilePath), 0755); err != nil {
			log.Fatalf("Error creating type directory: %v", err)
		}
		writeRSSFeed(typeFeed, typeRssFilePath)
		typeFeedLinks = append(typeFeedLinks, FeedLink{URL: fmt.Sprintf("%s/type/%s/rss.xml", config.Blog.PrefixURL, postType), Title: postType, Count: len(items)})
	}

	sort.Slice(typeFeedLinks, func(i, j int) bool {
		return typeFeedLinks[i].Count > typeFeedLinks[j].Count
	})

	// Create a page to list all RSS feeds
	feedsPagePath := filepath.Join(config.Blog.OutputDir, "feeds", "index.html")
	if err := os.MkdirAll(filepath.Dir(feedsPagePath), 0755); err != nil {
		log.Fatalf("Error creating feeds directory: %v", err)
	}

	tmpl, err := template.ParseGlob(filepath.Join(config.Blog.TemplatesDir, "*.html"))
	if err != nil {
		log.Fatalf("Error parsing feeds template: %v", err)
	}

	file, err := os.Create(feedsPagePath)
	if err != nil {
		log.Fatalf("Error creating feeds page: %v", err)
	}
	defer file.Close()

	data := struct {
		Config        *models.SSG_CONFIG
		TagFeedLinks  []FeedLink
		TypeFeedLinks []FeedLink
	}{
		Config:        config,
		TagFeedLinks:  tagFeedLinks,
		TypeFeedLinks: typeFeedLinks,
	}

	err = tmpl.ExecuteTemplate(file, "feeds_template.html", data)
	if err != nil {
		log.Fatalf("Error executing feeds template: %v", err)
	}

	fmt.Println("RSS feeds page generated:", feedsPagePath)
}

func writeRSSFeed(feed RSSFeed, filePath string) {
	xmlBytes, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		log.Printf("Error generating RSS XML for %s: %v", filePath, err)
		return
	}

	err = os.WriteFile(filePath, xmlBytes, 0666)
	if err != nil {
		log.Printf("Error writing RSS file for %s: %v", filePath, err)
		return
	}

	fmt.Println("RSS feed generated:", filePath)
}

func init() {
	RegisterPlugin("RSS", reflect.TypeOf(RSSPlugin{
		PluginName: "RSS",
	}))
}
