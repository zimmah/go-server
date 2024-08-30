package router

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	type StatusResponse struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, StatusResponse{Status: "ok"})
}
