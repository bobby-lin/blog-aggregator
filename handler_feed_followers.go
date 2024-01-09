package main

import (
	"context"
	"encoding/json"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) CreateFeedFollowerHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type requestBody struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	dec := json.NewDecoder(r.Body)
	reqBody := &requestBody{}
	err := dec.Decode(&reqBody)
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create feed follow")
		return
	}

	ff_id, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create feed follow")
		return
	}

	row := database.CreateFeedFollowerParams{
		ID:        ff_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    reqBody.FeedID,
		UserID:    user.ID,
	}

	ctx := context.Background()
	ff, err := cfg.DB.CreateFeedFollower(ctx, row)
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create feed follow")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	dat, _ := json.Marshal(ff)
	w.Write(dat)
}
