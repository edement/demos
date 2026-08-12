package storage

import (
	"context"
	"database/sql"
	"time"
)

type RefreshTokenStorage struct {
	db *sql.DB
}

func NewRefreshTokenStorage(db *sql.DB) *RefreshTokenStorage {
	return &RefreshTokenStorage{db: db}
}

func (s *RefreshTokenStorage) Save(ctx context.Context, userID int64, token string, expiresAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err := s.db.ExecContext(ctx, query, userID, token, expiresAt)
	return err
}

func (s *RefreshTokenStorage) GetUserIDByToken(ctx context.Context, token string) (int64, error) {
	query := `
		SELECT user_id FROM refresh_tokens
		WHERE token = $1 
		  AND revoked = FALSE 
		  AND expires_at > NOW()
	`
	var userID int64
	err := s.db.QueryRowContext(ctx, query, token).Scan(&userID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// Revoke удаляет/отзывает все refresh токены пользователя (при выходе)
func (s *RefreshTokenStorage) RevokeAllUserTokens(ctx context.Context, userID int64) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE user_id = $1 AND revoked = FALSE`
	_, err := s.db.ExecContext(ctx, query, userID)
	return err
}
