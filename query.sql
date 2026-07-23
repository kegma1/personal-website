-- name: CreatePost :one
INSERT INTO posts (
    title, subtitle, content, language, tags
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY created_at;
