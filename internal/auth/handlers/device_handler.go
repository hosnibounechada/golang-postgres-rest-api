package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hosnibounechada/go-api/internal/auth/services"
	"github.com/hosnibounechada/go-api/internal/auth/util"
)

type DeviceHandler struct {
	deviceService services.DeviceService
}

func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{
		deviceService: *services.NewDeviceService(),
	}
}

func (h *DeviceHandler) GetAll(c *gin.Context) {
	user := util.BuildUserFromToken(c)

	devices, err := h.deviceService.GetAll(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, devices)
}

func (h *DeviceHandler) DisconnectDevice(c *gin.Context) {
	deviceIdStr := c.Param("id")

	deviceId, err := strconv.ParseInt(deviceIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid device ID",
		})
		return
	}

	// user := util.BuildUserFromToken(c)

	h.deviceService.DisconnectDevice(deviceId)

	c.JSON(http.StatusNoContent, nil)
}

func (h *DeviceHandler) DisconnectAllDevices(c *gin.Context) {
	user := util.BuildUserFromToken(c)

	h.deviceService.DisconnectAllDevices(user.ID)

	c.JSON(http.StatusNoContent, nil)
}
