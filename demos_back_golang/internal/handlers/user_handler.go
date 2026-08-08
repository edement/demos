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

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	regReq := models.RegisterRequest{}
	// TODO: password -> password_hash
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		handler.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use storage
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	user, err := handler.storage.CreateUser(ctx, regReq)
	if err != nil {
		handler.logger.Error("Failed to create user:", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	// TODO: JWT tokens
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		handler.logger.Error("Error while encoding responce", sl.Err(err))
	}
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		handler.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	user, err := handler.storage.GetUser(ctx, loginReq.Email)
	if err != nil {
		handler.logger.Error("Failed to Get user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if loginReq.Password != user.Password {
		handler.logger.Error("Auth error", sl.Err(err))
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	// Response
	// TODO: JWT tokens
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsTrainer: user.IsTrainer,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handler.logger.Error("Error while encoding responce", sl.Err(err))
	}
}
