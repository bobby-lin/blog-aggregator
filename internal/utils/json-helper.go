package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	dat, _ := json.Marshal(payload)
	w.Write(dat)
}

func RespondWithError(w http.ResponseWriter, status int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	dat, _ := json.Marshal(errorResponse{Error: msg})
	w.Write(dat)
}
