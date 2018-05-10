package devicetypeservice

import (
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/model/device-type"
)

// DeviceTypeServiceImpl : implementation for services on device type
type DeviceTypeServiceImpl struct {
}

// FindDeviceTypes : retrieve device types
func (dts DeviceTypeServiceImpl) FindDeviceTypes() ([]devicetypemodel.DeviceType, error) {

	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceTypes, err := deviceTypeDao.FindDeviceTypes()
	return deviceTypes, err
}

// GetDeviceType : retrieve device type by id
func (dts DeviceTypeServiceImpl) GetDeviceType(id int) (devicetypemodel.DeviceType, error) {

	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceType, err := deviceTypeDao.GetDeviceType(id)
	return deviceType, err
}
