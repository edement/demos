package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type RefreshTokenStorage struct {
	db *sql.DB
}

func NewRefreshTokenStorage(db *sql.DB) *RefreshTokenStorage {
	return &RefreshTokenStorage{db: db}
}

func (s *RefreshTokenStorage) Save(ctx context.Context, userID int64, token string, expiresAt time.Time, familyID uuid.UUID) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at, family_id, used)
		VALUES ($1, $2, $3, $4, FALSE)
	`
	_, err := s.db.ExecContext(ctx, query, userID, token, expiresAt, familyID)
	return err
}

func (s *RefreshTokenStorage) GetUserIDByToken(ctx context.Context, token string) (int64, uuid.UUID, error) {
	query := `
		SELECT user_id, family_id FROM refresh_tokens
		WHERE token = $1 
		  AND revoked = FALSE 
		  AND used = FALSE
		  AND expires_at > NOW()
	`
	var userID int64
	var familyID uuid.UUID
	err := s.db.QueryRowContext(ctx, query, token).Scan(&userID, &familyID)
	if err == sql.ErrNoRows {
		return 0, uuid.Nil, nil
	}
	if err != nil {
		return 0, uuid.Nil, err
	}
	return userID, familyID, nil
}

func (s *RefreshTokenStorage) MarkTokenAsUsed(ctx context.Context, token string) error {
	query := `UPDATE refresh_tokens SET used = TRUE WHERE token = $1 AND used = FALSE`
	_, err := s.db.ExecContext(ctx, query, token)
	return err
}

func (s *RefreshTokenStorage) RevokeAllTokensInFamily(ctx context.Context, familyID uuid.UUID) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE family_id = $1 AND revoked = FALSE`
	_, err := s.db.ExecContext(ctx, query, familyID)
	return err
}

func (s *RefreshTokenStorage) RevokeAllUserTokens(ctx context.Context, userID int64) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE user_id = $1 AND revoked = FALSE`
	_, err := s.db.ExecContext(ctx, query, userID)
	return err
}
