package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AccountConfirmationMiddleware() gin.HandlerFunc {
	tokenSecretKey := []byte(os.Getenv("JWT_CONFIRMATION_KEY"))

	return func(c *gin.Context) {
		tokenString := c.Param("token")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Confirmation token is missing or invalid"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		id, _ := claims["i"].(float64)

		userID := int64(id)

		c.Set("userId", userID)

		c.Next()
	}
}
