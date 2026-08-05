package save

import (
	"context"
	"log/slog"
	"net/http"

	"demos_back_golang/internal/database"
)

type ClassCreate interface {
	Create(ctx context.Context, class *database.Class) error
}

func New(log *slog.Logger, classCreate ClassCreate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
