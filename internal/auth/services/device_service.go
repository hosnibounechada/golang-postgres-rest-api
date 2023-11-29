package services

import (
	"github.com/hosnibounechada/go-api/internal/auth/models"
	database "github.com/hosnibounechada/go-api/internal/db"
)

type DeviceService struct{}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (s *DeviceService) GetAll(userId int64) ([]models.DeviceRes, error) {
	var devices []models.Device
	var deviceRes []models.DeviceRes

	if err := database.DB.Where("user_id = ?", userId).Find(&devices).Error; err != nil {
		return nil, err
	}

	for _, device := range devices {
		deviceRes = append(deviceRes, models.DeviceRes{
			ID:      device.ID,
			Name:    device.Name,
			OS:      device.OS,
			Browser: device.Browser,
		})
	}

	return deviceRes, nil
}

func (s *DeviceService) DisconnectDevice(userId int64) {

	database.DB.Where("id = ?", userId).Delete(&models.Device{})
}

func (s *DeviceService) DisconnectAllDevices(userId int64) {

	database.DB.Where("user_id = ?", userId).Delete(&models.Device{})
}
