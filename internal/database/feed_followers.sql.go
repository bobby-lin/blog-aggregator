// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed_followers.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollower = `-- name: CreateFeedFollower :one
INSERT INTO feed_followers(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, feed_id, user_id, created_at, updated_at
`

type CreateFeedFollowerParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) CreateFeedFollower(ctx context.Context, arg CreateFeedFollowerParams) (FeedFollower, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollower,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i FeedFollower
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedFollower = `-- name: DeleteFeedFollower :one
DELETE FROM feed_followers
WHERE id = $1
RETURNING id, feed_id, user_id, created_at, updated_at
`

func (q *Queries) DeleteFeedFollower(ctx context.Context, id uuid.UUID) (FeedFollower, error) {
	row := q.db.QueryRowContext(ctx, deleteFeedFollower, id)
	var i FeedFollower
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedFollowers = `-- name: GetFeedFollowers :many
SELECT id, feed_id, user_id, created_at, updated_at
FROM feed_followers
WHERE user_id = $1
`

func (q *Queries) GetFeedFollowers(ctx context.Context, userID uuid.UUID) ([]FeedFollower, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowers, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollower
	for rows.Next() {
		var i FeedFollower
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}