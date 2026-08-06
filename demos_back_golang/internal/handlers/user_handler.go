package handlers

import (
	"context"
	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/storage"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type UserHandler struct {
	storage storage.UserRepository
	logger  *slog.Logger
}

func NewUserHandler(storage storage.UserRepository, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		storage: storage,
		logger:  logger,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler.logger.Debug("User object before additing:", "User", user) //
	// Use storage
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	user, err := handler.storage.CreateUser(ctx, user)
	if err != nil {
		handler.logger.Error("Failed to create user:", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.logger.Debug("User object after additing:", "User", user)
	// Response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		handler.logger.Error("Error while encoding responce:", sl.Err(err))
	}
}
