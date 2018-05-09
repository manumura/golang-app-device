package service

import (
	"github.com/manumura/golang-app-device/channel/dao"
	"github.com/manumura/golang-app-device/channel/model"
)

// ChannelServiceImpl : implementation for services on channel
type ChannelServiceImpl struct {
}

// FindChannels : retrieve channels
func (cs ChannelServiceImpl) FindChannels() ([]model.Channel, error) {

	channelDao := dao.NewChannelDao()
	channels, err := channelDao.FindChannels()
	return channels, err
}

// GetChannel : retrieve channel by id
func (cs ChannelServiceImpl) GetChannel(id int) (model.Channel, error) {

	channelDao := dao.NewChannelDao()
	channel, err := channelDao.GetChannel(id)
	return channel, err
}
