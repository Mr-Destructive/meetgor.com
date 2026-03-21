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
	"github.com/mr-destructive/meetgor.com/plugins"
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

func openDBWithUser(t *testing.T) func(string, string) (*sql.DB, error) {
	return func(string, string) (*sql.DB, error) {
		db := setupDB(t)
		createUser(t, db, "u", "n", "secret")
		return db, nil
	}
}

func setEnv(t *testing.T, key, value string) func() {
	t.Helper()
	old := os.Getenv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatal(err)
	}
	return func() { _ = os.Setenv(key, old) }
}

func TestOptions(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "OPTIONS"})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestTriggerMethodNotAllowed(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/api/trigger-build"})
	if resp.StatusCode != 405 {
		t.Fatalf("expected 405")
	}
}

func TestTriggerInvalidJSON(t *testing.T) {
	db := setupDB(t)
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: "{bad"})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestTriggerUserNotFound(t *testing.T) {
	db := setupDB(t)
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: `{"username":"u","password":"p"}`})
	if resp.StatusCode != 401 {
		t.Fatalf("expected 401")
	}
}

func TestTriggerAuthFail(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: `{"username":"u","password":"bad"}`})
	if resp.StatusCode != 401 {
		t.Fatalf("expected 401")
	}
}

func TestTriggerCallError(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	oldDo := httpDo
	httpDo = func(*http.Request) (*http.Response, error) { return nil, errors.New("boom") }
	defer func() { openDB = oldOpen; httpDo = oldDo }()
	set := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "x")
	defer set()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestTriggerCallNon2xx(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	oldDo := httpDo
	httpDo = func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 500, Body: ioNop("bad")}, nil
	}
	defer func() { openDB = oldOpen; httpDo = oldDo }()
	set := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "x")
	defer set()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestTriggerSuccess(t *testing.T) {
	db := setupDB(t)
	createUser(t, db, "u", "n", "secret")
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	oldDo := httpDo
	httpDo = func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioNop("ok")}, nil
	}
	defer func() { openDB = oldOpen; httpDo = oldDo }()
	set := setEnv(t, "GITHUB_PAT_FOR_TRIGGER", "x")
	defer set()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api/trigger-build", Body: `{"username":"u","password":"secret"}`})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
}

func TestEditGetNotFound(t *testing.T) {
	db := setupDB(t)
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/api",
		QueryStringParameters: map[string]string{
			"method": "edit",
			"type":   "posts",
			"slug":   "missing",
		},
	})
	if resp.StatusCode != 404 {
		t.Fatalf("expected 404")
	}
}

func TestEditPostInvalidJSON(t *testing.T) {
	db := setupDB(t)
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		QueryStringParameters: map[string]string{
			"method": "edit",
		},
		Body: "{bad",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestHTMXFormErrors(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return setupDB(t), nil }
	defer func() { openDB = oldOpen }()

	resp, _ := handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "%ZZ",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}

	resp, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "metadata={bad}",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}

	openDB = func(string, string) (*sql.DB, error) { return setupDB(t), nil }
	resp, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "metadata={}",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}

	openDB = func(string, string) (*sql.DB, error) { return setupDB(t), nil }
	resp, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "metadata={\"title\":\"t\"}",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}

	resp, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "metadata={\"title\":\"t\",\"post_dir\":\"posts\"}",
	})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestNewPostInvalidJSON(t *testing.T) {
	db := setupDB(t)
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = oldOpen }()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api", Body: "{bad"})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestCreatePostErrors(t *testing.T) {
	oldOpen := openDB
	openDB = openDBWithUser(t)
	defer func() { openDB = oldOpen }()

	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api", Body: `{"username":"u","password":"bad","title":"t","content":"c","metadata":{}}`})
	if resp.StatusCode != 401 {
		t.Fatalf("expected 401")
	}

	openDB = openDBWithUser(t)
	resp, _ = handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api", Body: `{"username":"u","password":"secret","title":"","content":"c","metadata":{}}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}

	// force CreatePost error by dropping table
	openDB = func(string, string) (*sql.DB, error) {
		db, _ := sql.Open("sqlite3", ":memory:")
		db.Exec(`CREATE TABLE IF NOT EXISTS authors (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			is_admin BOOLEAN DEFAULT 0
		);`)
		createUser(t, db, "u", "n", "secret")
		return db, nil
	}
	oldDDL := execDDL
	execDDL = func(ctx context.Context, db *sql.DB) error { return nil }
	resp, _ = handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api", Body: `{"username":"u","password":"secret","title":"t","content":"c","metadata":{}}`})
	execDDL = oldDDL
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestCreatePostSuccess(t *testing.T) {
	oldOpen := openDB
	openDB = openDBWithUser(t)
	defer func() { openDB = oldOpen }()
	nowTime = func() time.Time { return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) }
	defer func() { nowTime = time.Now }()

	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/api", Body: `{"username":"u","password":"secret","title":"t","content":"c","metadata":{"post_dir":"posts","type":"posts"}}`})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}

	// HTMX success
	openDB = openDBWithUser(t)
	resp, _ = handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/api",
		Headers:    map[string]string{"hx-request": "true"},
		Body:       "username=u&password=secret&title=t&content=c&metadata={\"title\":\"t\",\"post_dir\":\"posts\",\"type\":\"posts\",\"tags\":\"a,b\"}",
	})
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
