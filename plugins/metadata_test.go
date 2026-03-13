package plugins

import "testing"

type pluginMeta interface {
	Name() string
	Phase() Phase
	Requires() []string
	AdminPolicy() AdminPolicy
}

func TestPluginMetadataAccessors(t *testing.T) {
	tests := []struct {
		name   string
		plugin pluginMeta
	}{
		{"Tags", &TagsPlugin{PluginName: "Tags"}},
		{"Series", &SeriesPlugin{PluginName: "Series"}},
		{"YearWise", &YearPlugin{PluginName: "YearWise"}},
		{"Sitemap", &SitemapPlugin{PluginName: "Sitemap"}},
		{"RSS", &RSSPlugin{PluginName: "RSS"}},
		{"SeasonalEffects", &SeasonalEffectsPlugin{PluginName: "SeasonalEffects"}},
		{"SQLSearch", &SQLSearchPlugin{PluginName: "SQLSearch"}},
		{"StaticPageRenderer", &StaticPageRendererPlugin{PluginName: "StaticPageRenderer"}},
		{"SyncDB", &SyncDBPostsPlugin{PluginName: "SyncDB"}},
		{"LinkPreview", &LinkPreviewPlugin{PluginName: "LinkPreview"}},
		{"Db", &DbPlugin{PluginName: "Db"}},
		{"BeforeAfterPosts", &BeforeAfterPostsPlugin{PluginName: "BeforeAfterPosts"}},
	}

	for _, tt := range tests {
		if tt.plugin.Name() == "" {
			t.Fatalf("plugin %s returned empty Name", tt.name)
		}
		if tt.plugin.Phase() == "" {
			t.Fatalf("plugin %s returned empty Phase", tt.name)
		}
		_ = tt.plugin.Requires()
		_ = tt.plugin.AdminPolicy()
	}
}
