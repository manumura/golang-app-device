package devicetypedao

import (
	"errors"
	"log"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/model/device-type"
)

// DeviceTypeDaoImpl : implementation for DB operations on device type
type DeviceTypeDaoImpl struct {
}

// FindDeviceTypes : retrieve device types from the database
func (cd DeviceTypeDaoImpl) FindDeviceTypes() ([]devicetypemodel.DeviceType, error) {

	rows, err := config.Database.Query("SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deviceTypes := []devicetypemodel.DeviceType{}
	for rows.Next() {
		deviceType := devicetypemodel.DeviceType{}
		err := rows.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description) // order matters
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, deviceType)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Println(deviceTypes)
	return deviceTypes, nil
}

// GetDeviceType : retrieve device type by id from the database
func (cd DeviceTypeDaoImpl) GetDeviceType(id int) (devicetypemodel.DeviceType, error) {

	deviceType := devicetypemodel.DeviceType{}

	if id == 0 {
		return deviceType, errors.New("id cannot be empty")
	}

	row := config.Database.QueryRow("SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt WHERE dt.device_type_id = $1", id)

	err := row.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description)
	if err != nil {
		return deviceType, err
	}

	log.Println(deviceType)
	return deviceType, nil
}
