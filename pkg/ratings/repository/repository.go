package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/samber/do"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	defaultUnrateValue = iota - 1
	_
	defaultRateValue
)

const (
	QuestionKind Kind = "question"
	AnswerKind   Kind = "answer"
)

type Kind string

type RateRepository struct {
	db *sql.DB
}

func NewRepository(i *do.Injector) (RateRepository, error) {
	db, err := do.Invoke[*sql.DB](i)
	if err != nil {
		return RateRepository{}, err
	}
	return RateRepository{
		db: db,
	}, nil
}

func (rr RateRepository) Exists(ctx context.Context, k Kind, userID, recordID string) (bool, error) {
	return models.Ratings(rr.BuildRateFilter(k, userID, recordID)...).Exists(ctx, db.Get())
}

func (rr RateRepository) Rating(ctx context.Context, k Kind, userID, id string) (int64, error) {
	if _, err := rr.store(ctx, k, userID, id, defaultRateValue); err != nil {
		return 0, err
	}

	return rr.GetCurrentValue(ctx, k, userID, id)
}

func (rr RateRepository) Unrating(ctx context.Context, k Kind, userID, id string) (int64, error) {
	if _, err := rr.store(ctx, k, userID, id, defaultUnrateValue); err != nil {
		return 0, err
	}

	return rr.GetCurrentValue(ctx, k, userID, id)
}

func (rr RateRepository) store(ctx context.Context, k Kind, userID string, id string, value int16) (models.Rating, error) {
	exists, err := rr.Exists(ctx, k, userID, id)
	if err != nil {
		return models.Rating{}, err
	}

	if exists {
		r, err := models.Ratings(rr.BuildRateFilter(k, userID, id)...).One(ctx, db.Get())
		if err != nil {
			return models.Rating{}, err
		}

		r.RatedAt = time.Now()
		r.Value = value
		_, err = r.Update(ctx, db.Get(), boil.Infer())

		return *r, err
	}

	r := models.Rating{
		ID:       xid.New().String(),
		Kind:     string(k),
		RecordID: id,
		RatedBy:  userID,
		RatedAt:  time.Now(),
		Value:    value,
	}

	return r, r.Insert(ctx, db.Get(), boil.Infer())
}

func (rr RateRepository) GetCurrentValue(ctx context.Context, k Kind, userID, recordID string) (int64, error) {
	var data struct {
		Value int64 `boil:"value"`
	}

	if err := models.Ratings(append(rr.BuildRateFilter(k, userID, recordID),
		qm.Select("COALESCE(SUM(value), 0) as value"))...).Bind(ctx, db.Get(), &data); err != nil {
		return 0, err
	}
	return data.Value, nil
}

func (rr RateRepository) BuildRateFilter(k Kind, userID, recordID string) []qm.QueryMod {
	return []qm.QueryMod{
		qm.Where("kind=?", k),
		qm.And("rated_by=?", userID),
		qm.And("record_id=?", recordID),
	}
}

func (rr RateRepository) IsOwner(ctx context.Context, k Kind, userID, id string) bool {
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
