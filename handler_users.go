package main

import (
	"context"
	"encoding/json"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) GetUserHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	utils.RespondWithJSON(w, http.StatusOK, u)
}

func (cfg *apiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}

	dec := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := dec.Decode(&reqBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create user")
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
		utils.RespondWithError(w, http.StatusBadRequest, "fail to create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
