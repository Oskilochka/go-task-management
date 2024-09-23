package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"josk/task-management-system/models"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	t.Run("JWT generation", func(t *testing.T) {
		userID := uint(1)
		username := "username"

		tokenString, err := GenerateJWT(userID, username)

		assert.Nil(t, err)
		assert.NotEmpty(t, tokenString)

		claims, err := VerifyJWT(tokenString)
		assert.Nil(t, err)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, username, claims.Username)
	})

}

func TestVerifyJWT(t *testing.T) {
	t.Run("Valid token", func(t *testing.T) {
		tokenString, err := GenerateJWT(1, "testuser")
		assert.NoError(t, err)

		claims, err := VerifyJWT(tokenString)
		assert.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, uint(1), claims.UserID)
		assert.Equal(t, "testuser", claims.Username)
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalidTokenString"
		_, err := VerifyJWT(token)

		assert.NotNil(t, err)
		assert.Equal(t, "token is malformed: token contains an invalid number of segments", err.Error())
	})

	t.Run("Expired token", func(t *testing.T) {
		claims := &models.Claims{
			UserID:   1,
			Username: "username",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-24 * time.Hour)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(jwtKey)

		_, err := VerifyJWT(tokenString)
		assert.NotNil(t, err)
		assert.Equal(t, "token has invalid claims: token is expired", err.Error())
	})
}
