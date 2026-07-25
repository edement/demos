package database

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Class struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAT string   `json:"updated_at"`
}

type ClassStorage struct {
	db *sql.DB
}

func (s *ClassStorage) Create(ctx context.Context, post *Class) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		post.Tags,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAT,
	)

	if err != nil {
		return err
	}

	return nil
}
