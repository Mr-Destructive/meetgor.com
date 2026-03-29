package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	gmhtml "github.com/yuin/goldmark/renderer/html"
	nethtml "golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

type linkPostDraft struct {
	Title       string
	Slug        string
	Date        string
	Link        string
	Description string
	ImageURL    string
	Body        string
	Newsletter  string
	Hash        string
}

type linkMetadata struct {
	Title       string
	Description string
	ImageURL    string
}

var linkMetadataFetcher = fetchLinkMetadata

func materializeLinkPosts(outputDir, createdFilesPath string, newsletters []newsletterFeedPost) (int, error) {
	linkDir := defaultLinkOutputDir(outputDir)
	existing, existingHashes, existingLinks, err := readExistingLinkState(linkDir)
	if err != nil {
		return 0, err
	}

	if err := os.MkdirAll(linkDir, 0o755); err != nil {
		return 0, err
	}

	created := 0
	for _, newsletter := range newsletters {
		drafts, err := extractLinkDraftsFromNewsletter(newsletter)
		if err != nil {
			return created, err
		}

		for _, draft := range drafts {
			if draft.Slug == "" {
				continue
			}
			if existing[draft.Slug] {
				continue
			}
			if draft.Link != "" && existingLinks[draft.Link] {
				continue
			}
			if draft.Hash != "" && existingHashes[draft.Hash] {
				continue
			}
			if err := writeLinkPost(linkDir, createdFilesPath, draft); err != nil {
				return created, err
			}
			existing[draft.Slug] = true
			if draft.Link != "" {
				existingLinks[draft.Link] = true
			}
			if draft.Hash != "" {
				existingHashes[draft.Hash] = true
			}
			created++
		}
	}

	return created, nil
}

func extractLinkDraftsFromNewsletter(newsletter newsletterFeedPost) ([]linkPostDraft, error) {
	body := newsletter.Body
	if body == "" {
		body = newsletter.Description
	}
	if strings.TrimSpace(body) == "" {
		return nil, nil
	}

	rendered, err := renderNewsletterBodyToHTML(body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rendered))
	if err != nil {
		return nil, err
	}

	drafts := make([]linkPostDraft, 0)
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(_ int, heading *goquery.Selection) {
		if !isReadOrWatchedHeading(heading.Text()) {
			return
		}

		for sib := heading.Next(); sib.Length() > 0; sib = sib.Next() {
			if isHeadingNode(sib) {
				break
			}
			collectLinkDraftsFromNode(sib, newsletter, &drafts)
		}
	})

	return drafts, nil
}

func readExistingLinkState(linkDir string) (map[string]bool, map[string]bool, map[string]bool, error) {
	existing := make(map[string]bool)
	existingHashes := make(map[string]bool)
	existingLinks := make(map[string]bool)

	entries, err := os.ReadDir(linkDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, nil, nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		slug := strings.TrimSuffix(entry.Name(), ".md")
		existing[slug] = true

		content, err := os.ReadFile(filepath.Join(linkDir, entry.Name()))
		if err != nil {
			continue
		}

		frontmatter, _, err := splitNewsletterFrontmatter(string(content))
		if err != nil {
			continue
		}

		var metadata map[string]interface{}
		if err := yaml.Unmarshal([]byte(frontmatter), &metadata); err != nil {
			if err := json.Unmarshal([]byte(frontmatter), &metadata); err != nil {
				continue
			}
		}

		if hash := extractMetadataString(metadata, "hash"); hash != "" {
			existingHashes[hash] = true
		}
		if link := firstNonEmpty(
			extractMetadataString(metadata, "link"),
			extractMetadataString(metadata, "canonical_url"),
			extractMetadataString(metadata, "url"),
		); link != "" {
			existingLinks[link] = true
		}
	}

	return existing, existingHashes, existingLinks, nil
}

func collectLinkDraftsFromNode(node *goquery.Selection, newsletter newsletterFeedPost, drafts *[]linkPostDraft) {
	node.Find("li").Each(func(_ int, li *goquery.Selection) {
		if !isTopLevelListItem(li) {
			return
		}

		draft, ok := buildLinkDraftFromListItem(li, newsletter)
		if !ok {
			return
		}

		*drafts = append(*drafts, draft)
	})
}

