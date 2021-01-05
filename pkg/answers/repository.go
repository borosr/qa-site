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

func Insert(ctx context.Context, a models.Answer) (models.Answer, error) {
	a.ID = xid.New().String()
	a.CreatedAt.SetValid(time.Now())

	return a, a.Insert(ctx, db.Get(), boil.Infer())
}

func FindAnswersCreatedBy(ctx context.Context, ownerID string) (models.AnswerSlice, error) {
	return models.Answers(qm.Where("created_by=?", ownerID)).All(ctx, db.Get())
}

func FindMyAnswer(ctx context.Context, ownerID, id string) (*models.Answer, error) {
	return models.Answers(qm.Where("id=?", id), qm.And("created_by=?", ownerID)).One(ctx, db.Get())
}

func FindAnswersQuestionID(ctx context.Context, questionID string) (models.AnswerSlice, error) {
	return models.Answers(qm.Where("question_id=?", questionID)).All(ctx, db.Get())
}
