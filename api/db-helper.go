package api

import (
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt *time.Time
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	lastFetchedAt := time.Time{}

	if feed.LastFetchedAt.Valid {
		lastFetchedAt = feed.LastFetchedAt.Time
	}

	f := Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: &lastFetchedAt,
	}

	return f
}
