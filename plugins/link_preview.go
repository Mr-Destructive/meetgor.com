package plugins

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type LinkPreviewPlugin struct {
	PluginName string
}

func (p *LinkPreviewPlugin) Name() string {
	return p.PluginName
}

func (p *LinkPreviewPlugin) Execute(ssg *models.SSG) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := range ssg.Posts {
		if ssg.Posts[i].Frontmatter.Status == "draft" {
			continue
		}

		// Initialize Extras if nil
		if ssg.Posts[i].Frontmatter.Extras == nil {
			ssg.Posts[i].Frontmatter.Extras = make(map[string]interface{})
		}

		// Skip if post already has an image_url
		if ssg.Posts[i].Frontmatter.ImageUrl != "" {
			continue
		}

		// Check if post has a link field
		linkRaw, hasLink := ssg.Posts[i].Frontmatter.Extras["link"]
		if !hasLink {
			continue
		}

		link, ok := linkRaw.(string)
		if !ok || link == "" {
			continue
		}

		// Fetch preview image from the link
		previewURL := p.fetchOGImage(client, link)
		if previewURL != "" {
			ssg.Posts[i].Frontmatter.Extras["preview_image"] = previewURL
		}
	}

	fmt.Println("Link Preview plugin executed successfully")
}

// fetchOGImage fetches the Open Graph image from a URL
func (p *LinkPreviewPlugin) fetchOGImage(client *http.Client, url string) string {
	// Add http:// if no scheme
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	// Read only first 100KB to avoid large files
	limitedReader := io.LimitReader(resp.Body, 100*1024)
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		return ""
	}

	bodyStr := string(body)

	// Try to find og:image meta tag
	ogImageURL := p.extractMetaTag(bodyStr, "og:image")
	if ogImageURL != "" {
		return ogImageURL
	}

	// Fallback to twitter:image
	twitterImageURL := p.extractMetaTag(bodyStr, "twitter:image")
	if twitterImageURL != "" {
		return twitterImageURL
	}

	return ""
}

// extractMetaTag extracts content from a meta tag
func (p *LinkPreviewPlugin) extractMetaTag(html string, property string) string {
	// Pattern for og:image or twitter:image
	patterns := []string{
		`<meta\s+property=["']` + property + `["']\s+content=["']([^"']+)["']`,
		`<meta\s+content=["']([^"']+)["']\s+property=["']` + property + `["']`,
		`<meta\s+name=["']` + property + `["']\s+content=["']([^"']+)["']`,
		`<meta\s+content=["']([^"']+)["']\s+name=["']` + property + `["']`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(html)
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return ""
}

func init() {
	RegisterPlugin("LinkPreview", reflect.TypeOf(LinkPreviewPlugin{
		PluginName: "LinkPreview",
	}))
}
