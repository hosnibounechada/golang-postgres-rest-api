package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/handlers"
	"github.com/hosnibounechada/go-api/internal/auth/middleware"
)

func SetupDevicesRoutes(v1 *gin.RouterGroup, deviceHandler handlers.DeviceHandler) {
	deviceGroup := v1.Group("/devices")

	deviceGroup.GET("/", middleware.AuthMiddleware(), deviceHandler.GetAll)
	deviceGroup.DELETE("/", middleware.AuthMiddleware(), deviceHandler.DisconnectAllDevices)
	deviceGroup.DELETE("/:id", middleware.AuthMiddleware(), deviceHandler.DisconnectDevice)
}
