package middleware

import (
	"context"
	"demos_back_golang/internal/lib/jwt"
	"demos_back_golang/internal/models"
	"net/http"
	"strings"
)

type claimsKey string

const UserClaimsKey claimsKey = "user"

type AuthMiddleware struct {
	jwtService *jwt.JwtService
}

func NewAuthMiddleware(jwtService *jwt.JwtService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			claims, err := jwtService.ValidateAccessToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			userCtx := models.UserClaims{
				UserID:   claims.UserID,
				Username: claims.Username,
				Email:    claims.Email,
			}

			ctx := context.WithValue(r.Context(), UserClaimsKey, userCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
