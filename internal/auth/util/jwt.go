package util

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/models"
)

func GenerateAccessToken(user models.User) (string, error) {
	accessTokenKey := []byte(os.Getenv("JWT_ACCESS_KEY"))

	hours, err := strconv.Atoi(os.Getenv("JWT_ACCESS_DURATION"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"i":   user.ID,
		"u":   user.Username,
		"e":   user.Email,
		"fn":  user.FirstName,
		"ln":  user.LastName,
		"exp": time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	})

	tokenString, err := token.SignedString(accessTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(id int64, uuid string) (string, error) {
	refreshTokenKey := []byte(os.Getenv("JWT_REFRESH_KEY"))

	months, err := strconv.Atoi(os.Getenv("JWT_REFRESH_DURATION"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"i":   id,
		"u":   uuid,
		"exp": time.Now().AddDate(0, months, 0).Unix(),
	})

	tokenString, err := token.SignedString(refreshTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateConfirmationToken(id int64) (string, error) {
	refreshTokenKey := []byte(os.Getenv("JWT_CONFIRMATION_KEY"))

	months, err := strconv.Atoi(os.Getenv("JWT_CONFIRMATION_DURATION"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"i":   id,
		"exp": time.Now().AddDate(0, months, 0).Unix(),
	})

	tokenString, err := token.SignedString(refreshTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func BuildUserFromToken(c *gin.Context) models.User {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in the context"})
		c.Abort()
		return models.User{}
	}

	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user object in the context"})
		c.Abort()
		return models.User{}
	}
	return userObj
}
