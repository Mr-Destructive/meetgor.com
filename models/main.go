package models

import (
	"html/template"
	"io/fs"
)

const SSG_CONFIG_FILE_NAME string = "ssg.json"

type Author struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Github string `json:"github"`
}

type PageConfig struct {
	TemplatePath     string `json:"template"`
	FeedTemplatePath string `json:"feed_template"`
	Emoji            string `json:"emoji"`
}

type Theme struct {
	Bg            string `json:"bg"`
	Text          string `json:"text"`
	SecondaryText string `json:"secondary-text"`
	Link          struct {
		Normal string `json:"normal"`
		Hover  string `json:"hover"`
		Active string `json:"active"`
	} `json:"link"`
	Quotes     string `json:"quotes"`
	CodeBlocks struct {
		Bg     string `json:"bg"`
		Border string `json:"border"`
	} `json:"codeblocks"`
	Code struct {
		Text     string `json:"text"`
		Comment  string `json:"comment"`
		Keyword  string `json:"keyword"`
		String   string `json:"string"`
		Number   string `json:"number"`
		Variable string `json:"variable"`
		Function string `json:"function"`
	} `json:"code"`
}

type BlogConfig struct {
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	BaseUrl             string                `json:"base_url"`
	PostsDir            string                `json:"posts_dir"`
	TemplatesDir        string                `json:"templates_dir"`
	StaticDir           string                `json:"static_dir"`
	OutputDir           string                `json:"output_dir"`
	AdminDir            string                `json:"admin_dir"`
	DefaultFeedTemplate string                `json:"default_feed_template"`
	DefaultPostTemplate string                `json:"default_post_template"`
	PrefixURL           string                `json:"prefix_url"`
	PagesConfig         map[string]PageConfig `json:"pages"`
	Themes              map[string]Theme      `json:"themes"`
	Github              map[string]string     `json:"github"`
	CloudFunction       map[string]string     `json:"cloud_function"`
}

type SSG_CONFIG struct {
	Blog      BlogConfig `json:"blog"`
	Authors   []Author   `json:"authors"`
	Plugins   []string   `json:"plugins"`
	AdminMode bool
}

var config *SSG_CONFIG

type FrontMatter struct {
	Title       string                 `json:"title" yaml:"title"`
	Description string                 `json:"description" yaml:"description"`
	Status      string                 `json:"status" yaml:"status"`
	Type        string                 `json:"type" yaml:"type"`
	Date        string                 `json:"date" yaml:"date"`
	Slug        string                 `json:"slug" yaml:"slug"`
	Tags        []string               `json:"tags" yaml:"tags"`
	ImageUrl    string                 `json:"image_url" yaml:"image_url"`
	Extras      map[string]interface{} `json:",inline" yaml:",inline"`
}

type PostType int

const (
	POST PostType = iota
	TIL
	PROJECT
)

var PostTypes = map[PostType]string{
	POST:    "post",
	TIL:     "til",
	PROJECT: "project",
}

type Post struct {
	Frontmatter FrontMatter
	Content     template.HTML
	Markdown    string
}

type DBPost struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Slug        string `json:"slug"`
    Body        string `json:"body"`
    Metadata    string `json:"metadata"`
    Deleted     bool   `json:"deleted"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    AuthorID    int    `json:"author_id"`
}

type Feed struct {
	Title string
	Type  string
	Slug  string
	Posts []Post
}

type SSG struct {
	Config     SSG_CONFIG
	Posts      []Post
	FeedPosts  []Feed
	TemplateFS *template.Template
	FS         fs.FS
}

type ThemeCombo struct {
	Default   Theme
	Secondary Theme
}
type TemplateContext struct {
	Themes    ThemeCombo
	Config    SSG_CONFIG
	Post      Post
	FeedPosts []Feed
	FeedInfo  Feed
}
