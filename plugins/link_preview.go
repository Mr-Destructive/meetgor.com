package plugins

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
	"gopkg.in/yaml.v3"
)

type LinkPreviewPlugin struct {
	PluginName string
}

var linkPreviewClientFactory = func() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func (p *LinkPreviewPlugin) Name() string {
	return p.PluginName
}

func (p *LinkPreviewPlugin) Phase() Phase {
	return PhaseTransform
}

func (p *LinkPreviewPlugin) Requires() []string {
	return []string{"readPosts"}
}

func (p *LinkPreviewPlugin) AdminPolicy() AdminPolicy {
	return AdminRun
}

func (p *LinkPreviewPlugin) Execute(ssg *models.SSG) error {
	client := linkPreviewClientFactory()

	var wg sync.WaitGroup
	var mux sync.Mutex

	for i := range ssg.Posts {
		post := &ssg.Posts[i]
		if strings.EqualFold(post.Frontmatter.Status, "draft") {
			continue
		}
		if post.Frontmatter.Type != "links" {
			continue
		}
		if post.Frontmatter.Extras == nil {
			continue
		}
		linkRaw, ok := post.Frontmatter.Extras["link"]
		if !ok {
			continue
		}
		link, ok := linkRaw.(string)
		if !ok || link == "" {
			continue
		}

		if post.Frontmatter.ImageUrl != "" {
			continue
		}

		wg.Add(1)
		go func(post *models.Post, link string) {
			defer wg.Done()
			imageURL := p.fetchOGImage(client, link)
			if imageURL == "" {
				return
			}
			if err := writeImageURLOutside(post.SourcePath, imageURL); err != nil {
				log.Printf("LinkPreview: unable to persist image_url for %s: %v", post.SourcePath, err)
				return
			}
			mux.Lock()
			post.Frontmatter.ImageUrl = imageURL
			if post.Frontmatter.Extras == nil {
				post.Frontmatter.Extras = make(map[string]interface{})
			}
			post.Frontmatter.Extras["image_url"] = imageURL
			mux.Unlock()
		}(post, link)
	}

	wg.Wait()
	fmt.Println("Link Preview plugin executed successfully")
	return nil
}

func (p *LinkPreviewPlugin) fetchOGImage(client *http.Client, url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return p.platformPreviewImage(url)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; LinkPreviewBot/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := client.Do(req)
	if err != nil {
		return p.platformPreviewImage(url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return p.platformPreviewImage(url)
	}

	reader := io.LimitReader(resp.Body, 150*1024)
	body, err := io.ReadAll(reader)
	if err != nil {
		return p.platformPreviewImage(url)
	}

	bodyStr := string(body)
	candidates := []string{
		"og:image",
		"og:image:url",
		"og:image:secure_url",
		"twitter:image",
		"twitter:image:src",
	}
	for _, candidate := range candidates {
		if image := p.extractMetaTag(bodyStr, candidate); image != "" {
			if normalized := normalizeImageURL(url, image); normalized != "" {
				return normalized
			}
		}
	}
	return p.platformPreviewImage(url)
}

func normalizeImageURL(pageURL string, imageURL string) string {
	imageURL = strings.TrimSpace(imageURL)
	if imageURL == "" || strings.HasPrefix(imageURL, "data:") {
		return ""
	}
	if strings.HasPrefix(imageURL, "//") {
		if parsed, err := url.Parse(pageURL); err == nil && parsed.Scheme != "" {
			return parsed.Scheme + ":" + imageURL
		}
		return "https:" + imageURL
	}
	parsedImage, err := url.Parse(imageURL)
	if err != nil {
		return imageURL
	}
	if parsedImage.IsAbs() {
		return imageURL
	}
	parsedPage, err := url.Parse(pageURL)
	if err != nil {
		return imageURL
	}
	return parsedPage.ResolveReference(parsedImage).String()
}

func (p *LinkPreviewPlugin) extractMetaTag(html string, property string) string {
	patterns := []string{
		`<meta\s+property=["']` + property + `["']\s+content=["']([^"']+)["']`,
		`<meta\s+content=["']([^"']+)["']\s+property=["']` + property + `["']`,
		`<meta\s+name=["']` + property + `["']\s+content=["']([^"']+)["']`,
		`<meta\s+content=["']([^"']+)["']\s+name=["']` + property + `["']`,
	}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if matches := re.FindStringSubmatch(html); len(matches) > 1 {
			return strings.TrimSpace(matches[1])
		}
	}
	return ""
}

