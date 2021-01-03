package api

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/auth"
	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/questions"
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
		r.Get("/login/{media:(github)}", auth.SocialMediaRedirect)
		r.Get("/login/{media:(github)}/callback", auth.SocialMediaCallback)

		loggedIn := r.With(auth.Middleware)
		loggedIn.Get("/users", users.GetAll)
		loggedIn.Get("/users/{id}", users.Get)
		loggedIn.Put("/users/{id}", users.Update)
		loggedIn.Delete("/users/{id}", users.Delete)

		loggedIn.Get("/questions", questions.GetAll)
		loggedIn.Get("/questions/{id}", questions.Get)
		loggedIn.Delete("/questions/{id}", questions.Delete)
		loggedIn.Post("/questions", questions.Create)
		loggedIn.Put("/questions/{id}", questions.Update)

		loggedIn.Delete("/logout", auth.Logout)
		loggedIn.Post("/revoke", auth.Revoke)
	})

	config := settings.Get()
	log.Infof("Running the API on port: %s", config.Port)

	return http.ListenAndServe(":"+config.Port, r)
}
