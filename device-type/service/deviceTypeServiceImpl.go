package service

import (
	"github.com/manumura/golang-app-device/device-type/dao"
	"github.com/manumura/golang-app-device/device-type/model"
)

// DeviceTypeServiceImpl : implementation for services on device type
type DeviceTypeServiceImpl struct {
}

// FindDeviceTypes : retrieve device types
func (dts DeviceTypeServiceImpl) FindDeviceTypes() ([]model.DeviceType, error) {

	deviceTypeDao := dao.NewDeviceTypeDao()
	deviceTypes, err := deviceTypeDao.FindDeviceTypes()
	return deviceTypes, err
}

// GetDeviceType : retrieve device type by id
func (dts DeviceTypeServiceImpl) GetDeviceType(id int) (model.DeviceType, error) {

	deviceTypeDao := dao.NewDeviceTypeDao()
	deviceType, err := deviceTypeDao.GetDeviceType(id)
	return deviceType, err
}
