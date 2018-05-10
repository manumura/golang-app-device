package device

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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
func (dc DeviceController) FindDevices(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	devices, err := dc.deviceService.FindDevices()
	if err != nil {
		fmt.Println(err)
	}

	dj, err := json.Marshal(devices)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", dj)
}

// GetDevice : Get device by id as json
func (dc DeviceController) GetDevice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

	device, err := dc.deviceService.GetDevice(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	dj, err := json.Marshal(device)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", dj)
}

// FindDeviceStatuses : Get all device statuses as json
func (dc DeviceController) FindDeviceStatuses(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	deviceStatuses, err := dc.deviceService.FindDeviceStatuses()
	if err != nil {
		fmt.Println(err)
	}

	dsj, err := json.Marshal(deviceStatuses)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", dsj)
}
