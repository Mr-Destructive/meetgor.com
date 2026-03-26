package plugins

import "strings"

// NormalizePostType maps legacy or editor-facing aliases onto canonical post types.
func NormalizePostType(postType string) string {
	normalized := strings.ToLower(strings.TrimSpace(postType))
	switch normalized {
	case "", "post", "blog":
		return "posts"
	default:
		return normalized
	}
}
