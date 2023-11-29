package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/middleware"
	"github.com/hosnibounechada/go-api/internal/product/handlers"
)

func SetupProductsRoutes(v1 *gin.RouterGroup, productHandler handlers.ProductHandler) {
	productGroup := v1.Group("/products")

	productGroup.GET("/", middleware.AuthMiddleware(), productHandler.GetAll)
	productGroup.GET("/:id", middleware.AuthMiddleware(), productHandler.Get)
	productGroup.POST("/", middleware.AuthMiddleware(), productHandler.Create)
	productGroup.PUT("/:id", middleware.AuthMiddleware(), productHandler.Update)
	productGroup.DELETE("/:id", middleware.AuthMiddleware(), productHandler.Delete)
}
