package worker

import (
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"log"
)

type Worker struct {
	DB        *database.Queries
	FetchSize int32
}

func (w Worker) Start() {
	log.Println("starting worker")
	feedBatch, _ := w.GetNextFeedsToFetch(w.FetchSize)
	w.MarkFeedFetched(feedBatch[0])
}
