package main

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/mattn/go-sqlite3"
)

func TestSearchDBOpenError(t *testing.T) {
	old := openDB
	openDB = func(string, string) (*sql.DB, error) { return nil, errors.New("boom") }
	defer func() { openDB = old }()
	resp, _ := handler(events.APIGatewayProxyRequest{Body: `{}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestSearchInvalidJSON(t *testing.T) {
	old := openDB
	openDB = func(string, string) (*sql.DB, error) { return sql.Open("sqlite3", ":memory:") }
	defer func() { openDB = old }()
	resp, _ := handler(events.APIGatewayProxyRequest{Body: "{bad"})
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400")
	}
}

func TestSearchQueryError(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	old := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = old }()
	resp, _ := handler(events.APIGatewayProxyRequest{Body: `{"query":"select * from no_table"}`})
	if resp.StatusCode != 500 {
		t.Fatalf("expected 500")
	}
}

func TestSearchSuccess(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(`create table t (id integer, name text); insert into t values (1, 'a');`)
	old := openDB
	openDB = func(string, string) (*sql.DB, error) { return db, nil }
	defer func() { openDB = old }()
	resp, _ := handler(events.APIGatewayProxyRequest{Body: `{"query":"select id, name from t"}`})
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200")
	}
	if resp.Body == "" {
		t.Fatalf("expected body")
	}
}
