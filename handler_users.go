package main

import (
	"context"
	"encoding/json"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}

	dec := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := dec.Decode(&reqBody)
	if err != nil {
		return
	}

	id, err := uuid.NewUUID()

	userParams := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      reqBody.Name,
	}

	ctx := context.Background()

	user, err := cfg.DB.CreateUser(ctx, userParams)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	dat, _ := json.Marshal(user)
	w.Write(dat)

}
