package middlewares

import (
	"github.com/stretchr/testify/assert"
	"josk/task-management-system/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserKey)
	if userID == nil {
		utils.SendJSONResponse(w, map[string]string{"error": "No user in context"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"message": "Success", "userID": userID.(string)}, http.StatusOK)
}

func TestJWTMiddleware(t *testing.T) {
	t.Run("No Authorization header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/some-endpoint", nil)
		rr := httptest.NewRecorder()

		handler := JWTMiddleware(http.HandlerFunc(MockHandler))
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.JSONEq(t, `{"error": "Authorization header is missing"}`, rr.Body.String())
	})
}
