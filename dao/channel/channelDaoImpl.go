package channeldao

import (
	"errors"
	"log"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/model/channel"
)

// ChannelDaoImpl : implementation for DB operations on channel
type ChannelDaoImpl struct {
}

// FindChannels : retrieve channels from the database
func (cd ChannelDaoImpl) FindChannels() ([]channelmodel.Channel, error) {

	log.Println("ChannelDaoImpl")

	rows, err := config.Database.Query("SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	channels := []channelmodel.Channel{}
	for rows.Next() {
		channel := channelmodel.Channel{}
		err := rows.Scan(&channel.ID, &channel.Name, &channel.Description) // order matters
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Println(channels)
	return channels, nil
}

// GetChannel : retrieve channel by id from the database
func (cd ChannelDaoImpl) GetChannel(id int) (channelmodel.Channel, error) {

	log.Println("ChannelDaoImpl")
	channel := channelmodel.Channel{}

	if id == 0 {
		return channel, errors.New("id cannot be empty")
	}

	row := config.Database.QueryRow("SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c WHERE c.dist_channel_id = $1", id)

	err := row.Scan(&channel.ID, &channel.Name, &channel.Description)
	if err != nil {
		return channel, err
	}

	log.Println(channel)
	return channel, nil
}
