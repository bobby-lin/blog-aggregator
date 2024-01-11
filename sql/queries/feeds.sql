-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: SelectFeed :many
SELECT *
FROM feeds;

-- name: SelectNextFeedsToFetch :many
SELECT *
FROM  feeds
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT $1;

-- name: UpdateLastFetchedAt :one
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3
RETURNING *;