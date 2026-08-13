package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/middleware"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/services"
)

type UserHandler struct {
	service services.UserService
	logger  *slog.Logger
}

func NewUserHandler(service services.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	regReq := models.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		h.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, accessToken, refreshToken, err := h.service.Register(ctx, regReq)
	if err != nil {
		h.logger.Error("Failed to register user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response := map[string]interface{}{
		"user":   *user,
		"tokens": tokens,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		h.logger.Error("Failed to decode user", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, accessToken, refreshToken, err := h.service.Login(ctx, loginReq)
	if err != nil {
		h.logger.Error("Failed to login", sl.Err(err))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response := map[string]interface{}{
		"user":   *user,
		"tokens": tokens,
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	newAccessToken, newRefreshToken, err := h.service.Refresh(ctx, req.RefreshToken)
	if err != nil {
		h.logger.Error("Failed to refresh token", sl.Err(err))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(models.UserClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.Logout(ctx, claims.UserID); err != nil {
		h.logger.Error("Failed to revoke tokens", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(models.UserClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := h.service.Me(ctx, claims.UserID)
	if err != nil {
		h.logger.Error("Failed to get user", sl.Err(err))
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
