package data

import (
	"database/sql"
	"errors"

	"github.com/chefgoldbloom/devicemanager/internal/validator"
	"time"
)

type Camera struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	IpAddress  string    `json:"ip_address"`
	MacAddress string    `json:"mac_address"`
	Model      string    `json:"model"`
	Firmware   string    `json:"firmware"`
	Site       string    `json:"site"`
	Name       string    `json:"name"`
	Version    int32     `json:"version"`
}

func ValidateCamera(v *validator.Validator, camera *Camera) {
	v.Check(camera.MacAddress != "", "mac_address", "must be provided")
	v.Check(len(camera.MacAddress) == 12, "mac_address", "must be 12 characters")

	v.Check(camera.Firmware != "", "firmware", "must be provided")
	v.Check(camera.Model != "", "model", "must be provided")
	v.Check(camera.Name != "", "name", "must be provided")
	v.Check(camera.Site != "", "site", "must be provided")
}

// Camera DB Model
type CameraModel struct {
	DB *sql.DB
}

// CRUD
func (c CameraModel) Insert(camera *Camera) error {
	query := `
		INSERT INTO cameras (ip_address, mac_address, model, firmware, site, name)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at, version`

	args := []any{camera.IpAddress, camera.MacAddress, camera.Model, camera.Firmware, camera.Site, camera.Name}

	return c.DB.QueryRow(query, args...).Scan(&camera.ID, &camera.CreatedAt, &camera.Version)
}

func (c CameraModel) Get(id int64) (*Camera, error) {
	query := `
		SELECT id, created_at, ip_address, mac_address, model, firmware, site, name, version
		FROM cameras
		WHERE id = $1`

	var camera Camera

	err := c.DB.QueryRow(query, id).Scan(
		&camera.ID,
		&camera.CreatedAt,
		&camera.IpAddress,
		&camera.MacAddress,
		&camera.Model,
		&camera.Firmware,
		&camera.Site,
		&camera.Name,
		&camera.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &camera, nil
}

func (c CameraModel) Update(camera *Camera) error {
	query := `
		UPDATE cameras
		SET ip_address = $1, mac_address = $2, model = $3, firmware = $4, site = $5, name = $6, version = version + 1
		where id = $7
		returning version`

	args := []any{camera.IpAddress, camera.MacAddress, camera.Model, camera.Firmware, camera.Site, camera.Version, camera.ID}

	return c.DB.QueryRow(query, args...).Scan(&camera.Version)
}

func (c CameraModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM cameras where id = $1;`
	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsEffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
