-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4) -- we don't insert API key (autogenerate with default)
RETURNING *;

-- name: SelectUser :one
SELECT *
FROM users
WHERE api_key = $1;