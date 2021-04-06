package api

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/answers"
	"github.com/borosr/qa-site/pkg/auth"
	authRepository "github.com/borosr/qa-site/pkg/auth/repository"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/questions"
	"github.com/borosr/qa-site/pkg/ratings"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	userRepository "github.com/borosr/qa-site/pkg/users/repository"
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

	ur := userRepository.NewRepository(db.Get())
	ar := authRepository.NewRepository(db.GetBDB(), db.Get())
	uc := users.NewController(ur)
	ac := auth.NewController(ur, ar)

	r.Route("/api", func(r chi.Router) {
		loggedIn := r.With(ac.Middleware)

		r.Get("/status", healthcheck.Route)

		initAuth(r, ac, loggedIn)
		initUsers(r, uc, loggedIn)

		loggedIn.Get("/questions", questions.GetAll)
		loggedIn.Get("/questions/{id}", questions.Get)
		loggedIn.Delete("/questions/{id}", questions.Delete)
		loggedIn.Post("/questions", questions.Create)
		loggedIn.Put("/questions/{id}", questions.Update)

		loggedIn.Get("/questions/{questionID}/answers", answers.GetQuestionsAnswers)
		loggedIn.Put("/questions/{questionID}/answers/{answerID}/answered", answers.SetAnswered)

		loggedIn.Get("/answers", answers.GetMyAnswers)
		loggedIn.Post("/answers", answers.Create)
		loggedIn.Put("/answers/{id}", answers.Update)

		loggedIn.Put("/{kind:(answers|questions)}/{id}/rate", ratings.Rate)
		loggedIn.Put("/{kind:(answers|questions)}/{id}/unrate", ratings.Unrate)
		loggedIn.Put("/{kind:(answers|questions)}/{id}/rate/dismiss", ratings.Dismiss)
	})

	config := settings.Get()
	log.Infof("Running the API on port: %s", config.Port)

	return http.ListenAndServe(":"+config.Port, r)
}

func initAuth(r chi.Router, ac auth.Controller, loggedIn chi.Router) {
	r.Post("/login", ac.DefaultLogin)
	r.Get("/login/{media:(github)}", ac.SocialMediaRedirect)
	r.Get("/login/{media:(github)}/callback", ac.SocialMediaCallback)
	loggedIn.Delete("/logout", ac.Logout)
	r.Post("/revoke", ac.Revoke)
}

func initUsers(r chi.Router, uc users.UserController, loggedIn chi.Router) {
	r.Post("/users", uc.Create)
	loggedIn.Get("/users", uc.GetAll)
	loggedIn.Get("/users/{id}", uc.Get)
	loggedIn.Put("/users/{id}", uc.Update)
	loggedIn.Delete("/users/{id}", uc.Delete)
}