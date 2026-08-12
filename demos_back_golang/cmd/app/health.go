package main

import (
	"demos_back_golang/internal/lib/slogpretty/sl"
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"Message"`
}

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(HealthResponse{Message: "Health checked OK"})
	if err != nil {
		app.logger.Error("Failed to decode response", sl.Err(err))
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	//app.database.Posts.Create(r.Context(), nil)
}
