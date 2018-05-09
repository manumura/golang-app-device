package dao

import (
	"errors"
	"fmt"
	"time"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/device/model"
)

// DeviceDaoImpl : implementation for DB operations on device
type DeviceDaoImpl struct {
}

// FindDevices : retrieve devices from the database
func (dd DeviceDaoImpl) FindDevices() ([]model.Device, error) {

	sql := "SELECT dr.device_request_id, di.imei, dr.request_by, ds.name, dr.created_date_time "
	sql += "FROM device_request dr "
	sql += "LEFT JOIN device_status ds ON ds.device_status_id = dr. device_status_id "
	sql += "LEFT JOIN device_info di ON di.device_info_id = dr.device_info_id "
	sql += "LEFT JOIN device_type dt ON dt.device_type_id = di.device_type_id "

	rows, err := config.Database.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []model.Device{}
	for rows.Next() {
		device := model.Device{}
		err := rows.Scan(&device.ID, &device.Imei, &device.RequestBy, &device.Status, &device.Timestamp) // order matters
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(devices)
	return devices, nil
}

// GetDevice : retrieve device by id from the database
func (dd DeviceDaoImpl) GetDevice(id int) (model.Device, error) {

	device := model.Device{}

	if id == 0 {
		return device, errors.New("id cannot be empty")
	}

	sql := "SELECT dr.device_request_id, di.imei, dr.request_by, ds.name, dr.created_date_time "
	sql += "FROM device_request dr "
	sql += "LEFT JOIN device_status ds ON ds.device_status_id = dr. device_status_id "
	sql += "LEFT JOIN device_info di ON di.device_info_id = dr.device_info_id "
	sql += "LEFT JOIN device_type dt ON dt.device_type_id = di.device_type_id "
	sql += "WHERE dr.device_request_id = $1 "

	row := config.Database.QueryRow(sql, id)

	err := row.Scan(&device.ID, &device.Imei, &device.RequestBy, &device.Status, &device.Timestamp)
	if err != nil {
		return device, err
	}

	fmt.Println(device)
	return device, nil
}

// Delete : retrieve one device from the database
func (dd DeviceDaoImpl) Delete(id int) error {

	if id == 0 {
		return errors.New("id cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return err
	}

	// execute delete on device table.
	stmt, err := tx.Prepare("DELETE FROM device_request WHERE device_request_id = $1;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		tx.Rollback()
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	tx.Commit()
	if err != nil {
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	return nil
}

// Update : update one device in the database
func (dd DeviceDaoImpl) Update(device model.Device) (model.Device, error) {

	result := model.Device{}

	if device.ID == 0 || device.RequestBy == "" || device.Status == "" {
		return result, errors.New("parameters cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return result, err
	}

	// get status id from name
	var statusID int
	row := config.Database.QueryRow("SELECT ds.device_status_id FROM device_status ds WHERE ds.name = $1", device.Status)

	err = row.Scan(&statusID)
	if err != nil {
		return result, err
	}

	fmt.Println(statusID)

	// execute update on device table
	stmt, err := tx.Prepare("UPDATE device_request SET request_by = $1, device_status_id = $2 WHERE device_request_id = $3")

	if err != nil {
		return result, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(device.RequestBy, statusID, device.ID); err != nil {
		tx.Rollback()
		return result, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return result, err
	}

	result, err = dd.GetDevice(device.ID)
	return result, err

}

// Create : create one device in the database
func (dd DeviceDaoImpl) Create(device model.Device) (model.Device, error) {

	result := model.Device{}

	if device.ID == 0 || device.RequestBy == "" || device.Status == "" {
		return result, errors.New("parameters cannot be empty")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return result, err
	}

	// get status id from name
	var statusID int
	row := config.Database.QueryRow("SELECT ds.device_status_id FROM device_status ds WHERE ds.name = $1", device.Status)

	err = row.Scan(&statusID)
	if err != nil {
		return result, err
	}

	fmt.Println(statusID)

	// execute insert on device table
	stmt, err := tx.Prepare("INSERT INTO device_request (created_date_time, device_status_id, is_active, request_by) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return result, err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(time.Now, statusID, true, device.RequestBy); err != nil {
		tx.Rollback()
		return result, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return result, errors.New("device cannot be saved")
	}

	result, err = dd.GetDevice(device.ID)
	return result, err
}

// FindDeviceStatuses : retrieve device statuses from the database
func (dd DeviceDaoImpl) FindDeviceStatuses(id int) ([]model.DeviceStatus, error) {

	rows, err := config.Database.Query("SELECT ds.device_status_id, ds.name FROM device_status ds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deviceStatuses := []model.DeviceStatus{}
	for rows.Next() {
		deviceStatus := model.DeviceStatus{}
		err := rows.Scan(&deviceStatus.ID, &deviceStatus.Name) // order matters
		if err != nil {
			return nil, err
		}
		deviceStatuses = append(deviceStatuses, deviceStatus)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(deviceStatuses)
	return deviceStatuses, nil
}
