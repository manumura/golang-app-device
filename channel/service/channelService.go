package service

import "github.com/manumura/golang-app-device/channel/model"

// ChannelService : interface defining services on channel
type ChannelService interface {
	FindChannels() ([]model.Channel, error)
	GetChannel(id int) (model.Channel, error)
}

// NewChannelService : Create a new instance of ChannelService implemenation
func NewChannelService() ChannelService {
	return ChannelServiceImpl{}
}
