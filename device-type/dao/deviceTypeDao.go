package dao

import "github.com/manumura/golang-app-device/device-type/model"

// DeviceTypeDao : interface defining DB operations on device type
type DeviceTypeDao interface {
	FindDeviceTypes() ([]model.DeviceType, error)
	GetDeviceType(id int) (model.DeviceType, error)
}

// NewDeviceTypeDao : Create a new instance of DeviceTypeDao implemenation
func NewDeviceTypeDao() DeviceTypeDao {
	return DeviceTypeDaoImpl{}
}
