package main

import (
	"context"
	"encoding/json"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func (cfg *apiConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	dec := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := dec.Decode(&reqBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create feed")
		return
	}

	ctx := context.Background()

	apiKey := strings.Replace(r.Header.Get("Authorization"), "ApiKey ", "", 1)
	user, err := cfg.DB.SelectUser(ctx, apiKey)

	id, err := uuid.NewUUID()

	feedParam := database.CreateFeedParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      reqBody.Name,
		Url:       reqBody.Url,
		UserID:    user.ID,
	}

	f, err := cfg.DB.CreateFeed(ctx, feedParam)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create feed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	dat, _ := json.Marshal(f)
	w.Write(dat)
}
