package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/samber/do"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	StatusPublished = "Published"
	StatusArchived  = "Archived"
	StatusDeleted   = "Deleted"

	getAllSelect = "id, title, description, created_by, created_at, status"
	ratingSum    = "(SELECT COALESCE(SUM(value),0) as rating FROM ratings WHERE record_id=questions.id AND kind='question')"
)

var (
	questionRatingQuery = []qm.QueryMod{
		qm.Select(getAllSelect, ratingSum),
		qm.Where("status=?", StatusPublished),
	}
)

type QuestionRepository struct {
	db *sql.DB
}

func NewRepository(i *do.Injector) (QuestionRepository, error) {
	db, err := do.Invoke[*sql.DB](i)
	if err != nil {
		return QuestionRepository{}, err
	}
	return QuestionRepository{
		db: db,
	}, nil
}

func (qr QuestionRepository) Insert(ctx context.Context, q models.Question) (models.Question, error) {
	q.ID = xid.New().String()
	q.CreatedAt.SetValid(time.Now())
	return q, q.Insert(ctx, qr.db, boil.Infer())
}

func (qr QuestionRepository) Update(ctx context.Context, q models.Question) (models.Question, error) {
	_, err := q.Update(ctx, qr.db, boil.Infer())
	return q, err
}

func (qr QuestionRepository) Get(ctx context.Context, ID string) (QuestionWithRating, error) {
	var q QuestionWithRating
	err := models.Questions(
		append(
			questionRatingQuery,
			qm.And("id=?", ID))...).
		Bind(ctx, qr.db, &q)
	return q, err
}

func (qr QuestionRepository) GetCurrentUsers(ctx context.Context, ID, loggedInUserID string) (models.Question, error) {
	q, err := models.Questions(
		qm.Where("id=?", ID),
		qm.And("created_by=?", loggedInUserID),
		qm.And("status!=?", StatusDeleted)).One(ctx, qr.db)
	if err != nil {
		return models.Question{}, err
	}

	return *q, nil
}

func (qr QuestionRepository) Delete(ctx context.Context, ID, loggedInUserID string) error {
	question, err := models.Questions(
		qm.Where("id=?", ID),
		qm.And("created_by=?", loggedInUserID),
		qm.And("status!=?", StatusDeleted)).One(ctx, qr.db)

	if err != nil {
		return err
	}

	question.Status.SetValid(StatusDeleted)

	_, err = question.Update(ctx, qr.db, boil.Infer())
	return err
}

func (qr QuestionRepository) GetAll(ctx context.Context, sort string, limit, offset int) ([]QuestionWithRating, int64, error) {
	var resp []QuestionWithRating
	if err := models.Questions(
		append(questionRatingQuery,
			qm.OrderBy(sort),
			qm.Limit(limit),
			qm.Offset(offset))...).
		Bind(ctx, qr.db, &resp); err != nil {
		return nil, 0, err
	}

	count, err := models.Questions(questionRatingQuery...).Count(ctx, qr.db)
	return resp, count, err
}
