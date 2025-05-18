package plugins

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	libsqlssg "github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"

	models "github.com/mr-destructive/mr-destructive.github.io/models"
)

type DbPlugin struct {
	PluginName string
}

func (p *DbPlugin) Name() string {
	return p.PluginName
}

func (p *DbPlugin) Execute(ssg *models.SSG) {
	log.Println("------Executing DB plugin")

	buffer := bytes.Buffer{}
	templates, err := template.New("base").ParseFS(ssg.FS, "*.html")
	if err != nil {
		log.Fatal(err)
	}
	postContext := models.TemplateContext{
		Themes: models.ThemeCombo{
			Default:   ssg.Config.Blog.Themes["default"],
			Secondary: ssg.Config.Blog.Themes["secondary"],
		},
		Config: models.SSG_CONFIG{
			Blog: ssg.Config.Blog,
		},
	}
	err = templates.ExecuteTemplate(&buffer, "editor_template.html", postContext)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(filepath.Join(".", ssg.Config.Blog.OutputDir, "editor"), os.ModePerm)
	outputPath := filepath.Join(".", ssg.Config.Blog.OutputDir, "editor", "index.html")
	err = os.WriteFile(outputPath, buffer.Bytes(), 0660)
	if err != nil {
		log.Fatal(err)
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
				log.Fatal(err)
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
				log.Fatal(err)
			}
			createdPosts = append(createdPosts, createdPost)
		}
		log.Printf("Created %d posts", len(createdPosts))
	*/
}

func GetAllPostsSlug(posts []models.Post) []string {
	slugs := []string{}
	for _, post := range posts {
		slugs = append(slugs, post.Frontmatter.Slug)
	}
	return slugs
}

type Payload struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Title    string                 `json:"title"`
	Post     string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Author struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//go:embed db/schema.sql
var DDL string

func PostHandler(w http.ResponseWriter, r *http.Request) {

	dbURL := "http://127.0.0.1:8080" //os.Getenv("TURSO_DATABASE_URL")
	dbUrl := fmt.Sprintf("%s", dbURL)

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
	err = r.ParseForm()
	if err != nil {
		fmt.Errorf("error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//fmt.Println(string(body))
	//var payload Payload
	//err = json.Unmarshal(body, &payload)
	//if err != nil {
	//	fmt.Errorf("error: %s", err)
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//}
	metadata := make(map[string]interface{})
	err = json.Unmarshal([]byte(r.FormValue("metadata")), &metadata)
	payload := Payload{
		Title:    r.FormValue("title"),
		Metadata: metadata,
		Post:     r.FormValue("content"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	fmt.Println(payload)

	// authenticate
	fmt.Println(payload)
	author, err := queries.GetUser(ctx, payload.Username)
	err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(payload.Password))
	fmt.Println(err)
	if err != nil {
		fmt.Println("Authentication Failure")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// create post
	dbPost, err := CreatePostPayload(payload, int(author.ID), author.Name)
	fmt.Println("POST", dbPost)
	if err != nil {
		fmt.Println("Unable to construct the post", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("POST BODY:", dbPost.Body)
	createdPost, err := queries.CreatePost(ctx, dbPost)
	if err != nil {
		fmt.Println("Post Creation Failure")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(createdPost)
	// create file from post on github

}

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
		Title:    title,
		Slug:     slug,
		Body:     content,
		Metadata: string(metadataStr),
		AuthorID: int64(authorId),
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
	post.Frontmatter.Slug = fmt.Sprintf("%s%s/%s", ssg.Config.Blog.PrefixURL, post.Frontmatter.Type, post.Frontmatter.Slug)

	if post.Frontmatter.Date == "" {
		post.Frontmatter.Date = time.Now().Format("2006-01-02")
	}
}

func init() {
	RegisterPlugin("Db", reflect.TypeOf(DbPlugin{
		PluginName: "Db",
	}))
}
