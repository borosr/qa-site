package questions

import (
	"context"
	"time"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Insert(ctx context.Context, q models.Question) (models.Question, error) {
	q.ID = xid.New().String()
	q.CreatedAt.SetValid(time.Now())
	return q, q.Insert(ctx, db.Get(), boil.Infer())
}
