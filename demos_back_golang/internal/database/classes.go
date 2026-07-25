package database

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

type Class struct {
	ID        int64           `json:"id"`
	DateTime  string          `json:"datetime"`
	Location  string          `json:"location"`
	Price     decimal.Decimal `json:"price"`
	TrainerID int64           `json:"trainer_id"`
	CreatedAt string          `json:"created_at"`
	UpdatedAT string          `json:"updated_at"`
}

type ClassStorage struct {
	db *sql.DB
}

func (s *ClassStorage) Create(ctx context.Context, class *Class) error {
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
