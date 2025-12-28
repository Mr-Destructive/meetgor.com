package plugins

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type BeforeAfterPostsPlugin struct {
	PluginName string
}

func (p *BeforeAfterPostsPlugin) Name() string {
	return p.PluginName
}

// PostNavigation contains previous and next post information
type PostNavigation struct {
	Previous *models.Post
	Next     *models.Post
	Related  []models.Post
}

func (p *BeforeAfterPostsPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config

	// Add navigation data to each post
	for i := range ssg.Posts {
		if ssg.Posts[i].Frontmatter.Status == "draft" {
			continue
		}

		nav := p.findNavigationPosts(&ssg.Posts[i], ssg.Posts)

		// Store navigation in Extras for template access
		if ssg.Posts[i].Frontmatter.Extras == nil {
			ssg.Posts[i].Frontmatter.Extras = make(map[string]interface{})
		}

		// Add constructed slugs to navigation posts (since they haven't been set yet)
		if nav.Previous != nil {
			nav.Previous = p.enrichPostWithSlug(nav.Previous, config.Blog.PrefixURL, config.Blog.DefaultPostTemplate)
		}
		if nav.Next != nil {
			nav.Next = p.enrichPostWithSlug(nav.Next, config.Blog.PrefixURL, config.Blog.DefaultPostTemplate)
		}
		for j := range nav.Related {
			nav.Related[j] = *p.enrichPostWithSlug(&nav.Related[j], config.Blog.PrefixURL, config.Blog.DefaultPostTemplate)
		}

		ssg.Posts[i].Frontmatter.Extras["nav_previous"] = nav.Previous
		ssg.Posts[i].Frontmatter.Extras["nav_next"] = nav.Next
		ssg.Posts[i].Frontmatter.Extras["nav_related"] = nav.Related
	}

	fmt.Println("Before/After Posts plugin executed successfully")
}

// enrichPostWithSlug constructs the slug for a post since it hasn't been set yet during plugin execution
func (p *BeforeAfterPostsPlugin) enrichPostWithSlug(post *models.Post, prefixURL string, defaultPostTemplate string) *models.Post {
	if post == nil {
		return nil
	}

	// Make a copy to avoid modifying the original
	postCopy := *post

	// Determine the post type
	postType := postCopy.Frontmatter.Type
	if postType == "" {
		postType = "posts"
	}

	// Get the slug name (remove any existing path prefixes)
	slugName := postCopy.Frontmatter.Slug
	if slugName == "" {
		slugName = Slugify(postCopy.Frontmatter.Title)
	}

	// Construct the full slug with type prefix
	postCopy.Frontmatter.Slug = prefixURL + postType + "/" + slugName

	return &postCopy
}

// findNavigationPosts finds previous, next, and related posts
func (p *BeforeAfterPostsPlugin) findNavigationPosts(currentPost *models.Post, allPosts []models.Post) PostNavigation {
	nav := PostNavigation{
		Related: []models.Post{},
	}

	// Create a sorted copy of posts by date
	sortedPosts := make([]models.Post, 0)
	for _, post := range allPosts {
		if post.Frontmatter.Status != "draft" && post.Frontmatter.Date != "" {
			sortedPosts = append(sortedPosts, post)
		}
	}

	// Sort by date (newest first)
	sort.Slice(sortedPosts, func(i, j int) bool {
		date1, _ := parseDate(sortedPosts[i].Frontmatter.Date)
		date2, _ := parseDate(sortedPosts[j].Frontmatter.Date)
		return date1.After(date2)
	})

	// Find current post index
	currentIndex := -1
	for idx, post := range sortedPosts {
		if post.Frontmatter.Slug == currentPost.Frontmatter.Slug {
			currentIndex = idx
			break
		}
	}

	if currentIndex == -1 {
		// Post not found, return empty navigation
		return nav
	}

	// Check if post belongs to a series
	if currentPost.Frontmatter.Extras != nil && currentPost.Frontmatter.Extras["series"] != nil {
		seriesNav := p.findSeriesNavigation(currentPost, sortedPosts)
		if seriesNav.Previous != nil || seriesNav.Next != nil {
			nav.Previous = seriesNav.Previous
			nav.Next = seriesNav.Next
		}
	} else {
		// If no series, find posts of same type by date
		nav.Previous, nav.Next = p.findSameTypeNavigation(currentPost, sortedPosts, currentIndex)
	}

	// Find related posts by tags (always, regardless of series)
	nav.Related = p.findRelatedPostsByTags(currentPost, sortedPosts)

	return nav
}

