package deviceservice

import (
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/model/device"
)

// DeviceServiceImpl : implementation for services on device
type DeviceServiceImpl struct {
}

// FindDevices : retrieve devices
func (ds DeviceServiceImpl) FindDevices() ([]devicemodel.Device, error) {

	deviceDao := devicedao.NewDeviceDao()
	devices, err := deviceDao.FindDevices()
	return devices, err
}

// GetDevice : retrieve device by id
func (ds DeviceServiceImpl) GetDevice(id int) (devicemodel.Device, error) {

	deviceDao := devicedao.NewDeviceDao()
	device, err := deviceDao.GetDevice(id)
	return device, err
}

// Delete : retrieve one device
func (ds DeviceServiceImpl) Delete(id int) error {

	deviceDao := devicedao.NewDeviceDao()
	err := deviceDao.Delete(id)
	return err
}

// Update : update one device
func (ds DeviceServiceImpl) Update(device devicemodel.Device) (devicemodel.Device, error) {

	deviceDao := devicedao.NewDeviceDao()
	device, err := deviceDao.Update(device)
	return device, err
}

// Create : create one device
func (ds DeviceServiceImpl) Create(device devicemodel.Device) (devicemodel.Device, error) {

	deviceDao := devicedao.NewDeviceDao()
	device, err := deviceDao.Create(device)
	return device, err
}

// FindDeviceStatuses : retrieve device statuses
func (ds DeviceServiceImpl) FindDeviceStatuses() ([]devicemodel.DeviceStatus, error) {

	deviceDao := devicedao.NewDeviceDao()
	deviceStatuses, err := deviceDao.FindDeviceStatuses()
	return deviceStatuses, err
}
