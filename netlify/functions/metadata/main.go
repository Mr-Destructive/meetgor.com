package main

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type metadataRequest struct {
	URL string `json:"url"`
}

type metadataResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == http.MethodOptions {
		return optionsResponse()
	}

	var req metadataRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return jsonError(http.StatusBadRequest, "invalid request body")
	}

	if req.URL == "" {
		return jsonError(http.StatusBadRequest, "url is required")
	}

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		req.URL = "https://" + req.URL
	}

	httpReq, err := http.NewRequest(http.MethodGet, req.URL, nil)
	if err != nil {
		return jsonError(http.StatusBadRequest, "invalid url")
	}
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MetadataBot/1.0)")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return jsonError(http.StatusBadRequest, "failed to fetch url")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return jsonError(http.StatusBadRequest, "url returned non-200 status")
	}

	html, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to read response")
	}

	title := extractMeta(html, "og:title", "twitter:title", "title")
	description := extractMeta(html, "og:description", "twitter:description", "description")
	image := extractMeta(html, "og:image", "twitter:image")

	source := req.URL
	if idx := strings.Index(source, "//"); idx != -1 {
		source = source[idx+2:]
	}
	if idx := strings.Index(source, "/"); idx != -1 {
		source = source[:idx]
	}

	return jsonResponse(http.StatusOK, metadataResponse{
		Title:       title,
		Description: description,
		Image:       image,
		Source:      source,
	})
}

func extractMeta(html []byte, props ...string) string {
	for _, prop := range props {
		if prop == "title" {
			continue
		}
		if content := extractMetaByProp(string(html), prop); content != "" {
			return content
		}
	}
	return extractTagContent(html, "title")
}

func extractMetaByProp(html, prop string) string {
	patterns := []string{
		`meta[^>]*property=["']` + prop + `["'][^>]*content=["']([^"']*)["']`,
		`meta[^>]*content=["']([^"']*)["'][^>]*property=["']` + prop + `["']`,
		`meta[^>]*name=["']` + prop + `["'][^>]*content=["']([^"']*)["']`,
		`meta[^>]*content=["']([^"']*)["'][^>]*name=["']` + prop + `["']`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(`(?i)` + pattern)
		matches := re.FindStringSubmatch(html)
		if len(matches) > 1 {
			return matches[1]
		}
	}
	return ""
}

func extractTagContent(html []byte, tag string) string {
	open := `<` + tag + `>`
	close := `</` + tag + `>`
	idx := strings.Index(string(html), open)
	if idx == -1 {
		return ""
	}
	idx += len(open)
	end := strings.Index(string(html)[idx:], close)
	if end == -1 {
		return ""
	}
	return strings.TrimSpace(string(html)[idx : idx+end])
}

func awsHandler() {
	lambda.Start(handler)
}

func main() {
	awsHandler()
}

func jsonResponse(statusCode int, payload interface{}) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(payload)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
		},
		Body: string(body),
	}, nil
}

func jsonError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return jsonResponse(statusCode, map[string]string{"error": message})
}

func optionsResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body: "",
	}, nil
}
