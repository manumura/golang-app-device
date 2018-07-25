package channeldao

import (
	"github.com/manumura/golang-app-device/model/channel"
	"github.com/manumura/golang-app-device/config"
)

// ChannelDao : interface defining DB operations on channel
type ChannelDao interface {
	FindChannels() ([]channelmodel.Channel, error)
	GetChannel(id int) (channelmodel.Channel, error)
}

// NewChannelDao : Create a new instance of ChannelDao implemenation
func NewChannelDao(db *config.DB) ChannelDao {
	return ChannelDaoImpl{db}
	//return ChannelDaoImpl2{}
}
