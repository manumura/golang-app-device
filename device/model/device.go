package model

import "time"

// Device type
type Device struct {
	ID         int        `json:"id"`
	Imei       string     `json:"imei"`
	Status     string     `json:"status"`
	RequestBy  string     `json:"requestBy"`
	Timestamp  time.Time  `json:"timestamp"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}
