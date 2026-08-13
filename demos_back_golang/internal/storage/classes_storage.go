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

func (s *ClassStorage) CreateClass(ctx context.Context, class models.CreateClassRequest, trainerID int64) error {
	query := `
		INSERT INTO classes (datetime, location, price, trainer_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		class.DateTime,
		class.Location,
		class.Price,
		trainerID,
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

func (s *ClassStorage) DeleteClass(ctx context.Context, classId int64) error {
	query := `
		DELETE FROM classes
		WHERE classes.id = $1
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		classId,
	)

	return err
}

func (s *ClassStorage) UpdateClass(ctx context.Context, classId int64, classUpdate models.UpdateClassRequest) (models.ClassResponse, error) {
	query := `
		WITH updated AS (
			UPDATE classes
			SET 
				datetime = COALESCE($1, datetime),
				location = COALESCE($2, location),
				price = COALESCE($3, price)
			WHERE classes.id = $4
			RETURNING id, datetime, location, price, trainer_id
		)
		SELECT
			updated.id,
			updated.datetime,
			updated.location,
			updated.price,
			users.id,
			users.username
		FROM updated
		JOIN users ON updated.trainer_id = users.id 
	`
	class := models.ClassResponse{}

	err := s.db.QueryRowContext(
		ctx,
		query,
		classUpdate.DateTime,
		classUpdate.Location,
		classUpdate.Price,
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
