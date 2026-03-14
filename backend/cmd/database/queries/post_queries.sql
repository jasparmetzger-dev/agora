-- name: CreatePost :one
INSERT INTO posts (url, user_id, title, content, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING *;


-- name: GetPostById :one
SELECT *
FROM posts
WHERE id = $1;

-- name: GetPostByUrl :one
SELECT *
FROM posts
WHERE url = $1;

-- name: GetPostsByUserId :many
SELECT *
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetPostsByTitle :many
SELECT *
FROM posts
WHERE title ILIKE '%' || $1 || '%'
ORDER BY created_at DESC;


-- name: UpdatePostById :one
UPDATE posts
SET title = $2, content = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePostById :one
DELETE FROM posts
WHERE id = $1
RETURNING *;


-- name: ListPosts :many
SELECT *
FROM posts
ORDER BY created_at DESC;

-- name: ListNPosts :many
SELECT *
FROM posts
ORDER BY created_at DESC
LIMIT $1;

-- name: GetUserIdByPostId :one
SELECT user_id
FROM posts
WHERE id = $1;


-- name: CountPosts :one
SELECT COUNT(*) AS count
FROM posts;

-- name: CountPostsByUserId :one
SELECT COUNT(*) AS count
FROM posts
WHERE user_id = $1;