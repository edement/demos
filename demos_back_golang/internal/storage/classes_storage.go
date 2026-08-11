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

func (s *ClassStorage) GetClassById(ctx context.Context, classId int64) (models.ClassResponse, error) {
	query := `
		SELECT classes.id, datetime, location, price, users.id, users.username FROM classes
		JOIN users ON classes.trainer_id = users.id
		WHERE classes.id = $1
	`
	class := models.ClassResponse{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		classId,
	).Scan(
		&class.ID,
		&class.DateTime,
		&class.Location,
		&class.Price,
		&class.Trainer.ID,
		&class.Trainer.Username,
	)

	return class, err
}
