-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;


-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;


-- name: UpdateUserById :one
UPDATE users
SET username = $2, email = $3, password_hash = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUserById :one
DELETE FROM users
WHERE id = $1
RETURNING *;


-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;

-- name: ListNUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT $1;

-- name: CountUsers :one
SELECT COUNT(*) AS count
FROM users;