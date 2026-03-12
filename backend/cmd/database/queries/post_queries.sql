-- name: CreatePost :one
INSERT INTO posts (id, url,title, content, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING id, url, title, content, user_id, created_at, updated_at;


-- name: GetPostById :one
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
WHERE id = $1;

-- name: GetPostByUrl :one
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
WHERE url = $1;

-- name: GetPostsByUserId :many
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetPostsByTitle :many
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
WHERE title ILIKE '%' || $1 || '%'
ORDER BY created_at DESC;


-- name: UpdatePostById :one
UPDATE posts
SET title = $2, content = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, url, title, content, user_id, created_at, updated_at;

-- name: DeletePostById :one
DELETE FROM posts
WHERE id = $1
RETURNING id, url, title, content, user_id, created_at, updated_at;


-- name: ListPosts :many
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
ORDER BY created_at DESC;

-- name: ListNPosts :many
SELECT id, url, title, content, user_id, created_at, updated_at
FROM posts
ORDER BY created_at DESC
LIMIT $1;

-- name: CountPosts :one
SELECT COUNT(*) AS count
FROM posts;

-- name: CountPostsByUserId :one
SELECT COUNT(*) AS count
FROM posts
WHERE user_id = $1;

-- name: GetUserIdByPostId :one
SELECT user_id
FROM posts
WHERE id = $1;