package plugins

import (
	"bytes"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	libsqlssg "github.com/Mr-Destructive/meetgor.com/plugins/db/libsqlssg"

	models "github.com/Mr-Destructive/meetgor.com/models"
)

type DbPlugin struct {
	PluginName string
}

func (p *DbPlugin) Name() string {
	return p.PluginName
}

func (p *DbPlugin) Phase() Phase {
	return PhasePostProcess
}

func (p *DbPlugin) Requires() []string {
	return []string{"renderTemplates"}
}

func (p *DbPlugin) AdminPolicy() AdminPolicy {
	return AdminSkip
}

func (p *DbPlugin) Execute(ssg *models.SSG) error {
	log.Println("------Executing DB plugin")

	buffer := bytes.Buffer{}
	postContext := models.TemplateContext{
		Themes: models.ThemeCombo{
			Default:   ssg.Config.Blog.Themes["default"],
			Secondary: ssg.Config.Blog.Themes["secondary"],
		},
		Config: models.SSG_CONFIG{
			Blog: ssg.Config.Blog,
		},
		Years: YearsFromPosts(ssg.Posts),
	}
	err := ssg.TemplateFS.ExecuteTemplate(&buffer, "editor_template.html", postContext)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(".", ssg.Config.Blog.OutputDir, "editor"), os.ModePerm)
	outputPath := filepath.Join(".", ssg.Config.Blog.OutputDir, "editor", "index.html")
	err = os.WriteFile(outputPath, buffer.Bytes(), 0660)
	if err != nil {
		return err
	}
	/*
		dbURL := os.Getenv("TURSO_DATABASE_NAME")
		dbAuthToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
		dbUrl := fmt.Sprintf("%s?authToken=%s", dbURL, dbAuthToken)

		db, err := sql.Open("libsql", dbUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
			os.Exit(1)
		}
		ctx := context.Background()
		if s, err := db.ExecContext(ctx, DDL); err != nil {
			log.Println(s)
		}
		defer db.Close()
		queries := libsqlssg.New(db)
		dbPosts, err := queries.GetAllPosts(ctx)
		log.Printf("Found %d posts in DB", len(dbPosts))
		postsToCreate := []libsqlssg.CreatePostParams{}
		var authorId int64
		for _, post := range ssg.Posts {
			slug := post.Frontmatter.Slug
			postExists := false
			for _, dbPost := range dbPosts {
				authorId = dbPost.AuthorID
				if dbPost.Slug == slug {
					postExists = true
					continue
				} else {
				}
			}
			var metadata []byte
			if postExists {
				continue
			}
			metadata, err = json.Marshal(post.Frontmatter.Extras)
			if err != nil {
				fatalf(err)
			}

			postParams := libsqlssg.CreatePostParams{
				Slug:     slug,
				Title:    post.Frontmatter.Title,
				Body:     string(post.Content),
				Metadata: string(metadata),
				AuthorID: int64(authorId),
			}
			postsToCreate = append(postsToCreate, postParams)
		}
		createdPosts := []libsqlssg.Post{}
		for _, post := range postsToCreate {
			createdPost, err := queries.CreatePost(ctx, post)
			if err != nil {
				fatalf(err)
			}
			createdPosts = append(createdPosts, createdPost)
		}
		log.Printf("Created %d posts", len(createdPosts))
	*/
	return nil
}

type Payload struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Title    string                 `json:"title"`
	Post     string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

//go:embed db/schema.sql
var DDL string

func CreatePostPayload(payload Payload, authorId int, authorName string) (libsqlssg.CreatePostParams, error) {

	nilPost := libsqlssg.CreatePostParams{}
	metadata := payload.Metadata
	log.Printf("Recieved payload metadata: %+v", payload.Metadata)
	content := payload.Post
	if payload.Title == "" {
		return nilPost, fmt.Errorf("Post must contain a title")
	} else {
		metadata["title"] = payload.Title
	}
	title := metadata["title"].(string)
	var slug string
	if metadata["slug"] != nil {
		slug = metadata["slug"].(string)
	} else {
		slug = Slugify(title)
	}

	var postType string
	if val, ok := metadata["type"]; ok {
		postType = val.(string)
	} else {
		postType = "post"
	}
	var postDir string
	if val, ok := metadata["post_dir"]; ok {
		postDir = val.(string)
	} else {
		postDir = "posts"
	}
	var status string
	if val, ok := metadata["published"]; ok {
		status = val.(string)
	} else {
		status = "published"
	}
	metadata["type"] = postType
	metadata["published"] = status
	metadata["author"] = authorName
	metadata["date"] = time.Now().Format("2006-01-02")
	metadata["post_dir"] = postDir
	log.Printf("final metadata: %+v", metadata)
	metadataStr, err := json.Marshal(metadata)
	if err != nil {
		return nilPost, err
	}

	dbPost := libsqlssg.CreatePostParams{
		ID:       generatePostID(),
		TypeID:   metadata["type"].(string),
		Title:    title,
		Slug:     slug,
		Content:  content,
		Metadata: sql.NullString{String: string(metadataStr), Valid: true},
	}
	return dbPost, nil
}

func Slugify(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	processedString = strings.TrimSpace(processedString)
	slug := strings.ReplaceAll(processedString, " ", "-")
	slug = strings.ToLower(slug)
	return slug
}

func CleanPostFrontmatter(post *models.Post, ssg *models.SSG) {
	if post.Frontmatter.Type == "" {
		post.Frontmatter.Type = "posts"
	}

	if post.Frontmatter.Slug == "" || post.Frontmatter.Title == "" {
		if post.Frontmatter.Title == "" {
			if post.Frontmatter.Description == "" {
				post.Frontmatter.Description = string(post.Content)[:15]
			}
			post.Frontmatter.Title = post.Frontmatter.Description
			post.Frontmatter.Title = string(post.Frontmatter.Description)[:15]
		} else if post.Frontmatter.Slug == "" {
			post.Frontmatter.Slug = Slugify(post.Frontmatter.Title)
		}
	}
	slug := post.Frontmatter.Slug
	if ssg.Config.Blog.PrefixURL != "" {
		slug = strings.TrimPrefix(slug, ssg.Config.Blog.PrefixURL)
	}
	slug = strings.TrimPrefix(slug, "/")
	typePrefix := post.Frontmatter.Type + "/"
	if strings.HasPrefix(slug, typePrefix) {
		slug = strings.TrimPrefix(slug, typePrefix)
	}
	post.Frontmatter.Slug = fmt.Sprintf("%s%s/%s", ssg.Config.Blog.PrefixURL, post.Frontmatter.Type, slug)

	if post.Frontmatter.Date == "" {
		post.Frontmatter.Date = time.Now().Format("2006-01-02")
	}
}

func generatePostID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func init() {
	RegisterPlugin("Db", func() Plugin {
		return &DbPlugin{PluginName: "Db"}
	})
}
