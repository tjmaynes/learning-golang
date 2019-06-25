package handler

import (
	"net/http"
	jsonHandler "github.com/tjmaynes/learning-golang/handler/json"
)

// GetPingHandler ..
func GetPingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	jsonHandler.CreateResponse(w, http.StatusOK, map[string]string{"message": "PONG!"})
}
