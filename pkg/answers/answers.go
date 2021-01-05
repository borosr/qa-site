package answers

import (
	"database/sql"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Request struct {
	UpdateRequest
	QuestionID string `json:"question_id"`
}

type UpdateRequest struct {
	Answer string `json:"answer"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	var req Request
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	savedAnswer, err := Insert(ctx, models.Answer{
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

func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")

	var req UpdateRequest
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	answer, err := FindMyAnswer(ctx, loggedInUser.ID, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	answer.Answer = req.Answer
	if _, err := answer.Update(ctx, db.Get(), boil.Infer()); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answer)
}

func GetMyAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	answers, err := FindAnswersCreatedBy(ctx, loggedInUser.ID)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answers)
}

func GetQuestionsAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	questionID := chi.URLParam(r, "questionID")

	answers, err := FindAnswersQuestionID(ctx, questionID)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answers)
}

func SetAnswered(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	questionID := chi.URLParam(r, "questionID")
	answerID := chi.URLParam(r, "answerID")

	questionExist, err := models.Questions(
		qm.Where("id=?", questionID),
		qm.And("created_by=?", loggedInUser.ID)).Exists(ctx, db.Get())
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}
	if !questionExist {
		log.Warnf("User [%s] cannot set question [%s] answered with answer [%s]", loggedInUser.ID, questionID, answerID)
		api.BadRequest(w)

		return
	}

	answer, err := models.Answers(qm.Where("id=?", answerID), qm.And("question_id=?", questionID)).One(ctx, db.Get())
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	answer.Answered.SetValid(true)
	if _, err := answer.Update(ctx, db.Get(), boil.Infer()); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, answer)
}
