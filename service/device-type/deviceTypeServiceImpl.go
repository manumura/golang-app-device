package devicetypeservice

import (
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/model/device-type"
)

// DeviceTypeServiceImpl : implementation for services on device type
type DeviceTypeServiceImpl struct {
	deviceTypeDao devicetypedao.DeviceTypeDao
}

// NewDeviceTypeService : Create a new instance of DeviceTypeService implemenation
func NewDeviceTypeService(deviceTypeDao devicetypedao.DeviceTypeDao) DeviceTypeService {
	return DeviceTypeServiceImpl{deviceTypeDao}
}

// FindDeviceTypes : retrieve device types
func (dts DeviceTypeServiceImpl) FindDeviceTypes() ([]devicetypemodel.DeviceType, error) {

	deviceTypes, err := dts.deviceTypeDao.FindDeviceTypes()
	return deviceTypes, err
}

// GetDeviceType : retrieve device type by id
func (dts DeviceTypeServiceImpl) GetDeviceType(id int) (devicetypemodel.DeviceType, error) {

	deviceType, err := dts.deviceTypeDao.GetDeviceType(id)
	return deviceType, err
}
