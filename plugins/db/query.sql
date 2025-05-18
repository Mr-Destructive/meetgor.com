-- name: GetUser :one
SELECT id, name, password, is_admin FROM authors WHERE username = ?;

-- name: CreateAuthor :one
INSERT INTO authors (username, name, password)
VALUES (?, ?, ?) RETURNING id;

-- name: UpdateAuthor :one
UPDATE authors SET name = ?, password = ?, is_admin = ? WHERE id = ? RETURNING *;

-- name: GetAuthorByID :one
SELECT id, username, name, password, is_admin FROM authors WHERE id = ?;

-- name: CreatePost :one
INSERT INTO posts (title, slug, body, metadata, author_id) 
VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: GetPostsBySlugType :many
SELECT * FROM posts WHERE slug = ? AND deleted = 0;

-- name: GetAllPosts :many
SELECT * FROM posts WHERE deleted = 0;
