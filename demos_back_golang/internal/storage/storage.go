package storage

import (
	"context"
	"database/sql"
	"demos_back_golang/internal/config"
	"demos_back_golang/internal/models"
	"time"

	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type Storage struct {
	Users UserRepository
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		//Classes: &ClassStorage{db},
		Users: NewUserStorage(db),
	}
}

func NewDatabase(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.DSN())
	if err != nil {
		return nil, err
	}

	// Если будет необходимо можно настроить через config файл, так же
	// раскомментить в config.go
	/*
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)

		duration, err := time.ParseDuration(maxIdleTime)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxIdleTime(duration)
	*/

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
