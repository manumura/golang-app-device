package devicetypeservice

import (
	devicetypemodel "github.com/manumura/golang-app-device/model/device-type"
)

// DeviceTypeService : interface defining services on device type
type DeviceTypeService interface {
	FindDeviceTypes() ([]devicetypemodel.DeviceType, error)
	GetDeviceType(id int) (devicetypemodel.DeviceType, error)
}

// NewDeviceTypeService : Create a new instance of DeviceTypeService implemenation
func NewDeviceTypeService() DeviceTypeService {
	return DeviceTypeServiceImpl{}
}
