package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Mr-Destructive/meetgor.com/plugins"
	"golang.org/x/crypto/bcrypt"
)

func setupDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(plugins.DDL); err != nil {
		t.Fatal(err)
	}
	return db
}

func createUser(t *testing.T, db *sql.DB, username, name, password string) {
	t.Helper()
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := db.ExecContext(context.Background(), `INSERT INTO authors (username, name, password) VALUES (?, ?, ?)`, username, name, string(hashed))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOptions(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "OPTIONS"})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestValidateTriggerSecretBehaviors(t *testing.T) {
	req := events.APIGatewayProxyRequest{Headers: map[string]string{"X-Trigger-Secret": "foo"}}
	if !validateTriggerSecret(req, "foo") {
		t.Fatalf("expected secret to match")
	}
	if validateTriggerSecret(req, "other") {
		t.Fatalf("expected mismatch")
	}
	if !validateTriggerSecret(req, "") {
		t.Fatalf("expected pass when no secret configured")
	}
}

func TestGetHeaderValueCaseInsensitive(t *testing.T) {
	headers := map[string]string{"x-trigger-secret": "bar"}
	if v := getHeaderValue(headers, "X-Trigger-Secret"); v != "bar" {
		t.Fatalf("expected header value case insensitively")
	}
}

func TestGenerateSlug(t *testing.T) {
	if slug := generateSlug("Hello World!"); slug != "hello-world" {
		t.Fatalf("unexpected slug %s", slug)
	}
	if slug := generateSlug(""); slug != "" {
		t.Fatalf("expected empty slug for empty input")
	}
}

func TestDBOpenError(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return nil, errors.New("boom") }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{}"})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestUnmarshalError(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return setupDB(t), nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{bad"})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestUserNotFound(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return setupDB(t), nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"x","password":"y"}`})
	if resp.StatusCode != 401 {
		t.Fatalf("expected 401")
	}
}

func TestAuthFail(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"u","password":"bad"}`})
	if resp.StatusCode != 401 {
		t.Fatalf("expected 401")
	}
}

func TestLoadConfigFail(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	osSetenv := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "")
	defer osSetenv()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestTriggerGitHubActionFail(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	oldDo := httpDo
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	httpDo = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 500, Body: ioNop("bad")}, nil
	}
	defer func() { openDB = oldOpen; httpDo = oldDo }()
	set := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "x")
	defer set()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 502 {
		t.Fatalf("expected 502")
	}
}

func TestTriggerGitHubActionSuccess(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	oldDo := httpDo
	oldNow := nowTime
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	httpDo = func(req *http.Request) (*http.Response, error) {
		if !strings.Contains(req.URL.String(), "dispatches") {
			t.Fatalf("unexpected url")
		}
		return &http.Response{StatusCode: 204, Body: ioNop("")}, nil
	}
	nowTime = func() time.Time { return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) }
	defer func() { openDB = oldOpen; httpDo = oldDo; nowTime = oldNow }()
	set := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "x")
	defer set()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func ioNop(s string) *nopReadCloser {
	return &nopReadCloser{r: strings.NewReader(s)}
}

type nopReadCloser struct {
	r *strings.Reader
}

func (n *nopReadCloser) Read(p []byte) (int, error) { return n.r.Read(p) }
func (n *nopReadCloser) Close() error               { return nil }

func setEnv(t *testing.T, key, value string) func() {
	t.Helper()
	old := os.Getenv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatal(err)
	}
	return func() { _ = os.Setenv(key, old) }
}
