package devicemodel

import (
	// TODO https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267
	// https://stackoverflow.com/questions/24564619/nullable-time-time-in-golang
	// Timestamp  pq.NullTime `json:"timestamp"`
	// "database/sql"
	// "github.com/lib/pq"
	"time"
)

// Device type
type Device struct {
	ID         int        `json:"id"`
	Imei       string     `json:"imei"`
	StatusText string     `json:"statusText"`
	Status     int        `json:"status"`
	RequestBy  string     `json:"requestBy"`
	Timestamp  time.Time  `json:"timestamp"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}
