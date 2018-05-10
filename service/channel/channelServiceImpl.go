package channelservice

import (
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/model/channel"
)

// ChannelServiceImpl : implementation for services on channel
type ChannelServiceImpl struct {
}

// FindChannels : retrieve channels
func (cs ChannelServiceImpl) FindChannels() ([]channelmodel.Channel, error) {

	channelDao := channeldao.NewChannelDao()
	channels, err := channelDao.FindChannels()
	return channels, err
}

// GetChannel : retrieve channel by id
func (cs ChannelServiceImpl) GetChannel(id int) (channelmodel.Channel, error) {

	channelDao := channeldao.NewChannelDao()
	channel, err := channelDao.GetChannel(id)
	return channel, err
}
