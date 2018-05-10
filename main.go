package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manumura/golang-app-device/controller/channel"
	"github.com/manumura/golang-app-device/controller/device"
	"github.com/manumura/golang-app-device/controller/device-type"
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/service/channel"
	"github.com/manumura/golang-app-device/service/device"
	"github.com/manumura/golang-app-device/service/device-type"
)

// Application starts here.
func main() {
	r := httprouter.New()
	r.GET("/", index)

	// Get a ChannelController instance
	channelDao := channeldao.NewChannelDao()
	channelService := channelservice.NewChannelService(channelDao)
	channelController := channel.NewChannelController(channelService)
	r.GET("/dm/api/v1/channels", channelController.FindChannels)
	r.GET("/dm/api/v1/channels/:id", channelController.GetChannel)

	// Get a DeviceTypeController instance
	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceTypeService := devicetypeservice.NewDeviceTypeService(deviceTypeDao)
	deviceTypeController := devicetypecontroller.NewDeviceTypeController(deviceTypeService)
	r.GET("/dm/api/v1/deviceTypes", deviceTypeController.FindDeviceTypes)
	r.GET("/dm/api/v1/deviceTypes/:id", deviceTypeController.GetDeviceType)

	// Get a DeviceController instance
	deviceDao := devicedao.NewDeviceDao()
	deviceService := deviceservice.NewDeviceService(deviceDao)
	deviceController := device.NewDeviceController(deviceService)
	r.GET("/dm/api/v1/devices", deviceController.FindDevices)
	r.GET("/dm/api/v1/devices/:id", deviceController.GetDevice)
	r.GET("/dm/api/v1/status/devices", deviceController.FindDeviceStatuses)

	http.ListenAndServe(":17172", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
