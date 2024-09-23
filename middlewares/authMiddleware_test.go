package middlewares

import (
	"github.com/stretchr/testify/assert"
	"josk/task-management-system/auth"
	"josk/task-management-system/models"
	"josk/task-management-system/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockNextHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONResponse(w, map[string]string{"message": "success"}, http.StatusOK)
}

func TestJWTMiddleware(t *testing.T) {
	t.Run("Missing Authorization header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		w := httptest.NewRecorder()

		handler := JWTMiddleware(http.HandlerFunc(mockNextHandler))
		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Authorization header is missing")
	})

	t.Run("Invalid Token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer invalidToken")
		w := httptest.NewRecorder()

		handler := JWTMiddleware(http.HandlerFunc(mockNextHandler))
		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token")
	})

	t.Run("Valid Token", func(t *testing.T) {
		tokenString, _ := auth.GenerateJWT(1, "testuser")

		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		w := httptest.NewRecorder()

		var capturedRequest *http.Request
		handler := JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			capturedRequest = r
			mockNextHandler(w, r)
		}))

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")

		if capturedRequest == nil {
			t.Fatal("Request was not captured")
		}

		userClaims, ok := capturedRequest.Context().Value(UserKey).(*models.Claims)
		if !ok || userClaims == nil {
			t.Fatal("Expected claims in context, but got nil")
		}

		assert.Equal(t, uint(1), userClaims.UserID)
		assert.Equal(t, "testuser", userClaims.Username)
	})
}
