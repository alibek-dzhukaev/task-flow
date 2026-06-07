package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/alibek-dzhukaev/task-flow/internal/handler"
	"github.com/alibek-dzhukaev/task-flow/internal/util"
)

type contextKey string

const UserClaimsKey contextKey = "user_claims"

func Auth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				handler.Fail(w, http.StatusUnauthorized, "missing or invalid authorization header")
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := util.ParseToken(tokenStr, jwtSecret)
			if err != nil {
				handler.Fail(w, http.StatusUnauthorized, "invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