func (p *LinkPreviewPlugin) platformPreviewImage(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	host := strings.ToLower(parsed.Host)
	host = strings.TrimPrefix(host, "www.")

	switch host {
	case "youtu.be", "youtube.com", "m.youtube.com", "music.youtube.com", "youtube-nocookie.com":
		if id := extractYouTubeIDFromURL(parsed); id != "" {
			return "https://i.ytimg.com/vi/" + id + "/hqdefault.jpg"
		}
	case "vimeo.com", "player.vimeo.com":
		if id := extractVimeoID(parsed); id != "" {
			return "https://vumbnail.com/" + id + ".jpg"
		}
	case "dailymotion.com", "dai.ly":
		if id := extractDailymotionID(parsed); id != "" {
			return "https://www.dailymotion.com/thumbnail/video/" + id
		}
	}

	return ""
}

func extractYouTubeIDFromURL(u *url.URL) string {
	host := strings.ToLower(strings.TrimPrefix(u.Host, "www."))
	switch host {
	case "youtu.be":
		id := strings.TrimPrefix(u.Path, "/")
		if id == "" {
			return ""
		}
		return strings.Split(id, "/")[0]
	case "youtube.com", "m.youtube.com", "music.youtube.com", "youtube-nocookie.com":
		if strings.HasPrefix(u.Path, "/watch") {
			return u.Query().Get("v")
		}
		if strings.HasPrefix(u.Path, "/shorts/") {
			id := strings.TrimPrefix(u.Path, "/shorts/")
			return strings.Split(id, "/")[0]
		}
		if strings.HasPrefix(u.Path, "/embed/") {
			id := strings.TrimPrefix(u.Path, "/embed/")
			return strings.Split(id, "/")[0]
		}
		if strings.HasPrefix(u.Path, "/live/") {
			id := strings.TrimPrefix(u.Path, "/live/")
			return strings.Split(id, "/")[0]
		}
	}
	return ""
}

func extractVimeoID(u *url.URL) string {
	path := strings.Trim(u.Path, "/")
	if path == "" {
		return ""
	}
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if part == "" {
			continue
		}
		if isAllDigits(part) {
			return part
		}
	}
	return ""
}

func extractDailymotionID(u *url.URL) string {
	path := strings.Trim(u.Path, "/")
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "video/") {
		id := strings.TrimPrefix(path, "video/")
		return strings.Split(id, "/")[0]
	}
	if strings.HasPrefix(path, "embed/video/") {
		id := strings.TrimPrefix(path, "embed/video/")
		return strings.Split(id, "/")[0]
	}
	if strings.Contains(u.Host, "dai.ly") {
		return strings.Split(path, "/")[0]
	}
	return ""
}

func isAllDigits(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}
	return value != ""
}

func writeImageURLOutside(sourcePath, imageURL string) error {
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	if format, frontMatter, body, prefix, ok := splitFrontMatter(content); ok {
		var updated []byte
		switch format {
		case "yaml":
			var data map[string]interface{}
			if err := yaml.Unmarshal(frontMatter, &data); err != nil {
				return err
			}
			if data == nil {
				data = make(map[string]interface{})
			}
			data["image_url"] = imageURL
			newFM, err := yaml.Marshal(data)
			if err != nil {
				return err
			}
			newFM = bytes.TrimRight(newFM, "\n")
			var builder strings.Builder
			builder.Write(prefix)
			builder.WriteString("---\n")
			builder.Write(newFM)
			builder.WriteString("\n---")
			if len(body) > 0 && body[0] != '\n' {
				builder.WriteByte('\n')
			}
			builder.Write(body)
			updated = []byte(builder.String())
		default:
			// unsupported format
			return nil
		}
		return os.WriteFile(sourcePath, updated, 0644)
	}
	return fmt.Errorf("unable to locate YAML front matter for %s", sourcePath)
}

func splitFrontMatter(content []byte) (format string, frontMatter, body, prefix []byte, ok bool) {
	if parts := bytes.SplitN(content, []byte("---"), 3); len(parts) == 3 {
		return "yaml", parts[1], parts[2], parts[0], true
	}
	return "", nil, nil, nil, false
}

func init() {
	RegisterPlugin("LinkPreview", func() Plugin {
		return &LinkPreviewPlugin{PluginName: "LinkPreview"}
	})
}
