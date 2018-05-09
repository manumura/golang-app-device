package service

import "github.com/manumura/golang-app-device/device-type/model"

// DeviceTypeService : interface defining services on device type
type DeviceTypeService interface {
	FindDeviceTypes() ([]model.DeviceType, error)
	GetDeviceType(id int) (model.DeviceType, error)
}

// NewDeviceTypeService : Create a new instance of DeviceTypeService implemenation
func NewDeviceTypeService() DeviceTypeService {
	return DeviceTypeServiceImpl{}
}
