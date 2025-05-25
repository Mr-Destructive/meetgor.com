package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mr-destructive/mr-destructive.github.io/plugins/db/libsqlssg"
	"golang.org/x/crypto/bcrypt"
)

const (
	githubUsername   = "mr-destructive"
	githubRepoName   = "meetgor.com"
	githubBranch     = "main"
	timestampFileURL = "https://raw.githubusercontent.com/%s/%s/%s/.last_build_timestamp"
	githubApiUrlBase = "https://api.github.com/repos/%s/%s/dispatches"
	githubEventType  = "content-update"

	// Time constants
	githubTimeout = 20 * time.Second

	// Default old timestamp for fallback
	defaultOldTime = "2000-01-01T00:00:00Z"
)

type Payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Config holds all configuration values
type Config struct {
	GitHubPAT            string
	NetlifyTriggerSecret string
}

// loadConfig loads and validates environment variables
func loadConfig() (*Config, error) {
	config := &Config{
		GitHubPAT:            os.Getenv("GITHUB_PAT_FOR_TRIGGER"),
		NetlifyTriggerSecret: os.Getenv("NETLIFY_TRIGGER_SECRET"),
	}

	// Validate required fields
	if config.GitHubPAT == "" {
		return nil, fmt.Errorf("GITHUB_PAT_FOR_TRIGGER is required")
	}

	return config, nil
}

// createResponse creates standardized API Gateway response with CORS headers
func createResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
			"Content-Type":                 "application/json",
		},
		Body: body,
	}
}

// createErrorResponse creates a JSON error response
func createErrorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	errorBody := map[string]string{"error": message}
	body, _ := json.Marshal(errorBody)
	return createResponse(statusCode, string(body))
}

// createSuccessResponse creates a JSON success response
func createSuccessResponse(message string, data interface{}) events.APIGatewayProxyResponse {
	response := map[string]interface{}{
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}
	body, _ := json.Marshal(response)
	return createResponse(http.StatusOK, string(body))
}

// validateTriggerSecret checks the X-Trigger-Secret header if configured
func validateTriggerSecret(request events.APIGatewayProxyRequest, expectedSecret string) bool {
	if expectedSecret == "" {
		return true // No validation required
	}

	// Check for the header (case-insensitive)
	requestSecret := getHeaderValue(request.Headers, "X-Trigger-Secret")
	return requestSecret == expectedSecret
}

// getHeaderValue gets header value case-insensitively
func getHeaderValue(headers map[string]string, key string) string {
	// Try exact match first
	if val, ok := headers[key]; ok {
		return val
	}

	// Try lowercase
	if val, ok := headers[strings.ToLower(key)]; ok {
		return val
	}

	return ""
}

// generateSlug creates a URL-friendly slug from title
func generateSlug(title string) string {
	if title == "" {
		return ""
	}

	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

// triggerGitHubAction sends a repository dispatch event to GitHub
func triggerGitHubAction(githubPAT string) error {
	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	dispatchURL := fmt.Sprintf(githubApiUrlBase, githubUsername, githubRepoName)

	// Create dispatch payload
	dispatchPayload := map[string]interface{}{
		"event_type": githubEventType,
	}

	body, err := json.Marshal(dispatchPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", dispatchURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubPAT))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return nil
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

// handler is the main Netlify function handler
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Content sync function called: %s %s", request.HTTPMethod, request.Path)
	log.Printf("Request Body: %s", request.Body)
	reqBody := Payload{}
	ctx := context.Background()
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")
	var err error
	dbString := fmt.Sprintf("libsql://%s?authToken=%s", dbName, dbToken)
	db, err := sql.Open("libsql", dbString)
	if err != nil {
		log.Printf("Error in Connection: %v", err)
		return ErrorResponse(http.StatusInternalServerError, "Database connection failed"), nil
	}
	defer db.Close()
	queries := libsqlssg.New(db)
	err = json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		return createErrorResponse(http.StatusInternalServerError, "Failed to unmarshal request body"), nil
	}
	log.Printf("Request Body: %s", reqBody)

	log.Printf("Processing payload for user: %s", reqBody.Username)
	user, err := queries.GetUser(ctx, reqBody.Username)
	if err != nil {
		log.Printf("Error fetching user %s for post creation: %v", reqBody.Username, err)
		return ErrorResponse(http.StatusUnauthorized, "User not found or authentication failed for post creation"), nil
	}

	if !Authenticate(user.Name, user.Password, reqBody.Password) {
		log.Printf("Authentication failed for user %s during post creation", reqBody.Username)
		return ErrorResponse(http.StatusUnauthorized, "Authentication failed for post creation"), nil
	}

	log.Printf("User %s authenticated for triggering github workflow.", reqBody.Username)

	config, err := loadConfig()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return createErrorResponse(http.StatusInternalServerError, "Failed to load config"), nil
	}

	// Trigger GitHub Action
	if err := triggerGitHubAction(config.GitHubPAT); err != nil {
		log.Printf("Failed to trigger GitHub Action: %v", err)
		return createErrorResponse(http.StatusBadGateway, "Failed to trigger content build"), nil
	}
	currentTimestamp := time.Now().Format(time.RFC3339)

	return createSuccessResponse("Content update triggered successfully", map[string]interface{}{
		"last_build_time": currentTimestamp,
	}), nil
}

func main() {
	lambda.Start(handler)
}
