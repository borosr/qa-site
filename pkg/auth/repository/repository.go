package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/borosr/qa-site/pkg/auth"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/dgraph-io/badger/v2"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AuthRepository struct {
	bdb *badger.DB
	db  *sql.DB
}

func NewRepository(bdb *badger.DB,
	db *sql.DB) AuthRepository {
	return AuthRepository{
		bdb: bdb,
		db:  db,
	}
}

func (ar AuthRepository) StoreRevokeToken(ctx context.Context, ownerID, revokeToken string) error {
	token := models.RevokeToken{
		ID:      xid.New().String(),
		OwnerID: ownerID,
		Token:   revokeToken,
	}

	return token.Insert(ctx, ar.db, boil.Infer())
}

func (ar AuthRepository) GetRevokeToken(ctx context.Context, ownerID string) string {
	token, err := models.RevokeTokens(qm.Where("owner_id=?", ownerID)).One(ctx, ar.db)
	if err != nil {
		logrus.Error(err)

		return ""
	}

	return token.Token
}

func (ar AuthRepository) StoreJwtToken(ownerID, token string, expr time.Time) error {
	return ar.bdb.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
			value, _ := json.Marshal([]auth.TokenCache{{Token: token, Expr: expr}})

			return txn.Set([]byte(ownerID), value)
		} else if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []auth.TokenCache
			if err := json.Unmarshal(val, &tokens); err != nil {
				return err
			}
			tokens = append(tokens, auth.TokenCache{Token: token, Expr: expr})
			value, _ := json.Marshal(tokens)

			return txn.Set([]byte(ownerID), value)
		})
	})
}

func (ar AuthRepository) DeleteJwtToken(ownerID, token string) error {
	return ar.bdb.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []auth.TokenCache
			if err := json.Unmarshal(val, &tokens); err != nil {
				return err
			}

			var tokenIndex = -1
			for i, t := range tokens {
				if t.Token == token {
					tokenIndex = i

					break
				}
			}

			if tokenIndex == -1 {

				return nil
			}

			tokens[tokenIndex] = tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]

			value, _ := json.Marshal(tokens)

			return txn.Set([]byte(ownerID), value)
		})
	})
}

func (ar AuthRepository) ExistsAndNotExpired(ownerID, token string, now time.Time) error {
	return ar.bdb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []auth.TokenCache
			if err := json.Unmarshal(val, &tokens); err != nil {
				return err
			}

			for _, t := range tokens {
				if t.Token == token && now.Before(t.Expr) {
					return nil
				}
			}

			return fmt.Errorf("token [%s] is not found or expired", token)
		})
	})
}

func (ar AuthRepository) FindRevokeTokenBy(ctx context.Context, userID, token string) (*models.RevokeToken, error) {
	return models.RevokeTokens(
		qm.Where("owner_id=?", userID),
		qm.And("token=?", token),
	).One(ctx, ar.db)
}

func (ar AuthRepository) Update(ctx context.Context, token *models.RevokeToken) (*models.RevokeToken, error) {
	_, err := token.Update(ctx, db.Get(), boil.Infer())
	return token, err
}
