package users

import (
	"context"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Insert(ctx context.Context, u models.User) error {
	u.ID = xid.New().String()

	return u.Insert(ctx, db.Get(), boil.Infer())
}