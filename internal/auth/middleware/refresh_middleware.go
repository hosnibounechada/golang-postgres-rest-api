package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RefreshMiddleware() gin.HandlerFunc {
	tokenSecretKey := []byte(os.Getenv("JWT_REFRESH_KEY"))

	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")

		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is missing or invalid"})
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

		userID := int64(id)
		tokenStr, _ := claims["u"].(string)

		c.Set("userId", userID)
		c.Set("token", tokenStr)

		c.Next()
	}
}
