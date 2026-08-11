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

type ClassHandler struct {
	storage storage.ClassRepository
	logger  *slog.Logger
}

func NewClassHandler(storage storage.ClassRepository, logger *slog.Logger) *ClassHandler {
	return &ClassHandler{
		storage: storage,
		logger:  logger,
	}
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	// Structure Validation
	classReq := models.CreateClassRequest{}
	if err := json.NewDecoder(r.Body).Decode(&classReq); err != nil {
		h.logger.Error("Failed to create class", sl.Err(err))
		http.Error(w, "Failed to create class", http.StatusBadRequest)
		return
	}

	// Use repository
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	if err := h.storage.CreateClass(ctx, classReq); err != nil {
		h.logger.Error("Failed to create class", sl.Err(err))
		http.Error(w, "Failed to create class", http.StatusInternalServerError)
		return
	}
	// Response
	w.WriteHeader(http.StatusCreated)
}
