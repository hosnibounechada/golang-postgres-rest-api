package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/models"
)

func AuthMiddleware() gin.HandlerFunc {
	tokenSecretKey := []byte(os.Getenv("JWT_ACCESS_KEY"))

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		id, ok := claims["i"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid 'id' claim"})
			c.Abort()
			return
		}

		user := models.User{
			ID:        int64(claims["i"].(float64)),
			Username:  claims["u"].(string),
			Email:     claims["e"].(string),
			FirstName: claims["fn"].(string),
			LastName:  claims["ln"].(string),
		}

		userID := int64(id)

		c.Set("userID", userID)
		c.Set("user", user)

		c.Next()
	}
}
