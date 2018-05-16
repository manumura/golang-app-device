package devicedao

import (
	"errors"
	"log"
	"time"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/model/device"
)

// DeviceDaoImpl : implementation for DB operations on device
type DeviceDaoImpl struct {
}

// FindDevices : retrieve devices from the database
func (dd DeviceDaoImpl) FindDevices() ([]devicemodel.Device, error) {

	sql := "SELECT dr.device_request_id, di.imei, dr.request_by, ds.device_status_id, ds.name, dr.created_date_time "
	sql += "FROM device_request dr "
	sql += "LEFT JOIN device_status ds ON ds.device_status_id = dr. device_status_id "
	sql += "LEFT JOIN device_info di ON di.device_info_id = dr.device_info_id "
	sql += "LEFT JOIN device_type dt ON dt.device_type_id = di.device_type_id "

	rows, err := config.Database.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []devicemodel.Device{}
	for rows.Next() {
		device := devicemodel.Device{}
		err := rows.Scan(&device.ID, &device.Imei, &device.RequestBy, &device.Status, &device.StatusText, &device.Timestamp) // order matters
		if err != nil {
			log.Println(err)
			return nil, err
		}
		devices = append(devices, device)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(devices)
	return devices, nil
}

// GetDevice : retrieve device by id from the database
func (dd DeviceDaoImpl) GetDevice(id int) (devicemodel.Device, error) {

	device := devicemodel.Device{}

	if id == 0 {
		return device, errors.New("id cannot be empty")
	}

	sql := "SELECT dr.device_request_id, di.imei, dr.request_by, ds.device_status_id, ds.name, dr.created_date_time, "
	sql += "di.android_device_id, di.android_version, di.api_level, di.brand, di.build_number, "
	sql += "di.cpu_hardware, di.display_density, di.display_physical_size, di.display_resolution, "
	sql += "di.hardware_serial_no, di.instruction_sets, di.manufacturer, di.model "
	sql += "FROM device_request dr "
	sql += "LEFT JOIN device_status ds ON ds.device_status_id = dr. device_status_id "
	sql += "LEFT JOIN device_info di ON di.device_info_id = dr.device_info_id "
	sql += "LEFT JOIN device_type dt ON dt.device_type_id = di.device_type_id "
	sql += "WHERE dr.device_request_id = $1 "

	row := config.Database.QueryRow(sql, id)

	err := row.Scan(&device.ID, &device.Imei, &device.RequestBy, &device.Status, &device.StatusText, &device.Timestamp, &device.DeviceInfo.AndroidDeviceID, &device.DeviceInfo.AndroidVersion,
		&device.DeviceInfo.APILevel, &device.DeviceInfo.Brand, &device.DeviceInfo.BuildNumber, &device.DeviceInfo.CPUHardware, &device.DeviceInfo.DisplayDensity, &device.DeviceInfo.DisplayPhysicalSize,
		&device.DeviceInfo.DisplayResolution, &device.DeviceInfo.HardwareSerialNo, &device.DeviceInfo.InstructionSets, &device.DeviceInfo.Manufacturer, &device.DeviceInfo.Model)
	if err != nil {
		log.Println(err)
		return device, err
	}

	log.Println(device)
	return device, nil
}

// Delete : retrieve one device from the database
func (dd DeviceDaoImpl) Delete(id int) error {

	if id == 0 {
		return errors.New("id cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	// execute delete on device table.
	stmt, err := tx.Prepare("DELETE FROM device_request WHERE device_request_id = $1;")
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		log.Println(err)
		tx.Rollback()
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	tx.Commit()
	if err != nil {
		log.Println(err)
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	return nil
}

// Update : update one device in the database
func (dd DeviceDaoImpl) Update(d devicemodel.Device) (devicemodel.Device, error) {

	result := devicemodel.Device{}

	if d.ID == 0 || d.RequestBy == "" || d.Status == "" {
		return result, errors.New("parameters cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return result, err
	}

	// get status id from name
	var statusID int
	row := config.Database.QueryRow("SELECT ds.device_status_id FROM device_status ds WHERE ds.name = $1", d.Status)

	err = row.Scan(&statusID)
	if err != nil {
		log.Println(err)
		return result, err
	}

	log.Println(statusID)

	// execute update on device table
	stmt, err := tx.Prepare("UPDATE device_request SET request_by = $1, device_status_id = $2 WHERE device_request_id = $3")

	if err != nil {
		log.Println(err)
		return result, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(d.RequestBy, statusID, d.ID); err != nil {
		log.Println(err)
		tx.Rollback()
		return result, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return result, err
	}

	result, err = dd.GetDevice(d.ID)
	return result, err

}

// Create : create one device in the database
func (dd DeviceDaoImpl) Create(d devicemodel.Device) (devicemodel.Device, error) {

	result := devicemodel.Device{}

	if d.ID == 0 || d.RequestBy == "" || d.Status == "" {
		return result, errors.New("parameters cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return result, err
	}

	// get status id from name
	var statusID int
	row := config.Database.QueryRow("SELECT ds.device_status_id FROM device_status ds WHERE ds.name = $1", d.Status)

	err = row.Scan(&statusID)
	if err != nil {
		log.Println(err)
		return result, err
	}

	log.Println(statusID)

	// execute insert on device table
	stmt, err := tx.Prepare("INSERT INTO device_request (created_date_time, device_status_id, is_active, request_by) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println(err)
		return result, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(time.Now, statusID, true, d.RequestBy); err != nil {
		log.Println(err)
		tx.Rollback()
		return result, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return result, errors.New("device cannot be saved")
	}

	result, err = dd.GetDevice(d.ID)
	return result, err
}

// FindDeviceStatuses : retrieve device statuses from the database
func (dd DeviceDaoImpl) FindDeviceStatuses() ([]devicemodel.DeviceStatus, error) {

	rows, err := config.Database.Query("SELECT ds.device_status_id, ds.name FROM device_status ds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deviceStatuses := []devicemodel.DeviceStatus{}
	for rows.Next() {
		deviceStatus := devicemodel.DeviceStatus{}
		err := rows.Scan(&deviceStatus.ID, &deviceStatus.Name) // order matters
		if err != nil {
			log.Println(err)
			return nil, err
		}
		deviceStatuses = append(deviceStatuses, deviceStatus)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(deviceStatuses)
	return deviceStatuses, nil
}
