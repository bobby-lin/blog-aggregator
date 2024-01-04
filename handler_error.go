package main

import (
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusOK, "Internal Server Error")
}
