package deviceservice

import (
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/model/device"
)

// DeviceServiceImpl : implementation for services on device
type DeviceServiceImpl struct {
	deviceDao devicedao.DeviceDao
}

// NewDeviceService : Create a new instance of DeviceService implemenation
func NewDeviceService(deviceDao devicedao.DeviceDao) DeviceService {
	return DeviceServiceImpl{deviceDao}
}

// FindDevices : retrieve devices
func (ds DeviceServiceImpl) FindDevices() ([]devicemodel.Device, error) {

	devices, err := ds.deviceDao.FindDevices()
	return devices, err
}

// GetDevice : retrieve device by id
func (ds DeviceServiceImpl) GetDevice(id int) (devicemodel.Device, error) {

	device, err := ds.deviceDao.GetDevice(id)
	return device, err
}

// Delete : retrieve one device
func (ds DeviceServiceImpl) Delete(id int) error {

	err := ds.deviceDao.Delete(id)
	return err
}

// Update : update one device
func (ds DeviceServiceImpl) Update(device devicemodel.Device) (devicemodel.Device, error) {

	device, err := ds.deviceDao.Update(device)
	return device, err
}

// Create : create one device
func (ds DeviceServiceImpl) Create(device devicemodel.Device) (devicemodel.Device, error) {

	device, err := ds.deviceDao.Create(device)
	return device, err
}

// FindDeviceStatuses : retrieve device statuses
func (ds DeviceServiceImpl) FindDeviceStatuses() ([]devicemodel.DeviceStatus, error) {

	deviceStatuses, err := ds.deviceDao.FindDeviceStatuses()
	return deviceStatuses, err
}
