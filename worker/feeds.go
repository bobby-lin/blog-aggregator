package worker

import (
	"context"
	"database/sql"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"log"
	"time"
)

func (w Worker) GetNextFeedsToFetch(n int32) ([]database.Feed, error) {
	return w.DB.SelectNextFeedsToFetch(context.Background(), n)
}

func (w Worker) MarkFeedFetched(feed database.Feed) (database.Feed, error) {
	t := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	params := database.UpdateLastFetchedAtParams{LastFetchedAt: t, ID: feed.ID, UpdatedAt: time.Now()}
	updatedFeed, err := w.DB.UpdateLastFetchedAt(context.Background(), params)

	if err != nil {
		return database.Feed{}, err
	}

	log.Println("Updated feed " + updatedFeed.ID.String() + " at " + updatedFeed.LastFetchedAt.Time.GoString())
	return updatedFeed, nil
}