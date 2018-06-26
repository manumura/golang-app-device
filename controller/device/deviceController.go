package device

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/manumura/golang-app-device/model/device"
	"github.com/manumura/golang-app-device/service/device"
)

// DeviceController : Operations on device
type DeviceController struct {
	deviceService deviceservice.DeviceService
}

// NewDeviceController : Create a new instance of DeviceTypeController
func NewDeviceController(deviceService deviceservice.DeviceService) *DeviceController {
	return &DeviceController{deviceService}
}

// FindDevices : Get all devices as json
func (dc DeviceController) FindDevices(c echo.Context) error {

	log.Println(c.Cookie("SESSIONID"))
	log.Println(c.Cookies())

	devices, err := dc.deviceService.FindDevices()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, devices)
}

// GetDevice : Get device by id as json
func (dc DeviceController) GetDevice(c echo.Context) error {

	idAsString := c.Param("id")
	if idAsString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	device, err := dc.deviceService.GetDevice(id)
	// TODO : throw more appropriate error
	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusNotFound, "Page not found")
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, device)
}

// FindDeviceStatuses : Get all device statuses as json
func (dc DeviceController) FindDeviceStatuses(c echo.Context) error {

	deviceStatuses, err := dc.deviceService.FindDeviceStatuses()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, deviceStatuses)
}

// UpdateDevice : Update device
func (dc DeviceController) UpdateDevice(c echo.Context) error {

	log.Println("UpdateDevice")

	d := devicemodel.Device{}
	if err := c.Bind(&d); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	log.Println(d)

	if d.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	// Update device
	d, err := dc.deviceService.Update(d)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, d)
}

// DeleteDevice : Delete device by id
func (dc DeviceController) DeleteDevice(c echo.Context) error {

	log.Println("DeleteDevice")

	idAsString := c.Param("id")
	if idAsString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	log.Println(id)

	err = dc.deviceService.Delete(id)
	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusNotFound, "Page not found")
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.NoContent(http.StatusNoContent)
}
