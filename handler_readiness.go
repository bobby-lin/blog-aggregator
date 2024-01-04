package main

import (
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	utils.RespondWithJSON(w, http.StatusOK, response{Status: "ok"})
}
