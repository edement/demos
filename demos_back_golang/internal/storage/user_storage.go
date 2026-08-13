package storage

import (
	"context"
	"database/sql"
	"demos_back_golang/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(ctx context.Context, request models.RegisterRequest) (models.UserResponse, error) {
	query := `
		INSERT INTO users (username, password_hash, email)
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	user := models.UserResponse{
		Username:  request.Username,
		Email:     request.Email,
		IsTrainer: request.IsTrainer,
	}

	err := s.db.QueryRowContext(
		ctx,
		query,
		request.Username,
		request.Password,
		request.Email,
	).Scan(
		&user.ID,
	)

	return user, err
}

func (s *UserStorage) GetUser(ctx context.Context, email string) (models.User, error) {
	query := `
		SELECT id, username, password_hash, is_trainer FROM users 
		WHERE email = $1
	`
	user := models.User{
		Email: email,
	}

	err := s.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.IsTrainer,
	)
	// TODO: Проверка ошибки Системная или Пользователь не найден
	return user, err
}

func (s *UserStorage) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	query := `
        SELECT id, username, email, password_hash, is_trainer
        FROM users
        WHERE id = $1
    `
	user := models.User{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsTrainer,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
