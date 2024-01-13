package main

import (
	"context"
	"encoding/json"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) GetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := 25

	limitSize, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		limit = limitSize
	}

	queryParams := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := cfg.DB.GetPostsByUser(context.Background(), queryParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "fail to get user posts")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	dat, _ := json.Marshal(posts)
	w.Write(dat)
}
