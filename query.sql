-- name: CreatePost :one
INSERT INTO posts (
    title, subtitle, content, language, tags, created_at, edited_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: CleanPosts :exec
DELETE FROM posts;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPagedPosts :many
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;
