package data

import (
	"time"
)

type Camera struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	MacAddress string    `json:"mac_address"`
	Model      string    `json:"model"`
	Firmware   string    `json:"firmware"`
	Site       string    `json:"site"`
	Name       string    `json:"name"`
	Version    int32     `json:"version"`
}
