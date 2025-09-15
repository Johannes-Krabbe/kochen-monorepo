-- name: CreateUser :one
INSERT INTO users (id, email, name, username, password_hash)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, name, username, created_at, updated_at;

-- name: GetUser :one
SELECT id, email, name, username, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, name, username, password_hash, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT id, email, name, username, password_hash, created_at, updated_at
FROM users
WHERE username = $1;

-- name: GetUserByEmailOrUsername :one
SELECT id, email, name, username, password_hash, created_at, updated_at
FROM users
WHERE email = $1 OR username = $1;

-- name: ListUsers :many
SELECT id, email, name, username, created_at, updated_at
FROM users
ORDER BY created_at DESC;