package data

import (
	"database/sql"
	"errors"
)

// Custom error for when we try to Get() a camera that isn't there
var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Cameras CameraModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Cameras: CameraModel{DB: db},
	}
}
