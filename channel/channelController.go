package channel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ChannelController : Operations on channel
type ChannelController struct {
}

// NewChannelController : Create a new instance of ChannelController
func NewChannelController() *ChannelController {
	return &ChannelController{}
}

// GetChannels : Get all channels as json
func (cc ChannelController) GetChannels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Retrieve channels
	channelDao := NewChannelDao()
	channels, err := channelDao.FindChannels()
	if err != nil {
		fmt.Println(err)
	}

	cj, err := json.Marshal(channels)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", cj)
}
