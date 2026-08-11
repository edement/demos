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

func (s *ClassStorage) CreateClass(ctx context.Context, class models.CreateClassRequest) error {
	query := `
		INSERT INTO classes (datetime, location, price, trainer_id)
		VALUES ($1, $2, $3, $4)
	`
	trainerId := 1 // Заглушка до появления JWT

	_, err := s.db.ExecContext(
		ctx,
		query,
		class.DateTime,
		class.Location,
		class.Price,
		trainerId,
	)

	return err
}
