package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/models"
	"github.com/hosnibounechada/go-api/internal/auth/services"
	"github.com/hosnibounechada/go-api/internal/auth/util"
	pkgUtil "github.com/hosnibounechada/go-api/pkg/util"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: *services.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerPayload models.RegisterUserReq

	if err := c.ShouldBindJSON(&registerPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": pkgUtil.FormatValidationErrors(err)})
		return
	}

	createUserRes, err := h.authService.Register(registerPayload)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createUserRes)
}

func (h *AuthHandler) AccountConfirmation(c *gin.Context) {

	userID, _ := c.Get("userId")

	userId, _ := userID.(int64)

	h.authService.AccountConfirmation(userId)

	c.JSON(http.StatusNoContent, nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginPayload models.LoginUserReq

	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": pkgUtil.FormatValidationErrors(err)})
		return
	}

	loginRes, err := h.authService.Login(loginPayload)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", loginRes.Jwt.RefreshToken, 0, "/", "", false, true)

	c.JSON(http.StatusOK, loginRes)
}

func (h *AuthHandler) Logout(c *gin.Context) {

	userID, idExists := c.Get("userId")
	token, tokenExists := c.Get("token")

	if !idExists || !tokenExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID Or Token not found in the context"})
		c.Abort()
		return
	}

	userId, okUserId := userID.(int64)
	tokenStr, okToken := token.(string)

	if !okUserId || !okToken {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userID type"})
		return
	}

	h.authService.Logout(userId, tokenStr)

	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusNoContent, nil)
}

func (h *AuthHandler) Protected(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in the context"})
		c.Abort()
		return
	}

	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user object in the context"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userObj})
}

func (h *AuthHandler) Me(c *gin.Context) {

	user := util.BuildUserFromToken(c)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) Refresh(c *gin.Context) {

	userID, _ := c.Get("userId")
	token, _ := c.Get("token")

	userId, _ := userID.(int64)
	tokenStr, _ := token.(string)

	accessToken, err := h.authService.Refresh(userId, tokenStr)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
