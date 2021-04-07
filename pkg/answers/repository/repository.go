package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	getAllSelect = "id, question_id, created_by, answer, created_at, answered"
	ratingSum    = "(SELECT SUM(sum) as rating FROM (SELECT SUM(value) as sum FROM ratings WHERE record_id=answers.id AND kind='answer' UNION SELECT 0 as sum)) AS rating"
)

type AnswerRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) AnswerRepository {
	return AnswerRepository{
		db: db,
	}
}

func (ar AnswerRepository) Insert(ctx context.Context, a models.Answer) (models.Answer, error) {
	a.ID = xid.New().String()
	a.CreatedAt.SetValid(time.Now())

	return a, a.Insert(ctx, ar.db, boil.Infer())
}

func (ar AnswerRepository) Update(ctx context.Context, a models.Answer) (models.Answer, error) {
	_, err := a.Update(ctx, db.Get(), boil.Infer())

	return a, err
}

func (ar AnswerRepository) FindAnswersCreatedBy(ctx context.Context, ownerID string) ([]AnswerWithRating, error) {
	var resp []AnswerWithRating
	if err := models.Answers(
		qm.Select(getAllSelect, ratingSum),
		qm.Where("created_by=?", ownerID),
	).Bind(ctx, ar.db, &resp);
		err != nil {
		return nil, err
	}

	return resp, nil
}

func (ar AnswerRepository) FindMyAnswer(ctx context.Context, ownerID, id string) (*models.Answer, error) {
	return models.Answers(qm.Where("id=?", id), qm.And("created_by=?", ownerID)).One(ctx, ar.db)
}

func (ar AnswerRepository) FindAnswersQuestionID(ctx context.Context, questionID string) ([]AnswerWithRating, error) {
	var resp []AnswerWithRating
	if err := models.Answers(
		qm.Select(getAllSelect, ratingSum),
		qm.Where("question_id=?", questionID),
	).Bind(ctx, ar.db, &resp);
		err != nil {
		return nil, err
	}

	return resp, nil
}

func (ar AnswerRepository) SetAnswered(ctx context.Context, ID, questionID string) (models.Answer, error) {
	answer, err := models.Answers(qm.Where("id=?", ID), qm.And("question_id=?", questionID)).One(ctx, ar.db)
	if err != nil {
		return models.Answer{}, err
	}

	answer.Answered.SetValid(true)
	if _, err := answer.Update(ctx, ar.db, boil.Infer()); err != nil {
		return models.Answer{}, err
	}

	return *answer, nil
}

