package device

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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

	devices, err := dc.deviceService.FindDevices()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, deviceStatuses)
}
