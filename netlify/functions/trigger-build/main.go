package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Helper function to create standardized APIGatewayProxyResponse with CORS headers
func createResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
		"Access-Control-Allow-Methods": "POST, OPTIONS",
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Trigger function called")

	// Handle OPTIONS request for CORS preflight
	if request.HTTPMethod == "OPTIONS" {
		log.Println("Handling OPTIONS request")
		return createResponse(http.StatusOK, ""), nil
	}

	// Method Check: Ensure the request method is POST
	if request.HTTPMethod != "POST" {
		log.Printf("Method not allowed: %s", request.HTTPMethod)
		return createResponse(http.StatusMethodNotAllowed, "Method Not Allowed"), nil
	}

	// Retrieve Environment Variables
	githubPat := os.Getenv("GITHUB_PAT_FOR_TRIGGER")
	netlifyTriggerSecret := os.Getenv("NETLIFY_TRIGGER_SECRET") // Optional
	ghOwner := os.Getenv("GH_OWNER")
	ghRepo := os.Getenv("GH_REPO")
	ghBranch := os.Getenv("GH_BRANCH") // New environment variable for branch

	if ghBranch == "" {
		ghBranch = "main" // Default to "main" if not set
		log.Println("GH_BRANCH not set, defaulting to 'main'")
	}

	if githubPat == "" || ghOwner == "" || ghRepo == "" {
		log.Println("Error: Missing required environment variables (GITHUB_PAT_FOR_TRIGGER, GH_OWNER, or GH_REPO)")
		return createResponse(http.StatusInternalServerError, "Internal Server Error: Missing configuration"), nil
	}

	// Security Check (if NETLIFY_TRIGGER_SECRET is set)
	if netlifyTriggerSecret != "" {
		log.Println("NETLIFY_TRIGGER_SECRET is set, checking X-Trigger-Secret header")
		requestSecret := ""
		// Headers can be multi-value, but for X-Trigger-Secret we expect a single value.
		// Also, header names are case-insensitive. APIGatewayProxyRequest normalizes them to lowercase.
		if val, ok := request.Headers["x-trigger-secret"]; ok {
			requestSecret = val
		} else if val, ok := request.Headers["X-Trigger-Secret"]; ok { // Check uppercase just in case
			requestSecret = val
		}


		if requestSecret != netlifyTriggerSecret {
			log.Printf("Unauthorized: Invalid or missing X-Trigger-Secret. Expected: %s, Got: %s", netlifyTriggerSecret, requestSecret)
			return createResponse(http.StatusUnauthorized, "Unauthorized: Invalid trigger secret"), nil
		}
		log.Println("X-Trigger-Secret validated successfully")
	}

	// Construct GitHub API URL
	githubApiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/deploy.yml/dispatches", ghOwner, ghRepo)
	log.Printf("GitHub API URL: %s", githubApiUrl)

	// Make API Call to GitHub
	// Create a JSON request body
	requestBodyMap := map[string]string{"ref": ghBranch}
	jsonBody, err := json.Marshal(requestBodyMap)
	if err != nil {
		log.Printf("Error marshalling JSON request body: %v", err)
		return createResponse(http.StatusInternalServerError, "Internal Server Error: Could not create request to GitHub"), nil
	}

	// Create a new HTTP POST request to the GitHub API URL
	req, err := http.NewRequestWithContext(context.Background(), "POST", githubApiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Error creating GitHub API request: %v", err)
		return createResponse(http.StatusInternalServerError, "Internal Server Error: Could not create request to GitHub"), nil
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubPat))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	log.Printf("Dispatching GitHub Action for branch: %s", ghBranch)

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to GitHub API: %v", err)
		// Check if the error is due to context deadline exceeded or cancellation, which might indicate a timeout
		if err == context.DeadlineExceeded {
			return createResponse(http.StatusGatewayTimeout, "Bad Gateway: GitHub API request timed out"), nil
		}
		return createResponse(http.StatusBadGateway, "Bad Gateway: Error communicating with GitHub"), nil
	}
	defer resp.Body.Close()

	// If the GitHub API returns a non-successful status code (e.g., not 204 No Content)
	if resp.StatusCode != http.StatusNoContent {
		responseBodyBytes, _ := os.ReadAll(resp.Body) // Read body for logging
		log.Printf("Error from GitHub API: StatusCode %d, Body: %s", resp.StatusCode, string(responseBodyBytes))
		return createResponse(http.StatusBadGateway, fmt.Sprintf("Bad Gateway: GitHub API responded with %d", resp.StatusCode)), nil
	}

	log.Println("Build triggered successfully via GitHub API.")
	return createResponse(http.StatusOK, "Build triggered successfully."), nil
}

func main() {
	// Check if running locally
	if _, ok := os.LookupEnv("NETLIFY_LOCAL"); ok {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Simulate APIGatewayProxyRequest for local testing
			var body []byte
			if r.Body != nil {
				body, _ = os.ReadAll(r.Body)
				r.Body.Close() // Close the original body
				r.Body = http.MaxBytesReader(w, bytes.NewReader(body), 1024*1024) // Restore body for next read if any
			}

			// Convert http.Header to map[string]string and map[string][]string
			headers := make(map[string]string)
			multiValueHeaders := make(map[string][]string)
			for k, v := range r.Header {
				headers[strings.ToLower(k)] = v[0] // APIGatewayProxyRequest headers are typically lowercase
				multiValueHeaders[strings.ToLower(k)] = v
			}


			req := events.APIGatewayProxyRequest{
				HTTPMethod:        r.Method,
				Path:              r.URL.Path,
				Headers:           headers,
				MultiValueHeaders: multiValueHeaders,
				Body:              string(body),
			}
			// Call the handler
			resp, err := handler(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set headers
			for k, v := range resp.Headers {
				w.Header().Set(k, v)
			}
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, resp.Body)
		})
		log.Println("Running trigger function locally on port 9999")
		log.Fatal(http.ListenAndServe(":9999", nil))
	} else {
		lambda.Start(handler)
	}
}

// Helper function to convert http.Header to map[string]string (no longer directly used by handler)
// func convertHeaders(h http.Header) map[string]string {
// 	headers := make(map[string]string)
// 	for k, v := range h {
// 		headers[k] = v[0] // Take the first value for simplicity
// 	}
// 	return headers
// }

// Helper function to read request body (no longer directly used by handler)
// func getBody(r *http.Request) string {
// 	body, err := os.ReadAll(r.Body)
// 	if err != nil {
// 		log.Printf("Error reading request body: %v", err)
// 		return ""
// 	}
// 	defer r.Body.Close()
// 	return string(body)
// }
