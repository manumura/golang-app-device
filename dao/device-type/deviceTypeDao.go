package devicetypedao

import devicetype "github.com/manumura/golang-app-device/model/device-type"

// DeviceTypeDao : interface defining DB operations on device type
type DeviceTypeDao interface {
	FindDeviceTypes() ([]devicetype.DeviceType, error)
	GetDeviceType(id int) (devicetype.DeviceType, error)
}

// NewDeviceTypeDao : Create a new instance of DeviceTypeDao implemenation
func NewDeviceTypeDao() DeviceTypeDao {
	return DeviceTypeDaoImpl{}
}
