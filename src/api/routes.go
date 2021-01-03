package api

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/auth"
	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	"github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func Init() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.Logger("router", log.New()))
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/status", healthcheck.Route)

		r.Post("/users", users.Create)
		r.Post("/login", auth.DefaultLogin)

		loggedIn := r.With(auth.Middleware)
		loggedIn.Get("/users", users.GetAll)
		loggedIn.Get("/users/{id}", users.Get)
		loggedIn.Delete("/users/{id}", users.Delete)
		loggedIn.Put("/users/{id}", users.Update)
	})

	config := settings.Get()
	log.Infof("Running the API on port: %s", config.Port)

	return http.ListenAndServe(":"+config.Port, r)
}
