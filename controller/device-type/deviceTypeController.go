package devicetypecontroller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/manumura/golang-app-device/service/device-type"
)

// DeviceTypeController : Operations on device type
type DeviceTypeController struct {
	deviceTypeService devicetypeservice.DeviceTypeService
}

// NewDeviceTypeController : Create a new instance of DeviceTypeController
func NewDeviceTypeController(deviceTypeService devicetypeservice.DeviceTypeService) *DeviceTypeController {
	return &DeviceTypeController{deviceTypeService}
}

// FindDeviceTypes : Get all device types as json
// func (dtc DeviceTypeController) FindDeviceTypes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func (dtc DeviceTypeController) FindDeviceTypes(c echo.Context) error {

	deviceTypes, err := dtc.deviceTypeService.FindDeviceTypes()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, deviceTypes)
}

// GetDeviceType : Get device type by id as json
// func (dtc DeviceTypeController) GetDeviceType(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func (dtc DeviceTypeController) GetDeviceType(c echo.Context) error {

	// idAsString := p.ByName("id")
	idAsString := c.Param("id")
	if idAsString == "" {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		// return
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		// return
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	deviceType, err := dtc.deviceTypeService.GetDeviceType(id)
	switch {
	case err == sql.ErrNoRows:
		// http.NotFound(w, r)
		// return
		return echo.NewHTTPError(http.StatusNotFound, "Page not found")
	case err != nil:
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		// return
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// dtj, err := json.Marshal(deviceType)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK) // 200
	// fmt.Fprintf(w, "%s\n", dtj)
	return c.JSON(http.StatusOK, deviceType)
}
