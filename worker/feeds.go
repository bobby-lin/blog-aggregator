package worker

import (
	"context"
	"github.com/bobby-lin/blog-aggregator/internal/database"
)

func (w Worker) GetNextFeedsToFetch(n int32) ([]database.Feed, error) {
	return w.DB.SelectNextFeedsToFetch(context.Background(), n)
}
