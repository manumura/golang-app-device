package devicetypedao

import (
	"github.com/manumura/golang-app-device/model/device-type"
	"github.com/manumura/golang-app-device/config"
)

// DeviceTypeDao : interface defining DB operations on device type
type DeviceTypeDao interface {
	FindDeviceTypes() ([]devicetypemodel.DeviceType, error)
	GetDeviceType(id int) (devicetypemodel.DeviceType, error)
}

// NewDeviceTypeDao : Create a new instance of DeviceTypeDao implemenation
func NewDeviceTypeDao(db *config.DB) DeviceTypeDao {
	return DeviceTypeDaoImpl{db}
}
