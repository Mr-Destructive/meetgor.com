package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mr-destructive/mr-destructive.github.io/plugins"
	_ "github.com/mattn/go-sqlite3"
)

func TestHandlerMethodNotAllowed(t *testing.T) {
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "GET"})
	if resp.StatusCode != 405 {
		t.Fatalf("expected 405")
	}
}

func TestHandlerDBOpenError(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) {
		return nil, errors.New("boom")
	}
	defer func() { openDB = oldOpen }()

	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{}"})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestHandlerDDLFailure(t *testing.T) {
	oldOpen := openDB
	oldDDL := execDDL
	openDB = func(string, string) (*sql.DB, error) {
		return sql.Open("sqlite3", ":memory:")
	}
	execDDL = func(ctx context.Context, db *sql.DB) error {
		return errors.New("ddl")
	}
	defer func() {
		openDB = oldOpen
		execDDL = oldDDL
	}()
	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{}"})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestHandlerInvalidJSON(t *testing.T) {
	oldOpen := openDB
	openDB = func(string, string) (*sql.DB, error) {
		return sql.Open("sqlite3", ":memory:")
	}
	defer func() { openDB = oldOpen }()

	oldDDL := execDDL
	execDDL = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, plugins.DDL)
		return err
	}
	defer func() { execDDL = oldDDL }()

	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{bad"})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestHandlerHashError(t *testing.T) {
	oldOpen := openDB
	oldHash := hashPassword
	oldDDL := execDDL
	openDB = func(string, string) (*sql.DB, error) {
		return sql.Open("sqlite3", ":memory:")
	}
	execDDL = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, plugins.DDL)
		return err
	}
	hashPassword = func([]byte, int) ([]byte, error) {
		return nil, errors.New("hash")
	}
	defer func() {
		openDB = oldOpen
		hashPassword = oldHash
		execDDL = oldDDL
	}()

	resp, _ := handler(events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: `{"username":"u","name":"n","password":"p"}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestHandlerSuccess(t *testing.T) {
	oldOpen := openDB
	oldDDL := execDDL
	openDB = func(string, string) (*sql.DB, error) {
		return sql.Open("sqlite3", ":memory:")
	}
	execDDL = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, plugins.DDL)
		return err
	}
	defer func() {
		openDB = oldOpen
		execDDL = oldDDL
	}()

	resp, _ := handler(events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       `{"username":"u","name":"n","password":"p"}`,
	})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestAuthenticate(t *testing.T) {
	if Authenticate("u", "a", "b") {
		t.Fatalf("expected false")
	}
}
