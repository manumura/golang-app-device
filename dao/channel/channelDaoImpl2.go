package channeldao

import (
	"fmt"

	"github.com/manumura/golang-app-device/model/channel"
	"github.com/manumura/golang-app-device/config"
	"log"
)

// ChannelDaoImpl2 : test / implementation for DB operations on channel
type ChannelDaoImpl2 struct {
}

// FindChannels : test / retrieve channels from the database
func (cd ChannelDaoImpl2) FindChannels() ([]channelmodel.Channel, error) {

	fmt.Println("ChannelDaoImpl2")

	sql := "SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c"
	//rows, err := config.Database.Query(sql)

	stmt, err := config.Database.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var channels []channelmodel.Channel
	for rows.Next() {
		channel := channelmodel.Channel{}
		err := rows.Scan(&channel.ID, &channel.Name, &channel.Description) // order matters
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(channels)
	return channels, nil
}
