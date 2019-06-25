package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HealthCheckHandlerTakesRequests(t *testing.T) {
	request, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// recorder := httptest
}