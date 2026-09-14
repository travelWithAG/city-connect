package middleware

import (
	"connect/core/response"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	RoleKey ContextKey = "role"
)
type CustomClaims struct {
	UserID string `json:"user_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleWare(secretKey string) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == ""{
				response.ErrorOccured(w, http.StatusUnauthorized, false, "You can't access this part without being authorized, this will be reported", nil)
				return
			}

			authHeaderParts := strings.Split(authHeader, " ")

			if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer"{
				response.JSON(w, http.StatusUnauthorized, false, "Invalid Request Performed This Will Be Reported", nil)
				return 
			}

			tokenString := authHeaderParts[1]

			claims := &CustomClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token)(interface{}, error)  {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Unexpecting signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid{
				response.ErrorOccured(w, http.StatusUnauthorized, false, "Invalid or Token Expired, Try to reflesh page.", nil)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

			 next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func LoggingInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}