func buildLinkDraftFromListItem(li *goquery.Selection, newsletter newsletterFeedPost) (linkPostDraft, bool) {
	anchor := li.Find("a").First()
	if anchor.Length() == 0 {
		return linkPostDraft{}, false
	}

	rawLink, _ := anchor.Attr("href")
	rawTitle := strings.TrimSpace(anchor.Text())
	if rawLink == "" {
		return linkPostDraft{}, false
	}

	metadata, err := linkMetadataFetcher(rawLink)
	if err != nil {
		metadata = linkMetadata{}
	}

	baseTitle := firstNonEmpty(rawTitle, metadata.Title)
	if baseTitle == "" {
		baseTitle = rawLink
	}

	commentary := extractCommentaryFromListItem(li)
	if len(commentary) == 0 {
		commentary = []string{"Mentioned in the newsletter."}
	}

	body := buildLinkBody(baseTitle, rawLink, commentary)
	imageURL := firstNonEmpty(metadata.ImageURL, youtubeThumbnailURL(rawLink))
	description := firstNonEmpty(metadata.Description, rawTitle)
	slug := slugFromTitle(baseTitle)
	hash := computeLinkHash(baseTitle, body, rawLink)

	return linkPostDraft{
		Title:       baseTitle,
		Slug:        slug,
		Date:        newsletter.Date,
		Link:        rawLink,
		Description: description,
		ImageURL:    imageURL,
		Body:        body,
		Newsletter:  newsletter.Slug,
		Hash:        hash,
	}, true
}

func extractCommentaryFromListItem(li *goquery.Selection) []string {
	commentary := make([]string, 0)

	li.ChildrenFiltered("p").Each(func(_ int, p *goquery.Selection) {
		if strings.TrimSpace(p.Text()) != "" {
			commentary = append(commentary, normalizeCommentaryText(p.Text()))
		}
	})

	li.ChildrenFiltered("ul, ol").Each(func(_ int, list *goquery.Selection) {
		list.ChildrenFiltered("li").Each(func(_ int, note *goquery.Selection) {
			if text := normalizeCommentaryText(note.Text()); text != "" {
				commentary = append(commentary, text)
			}
		})
	})

	return commentary
}

func buildLinkBody(title, rawLink string, lines []string) string {
	var buf strings.Builder
	if strings.TrimSpace(title) != "" && strings.TrimSpace(rawLink) != "" {
		buf.WriteString("My thoughts on [")
		buf.WriteString(strings.TrimSpace(title))
		buf.WriteString("](")
		buf.WriteString(strings.TrimSpace(rawLink))
		buf.WriteString("): ")
		buf.WriteString(strings.TrimSpace(title))
		buf.WriteString("\n\n")
	}
	buf.WriteString("## Commentary\n\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		buf.WriteString("- ")
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	return buf.String()
}

func renderNewsletterBodyToHTML(body string) (string, error) {
	var out bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(gmhtml.WithUnsafe()),
	)
	if err := md.Convert([]byte(body), &out); err != nil {
		return "", err
	}
	return out.String(), nil
}

