package channeldao

import "github.com/manumura/golang-app-device/model/channel"

// ChannelDao : interface defining DB operations on channel
type ChannelDao interface {
	FindChannels() ([]channelmodel.Channel, error)
	GetChannel(id int) (channelmodel.Channel, error)
}

// NewChannelDao : Create a new instance of ChannelDao implemenation
func NewChannelDao() ChannelDao {
	return ChannelDaoImpl{}
	//return ChannelDaoImpl2{}
}
