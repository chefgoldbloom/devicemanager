package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/v1/healthcheck", app.healthcheckHandler)
	r.Post("/v1/cameras", app.createCameraHandler)
	r.Get("/v1/cameras/{id}", app.showCameraHandler)

	return r
}
