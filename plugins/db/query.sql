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
INSERT INTO posts (id, type_id, title, slug, content, metadata, tags, status) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, type_id, title, slug, content, metadata, tags, status;

-- name: GetPostsBySlugType :many
SELECT id, type_id, title, slug, content, metadata, tags, status FROM posts WHERE slug = ? AND status != 'deleted';

-- name: GetAllPosts :many
SELECT id, type_id, title, slug, content, metadata, tags, status, created_at, updated_at FROM posts WHERE status != 'deleted';

-- name: UpdatePost :exec
UPDATE posts SET title = ?, content = ?, metadata = ?, tags = ? WHERE slug = ?;

-- name: GetPostBySlug :one
SELECT id, type_id, title, slug, content, metadata, tags, status FROM posts WHERE slug = ? AND status != 'deleted';

-- name: DeletePost :exec
UPDATE posts SET status = 'deleted' WHERE slug = ?;

-- name: EnsurePostType :exec
INSERT OR IGNORE INTO post_types (id, name, slug)
VALUES (?, ?, ?);
