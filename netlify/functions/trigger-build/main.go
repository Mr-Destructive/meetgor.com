package main

import (
	"bytes"
	"context"
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
)

const (
	githubUsername   = "mr-destructive"
	githubRepoName   = "meetgor.com"
	githubBranch     = "main"
	timestampFileURL = "https://raw.githubusercontent.com/%s/%s/%s/.last_build_timestamp"
	githubApiUrlBase = "https://api.github.com/repos/%s/%s/dispatches"
	githubEventType  = "deploy-site"

	// Time constants
	githubTimeout = 20 * time.Second

	// Default old timestamp for fallback
	defaultOldTime = "2000-01-01T00:00:00Z"
)

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

// handler is the main Netlify function handler
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Content sync function called: %s %s", request.HTTPMethod, request.Path)

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
