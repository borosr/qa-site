package questions

import (
	"database/sql"
	"net/http"
	"strconv"

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

const (
	StatusPublished = "Published"
	StatusArchived  = "Archived"
	StatusDeleted   = "Deleted"

	DefaultOffset = 0
	DefaultLimit  = 10

	getAllSelect = "id, title, description, created_by, created_at, status"
	ratingSum    = "(SELECT SUM(sum) as rating FROM (SELECT SUM(value) as sum FROM ratings WHERE record_id=questions.id AND kind='question' UNION SELECT 0 as sum)) AS rating"
)

func Create(w http.ResponseWriter, r *http.Request) {
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
	if question, err = Insert(ctx, models.Question{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   loggedInUser.ID,
		Status:      null.StringFrom(StatusPublished),
	}); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, question)
}

func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")

	var req Request
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	question, err := models.Questions(
		qm.Where("id=?", id),
		qm.And("created_by=?", loggedInUser.ID),
		qm.And("status!=?", StatusDeleted)).One(ctx, db.Get())
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

	if _, err := question.Update(ctx, db.Get(), boil.Infer()); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, question)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
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

	var resp []Response
	if err := models.Questions(
		append(buildQuestionsRatingQuery(),
			qm.Limit(limit),
			qm.Offset(offset))...).
		Bind(ctx, db.Get(), &resp);
		err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	var count int64
	if count, err = models.Questions(buildQuestionsRatingQuery()...).Count(ctx, db.Get()); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, PageableResponse{
		Data:  resp,
		Count: count,
	})
}

func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var resp Response
	if err := models.Questions(
		append(
			buildQuestionsRatingQuery(),
			qm.And("id=?", id))...).
		Bind(ctx, db.Get(), &resp);
		err != nil && errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, resp)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)

	question, err := models.Questions(
		qm.Where("id=?", id),
		qm.And("created_by=?", loggedInUser.ID),
		qm.And("status!=?", StatusDeleted)).One(ctx, db.Get())
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	question.Status.SetValid(StatusDeleted)

	if _, err := question.Update(ctx, db.Get(), boil.Infer()); err != nil {
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

func buildQuestionsRatingQuery() []qm.QueryMod {
	return []qm.QueryMod{
		qm.Select(getAllSelect, ratingSum),
		qm.Where("status=?", StatusPublished),
	}
}
