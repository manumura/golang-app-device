package devicedao

import (
	"github.com/manumura/golang-app-device/model/device"
	"github.com/manumura/golang-app-device/config"
)

// DeviceDao : interface defining DB operations on device
type DeviceDao interface {
	FindDevices() ([]devicemodel.Device, error)
	GetDevice(id int) (devicemodel.Device, error)
	Delete(id int) error
	Update(devicemodel.Device) (devicemodel.Device, error)
	Create(devicemodel.Device) (devicemodel.Device, error)
	FindDeviceStatuses() ([]devicemodel.DeviceStatus, error)
}

// NewDeviceDao : Create a new instance of DeviceDao implemenation
func NewDeviceDao(db *config.DB) DeviceDao {
	return DeviceDaoImpl{db}
}
