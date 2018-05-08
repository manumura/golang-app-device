package channel

// ChannelDao : interface defining DB operations on channel
type ChannelDao interface {
	FindChannels() ([]Channel, error)
}

// NewChannelDao : Create a new instance of ChannelDao implemenation
func NewChannelDao() ChannelDao {
	return ChannelDaoImpl{}
	//return ChannelDaoImpl2{}
}
