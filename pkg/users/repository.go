package users

import (
	"context"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Insert(ctx context.Context, u models.User) error {
	u.ID = xid.New().String()

	return u.Insert(ctx, db.Get(), boil.Infer())
}

func FindByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := models.Users(qm.Where("username=?", username)).One(ctx, db.Get())
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}
