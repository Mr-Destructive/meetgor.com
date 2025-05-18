package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	dbURL := os.Getenv("TURSO_DATABASE_NAME")
	dbAuthToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	dbUrl := fmt.Sprintf("%s?authToken=%s", dbURL, dbAuthToken)
	args := os.Args
	sync_post := false
	if len(args) > 1 {
		if args[1] == "sync_post" {
			sync_post = true
		}
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()
	var query string
	if sync_post {
		query = "SELECT * FROM posts;"
	} else {
		onehourBackTime := time.Now().Add(time.Hour * -1).Format("2006-01-02 15:04:05")
		query = fmt.Sprintf("SELECT * FROM posts WHERE created_at > '%s';", onehourBackTime)
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var slug string
		var body string
		var created string
		var updated string
		var metadata string
		var authorId int64
		var deleted bool
		err := rows.Scan(&id, &title, &slug, &body, &metadata, &created, &updated, &authorId, &deleted)
		if err != nil {
			fmt.Println(err)
		}
		bodyMd, err := htmltomarkdown.ConvertString(body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, title, slug, body, created, updated, metadata, authorId)
		writePostFile(title, slug, bodyMd, metadata, created, updated)
	}
}

func writePostFile(title, slug, body, metadataStr, created, updated string) {
	metadata := make(map[string]interface{})
	fmt.Println(metadataStr)
	err := json.Unmarshal([]byte(metadataStr), &metadata)
	fmt.Println(metadata)
	if err != nil {
		panic(err)
	}
	var postDir string
	var baseDir string = "posts/"
	if val, ok := metadata["post_dir"]; ok {
		postDir = val.(string)
	} else {
		postDir = "posts"
	}
	postDir = baseDir + postDir
	//create folder if not exists
	if _, err := os.Stat(postDir); os.IsNotExist(err) {
		os.Mkdir(postDir, 0777)
	}
	_, ok := metadata["slug"]
	if !ok {
		slug = Slugify(title)
	}
	filePath := fmt.Sprintf("%s/%s.md", postDir, slug)
	fileContent := fmt.Sprintf("%s\n\n%s", metadataStr, body)
	os.WriteFile(filePath, []byte(fileContent), 0660)

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
