package handler

import (
	"net/http"
)

// GetPingHandler ..
func GetPingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "PONG!"})
}
