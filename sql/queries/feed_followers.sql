-- name: CreateFeedFollower :one
INSERT INTO feed_followers(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollower :one
DELETE FROM feed_followers
WHERE id = $1
RETURNING *;

-- name: GetFeedFollowers :many
SELECT *
FROM feed_followers
WHERE user_id = $1;