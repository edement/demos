package handlers

import (
	"context"
	"demos_back_golang/internal/lib/jwt"
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
	storage           storage.UserRepository
	logger            *slog.Logger
	refreshTokenStore *storage.RefreshTokenStorage
	jwtSecret         string
}

func NewUserHandler(
	storage storage.UserRepository,
	refreshTokenStore *storage.RefreshTokenStorage,
	logger *slog.Logger,
	jwtSecret string,
) *UserHandler {
	return &UserHandler{
		storage:           storage,
		refreshTokenStore: refreshTokenStore,
		logger:            logger,
		jwtSecret:         jwtSecret,
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

	accessToken, err := jwt.GenerateAccessToken(user.ID, user.Username, user.Email, h.jwtSecret)
	if err != nil {
		h.logger.Error("Failed to generate access token", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID, h.jwtSecret)
	if err != nil {
		h.logger.Error("Failed to generate refresh token", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	if err := h.refreshTokenStore.Save(ctx, user.ID, refreshToken, expiresAt); err != nil {
		h.logger.Error("Failed to save refresh token", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			IsTrainer: user.IsTrainer,
		},
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.RefreshToken == "" {
		http.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}

	claims, err := jwt.ValidateRefreshToken(req.RefreshToken, h.jwtSecret)
	if err != nil {
		h.logger.Error("Invalid refresh token", sl.Err(err))
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := h.refreshTokenStore.GetUserIDByToken(ctx, req.RefreshToken)
	if err != nil {
		h.logger.Error("Failed to check refresh token in DB", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if userID == 0 {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	if userID != claims.UserID {
		h.logger.Error("Refresh token user_id mismatch")
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	user, err := h.storage.GetUserByID(ctx, claims.UserID)
	if err != nil {
		h.logger.Error("Failed to get user", sl.Err(err))
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	newAccessToken, err := jwt.GenerateAccessToken(user.ID, user.Username, user.Email, h.jwtSecret)
	if err != nil {
		h.logger.Error("Failed to generate new access token", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"access_token": newAccessToken,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[len(prefix):]

	claims, err := jwt.ValidateAccessToken(tokenString, h.jwtSecret)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.refreshTokenStore.RevokeAllUserTokens(ctx, claims.UserID); err != nil {
		h.logger.Error("Failed to revoke tokens", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[len(prefix):]

	claims, err := jwt.ValidateAccessToken(tokenString, h.jwtSecret)
	if err != nil {
		h.logger.Error("Invalid token", sl.Err(err))
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.storage.GetUserByID(ctx, claims.UserID)
	if err != nil {
		h.logger.Error("Failed to get user", sl.Err(err))
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsTrainer: user.IsTrainer,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
}
