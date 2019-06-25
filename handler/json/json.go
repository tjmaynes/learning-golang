package json

import (
	"encoding/json"
	"net/http"
)

// CreateErrorResponse ..
func CreateErrorResponse(w http.ResponseWriter, code int, msg string) {
	CreateResponse(w, code, map[string]string{"message": msg})
}

// CreateResponse ..
func CreateResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
