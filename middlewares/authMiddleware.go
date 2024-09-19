package middlewares

import (
	"context"
	"josk/task-management-system/auth"
	"josk/task-management-system/utils"
	"net/http"
	"strings"
)

type key string

const UserKey key = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.SendJSONResponse(w, map[string]string{"error": "Authorization header is missing"}, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.VerifyJWT(tokenString)
		if err != nil {
			utils.SendJSONResponse(w, map[string]string{"error": "Invalid token"}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
