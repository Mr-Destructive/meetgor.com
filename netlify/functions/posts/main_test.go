package main

import "testing"

func TestResolvePostTypeNormalizesLegacyAliases(t *testing.T) {
	if got := resolvePostType(postPayload{TypeID: "blog"}); got != "posts" {
		t.Fatalf("expected blog to normalize to posts, got %q", got)
	}
	if got := resolvePostType(postPayload{Metadata: map[string]interface{}{"type": "post"}}); got != "posts" {
		t.Fatalf("expected post to normalize to posts, got %q", got)
	}
	if got := resolvePostType(postPayload{TypeID: "til"}); got != "til" {
		t.Fatalf("expected til to stay til, got %q", got)
	}
}

func TestValidateCreateAllowsLinksWithoutContent(t *testing.T) {
	payload := postPayload{
		Username: "u",
		Password: "p",
		Title:    "Link",
		Metadata: map[string]interface{}{
			"type": "links",
		},
	}
	if err := validateCreate(payload); err != nil {
		t.Fatalf("expected links post to validate without content: %v", err)
	}
}
