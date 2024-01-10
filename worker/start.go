package worker

import (
	"fmt"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"log"
)

type Worker struct {
	DB        *database.Queries
	FetchSize int32
}

func (w Worker) Start() {
	log.Println("starting worker")
	fmt.Println(w.GetNextFeedsToFetch(w.FetchSize))
}
