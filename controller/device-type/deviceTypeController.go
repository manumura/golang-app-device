package devicetypecontroller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	devicetypeservice "github.com/manumura/golang-app-device/service/device-type"
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
func (dtc DeviceTypeController) FindDeviceTypes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	deviceTypes, err := dtc.deviceTypeService.FindDeviceTypes()
	if err != nil {
		fmt.Println(err)
	}

	dtj, err := json.Marshal(deviceTypes)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", dtj)
}

// GetDeviceType : Get device type by id as json
func (dtc DeviceTypeController) GetDeviceType(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	idAsString := p.ByName("id")
	if idAsString == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	deviceType, err := dtc.deviceTypeService.GetDeviceType(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	dtj, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", dtj)
}
