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
	CreateUser(ctx context.Context, request models.RegisterRequest) (models.UserResponse, error)
	GetUser(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
}

type ClassRepository interface {
	CreateClass(ctx context.Context, request models.CreateClassRequest) error
	GetClassById(ctx context.Context, classId int64) (models.ClassResponse, error)
	//GetAllClasses(ctx context.Context) (models.ClassResponse[], error)
	DeleteClass(ctx context.Context, classId int64) error
	UpdateClass(ctx context.Context, classId int64, class models.UpdateClassRequest) (models.ClassResponse, error)
}

type Storage struct {
	Users         UserRepository
	Classes       ClassRepository
	RefreshTokens *RefreshTokenStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Classes:       NewClassStorage(db),
		Users:         NewUserStorage(db),
		RefreshTokens: NewRefreshTokenStorage(db),
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
