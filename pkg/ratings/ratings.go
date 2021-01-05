package ratings

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	defaultUnrateValue = iota - 1
	_
	defaultRateValue
)

func Rate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if isOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	log.Infof("Rating %s with id [%s]", k, id)
	result, err := Rating(ctx, k, loggedInUser.ID, id)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, result)
}

func Unrate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if isOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	log.Infof("Unrating %s with id [%s]", k, id)
	result, err := Unrating(ctx, k, loggedInUser.ID, id)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, result)
}

func Dismiss(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if isOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	rating, err := models.Ratings(buildRateFilter(k, loggedInUser.ID, id)...).One(ctx, db.Get())
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Errorf("rating not found for user [%s] and [%s][%s]", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	if rating.RatedBy != loggedInUser.ID {
		log.Errorf("User [%s] can't change other user's rating [%s]", loggedInUser.ID, rating.ID)
		api.BadRequest(w)

		return
	}

	if _, err := rating.Delete(ctx, db.Get()); err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "Rate dismissed"})
}

func isOwner(ctx context.Context, k kind, userID, id string) bool {
	var err error
	var exists bool
	switch k {
	case QuestionKind:
		exists, err = models.Questions(qm.Where("id=?", id), qm.And("created_by=?", userID)).Exists(ctx, db.Get())
	case AnswerKind:
		exists, err = models.Answers(qm.Where("id=?", id), qm.And("created_by=?", userID)).Exists(ctx, db.Get())
	}
	if err != nil {
		log.Error(err)

		return true // to prevent rating
	}

	return exists
}

func getKind(k string) (kind, error) {
	switch k {
	case "questions":
		return QuestionKind, nil
	case "answers":
		return AnswerKind, nil
	}

	return "", fmt.Errorf("invalid kind: [%s]", k)
}
