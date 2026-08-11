package handlers

import (
	"context"
	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/storage"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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
		h.logger.Error("Failed to decode request", sl.Err(err))
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

func (h *ClassHandler) GetClassById(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	classID, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse class id", sl.Err(err))
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	// Use repository
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	class, err := h.storage.GetClassById(ctx, classID)
	if err != nil {
		h.logger.Error("Failed to get class", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Response
	response, err := json.Marshal(class)
	if err != nil {
		h.logger.Error("Failed to decode response", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	classId, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse class id", sl.Err(err))
		http.Error(w, "Failed to delete class", http.StatusNotFound)
		return
	}

	// Use repository
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	err = h.storage.DeleteClass(ctx, classId)
	if err != nil {
		h.logger.Error("Failed to delete class")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Response
	w.WriteHeader(http.StatusNoContent)
}

func (h *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	// Structure validation
	classUpdate := models.UpdateClassRequest{}
	if err := json.NewDecoder(r.Body).Decode(&classUpdate); err != nil {
		h.logger.Error("Failed to decode request")
		http.Error(w, "Failed to update class", http.StatusBadRequest)
		return
	}

	classId, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse class id", sl.Err(err))
		http.Error(w, "Failed to update class", http.StatusNotFound)
		return
	}

	// Use repository
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	class, err := h.storage.UpdateClass(ctx, classId, classUpdate)
	if err != nil {
		h.logger.Error("Failed to update class", sl.Err(err))
		http.Error(w, "Failed to update class", http.StatusInternalServerError)
		return
	}

	// Response
	response, err := json.Marshal(class)
	if err != nil {
		h.logger.Error("Failed to decode response", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
