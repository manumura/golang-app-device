package channel

// ChannelDao : interface defining DB operations on channel
type ChannelDao interface {
	FindChannels() ([]Channel, error)
}
