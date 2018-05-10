package channelservice

import (
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/model/channel"
)

// ChannelServiceImpl : implementation for services on channel
type ChannelServiceImpl struct {
	channelDao channeldao.ChannelDao
}

// NewChannelService : Create a new instance of ChannelService implemenation
func NewChannelService(channelDao channeldao.ChannelDao) ChannelService {
	return ChannelServiceImpl{channelDao}
}

// FindChannels : retrieve channels
func (cs ChannelServiceImpl) FindChannels() ([]channelmodel.Channel, error) {

	channels, err := cs.channelDao.FindChannels()
	return channels, err
}

// GetChannel : retrieve channel by id
func (cs ChannelServiceImpl) GetChannel(id int) (channelmodel.Channel, error) {

	channel, err := cs.channelDao.GetChannel(id)
	return channel, err
}
