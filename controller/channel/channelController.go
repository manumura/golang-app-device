package channel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/manumura/golang-app-device/service/channel"
)

// ChannelController : Operations on channel
type ChannelController struct {
	channelService channelservice.ChannelService
}

// NewChannelController : Create a new instance of ChannelController
func NewChannelController(channelService channelservice.ChannelService) *ChannelController {
	return &ChannelController{channelService}
}

// FindChannels : Get all channels as json
func (cc ChannelController) FindChannels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Retrieve channels
	channels, err := cc.channelService.FindChannels()
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

// GetChannel : Get channel by id as json
func (cc ChannelController) GetChannel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

	// Retrieve channel
	channel, err := cc.channelService.GetChannel(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	cj, err := json.Marshal(channel)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", cj)
}
