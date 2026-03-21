package plugins

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Mr-Destructive/meetgor.com/models"
)

type MissingLinksFiller struct {
	PluginName string
}

func (p *MissingLinksFiller) Name() string {
	return p.PluginName
}

func (p *MissingLinksFiller) Phase() Phase {
	return PhaseTransform
}

func (p *MissingLinksFiller) Requires() []string {
	return []string{"readPosts"}
}

func (p *MissingLinksFiller) AdminPolicy() AdminPolicy {
	return AdminSkip
}

// Execute fills in missing links for link-type posts from newsletters
func (p *MissingLinksFiller) Execute(ssg *models.SSG) error {
	missingLinkPosts := make([]*models.Post, 0)

	// Find all link posts missing the link field
	for i := range ssg.Posts {
		post := &ssg.Posts[i]
		if post.Frontmatter.Type != "links" {
			continue
		}

		// Check if link field is missing
		_, hasLink := post.Frontmatter.Extras["link"]
		if hasLink {
			continue
		}

		missingLinkPosts = append(missingLinkPosts, post)
	}

	if len(missingLinkPosts) == 0 {
		return nil
	}

	// Build newsletters map
	newslettersMap := make(map[string]string)
	newsletterDir := filepath.Join("posts", "newsletter")
	entries, err := os.ReadDir(newsletterDir)
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(newsletterDir, entry.Name())
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		content := string(fileBytes)
		name := strings.TrimSuffix(entry.Name(), ".md")
		newslettersMap[name] = content
	}

	// Try to fill missing links using newsletter sources
	filledCount := 0
	for _, post := range missingLinkPosts {
		// Extract source from post content if available
		source := p.extractSourceFromPost(post)
		if source == "" {
			continue
		}

		// Look for the newsletter
		if newsletterContent, exists := newslettersMap[source]; exists {
			link := p.extractLinkForPostFromNewsletter(post, newsletterContent)
			if link != "" {
				if post.Frontmatter.Extras == nil {
					post.Frontmatter.Extras = make(map[string]interface{})
				}
				post.Frontmatter.Extras["link"] = link
				filledCount++
			}
		}
	}

	if filledCount > 0 {
		fmt.Printf("Missing Links Filler: Filled %d missing links from newsletters\n", filledCount)
	}

	return nil
}

// extractSourceFromPost extracts the source newsletter reference from post content
func (p *MissingLinksFiller) extractSourceFromPost(post *models.Post) string {
	sourcePattern := `\*\*Source:\*\*\s+([\w\-]+)`
	re := regexp.MustCompile(sourcePattern)
	matches := re.FindStringSubmatch(string(post.Content))
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// extractLinkForPostFromNewsletter finds the URL for a post in its source newsletter
// Only searches within "Read/Watched/Links" type sections
func (p *MissingLinksFiller) extractLinkForPostFromNewsletter(post *models.Post, newsletterContent string) string {
	// Extract sections that are "Read", "Watched", "Links", or similar
	sections := p.extractReadWatchedSections(newsletterContent)
	
	if len(sections) == 0 {
		return ""
	}
	
	titleLower := strings.ToLower(post.Frontmatter.Title)
	slugLower := strings.ToLower(post.Frontmatter.Slug)
	
	// Search only within these sections
	for _, section := range sections {
		lines := strings.Split(section.Content, "\n")
		for i, line := range lines {
			lineLower := strings.ToLower(line)
			
			// Check if this line mentions the post title or slug
			if strings.Contains(lineLower, titleLower) || strings.Contains(lineLower, slugLower) {
				// Look for URL in nearby lines within the section
				for j := i - 2; j <= i+2; j++ {
					if j < 0 || j >= len(lines) {
						continue
					}
					
					url := p.extractURLFromLine(lines[j])
					if url != "" {
						return url
					}
				}
			}
		}
	}
	
	return ""
}

// Section represents a bounded section of the newsletter
type Section struct {
	Name    string
	Content string
	Start   int
	End     int
}

// extractReadWatchedSections extracts sections marked as Read, Watched, Links, etc.
func (p *MissingLinksFiller) extractReadWatchedSections(content string) []Section {
	var sections []Section
	lines := strings.Split(content, "\n")
	
	// Section header patterns to look for
	sectionPatterns := []string{
		"read",
		"watched",
		"links",
		"resources",
		"interesting",
		"articles",
		"bookmarks",
		"findings",
	}
	
	var currentSection *Section
	
	for i, line := range lines {
		lineLower := strings.ToLower(line)
		
		// Check if this is a markdown header (## or ###)
		isHeader := strings.HasPrefix(strings.TrimSpace(line), "##")
		
		if isHeader {
			// Check if header matches our section patterns
			isRelevantSection := false
			var sectionName string
			
			for _, pattern := range sectionPatterns {
				if strings.Contains(lineLower, pattern) {
					isRelevantSection = true
					sectionName = pattern
					break
				}
			}
			
			// Save previous section if exists
			if currentSection != nil {
				currentSection.End = i - 1
				sections = append(sections, *currentSection)
			}
			
			// Start new section if relevant
			if isRelevantSection {
				currentSection = &Section{
					Name:  sectionName,
					Start: i,
					End:   len(lines) - 1,
				}
			} else {
				currentSection = nil
			}
		}
	}
	
	// Save the last section if exists
	if currentSection != nil {
		currentSection.End = len(lines) - 1
		// Extract actual content for the section
		startLine := currentSection.Start + 1 // Skip the header line
		endLine := currentSection.End
		
		if startLine <= endLine {
			currentSection.Content = strings.Join(lines[startLine : endLine+1], "\n")
			sections = append(sections, *currentSection)
		}
	}
	
	// Extract content for all sections
	for i := range sections {
		startLine := sections[i].Start + 1 // Skip header
		endLine := sections[i].End
		
		if startLine <= endLine {
			sections[i].Content = strings.Join(lines[startLine : endLine+1], "\n")
		}
	}
	
	return sections
}

// extractURLFromLine extracts a URL from a line of text
func (p *MissingLinksFiller) extractURLFromLine(line string) string {
	// Pattern for markdown link [text](url)
	markdownPattern := `\[([^\]]+)\]\(([^)]+)\)`
	re := regexp.MustCompile(markdownPattern)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 2 && strings.HasPrefix(matches[2], "http") {
		return matches[2]
	}
	
	// Pattern for plain URL
	urlPattern := `(https?://[^\s\)>\]]+)`
	re = regexp.MustCompile(urlPattern)
	matches = re.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	
	return ""
}


func init() {
	RegisterPlugin("MissingLinksFiller", func() Plugin {
		return &MissingLinksFiller{PluginName: "MissingLinksFiller"}
	})
}
