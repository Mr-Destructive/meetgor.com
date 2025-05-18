package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mr-destructive/mr-destructive.github.io/plugins"
	"github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

var editForm string = `
<form id="postForm">
    <div class="mb-4">
        <label for="title" class="block text-lg font-medium">Title:</label>
        <input type="text" value="{{ .Title }}" name="title" id="title" class="w-full p-2 border rounded-md shadow-sm" required>
    </div>

    <div class="mb-4">
        <label for="metadata" class="block text-lg font-medium">Metadata (JSON):</label>
        <textarea name="metadata" id="metadata" rows="4" class="w-full p-2 border rounded-md shadow-sm" placeholder='{"key": "value"}'> {{ .Metadata }}</textarea>
    </div>

    <div class="mb-4">
        <label for="content" class="block text-lg font-medium">Body (Markdown):</label>
        <textarea name="content" id="content" rows="6" class="w-full p-2 border rounded-md shadow-sm">{{ .Post }}</textarea>
    </div>

    <div class="mb-4">
        <label for="username" class="block text-lg font-medium">Username:</label>
        <input type="text" name="username" id="username" class="w-full p-2 border rounded-md shadow-sm" required>
    </div>

    <div class="mb-6">
        <label for="password" class="block text-lg font-medium">Password:</label>
        <input type="password" name="password" id="password" class="w-full p-2 border rounded-md shadow-sm" required>
    </div>

    <button type="submit" class="w-full p-3 bg-blue-500 text-white rounded-md shadow-lg focus:outline-none focus:ring-2 hover:bg-blue-600">Submit</button>
</form>
`

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
			},
			Body: "",
		}, nil
	}
	ctx := context.Background()
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	var err error
	dbString := fmt.Sprintf("libsql://%s?authToken=%s", dbName, dbToken)
	db, err := sql.Open("libsql", dbString)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}
	defer db.Close()

	queries := libsqlssg.New(db)
	if _, err := db.ExecContext(ctx, plugins.DDL); err != nil {
		return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}

	queryParams := make(map[string]string)
	if req.QueryStringParameters != nil {
		queryParams = req.QueryStringParameters
	}

	var payload plugins.Payload
	log.Printf("Headers: %v", req.Headers)
	log.Printf("hx-request??? %v", req.Headers["hx-request"])
	if queryParams["method"] == "edit" {
		if req.HTTPMethod == http.MethodGet {
			var prefixURL string = ""
			if prefixURL, ok := queryParams["prefixURL"]; ok {
				prefixURL = prefixURL
			}
			slug := fmt.Sprintf("%s%s/%s", prefixURL, queryParams["type"], queryParams["slug"])
			log.Printf("Slug: %s", slug)
			postsBySlug, err := queries.GetPostsBySlugType(ctx, slug)
			postType := queryParams["type"]
			if len(postsBySlug) == 0 {
				slug = strings.TrimPrefix(queryParams["slug"], "/")
				slug = strings.TrimPrefix(slug, postType)
				slug = strings.TrimPrefix(slug, "/")
				postsBySlug, err = queries.GetPostsBySlugType(ctx, slug)
				log.Printf("Slug: %s", slug)
			}
			for _, post := range postsBySlug {
				metadataObj := make(map[string]interface{})
				err = json.Unmarshal([]byte(post.Metadata), &metadataObj)
				log.Printf("Metadata: %v", metadataObj)
				log.Printf("Error: %v", err)
				if err != nil {
					return errorResponse(http.StatusInternalServerError, "Invalid metadata Payload"), nil
				}
				markdown, err := htmltomarkdown.ConvertString(post.Body)
				log.Printf("Markdown: %v", markdown)
				log.Printf("Error: %v", err)
				if err != nil {
					return errorResponse(http.StatusInternalServerError, "Invalid metadata Payload"), nil
				}
				//if metadataObj["type"] == postType {
				payload = plugins.Payload{
					Title:    post.Title,
					Metadata: metadataObj,
					Post:     markdown,
				}
				templ := template.Must(template.New("editForm").Parse(editForm))
				buffer := new(bytes.Buffer)
				templ.Execute(buffer, payload)
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusOK,
					Headers: map[string]string{
						"Access-Control-Allow-Origin":  "*",
						"Access-Control-Allow-Methods": "POST, OPTIONS",
						"Access-Control-Allow-Headers": "Content-Type, Authorization",
					},
					Body:            buffer.String(),
					IsBase64Encoded: false,
				}, nil
			}
			if err != nil {
				return errorResponse(http.StatusInternalServerError, "Post Not Found"), nil
			}
		} else if req.HTTPMethod == http.MethodPost {
			err = json.Unmarshal([]byte(req.Body), &payload)
			if err != nil {
				return errorResponse(http.StatusInternalServerError, "Invalid Payload"), nil
			}
		}
	}
	if req.Headers["hx-request"] == "true" {
		formData, err := url.ParseQuery(req.Body)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Invalid form Payload"), nil
		}
		metadata := make(map[string]interface{})
		err = json.Unmarshal([]byte(formData.Get("metadata")), &metadata)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Invalid metadata Payload"), nil
		}
		payload = plugins.Payload{
			Username: formData.Get("username"),
			Password: formData.Get("password"),
			Title:    formData.Get("title"),
			Post:     formData.Get("content"),
			Metadata: metadata,
		}
	} else {
		err = json.Unmarshal([]byte(req.Body), &payload)
		if err != nil {
			return errorResponse(http.StatusInternalServerError, "Invalid Payload"), nil
		}
	}
	log.Printf("Payload: %v", payload)
	user, err := queries.GetUser(ctx, payload.Username)
	log.Printf("User: %v", user)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "User Not Found"), nil
	}
	if !Authenticate(payload.Username, user.Password, payload.Password) {
		return errorResponse(http.StatusInternalServerError, "Authentication Failed"), nil
	}

	post, err := plugins.CreatePostPayload(payload, int(user.ID), user.Name)
	log.Printf("Post: %v", post)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}
	_, err = queries.CreatePost(ctx, post)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}

	if req.Headers["hx-request"] == "true" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type":                 "text/html",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
			},
			Body: `<div class="success-message">Post created successfully!</div>`,
		}, nil
	}

	return jsonResponse(http.StatusOK, post), nil
}

func Authenticate(username, hashedPassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	fmt.Println(err)
	if err != nil {
		fmt.Println("Authentication Failure")
		return false
	}
	return true
}
func jsonResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
		},
		Body: string(body),
	}
}

func errorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return jsonResponse(statusCode, map[string]string{"error": message})
}
