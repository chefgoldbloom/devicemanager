package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chefgoldbloom/devicemanager/internal/data"
	"github.com/chefgoldbloom/devicemanager/internal/validator"
)

// createCameraHandler handles the creation of a new camera by decoding the request body
// and returning the parsed input as a response.
func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
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

	fmt.Fprintf(w, "%+v\n", input)
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
