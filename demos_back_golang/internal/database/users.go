package database

import (
	"context"
	"database/sql"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at"`
}

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, paswordhash, email)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.PasswordHash,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
