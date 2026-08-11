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

	"golang.org/x/crypto/bcrypt"
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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	regReq := models.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		h.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use storage
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	h.logger.Debug("Request Password from client", "password", regReq.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error("Failed to hash password", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	regReq.Password = string(hashedPassword)

	user, err := h.storage.CreateUser(ctx, regReq)
	if err != nil {
		h.logger.Error("Failed to create user:", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	// TODO: JWT tokens
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		h.logger.Error("Error while encoding responce", sl.Err(err))
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		h.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	user, err := h.storage.GetUser(ctx, loginReq.Email)
	// TODO: Проверка ошибки Системная или Пользователь не найден
	if err != nil {
		h.logger.Error("Failed to Get user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		h.logger.Error("Auth error", sl.Err(err))
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
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
		h.logger.Error("Error while encoding responce", sl.Err(err))
	}
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	/*h.storage.GetUser()

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
		h.logger.Error("Error while encoding responce", sl.Err(err))
	}*/
}
