package dao

import (
	"errors"
	"fmt"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/device-type/model"
)

// DeviceTypeDaoImpl : implementation for DB operations on device type
type DeviceTypeDaoImpl struct {
}

// FindDeviceTypes : retrieve device types from the database
func (cd DeviceTypeDaoImpl) FindDeviceTypes() ([]model.DeviceType, error) {

	rows, err := config.Database.Query("SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deviceTypes := []model.DeviceType{}
	for rows.Next() {
		deviceType := model.DeviceType{}
		err := rows.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description) // order matters
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, deviceType)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(deviceTypes)
	return deviceTypes, nil
}

// GetDeviceType : retrieve device type by id from the database
func (cd DeviceTypeDaoImpl) GetDeviceType(id int) (model.DeviceType, error) {

	deviceType := model.DeviceType{}

	if id == 0 {
		return deviceType, errors.New("id cannot be empty")
	}

	row := config.Database.QueryRow("SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt WHERE dt.device_type_id = $1", id)

	err := row.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description)
	if err != nil {
		return deviceType, err
	}

	fmt.Println(deviceType)
	return deviceType, nil
}
