package answers

import (
	"context"
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

func Insert(ctx context.Context, a models.Answer) (models.Answer, error) {
	a.ID = xid.New().String()
	a.CreatedAt.SetValid(time.Now())

	return a, a.Insert(ctx, db.Get(), boil.Infer())
}

func FindAnswersCreatedBy(ctx context.Context, ownerID string) ([]Response, error) {
	var resp []Response
	if err := models.Answers(
		qm.Select(getAllSelect, ratingSum),
		qm.Where("created_by=?", ownerID),
	).Bind(ctx, db.Get(), &resp);
		err != nil {
		return nil, err
	}

	return resp, nil
}

func FindMyAnswer(ctx context.Context, ownerID, id string) (*models.Answer, error) {
	return models.Answers(qm.Where("id=?", id), qm.And("created_by=?", ownerID)).One(ctx, db.Get())
}

func FindAnswersQuestionID(ctx context.Context, questionID string) ([]Response, error) {
	var resp []Response
	if err := models.Answers(
		qm.Select(getAllSelect, ratingSum),
		qm.Where("question_id=?", questionID),
	).Bind(ctx, db.Get(), &resp);
		err != nil {
		return nil, err
	}

	return resp, nil
}
