package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chefgoldbloom/devicemanager/internal/data"
)

func (app *application) createCameraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new camera")
}

func (app *application) showCameraHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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

	err = app.writeJSON(w, http.StatusOK, camera, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
