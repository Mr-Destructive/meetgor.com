package plugins

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
	"net/url"
	"strings"

	"github.com/Mr-Destructive/meetgor.com/models"
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

var rssNow = time.Now

func parsePostDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date")
	}
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

func (p *RSSPlugin) Name() string {
	return p.PluginName
}

func (p *RSSPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *RSSPlugin) Requires() []string {
	return []string{"readPosts"}
}

func (p *RSSPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (p *RSSPlugin) Execute(ssg *models.SSG) error {
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
			pubDate, err := parsePostDate(post.Frontmatter.Date)
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
			PubDate:     rssNow().Format(time.RFC1123),
			Items:       rssItems,
		},
	}

	// Generate XML
	xmlBytes, err := xml.MarshalIndent(rssFeed, "", "  ")
	if err != nil {
		return fmt.Errorf("error generating RSS XML: %w", err)
	}

	// Write RSS XML to file
	err = os.WriteFile(rssFilePath, xmlBytes, 0666)
	if err != nil {
		return fmt.Errorf("error writing RSS file: %w", err)
	}

	fmt.Println("RSS feed generated:", rssFilePath)

	// Create separate RSS feeds for each tag and type
	postsByTag := make(map[string][]RSSItem)
	postsByType := make(map[string][]RSSItem)

	for _, post := range ssg.Posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}

		pubDate, err := parsePostDate(post.Frontmatter.Date)
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
				PubDate:     rssNow().Format(time.RFC1123),
				Items:       items,
			},
		}
		tagRssFilePath := filepath.Join(config.Blog.OutputDir, "tags", tag, "rss.xml")
		if err := os.MkdirAll(filepath.Dir(tagRssFilePath), 0755); err != nil {
			return fmt.Errorf("error creating tag directory: %w", err)
		}
		writeRSSFeed(tagFeed, tagRssFilePath)
		escapedTag := url.PathEscape(tag)
		escapedTag = strings.ReplaceAll(escapedTag, "+", "%2B")
		tagFeedLinks = append(tagFeedLinks, FeedLink{
			URL:   fmt.Sprintf("%s/tags/%s/rss.xml", config.Blog.PrefixURL, escapedTag),
			Title: tag,
			Count: len(items),
		})
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
				PubDate:     rssNow().Format(time.RFC1123),
				Items:       items,
			},
		}
		typeRssFilePath := filepath.Join(config.Blog.OutputDir, "type", postType, "rss.xml")
		if err := os.MkdirAll(filepath.Dir(typeRssFilePath), 0755); err != nil {
			return fmt.Errorf("error creating type directory: %w", err)
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
		return fmt.Errorf("error creating feeds directory: %w", err)
	}

	file, err := os.Create(feedsPagePath)
	if err != nil {
		return fmt.Errorf("error creating feeds page: %w", err)
	}
	defer file.Close()

	data := struct {
		Config        *models.SSG_CONFIG
		TagFeedLinks  []FeedLink
		TypeFeedLinks []FeedLink
		Years         []string
	}{
		Config:        config,
		TagFeedLinks:  tagFeedLinks,
		TypeFeedLinks: typeFeedLinks,
		Years:         YearsFromPosts(ssg.Posts),
	}

	err = ssg.TemplateFS.ExecuteTemplate(file, "feeds_template.html", data)
	if err != nil {
		return fmt.Errorf("error executing feeds template: %w", err)
	}

	fmt.Println("RSS feeds page generated:", feedsPagePath)
	return nil
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
	RegisterPlugin("RSS", func() Plugin {
		return &RSSPlugin{PluginName: "RSS"}
	})
}
