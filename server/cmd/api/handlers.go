package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status      string `json:"status"`      // [up, down] or [available, unavailable]
		Environment string `json:"environment"` // [development, staging, production]
		Version     string `json:"version"`     // major.minor.patch
		Timestamp   string `json:"timestamp"`   // ISO 8601
	}

	res := response{
		"up",
		"development",
		version,
		time.Now().UTC().Format(time.RFC3339),
	}

	b, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