// findSeriesNavigation finds previous and next posts in the same series
func (p *BeforeAfterPostsPlugin) findSeriesNavigation(currentPost *models.Post, sortedPosts []models.Post) PostNavigation {
	nav := PostNavigation{}

	// Get series from current post
	var currentSeries string
	if seriesRaw, ok := currentPost.Frontmatter.Extras["series"].(string); ok {
		currentSeries = Slugify(seriesRaw)
	} else if seriesArray, ok := currentPost.Frontmatter.Extras["series"].([]interface{}); ok && len(seriesArray) > 0 {
		currentSeries = Slugify(seriesArray[0].(string))
	}

	if currentSeries == "" {
		return nav
	}

	// Filter posts in same series
	var seriesPosts []models.Post
	for _, post := range sortedPosts {
		if post.Frontmatter.Extras != nil && post.Frontmatter.Extras["series"] != nil {
			var postSeries string
			if seriesRaw, ok := post.Frontmatter.Extras["series"].(string); ok {
				postSeries = Slugify(seriesRaw)
			} else if seriesArray, ok := post.Frontmatter.Extras["series"].([]interface{}); ok && len(seriesArray) > 0 {
				postSeries = Slugify(seriesArray[0].(string))
			}

			if postSeries == currentSeries && post.Frontmatter.Slug != currentPost.Frontmatter.Slug {
				seriesPosts = append(seriesPosts, post)
			}
		}
	}

	// Sort series posts by date (oldest first for chronological order)
	sort.Slice(seriesPosts, func(i, j int) bool {
		date1, _ := parseDate(seriesPosts[i].Frontmatter.Date)
		date2, _ := parseDate(seriesPosts[j].Frontmatter.Date)
		return date1.Before(date2)
	})

	// Find current post position in series
	currentDate, _ := parseDate(currentPost.Frontmatter.Date)
	for _, post := range seriesPosts {
		postDate, _ := parseDate(post.Frontmatter.Date)

		// Previous: last post before current
		if postDate.Before(currentDate) {
			nav.Previous = &post
		}
		// Next: first post after current
		if postDate.After(currentDate) && nav.Next == nil {
			nav.Next = &post
		}
	}

	return nav
}

// findSameTypeNavigation finds previous and next posts of the same type by date
func (p *BeforeAfterPostsPlugin) findSameTypeNavigation(currentPost *models.Post, sortedPosts []models.Post, currentIndex int) (*models.Post, *models.Post) {
	var previous, next *models.Post

	currentType := currentPost.Frontmatter.Type
	if currentType == "" {
		currentType = "posts"
	}

	// Find previous post of same type
	for i := currentIndex + 1; i < len(sortedPosts); i++ {
		postType := sortedPosts[i].Frontmatter.Type
		if postType == "" {
			postType = "posts"
		}
		if postType == currentType {
			previous = &sortedPosts[i]
			break
		}
	}

	// Find next post of same type
	for i := currentIndex - 1; i >= 0; i-- {
		postType := sortedPosts[i].Frontmatter.Type
		if postType == "" {
			postType = "posts"
		}
		if postType == currentType {
			next = &sortedPosts[i]
			break
		}
	}

	return previous, next
}

// findRelatedPostsByTags finds posts with similar tags
func (p *BeforeAfterPostsPlugin) findRelatedPostsByTags(currentPost *models.Post, allPosts []models.Post) []models.Post {
	if len(currentPost.Frontmatter.Tags) == 0 {
		return []models.Post{}
	}

	type PostScore struct {
		Post  models.Post
		Score int
	}

	var scoredPosts []PostScore

	for _, post := range allPosts {
		if post.Frontmatter.Slug == currentPost.Frontmatter.Slug {
			continue
		}

		// Calculate tag match score
		score := 0
		for _, currentTag := range currentPost.Frontmatter.Tags {
			for _, postTag := range post.Frontmatter.Tags {
				if strings.EqualFold(currentTag, postTag) {
					score++
				}
			}
		}

		if score > 0 {
			scoredPosts = append(scoredPosts, PostScore{Post: post, Score: score})
		}
	}

	// Sort by score (highest first), then by date (newest first)
	sort.Slice(scoredPosts, func(i, j int) bool {
		if scoredPosts[i].Score != scoredPosts[j].Score {
			return scoredPosts[i].Score > scoredPosts[j].Score
		}
		date1, _ := parseDate(scoredPosts[i].Post.Frontmatter.Date)
		date2, _ := parseDate(scoredPosts[j].Post.Frontmatter.Date)
		return date1.After(date2)
	})

	// Return top 5 related posts
	maxRelated := 5
	if len(scoredPosts) < maxRelated {
		maxRelated = len(scoredPosts)
	}

	related := make([]models.Post, maxRelated)
	for i := 0; i < maxRelated; i++ {
		related[i] = scoredPosts[i].Post
	}

	return related
}

// parseDate parses a date string in format "YYYY-MM-DD"
func parseDate(dateStr string) (time.Time, error) {
	if len(dateStr) < 10 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}
	return time.Parse("2006-01-02", dateStr[:10])
}

func init() {
	RegisterPlugin("BeforeAfterPosts", reflect.TypeOf(BeforeAfterPostsPlugin{
		PluginName: "BeforeAfterPosts",
	}))
}
