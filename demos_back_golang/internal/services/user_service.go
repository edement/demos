package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"demos_back_golang/internal/lib/jwt"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/storage"
)

type UserService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, string, string, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.UserResponse, string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
	Logout(ctx context.Context, userID int64) error
	Me(ctx context.Context, userID int64) (*models.UserResponse, error)
}

type userService struct {
	userRepo      storage.UserRepository
	refreshTokens *storage.RefreshTokenStorage
	jwtService    *jwt.JwtService
}

func NewUserService(
	userRepo storage.UserRepository,
	refreshTokens *storage.RefreshTokenStorage,
	jwtService *jwt.JwtService,
) UserService {
	return &userService{
		userRepo:      userRepo,
		refreshTokens: refreshTokens,
		jwtService:    jwtService,
	}
}

func (s *userService) Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, string, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", err
	}
	req.Password = string(hashedPassword)

	user, err := s.userRepo.CreateUser(ctx, req)
	if err != nil {
		return nil, "", "", err
	}

	accessToken, refreshToken, err := s.generateTokens(ctx, user.ID, user.Username, user.Email)
	if err != nil {
		return nil, "", "", err
	}

	return &user, accessToken, refreshToken, nil
}

func (s *userService) Login(ctx context.Context, req models.LoginRequest) (*models.UserResponse, string, string, error) {
	user, err := s.userRepo.GetUser(ctx, req.Email)
	if err != nil {
		return nil, "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("invalid email or password")
	}

	accessToken, refreshToken, err := s.generateTokens(ctx, user.ID, user.Username, user.Email)
	if err != nil {
		return nil, "", "", err
	}

	resp := &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsTrainer: user.IsTrainer,
	}

	return resp, accessToken, refreshToken, nil
}

func (s *userService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := s.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	userID, familyID, err := s.refreshTokens.GetUserIDByToken(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	if userID == 0 {
		return "", "", errors.New("invalid refresh token")
	}

	if userID != claims.UserID {
		return "", "", errors.New("refresh token user_id mismatch")
	}

	if err := s.refreshTokens.MarkTokenAsUsed(ctx, refreshToken); err != nil {
		return "", "", err
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	newAccessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Username, user.Email)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, familyID)
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	if err := s.refreshTokens.Save(ctx, user.ID, newRefreshToken, expiresAt, familyID); err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *userService) Logout(ctx context.Context, userID int64) error {
	return s.refreshTokens.RevokeAllUserTokens(ctx, userID)
}

func (s *userService) Me(ctx context.Context, userID int64) (*models.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsTrainer: user.IsTrainer,
	}, nil
}

func (s *userService) generateTokens(ctx context.Context, userID int64, username string, email string) (string, string, error) {
	familyID := uuid.New()

	accessToken, err := s.jwtService.GenerateAccessToken(userID, username, email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(userID, familyID)
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	if err := s.refreshTokens.Save(ctx, userID, refreshToken, expiresAt, familyID); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
