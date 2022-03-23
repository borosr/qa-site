package api

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/answers"
	answerRepository "github.com/borosr/qa-site/pkg/answers/repository"
	"github.com/borosr/qa-site/pkg/auth"
	authRepository "github.com/borosr/qa-site/pkg/auth/repository"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/questions"
	questionRepository "github.com/borosr/qa-site/pkg/questions/repository"
	"github.com/borosr/qa-site/pkg/ratings"
	rateRepository "github.com/borosr/qa-site/pkg/ratings/repository"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	userRepository "github.com/borosr/qa-site/pkg/users/repository"
	logger "github.com/chi-middleware/logrus-logger"
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
	r.Use(healthcheck.Middleware)

	ur := userRepository.NewRepository(db.Get())
	aur := authRepository.NewRepository(db.GetBDB(), db.Get())
	qr := questionRepository.NewRepository(db.Get())
	anr := answerRepository.NewRepository(db.Get())
	rr := rateRepository.NewRepository(db.Get())

	uc := users.NewController(ur)
	auc := auth.NewController(ur, aur)
	qc := questions.NewController(qr)
	anc := answers.NewController(anr, qr)
	rc := ratings.NewController(rr)
	hcc := healthcheck.NewController()

	r.Route("/api", func(r chi.Router) {
		loggedIn := r.With(auc.Middleware)

		r.Get("/status", hcc.Route)

		r.Get("/info", hcc.Info)

		initAuth(r, loggedIn, auc)
		initUsers(r, loggedIn, uc)
		initQuestions(r, loggedIn, qc)
		initAnswers(r, loggedIn, anc)
		initRatings(loggedIn, rc)

	})

	config := settings.Get()
	log.Infof("Running the API on port: %s", config.Port)

	if hchk := healthcheck.Instance(); hchk.Healthy() {
		hchk.Ok()
	}

	return http.ListenAndServe(":"+config.Port, r)
}

func initRatings(loggedIn chi.Router, rc ratings.RateController) {
	loggedIn.Put("/{kind:(answers|questions)}/{id}/rate", rc.Rate)
	loggedIn.Put("/{kind:(answers|questions)}/{id}/unrate", rc.Unrate)
	loggedIn.Put("/{kind:(answers|questions)}/{id}/rate/dismiss", rc.Dismiss)
}

func initAnswers(r, loggedIn chi.Router, anc answers.AnswerController) {
	readRouter := loggedIn
	if settings.Get().Visibility == settings.VisibilityPublic {
		readRouter = r
	}
	readRouter.Get("/questions/{questionID}/answers", anc.GetQuestionsAnswers)
	loggedIn.Put("/questions/{questionID}/answers/{answerID}/answered", anc.SetAnswered)
	loggedIn.Get("/answers", anc.GetMyAnswers)
	loggedIn.Post("/answers", anc.Create)
	loggedIn.Put("/answers/{id}", anc.Update)
}

func initQuestions(r, loggedIn chi.Router, qc questions.QuestionController) {
	readRouter := loggedIn
	if settings.Get().Visibility == settings.VisibilityPublic {
		readRouter = r
	}
	readRouter.Get("/questions", qc.GetAll)
	readRouter.Get("/questions/{id}", qc.Get)
	loggedIn.Delete("/questions/{id}", qc.Delete)
	loggedIn.Post("/questions", qc.Create)
	loggedIn.Put("/questions/{id}", qc.Update)
}

func initAuth(r, loggedIn chi.Router, ac auth.Controller) {
	r.Post("/login", ac.DefaultLogin)
	r.Get("/login/{media:(github)}", ac.SocialMediaRedirect)
	r.Get("/login/{media:(github)}/callback", ac.SocialMediaCallback)
	loggedIn.Delete("/logout", ac.Logout)
	r.Post("/revoke", ac.Revoke)
}

func initUsers(r, loggedIn chi.Router, uc users.UserController) {
	r.Post("/users", uc.Create)
	loggedIn.Get("/users", uc.GetAll)
	loggedIn.Get("/users/{id}", uc.Get)
	loggedIn.Put("/users/{id}", uc.Update)
	loggedIn.Delete("/users/{id}", uc.Delete)
}
