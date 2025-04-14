package data

import (
	"time"

	"github.com/chefgoldbloom/devicemanager/internal/validator"
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

func ValidateCamera(v *validator.Validator, camera *Camera) {
	v.Check(camera.MacAddress != "", "mac_address", "must be provided")
	v.Check(len(camera.MacAddress) != 12, "mac_address", "must be 12 characters")

	v.Check(camera.Firmware != "", "firmware", "must be provided")
	v.Check(camera.Model != "", "model", "must be provided")
	v.Check(camera.Name != "", "name", "must be provided")
	v.Check(camera.Site != "", "site", "must be provided")
}
