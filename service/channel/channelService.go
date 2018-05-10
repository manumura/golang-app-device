package channelservice

import "github.com/manumura/golang-app-device/model/channel"

// ChannelService : interface defining services on channel
type ChannelService interface {
	FindChannels() ([]channelmodel.Channel, error)
	GetChannel(id int) (channelmodel.Channel, error)
}

// NewChannelService : Create a new instance of ChannelService implemenation
func NewChannelService() ChannelService {
	return ChannelServiceImpl{}
}
