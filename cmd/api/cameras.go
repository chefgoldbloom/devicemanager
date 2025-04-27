package main

import (
	"fmt"
	"net/http"
	"time"

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
			app.postgresError(w, r, pgErr)
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

	camera := data.Camera{
		ID:         id,
		CreatedAt:  time.Now(),
		IpAddress:  "192.168.90.1",
		MacAddress: "ACCC12345678",
		Model:      "P1234-VE",
		Firmware:   "1.2.3",
		Site:       "Dummy-Site-OPS",
		Name:       "Dummy Camera",
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"camera": camera}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
