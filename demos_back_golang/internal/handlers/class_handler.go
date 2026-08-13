package handlers

import (
	"context"
	"demos_back_golang/internal/lib/slogpretty/sl"
	"demos_back_golang/internal/middleware"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/services"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type ClassHandler struct {
	service services.ClassService
	logger  *slog.Logger
}

func NewClassHandler(service services.ClassService, logger *slog.Logger) *ClassHandler {
	return &ClassHandler{
		service: service,
		logger:  logger,
	}
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	classReq := models.CreateClassRequest{}
	if err := json.NewDecoder(r.Body).Decode(&classReq); err != nil {
		h.logger.Error("Failed to decode request", sl.Err(err))
		http.Error(w, "Failed to create class", http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(middleware.UserClaimsKey).(models.UserClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	class, err := h.service.CreateClass(ctx, classReq, claims.UserID)
	if err != nil {
		h.logger.Error("Failed to create class", sl.Err(err))
		http.Error(w, "Failed to create class", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(class)
	if err != nil {
		h.logger.Error("Failed to encode response", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Error while encoding response", sl.Err(err))
	}
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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	class, err := h.service.GetClassById(ctx, classID)
	if err != nil {
		h.logger.Error("Failed to get class", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(class)
	if err != nil {
		h.logger.Error("Failed to encode response", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *ClassHandler) GetClasses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	classes, err := h.service.GetClasses(ctx)
	if err != nil {
		h.logger.Error("Failed to get classes", sl.Err(err))
		http.Error(w, "Failed to get classes", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(classes)
	if err != nil {
		h.logger.Error("Failed to encode response", sl.Err(err))
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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Learn about context (for what, why, etc...)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	if err := h.service.DeleteClass(ctx, classId); err != nil {
		h.logger.Error("Failed to delete class", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	classUpdate := models.UpdateClassRequest{}
	if err := json.NewDecoder(r.Body).Decode(&classUpdate); err != nil {
		h.logger.Error("Failed to decode request", sl.Err(err))
		http.Error(w, "Failed to update class", http.StatusBadRequest)
		return
	}

	classId, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse class id", sl.Err(err))
		http.Error(w, "Failed to update class", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// TODO: Проверка ошибки Системная или Класс не найден
	class, err := h.service.UpdateClass(ctx, classId, classUpdate)
	if err != nil {
		h.logger.Error("Failed to update class", sl.Err(err))
		http.Error(w, "Failed to update class", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(class)
	if err != nil {
		h.logger.Error("Failed to encode response", sl.Err(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
