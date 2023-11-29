package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/handlers"
	"github.com/hosnibounechada/go-api/internal/auth/middleware"
)

func SetupAuthRoutes(v1 *gin.RouterGroup, authHandler handlers.AuthHandler) {
	authGroup := v1.Group("/auth")

	authGroup.POST("/register", authHandler.Register)
	authGroup.GET("/account-confirmation/:token", middleware.AccountConfirmationMiddleware(), authHandler.AccountConfirmation)
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/me", middleware.AuthMiddleware(), authHandler.Me)
	authGroup.GET("/refresh", middleware.RefreshMiddleware(), authHandler.Refresh)
	authGroup.GET("/logout", middleware.RefreshMiddleware(), authHandler.Logout)
	authGroup.GET("/protected", middleware.AuthMiddleware(), authHandler.Protected)

}
