package main

import (
	"net/http"
)

func (app *Config) SolanaService(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the solana service",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
