package main

import (
	"net/http"
)

// Application starts here.
import (
	"fmt"
	//"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manumura/golang-app-device/channel"
	"github.com/manumura/golang-app-device/device-type"
)

func main() {
	r := httprouter.New()

	// Get a ChannelController instance
	channelController := channel.NewChannelController()

	// Get a DeviceTypeController instance
	deviceTypeController := deviceType.NewDeviceTypeController()

	r.GET("/", index)
	r.GET("/dm/api/v1/channels", channelController.FindChannels)
	r.GET("/dm/api/v1/channels/:id", channelController.GetChannel)
	r.GET("/dm/api/v1/deviceTypes", deviceTypeController.FindDeviceTypes)
	r.GET("/dm/api/v1/deviceTypes/:id", deviceTypeController.GetDeviceType)
	http.ListenAndServe(":17172", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
