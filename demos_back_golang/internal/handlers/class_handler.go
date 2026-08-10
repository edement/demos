package handlers

import (
	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/storage"
	"encoding/json"
	"log/slog"
	"net/http"
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
	class := models.CreateClassRequest{}
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		h.logger.Error("Failed to create class", sl.Err(err))
		http.Error(w, "Failed to create class", http.StatusBadRequest)
		return
	}

	// Use repository
	// Response
}
