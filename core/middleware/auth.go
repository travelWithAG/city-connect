package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"connect/core/response"

	"github.com/golang-jwt/jwt/v5"
)

// AuthUser aggregates context payload to allow single-allocation context storage.
type AuthUser struct {
	UserID string
	Role   string
}

// Private context key type avoids memory/collision collisions across packages.
type contextKey struct{}

var authUserKey = contextKey{}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthMiddleware validates JWT Bearer tokens and attaches extracted identities to context.
func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	secretBytes := []byte(secretKey) // Pre-allocate byte slice once at initialization

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.ErrorOccured(w, http.StatusUnauthorized, false, "Authorization header required", nil)
				return
			}

			// strings.Cut avoids slice allocation from strings.Split
			prefix, tokenString, found := strings.Cut(authHeader, " ")
			if !found || !strings.EqualFold(prefix, "bearer") || tokenString == "" {
				response.JSON(w, http.StatusUnauthorized, false, "Invalid Authorization header format", nil)
				return
			}

			claims := &CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				// Prevent Algorithm Confusion Attacks (e.g., None or RS256 spoofing)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return secretBytes, nil
			})

			if err != nil || !token.Valid {
				response.ErrorOccured(w, http.StatusUnauthorized, false, "Invalid or expired token", nil)
				return
			}

			// Single-allocation context value node instead of chained context.WithValue calls
			user := AuthUser{
				UserID: claims.UserID,
				Role:   claims.Role,
			}
			ctx := context.WithValue(r.Context(), authUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// LoggingMiddleware logs incoming HTTP requests for diagnostic and auditing purposes
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call next handler in chain
		next.ServeHTTP(w, r)
	})
}

// --- Context Helper Functions ---

// GetAuthUser retrieves the authenticated user struct from the request context.
func GetAuthUser(ctx context.Context) (AuthUser, bool) {
	user, ok := ctx.Value(authUserKey).(AuthUser)
	return user, ok
}

// GetUserID retrieves the User ID directly from the request context.
func GetUserID(ctx context.Context) (string, bool) {
	if user, ok := GetAuthUser(ctx); ok {
		return user.UserID, true
	}
	return "", false
}

// GetRole retrieves the User Role directly from the request context.
func GetRole(ctx context.Context) (string, bool) {
	if user, ok := GetAuthUser(ctx); ok {
		return user.Role, true
	}
	return "", false
}