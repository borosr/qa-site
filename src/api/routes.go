package api

import (
	"log"
	"net/http"

	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Init() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/status", healthcheck.Route)

		r.Post("/users", users.Create)
		// TODO add other endpoints
	})

	config := settings.Get()
	log.Printf("Running the API on port: %s", config.Port)

	return http.ListenAndServe(":"+config.Port, r)
}
