package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//go:embed ui/*
var ui embed.FS

type createPayload struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Title    string                 `json:"title"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

type authPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	lambda.Start(handler)
}

var httpDo = func(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod == http.MethodOptions {
		return response(http.StatusOK, "", "application/json"), nil
	}

	path := normalizePath(req.Path)
	switch {
	case path == "" || path == "index.html":
		return serveUI("ui/index.html")
	case path == "app.js":
		return serveUI("ui/app.js")
	case path == "api/create":
		if req.HTTPMethod != http.MethodPost {
			return jsonError(http.StatusMethodNotAllowed, "method not allowed"), nil
		}
		return handleCreate(req)
	case path == "api/trigger-build":
		if req.HTTPMethod != http.MethodPost {
			return jsonError(http.StatusMethodNotAllowed, "method not allowed"), nil
		}
		return proxy(req, "/.netlify/functions/trigger-build", req.Body)
	case path == "api/sync-db":
		if req.HTTPMethod != http.MethodPost {
			return jsonError(http.StatusMethodNotAllowed, "method not allowed"), nil
		}
		return proxy(req, "/.netlify/functions/sync-db", req.Body)
	default:
		return jsonError(http.StatusNotFound, "not found"), nil
	}
}

func normalizePath(p string) string {
	p = strings.TrimPrefix(p, "/.netlify/functions/cms")
	p = strings.TrimPrefix(p, "/cms")
	return strings.Trim(strings.TrimSpace(p), "/")
}

func serveUI(path string) (events.APIGatewayProxyResponse, error) {
	b, err := ui.ReadFile(path)
	if err != nil {
		return jsonError(http.StatusNotFound, "ui not found"), nil
	}
	contentType := "text/html"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	}
	return response(http.StatusOK, string(b), contentType), nil
}

func handleCreate(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var p createPayload
	if err := json.Unmarshal([]byte(req.Body), &p); err != nil {
		return jsonError(http.StatusBadRequest, "invalid payload"), nil
	}
	if p.Metadata == nil {
		p.Metadata = map[string]interface{}{}
	}
	if _, ok := p.Metadata["title"]; !ok && p.Title != "" {
		p.Metadata["title"] = p.Title
	}
	if _, ok := p.Metadata["post_dir"]; !ok {
		p.Metadata["post_dir"] = "posts"
	}
	if _, ok := p.Metadata["type"]; !ok {
		p.Metadata["type"] = "posts"
	}
	if _, ok := p.Metadata["published"]; !ok {
		p.Metadata["published"] = "published"
	}

	payload := map[string]interface{}{
		"username": p.Username,
		"password": p.Password,
		"title":    p.Title,
		"content":  p.Content,
		"metadata": p.Metadata,
	}
	b, _ := json.Marshal(payload)
	return proxy(req, "/.netlify/functions/api", string(b))
}

func proxy(req events.APIGatewayProxyRequest, targetPath, body string) (events.APIGatewayProxyResponse, error) {
	host := req.Headers["Host"]
	if host == "" {
		host = req.Headers["host"]
	}
	if host == "" {
		return jsonError(http.StatusBadRequest, "missing host header"), nil
	}

	scheme := req.Headers["X-Forwarded-Proto"]
	if scheme == "" {
		scheme = req.Headers["x-forwarded-proto"]
	}
	if scheme == "" {
		scheme = "https"
	}

	url := fmt.Sprintf("%s://%s%s", scheme, host, targetPath)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(body))
	if err != nil {
		return jsonError(http.StatusInternalServerError, "failed to create upstream request"), nil
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if v := req.Headers["X-Trigger-Secret"]; v != "" {
		httpReq.Header.Set("X-Trigger-Secret", v)
	} else if v := req.Headers["x-trigger-secret"]; v != "" {
		httpReq.Header.Set("X-Trigger-Secret", v)
	}

	resp, err := httpDo(httpReq)
	if err != nil {
		return jsonError(http.StatusBadGateway, "upstream request failed"), nil
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	ct := resp.Header.Get("Content-Type")
	if ct == "" {
		ct = "application/json"
	}
	return response(resp.StatusCode, string(respBody), ct), nil
}

func jsonError(code int, msg string) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(map[string]string{"error": msg})
	return response(code, string(body), "application/json")
}

func response(code int, body, contentType string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":                 contentType,
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Trigger-Secret",
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
		},
	}
}
