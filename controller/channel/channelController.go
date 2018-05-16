package channel

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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
func (cc ChannelController) FindChannels(c echo.Context) error {

	// Retrieve channels
	channels, err := cc.channelService.FindChannels()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, channels)
}

// GetChannel : Get channel by id as json
func (cc ChannelController) GetChannel(c echo.Context) error {

	idAsString := c.Param("id")
	if idAsString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	// Retrieve channel
	channel, err := cc.channelService.GetChannel(id)
	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusNotFound, "Page not found")
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, channel)
}
