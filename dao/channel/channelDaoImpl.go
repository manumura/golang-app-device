package channeldao

import (
	"errors"
	"log"

	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/model/channel"
)

// ChannelDaoImpl : implementation for DB operations on channel
type ChannelDaoImpl struct {
	db *config.DB
}

// FindChannels : retrieve channels from the database
func (cd ChannelDaoImpl) FindChannels() ([]channelmodel.Channel, error) {

	log.Println("ChannelDaoImpl")

	sql := "SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c"
	//rows, err := config.Database.Query(sql)

	stmt, err := cd.db.Prepare(sql)
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
			log.Println(err)
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

// GetChannel : retrieve channel by id from the database
func (cd ChannelDaoImpl) GetChannel(id int) (channelmodel.Channel, error) {

	log.Println("ChannelDaoImpl")
	channel := channelmodel.Channel{}

	if id == 0 {
		return channel, errors.New("id cannot be empty")
	}

	sql := "SELECT c.dist_channel_id, c.name, c.description FROM app_dist_channel c WHERE c.dist_channel_id = $1"
	//row := config.Database.QueryRow(sql, id)

	//err := row.Scan(&channel.ID, &channel.Name, &channel.Description)

	stmt, err := cd.db.Prepare(sql)
	if err != nil {
		log.Println(err)
		return channel, err
	}

	err = stmt.QueryRow(id).Scan(&channel.ID, &channel.Name, &channel.Description)
	if err != nil {
		log.Println(err)
		return channel, err
	}

	log.Println(channel)
	return channel, nil
}
