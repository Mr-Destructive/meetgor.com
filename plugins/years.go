package plugins

import (
	"sort"
	"time"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

func YearsFromPosts(posts []models.Post) []string {
	yearSet := make(map[string]struct{})
	for _, post := range posts {
		if post.Frontmatter.Status == "draft" {
			continue
		}
		if post.Frontmatter.Date == "" {
			continue
		}
		parsed, err := time.Parse("2006-01-02", post.Frontmatter.Date)
		if err != nil {
			continue
		}
		if parsed.Year() < 2000 {
			continue
		}
		yearSet[parsed.Format("2006")] = struct{}{}
	}

	years := make([]string, 0, len(yearSet))
	for year := range yearSet {
		years = append(years, year)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(years)))
	return years
}
