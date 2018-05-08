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
)

func main() {
	r := httprouter.New()

	// Get a ChannelController instance
	channelController := channel.NewChannelController()

	r.GET("/", index)
	r.GET("/dm/api/v1/channels", channelController.GetChannels)
	http.ListenAndServe(":17172", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
