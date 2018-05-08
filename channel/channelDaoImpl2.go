package channel

import (
	"fmt"

	"github.com/manumura/golang-app-device/config"
)

// ChannelDaoImpl2 : test / implementation for DB operations on channel
type ChannelDaoImpl2 struct {
}

// FindChannels : test / retrieve channels from the database
func (cd ChannelDaoImpl2) FindChannels() ([]Channel, error) {

	fmt.Println("ChannelDaoImpl2")

	rows, err := config.Database.Query("SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	channels := []Channel{}
	for rows.Next() {
		channel := Channel{}
		err := rows.Scan(&channel.ID, &channel.Name, &channel.Description) // order matters
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(channels)
	return channels, nil
}