package routes

import (
	"github.com/alimohseni/fem_project/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
		r := chi.NewRouter()

		r.Get("/health", app.HealthCheck)

		return r
}