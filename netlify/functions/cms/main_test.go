package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCMSOptions(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: http.MethodOptions})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestCMSServeUI(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: http.MethodGet, Path: "/.netlify/functions/cms"})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestCMSCreateInvalid(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: http.MethodPost, Path: "/.netlify/functions/cms/api/create", Body: "{bad"})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestCMSProxyMissingHost(t *testing.T) {
	resp, _ := proxy(events.APIGatewayProxyRequest{Headers: map[string]string{}}, "/x", "{}")
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestCMSProxyError(t *testing.T) {
	old := httpDo
	httpDo = func(*http.Request) (*http.Response, error) { return nil, errors.New("boom") }
	defer func() { httpDo = old }()
	resp, _ := proxy(events.APIGatewayProxyRequest{Headers: map[string]string{"Host": "example.com"}}, "/x", "{}")
	if resp.StatusCode != 502 {
		t.Fatalf("expected 502")
	}
}

func TestCMSProxySuccess(t *testing.T) {
	old := httpDo
	httpDo = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"ok":1}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, nil
	}
	defer func() { httpDo = old }()
	resp, _ := proxy(events.APIGatewayProxyRequest{Headers: map[string]string{"Host": "example.com"}}, "/x", "{}")
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestCMSHandleCreate(t *testing.T) {
	old := httpDo
	httpDo = func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{}`)), Header: http.Header{}}, nil
	}
	defer func() { httpDo = old }()
	req := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Path:       "/.netlify/functions/cms/api/create",
		Headers:    map[string]string{"Host": "example.com"},
		Body:       `{"username":"u","password":"p","title":"t","content":"c","metadata":{}}`,
	}
	resp, _ := handler(req)
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}
