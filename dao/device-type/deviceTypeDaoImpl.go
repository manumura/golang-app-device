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
func (dtd DeviceTypeDaoImpl) FindDeviceTypes() ([]devicetypemodel.DeviceType, error) {

	sql := "SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt"
	//rows, err := config.Database.Query(sql)

	stmt, err := config.Database.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var deviceTypes []devicetypemodel.DeviceType
	for rows.Next() {
		deviceType := devicetypemodel.DeviceType{}
		err := rows.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description) // order matters
		if err != nil {
			log.Println(err)
			return nil, err
		}
		deviceTypes = append(deviceTypes, deviceType)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(deviceTypes)
	return deviceTypes, nil
}

// GetDeviceType : retrieve device type by id from the database
func (dtd DeviceTypeDaoImpl) GetDeviceType(id int) (devicetypemodel.DeviceType, error) {

	deviceType := devicetypemodel.DeviceType{}

	if id == 0 {
		return deviceType, errors.New("id cannot be empty")
	}

	sql := "SELECT dt.device_type_id, dt.name, dt.description FROM device_type dt WHERE dt.device_type_id = $1"
	//row := config.Database.QueryRow(sql, id)

	//err := row.Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description)

	stmt, err := config.Database.Prepare(sql)
	if err != nil {
		log.Println(err)
		return deviceType, err
	}

	err = stmt.QueryRow(id).Scan(&deviceType.ID, &deviceType.Name, &deviceType.Description)
	if err != nil {
		log.Println(err)
		return deviceType, err
	}

	log.Println(deviceType)
	return deviceType, nil
}
