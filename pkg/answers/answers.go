package answers

import (
	"database/sql"
	"net/http"

	"github.com/borosr/qa-site/pkg/answers/repository"
	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/models"
	questionRepository "github.com/borosr/qa-site/pkg/questions/repository"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

type AnswerController struct {
	answerRepository   repository.AnswerRepository
	questionRepository questionRepository.QuestionRepository
}

func NewController(answerRepository repository.AnswerRepository,
	questionRepository questionRepository.QuestionRepository) AnswerController {
	return AnswerController{
		answerRepository:   answerRepository,
		questionRepository: questionRepository,
	}
}

func (ac AnswerController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	var req Request
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	savedAnswer, err := ac.answerRepository.Insert(ctx, models.Answer{
		Answer:     req.Answer,
		QuestionID: req.QuestionID,
		CreatedBy:  loggedInUser.ID,
		Answered:   null.BoolFrom(false),
	})
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, savedAnswer)
}

func (ac AnswerController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")

	var req UpdateRequest
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	answer, err := ac.answerRepository.FindMyAnswer(ctx, loggedInUser.ID, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	answer.Answer = req.Answer
	if _, err := ac.answerRepository.Update(ctx, *answer); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answer)
}

func (ac AnswerController) GetMyAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	answers, err := ac.answerRepository.FindAnswersCreatedBy(ctx, loggedInUser.ID)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answers)
}

func (ac AnswerController) GetQuestionsAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	questionID := chi.URLParam(r, "questionID")

	answers, err := ac.answerRepository.FindAnswersQuestionID(ctx, questionID)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answers)
}

func (ac AnswerController) SetAnswered(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	questionID := chi.URLParam(r, "questionID")
	answerID := chi.URLParam(r, "answerID")

	_, err := ac.questionRepository.GetCurrentUsers(ctx, questionID, loggedInUser.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Warnf("User [%s] cannot set question [%s] answered with answer [%s]", loggedInUser.ID, questionID, answerID)
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	a, err := ac.answerRepository.SetAnswered(ctx, answerID, questionID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, a)
}
