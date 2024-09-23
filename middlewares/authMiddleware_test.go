package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"josk/task-management-system/auth"
	"josk/task-management-system/models"
	"josk/task-management-system/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var jwtKey = []byte("your_secret_key")

func generateTestToken(claims *models.Claims, signingMethod jwt.SigningMethod, key []byte) (string, error) {
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString(key)
}

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

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalidTokenString"
		_, err := auth.VerifyJWT(token)

		assert.NotNil(t, err)
		assert.Equal(t, "token is malformed: token contains an invalid number of segments", err.Error())
	})

	t.Run("Valid token", func(t *testing.T) {
		expirationTime := time.Now().Add(24 * time.Hour)

		claims := &models.Claims{
			UserID: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			}}

		token, _ := generateTestToken(claims, jwt.SigningMethodHS256, jwtKey)
		resultClaims, err := auth.VerifyJWT(token)

		assert.Nil(t, err)
		assert.Equal(t, claims.UserID, resultClaims.UserID)
	})

	t.Run("Expired token", func(t *testing.T) {
		expirationTime := time.Now().Add(-1 * time.Hour)

		claims := &models.Claims{
			UserID: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			}}

		token, _ := generateTestToken(claims, jwt.SigningMethodHS256, jwtKey)
		_, err := auth.VerifyJWT(token)

		assert.NotNil(t, err)
		assert.Equal(t, "token has invalid claims: token is expired", err.Error())
	})
}
