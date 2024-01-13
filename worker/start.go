package worker

import (
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"log"
	"time"
)

type Worker struct {
	DB        *database.Queries
	FetchSize int32
}

func (w Worker) Start() {
	log.Println("starting worker")
	interval := 10 * time.Second
	go w.Loop(interval)
}
