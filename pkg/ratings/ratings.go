package ratings

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/ratings/repository"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type RateController struct {
	rateRepository repository.RateRepository
}

func NewController(rateRepository repository.RateRepository) RateController {
	return RateController{
		rateRepository: rateRepository,
	}
}

func (rc RateController) Rate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if rc.rateRepository.IsOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	log.Infof("Rating %s with id [%s]", k, id)
	result, err := rc.rateRepository.Rating(ctx, k, loggedInUser.ID, id)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, Response{Value: result})
}

func (rc RateController) Unrate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if rc.rateRepository.IsOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	log.Infof("Unrating %s with id [%s]", k, id)
	result, err := rc.rateRepository.Unrating(ctx, k, loggedInUser.ID, id)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, Response{Value: result})
}

func (rc RateController) Dismiss(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)
	id := chi.URLParam(r, "id")
	k, err := getKind(chi.URLParam(r, "kind"))
	if err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	if rc.rateRepository.IsOwner(ctx, k, loggedInUser.ID, id) {
		log.Errorf("owner [%s] of the record [%s]:[%s] trying to rate", loggedInUser.ID, k, id)
		api.BadRequest(w)

		return
	}

	rating, err := models.Ratings(rc.rateRepository.BuildRateFilter(k, loggedInUser.ID, id)...).One(ctx, db.Get())
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

	result, err := rc.rateRepository.GetCurrentValue(ctx, k, loggedInUser.ID, id)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, Response{Value: result})
}

func getKind(k string) (repository.Kind, error) {
	switch k {
	case "questions":
		return repository.QuestionKind, nil
	case "answers":
		return repository.AnswerKind, nil
	}

	return "", fmt.Errorf("invalid kind: [%s]", k)
}
