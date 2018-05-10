package deviceservice

import (
	"github.com/manumura/golang-app-device/model/device"
)

// DeviceService : interface defining services on device
type DeviceService interface {
	FindDevices() ([]devicemodel.Device, error)
	GetDevice(id int) (devicemodel.Device, error)
	Delete(id int) error
	Update(devicemodel.Device) (devicemodel.Device, error)
	Create(devicemodel.Device) (devicemodel.Device, error)
	FindDeviceStatuses() ([]devicemodel.DeviceStatus, error)
}

// NewDeviceService : Create a new instance of DeviceService implemenation
func NewDeviceService() DeviceService {
	return DeviceServiceImpl{}
}
