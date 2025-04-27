package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chefgoldbloom/devicemanager/internal/data"
	"github.com/chefgoldbloom/devicemanager/internal/validator"
	"github.com/lib/pq"
)

// createCameraHandler handles the creation of a new camera by decoding the request body
// and returning the parsed input as a response.
func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		IpAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
		Model      string `json:"model"`
		Firmware   string `json:"firmware"`
		Site       string `json:"site"`
		Name       string `json:"name"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	camera := &data.Camera{
		IpAddress:  input.IpAddress,
		MacAddress: input.MacAddress,
		Model:      input.Model,
		Firmware:   input.Firmware,
		Site:       input.Site,
		Name:       input.Name,
	}

	v := validator.New()

	if data.ValidateCamera(v, camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Cameras.Insert(camera)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			app.postgresErrorResponse(w, r, pgErr)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Send location of newly created resource
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/cameras/%d", camera.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"camera": camera}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	camera, err := app.models.Cameras.Get(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	camera, err := app.models.Cameras.Get(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	var input struct {
		IpAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
		Model      string `json:"model"`
		Firmware   string `json:"firmware"`
		Site       string `json:"site"`
		Name       string `json:"name"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	camera.IpAddress = input.IpAddress
	camera.MacAddress = input.MacAddress
	camera.Model = input.Model
	camera.Firmware = input.Firmware
	camera.Site = input.Site
	camera.Name = input.Name

	v := validator.New()

	if data.ValidateCamera(v, camera); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Cameras.Update(camera)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			app.postgresErrorResponse(w, r, pgErr)
		} else {
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	err = app.models.Cameras.Delete(int64(id))
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			app.postgresErrorResponse(w, r, pgErr)
		} else {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.notFoundResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "camera successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
