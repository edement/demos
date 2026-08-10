package storage

import (
	"context"
	"database/sql"
	"demos_back_golang/internal/models"
)

type ClassStorage struct {
	db *sql.DB
}

func NewClassStorage(db *sql.DB) *ClassStorage {
	return &ClassStorage{db: db}
}

func (s *ClassStorage) Create(ctx context.Context, class *models.Class) error {
	query := `
		INSERT INTO classes (datetime, location, price, trainer_id)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		class.DateTime,
		class.Location,
		class.Price,
		class.TrainerID,
	).Scan(
		&class.ID,
		&class.CreatedAt,
		&class.UpdatedAT,
	)

	if err != nil {
		return err
	}

	return nil
}
