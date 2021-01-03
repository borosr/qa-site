package auth

import (
	"context"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func StoreRevokeToken(ctx context.Context, ownerID, revokeToken string) error {
	token := models.RevokeToken{
		ID:      xid.New().String(),
		OwnerID: ownerID,
		Token:   revokeToken,
	}

	return token.Insert(ctx, db.Get(), boil.Infer())
}

func GetRevokeToken(ctx context.Context, ownerID string) string {
	token, err := models.RevokeTokens(qm.Where("owner_id=?", ownerID)).One(ctx, db.Get())
	if err != nil {
		logrus.Error(err)

		return ""
	}

	return token.Token
}
