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
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mr-destructive/mr-destructive.github.io/plugins"
	"github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

// AuthRequest defines the structure for authentication requests.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	ctx := context.Background()
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	netlifyTriggerSecret := os.Getenv("NETLIFY_TRIGGER_SECRET")
	appBaseURL := os.Getenv("APP_BASE_URL")

	if req.HTTPMethod == "OPTIONS" {
		log.Println("Handling OPTIONS request")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*", // Consider restricting this in production
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
			},
			Body: "",
		}, nil
	}

	var err error
	dbString := fmt.Sprintf("libsql://%s?authToken=%s", dbName, dbToken)
	db, err := sql.Open("libsql", dbString)
	if err != nil {
		log.Printf("Error in Connection: %v", err)
		return ErrorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}
	defer db.Close()

	queries := libsqlssg.New(db)
	if _, errDb := db.ExecContext(ctx, plugins.DDL); errDb != nil {
		log.Printf("Error in DDL: %v", errDb)
		// Not returning error here as table might already exist
	}

	log.Printf("Request Path: %s", req.Path)

	// Handle /api/trigger-build and /api/sync-db
	if strings.HasSuffix(req.Path, "/trigger-build") || strings.HasSuffix(req.Path, "/sync-db") {
		if req.HTTPMethod == http.MethodPost {
			var authReq AuthRequest
			err := json.Unmarshal([]byte(req.Body), &authReq)
			if err != nil {
				log.Printf("Error unmarshalling auth request: %v", err)
				return ErrorResponse(http.StatusBadRequest, "Invalid request payload"), nil
			}

			log.Printf("AuthRequest received for user: %s", authReq.Username)
			user, err := queries.GetUser(ctx, authReq.Username)
			if err != nil {
				log.Printf("Error fetching user %s: %v", authReq.Username, err)
				return ErrorResponse(http.StatusUnauthorized, "User not found or authentication failed"), nil
			}

			if !Authenticate(authReq.Username, user.Password, authReq.Password) {
				log.Printf("Authentication failed for user: %s", authReq.Username)
				return ErrorResponse(http.StatusUnauthorized, "Authentication failed"), nil
			}

			log.Printf("User %s authenticated successfully. Calling trigger build function.", authReq.Username)
			resp, err := callTriggerBuildFunction(appBaseURL, netlifyTriggerSecret)
			if err != nil {
				log.Printf("Error calling trigger build function: %v", err)
				return ErrorResponse(http.StatusInternalServerError, "Failed to trigger build"), nil
			}
			defer resp.Body.Close()

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				log.Printf("Trigger build function called successfully for user: %s, status: %d", authReq.Username, resp.StatusCode)
				return jsonResponse(http.StatusOK, map[string]string{"message": "Action triggered successfully"}), nil
			} else {
				log.Printf("Trigger build function call failed for user: %s, status: %d", authReq.Username, resp.StatusCode)
				return ErrorResponse(resp.StatusCode, "Failed to trigger build (upstream error)"), nil
			}
		} else {
			return ErrorResponse(http.StatusMethodNotAllowed, fmt.Sprintf("Method %s not allowed for this path", req.HTTPMethod)), nil
		}
	}

	// Existing logic for post creation
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
			if prefixURLVal, ok := queryParams["prefixURL"]; ok {
				prefixURL = prefixURLVal
			}
			slug := fmt.Sprintf("%s%s/%s", prefixURL, queryParams["type"], queryParams["slug"])
			log.Printf("Slug for edit: %s", slug)
			postsBySlug, err := queries.GetPostsBySlugType(ctx, slug)
			if err != nil {
				log.Printf("Error fetching post by slug %s: %v", slug, err)
				// Attempt fallback if needed
			}
			postType := queryParams["type"]
			if len(postsBySlug) == 0 {
				originalSlug := queryParams["slug"]
				newSlug := strings.TrimPrefix(originalSlug, "/")
				newSlug = strings.TrimPrefix(newSlug, postType)
				newSlug = strings.TrimPrefix(newSlug, "/")
				log.Printf("Attempting fallback slug: %s", newSlug)
				postsBySlug, err = queries.GetPostsBySlugType(ctx, newSlug)
				if err != nil {
					log.Printf("Error fetching post by fallback slug %s: %v", newSlug, err)
				}
			}

			if len(postsBySlug) > 0 {
				post := postsBySlug[0] // Assuming first one is the target
				metadataObj := make(map[string]interface{})
				err = json.Unmarshal([]byte(post.Metadata), &metadataObj)
				if err != nil {
					log.Printf("Error unmarshalling metadata for post %s: %v", post.Title, err)
					return ErrorResponse(http.StatusInternalServerError, "Invalid metadata in stored post"), nil
				}
				markdown := post.Body
				//markdown, errMd := htmltomarkdown.ConvertString(post.Body)
				//if errMd != nil {
				//	log.Printf("Error converting HTML to markdown for post %s: %v", post.Title, errMd)
				//	return ErrorResponse(http.StatusInternalServerError, "Error processing post content"), nil
				//}
				payload = plugins.Payload{
					Title:    post.Title,
					Metadata: metadataObj,
					Post:     markdown,
				}
				templ := template.Must(template.New("editForm").Parse(editForm))
				buffer := new(bytes.Buffer)
				if err := templ.Execute(buffer, payload); err != nil {
					log.Printf("Error executing template for edit form: %v", err)
					return ErrorResponse(http.StatusInternalServerError, "Error generating edit form"), nil
				}
				return events.APIGatewayProxyResponse{
					StatusCode: http.StatusOK,
					Headers: map[string]string{
						"Access-Control-Allow-Origin":  "*",
						"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
						"Access-Control-Allow-Headers": "Content-Type, Authorization",
						"Content-Type":                 "text/html",
					},
					Body:            buffer.String(),
					IsBase64Encoded: false,
				}, nil
			}
			log.Printf("No post found for slug: %s (type: %s)", queryParams["slug"], postType)
			return ErrorResponse(http.StatusNotFound, "Post not found"), nil

		} else if req.HTTPMethod == http.MethodPost { // This is for editing an existing post, typically from the edit form
			err = json.Unmarshal([]byte(req.Body), &payload)
			if err != nil {
				log.Printf("Error unmarshalling payload for edit post: %v", err)
				return ErrorResponse(http.StatusBadRequest, "Invalid payload for edit"), nil
			}
		}
	} else if req.Headers["hx-request"] == "true" { // HTMX request for new post
		formData, err := url.ParseQuery(req.Body)
		if err != nil {
			log.Printf("Error parsing form data from HTMX request: %v", err)
			return ErrorResponse(http.StatusBadRequest, "Invalid form payload"), nil
		}
		metadataStr := formData.Get("metadata")
		metadata := make(map[string]interface{})
		if metadataStr != "" {
			err = json.Unmarshal([]byte(metadataStr), &metadata)
			if err != nil {
				log.Printf("Error unmarshalling metadata from HTMX request: %v", err)
				return ErrorResponse(http.StatusBadRequest, "Invalid metadata format"), nil
			}
			//check for important fields like title, post_dir, type and status
			if _, ok := metadata["title"]; !ok {
				if formData.Get("title") == "" {
					return ErrorResponse(http.StatusBadRequest, "Missing title in metadata"), nil
				}
				metadata["title"] = formData.Get("title")
			}
			if _, ok := metadata["post_dir"]; !ok {
				return ErrorResponse(http.StatusBadRequest, "Missing post_dir in metadata"), nil
			}
			if _, ok := metadata["type"]; !ok {
				return ErrorResponse(http.StatusBadRequest, "Missing type in metadata"), nil
			}
			if _, ok := metadata["date"]; !ok {
				metadata["date"] = time.Now().Format("2006-01-02")
			}
			if _, ok := metadata["tags"]; ok {
				tags := strings.Split(metadata["tags"].(string), ",")
				metadata["tags"] = tags
			}
		}
		payload = plugins.Payload{
			Username: formData.Get("username"),
			Password: formData.Get("password"),
			Title:    formData.Get("title"),
			Post:     formData.Get("content"),
			Metadata: metadata,
		}
	} else { // Standard API request for new post
		err = json.Unmarshal([]byte(req.Body), &payload)
		if err != nil {
			log.Printf("Error unmarshalling JSON payload for new post: %v", err)
			return ErrorResponse(http.StatusBadRequest, "Invalid JSON payload"), nil
		}
	}

	log.Printf("Processing payload for user: %s", payload.Username)
	user, err := queries.GetUser(ctx, payload.Username)
	if err != nil {
		log.Printf("Error fetching user %s for post creation: %v", payload.Username, err)
		return ErrorResponse(http.StatusUnauthorized, "User not found or authentication failed for post creation"), nil
	}

	if !Authenticate(payload.Username, user.Password, payload.Password) {
		log.Printf("Authentication failed for user %s during post creation", payload.Username)
		return ErrorResponse(http.StatusUnauthorized, "Authentication failed for post creation"), nil
	}

	log.Printf("User %s authenticated for post creation.", payload.Username)
	postToCreate, err := plugins.CreatePostPayload(payload, int(user.ID), user.Name)
	if err != nil {
		log.Printf("Error in CreatePostPayload: %v", err)
		return ErrorResponse(http.StatusInternalServerError, "Error preparing post data"), nil
	}

	_, err = queries.CreatePost(ctx, postToCreate)
	if err != nil {
		log.Printf("Error in CreatePost: %v", err)
		return ErrorResponse(http.StatusInternalServerError, "Failed to create post in database"), nil
	}

	log.Printf("Post '%s' created successfully by user %s", postToCreate.Title, payload.Username)
	if req.Headers["hx-request"] == "true" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type":                 "text/html",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
			},
			Body: `<div class="success-message">Post created successfully!</div>`,
		}, nil
	}

	return jsonResponse(http.StatusOK, postToCreate), nil
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

// callTriggerBuildFunction makes an HTTP POST request to the Netlify trigger-build function.
func callTriggerBuildFunction(baseURL, triggerSecret string) (*http.Response, error) {
	triggerURL := fmt.Sprintf("%s/.netlify/functions/trigger-build", baseURL)
	log.Printf("Calling Netlify trigger function at: %s", triggerURL)

	req, err := http.NewRequest("POST", triggerURL, nil)
	if err != nil {
		log.Printf("Error creating request to trigger function: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Trigger-Secret", triggerSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to trigger function: %v", err)
		return nil, err
	}

	log.Printf("Response status from trigger function: %s", resp.Status)
	return resp, nil
}

func jsonResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
		},
		Body: string(body),
	}
}

func ErrorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	// Log the error being sent back to the client
	log.Printf("Responding with error: %d - %s", statusCode, message)
	return jsonResponse(statusCode, map[string]string{"error": message})
}
