package dao

import "github.com/manumura/golang-app-device/channel/model"

// ChannelDao : interface defining DB operations on channel
type ChannelDao interface {
	FindChannels() ([]model.Channel, error)
	GetChannel(id int) (model.Channel, error)
}

// NewChannelDao : Create a new instance of ChannelDao implemenation
func NewChannelDao() ChannelDao {
	return ChannelDaoImpl{}
	//return ChannelDaoImpl2{}
}
