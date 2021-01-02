package api

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/settings"
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
		// TODO add other endpoints
	})

	return http.ListenAndServe(":"+settings.Get().Port, r)
}
