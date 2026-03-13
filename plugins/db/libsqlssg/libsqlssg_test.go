package libsqlssg

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestQueries(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	schema := `
CREATE TABLE IF NOT EXISTS authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    slug TEXT NOT NULL,
    body TEXT NOT NULL,
    metadata TEXT NOT NULL,
    deleted BOOLEAN DEFAULT 0,
    created_at DEFAULT CURRENT_TIMESTAMP,
    updated_at DEFAULT CURRENT_TIMESTAMP,
    author_id INTEGER REFERENCES authors(id) NOT NULL
);
`
	if _, err := db.Exec(schema); err != nil {
		t.Fatal(err)
	}
	q := New(db)
	ctx := context.Background()

	authorID, err := q.CreateAuthor(ctx, CreateAuthorParams{
		Username: "u",
		Name:     "n",
		Password: "p",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = q.CreatePost(ctx, CreatePostParams{
		Title:    "t",
		Slug:     "posts/t",
		Body:     "b",
		Metadata: "{}",
		AuthorID: authorID,
	})
	if err != nil {
		t.Fatal(err)
	}

	posts, err := q.GetAllPosts(ctx)
	if err != nil || len(posts) == 0 {
		t.Fatal("expected posts")
	}

	bySlug, err := q.GetPostsBySlugType(ctx, "posts/t")
	if err != nil || len(bySlug) == 0 {
		t.Fatal("expected posts by slug")
	}

	user, err := q.GetUser(ctx, "u")
	if err != nil || user.Name != "n" {
		t.Fatal("expected user")
	}

	author, err := q.GetAuthorByID(ctx, authorID)
	if err != nil || author.Username != "u" {
		t.Fatal("expected author")
	}

	_, err = q.UpdateAuthor(ctx, UpdateAuthorParams{
		Name:     "n2",
		Password: "p2",
		IsAdmin:  sql.NullBool{Bool: true, Valid: true},
		ID:       authorID,
	})
	if err != nil {
		t.Fatal(err)
	}
}
