package dao

import "github.com/manumura/golang-app-device/device/model"

// DeviceDao : interface defining DB operations on device
type DeviceDao interface {
	FindDevices() ([]model.Device, error)
	GetDevice(id int) (model.Device, error)
	Delete(id int) error
	Update(model.Device) (model.Device, error)
	Create(model.Device) (model.Device, error)
	FindDeviceStatuses(id int) ([]model.DeviceStatus, error)
}

// NewDeviceDao : Create a new instance of DeviceDao implemenation
func NewDeviceDao() DeviceDao {
	return DeviceDaoImpl{}
}
