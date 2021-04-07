package questions

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/questions/repository"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

const (
	DefaultOffset = 0
	DefaultLimit  = 10
	DefaultSort = "created_at"
)

type QuestionController struct {
	questionRepository repository.QuestionRepository
}

func NewController(questionRepository repository.QuestionRepository) QuestionController {
	return QuestionController{
		questionRepository: questionRepository,
	}
}

func (qc QuestionController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	var req Request
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	var err error
	var question models.Question
	if question, err = qc.questionRepository.Insert(ctx, models.Question{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   loggedInUser.ID,
		Status:      null.StringFrom(repository.StatusPublished),
	}); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, question)
}

func (qc QuestionController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")

	var req Request
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	question, err := qc.questionRepository.GetCurrentUsers(ctx, id, loggedInUser.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	question.Title = req.Title
	question.Description = req.Description

	if _, err := qc.questionRepository.Update(ctx, question); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, question)
}

func (qc QuestionController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit, err := getLimit(r)
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}
	offset, err := getOffset(r)
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	questions, count, err := qc.questionRepository.GetAll(ctx, getSort(r), limit, offset)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, PageableResponse{
		Data:  questions,
		Count: count,
	})
}

func getSort(r *http.Request) string {
	if queryValue := r.URL.Query().Get("sort"); queryValue != "" {
		return queryValue
	}

	return DefaultSort
}

func (qc QuestionController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	q, err := qc.questionRepository.Get(ctx, id)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, q)
}

func (qc QuestionController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)

	err := qc.questionRepository.Delete(ctx, id, loggedInUser.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "OK"})
}

func getLimit(r *http.Request) (int, error) {
	queryValue := r.URL.Query().Get("limit")
	if queryValue == "" {
		return DefaultLimit, nil
	}
	limit, err := strconv.Atoi(queryValue)
	if err != nil {
		return 0, err
	}

	return limit, err
}

func getOffset(r *http.Request) (int, error) {
	queryValue := r.URL.Query().Get("offset")
	if queryValue == "" {
		return DefaultOffset, nil
	}
	limit, err := strconv.Atoi(queryValue)
	if err != nil {
		return 0, err
	}

	return limit, err
}
