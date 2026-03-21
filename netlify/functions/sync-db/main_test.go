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
	"github.com/Mr-Destructive/meetgor.com/plugins"
	_ "github.com/mattn/go-sqlite3"
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
	old := os.Getenv("GITHUB_PAT_FOR_TRIGGER")
	os.Setenv("GITHUB_PAT_FOR_TRIGGER", "")
	defer os.Setenv("GITHUB_PAT_FOR_TRIGGER", old)
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
	old := os.Getenv("GITHUB_PAT_FOR_TRIGGER")
	os.Setenv("GITHUB_PAT_FOR_TRIGGER", "x")
	defer os.Setenv("GITHUB_PAT_FOR_TRIGGER", old)
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
	old := os.Getenv("GITHUB_PAT_FOR_TRIGGER")
	os.Setenv("GITHUB_PAT_FOR_TRIGGER", "x")
	defer os.Setenv("GITHUB_PAT_FOR_TRIGGER", old)
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
func (n *nopReadCloser) Close() error              { return nil }
