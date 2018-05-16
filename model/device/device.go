package devicemodel

import "time"

// Device type
type Device struct {
	ID         int        `json:"id"`
	Imei       string     `json:"imei"`
	StatusText string     `json:"statusText"`
	Status     string     `json:"status"`
	RequestBy  string     `json:"requestBy"`
	Timestamp  time.Time  `json:"timestamp"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}
