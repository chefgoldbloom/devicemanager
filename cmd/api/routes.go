package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/healthcheck", app.healthcheckHandler)
	r.Post("/v1/cameras", app.createCameraHandler)
	r.Get("/v1/cameras/{id}", app.showCameraHandler)
	r.Put("/v1/cameras/{id}", app.updateCameraHandler)
	r.Delete("/v1/cameras/{id}", app.deleteCameraHandler)

	return r
}