func fetchLinkMetadata(rawURL string) (linkMetadata, error) {
	if rawURL == "" {
		return linkMetadata{}, fmt.Errorf("empty url")
	}

	if thumb := youtubeThumbnailURL(rawURL); thumb != "" {
		return linkMetadata{ImageURL: thumb}, nil
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return linkMetadata{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return linkMetadata{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return linkMetadata{}, err
	}

	meta := linkMetadata{
		Title:       cleanMetaValue(doc.Find(`meta[property="og:title"]`).AttrOr("content", "")),
		Description: cleanMetaValue(doc.Find(`meta[property="og:description"]`).AttrOr("content", "")),
		ImageURL:    cleanMetaValue(doc.Find(`meta[property="og:image"]`).AttrOr("content", "")),
	}

	if meta.Title == "" {
		meta.Title = cleanMetaValue(doc.Find("title").Text())
	}
	if meta.Description == "" {
		meta.Description = cleanMetaValue(doc.Find(`meta[name="description"]`).AttrOr("content", ""))
	}
	if meta.ImageURL == "" {
		meta.ImageURL = cleanMetaValue(doc.Find(`meta[name="twitter:image"]`).AttrOr("content", ""))
	}

	return meta, nil
}

func youtubeThumbnailURL(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	switch {
	case strings.Contains(parsed.Host, "youtu.be"):
		id := strings.Trim(strings.TrimPrefix(parsed.Path, "/"), "/")
		if id == "" {
			return ""
		}
		return "https://i.ytimg.com/vi/" + id + "/hqdefault.jpg"
	case strings.Contains(parsed.Host, "youtube.com"):
		id := parsed.Query().Get("v")
		if id == "" {
			return ""
		}
		return "https://i.ytimg.com/vi/" + id + "/hqdefault.jpg"
	default:
		return ""
	}
}

func extractYouTubeTitle(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	if strings.Contains(parsed.Host, "youtu.be") {
		return strings.TrimSpace(strings.TrimPrefix(parsed.Path, "/"))
	}
	return ""
}

func cleanMetaValue(value string) string {
	return strings.TrimSpace(value)
}

func normalizeCommentaryText(text string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(text)), " ")
}

func isReadOrWatchedHeading(text string) bool {
	normalized := strings.ToLower(strings.TrimSpace(text))
	normalized = strings.TrimSuffix(normalized, ":")
	switch normalized {
	case "read", "watch", "watched":
		return true
	default:
		return false
	}
}

func isHeadingNode(sel *goquery.Selection) bool {
	if sel.Length() == 0 {
		return false
	}
	switch strings.ToLower(sel.Nodes[0].Data) {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}

func isTopLevelListItem(sel *goquery.Selection) bool {
	if sel.Length() == 0 || len(sel.Nodes) == 0 {
		return false
	}
	node := sel.Nodes[0]
	if node.Parent == nil {
		return false
	}
	parentName := strings.ToLower(node.Parent.Data)
	if parentName != "ul" && parentName != "ol" {
		return false
	}
	if node.Parent.Parent != nil && strings.ToLower(node.Parent.Parent.Data) == "li" {
		return false
	}
	return true
}

func computeLinkHash(title, body, rawURL string) string {
	sum := sha256.Sum256([]byte(title + "\n" + body + "\n" + rawURL))
	return hex.EncodeToString(sum[:])
}

func defaultLinkOutputDir(newsletterOutputDir string) string {
	parent := filepath.Dir(newsletterOutputDir)
	if parent == "" {
		return "posts/links"
	}
	return filepath.Join(parent, "links")
}

func writeLinkPost(outputDir, createdFilesPath string, draft linkPostDraft) error {
	content := generateLinkFrontmatter(draft) + draft.Body
	filePath := filepath.Join(outputDir, draft.Slug+".md")
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return err
	}

	file, err := os.OpenFile(createdFilesPath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, filePath); err != nil {
		return err
	}

	fmt.Printf("✨ Created link post: %s\n", filePath)
	return nil
}

func generateLinkFrontmatter(draft linkPostDraft) string {
	return fmt.Sprintf(`---
title: "%s"
date: %s
link: "%s"
status: published
image_url: "%s"
source: newsletter
newsletter: %s
type: links
slug: %s
tags:
description: "%s"
hash: %s
---
`,
		escapeFrontmatterValue(draft.Title),
		draft.Date,
		escapeFrontmatterValue(draft.Link),
		escapeFrontmatterValue(draft.ImageURL),
		draft.Newsletter,
		draft.Slug,
		escapeFrontmatterValue(draft.Description),
		draft.Hash,
	)
}

func escapeFrontmatterValue(value string) string {
	return escapeQuotes(value)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func isHeadingElement(node *nethtml.Node) bool {
	if node == nil || node.Type != nethtml.ElementNode {
		return false
	}
	switch strings.ToLower(node.Data) {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}
