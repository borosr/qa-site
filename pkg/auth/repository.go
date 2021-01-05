package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/dgraph-io/badger/v2"
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

func StoreJwtToken(ownerID, token string, expr time.Time) error {
	return db.GetBDB().Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil && errors.Is(err, badger.ErrKeyNotFound) {
			value, _ := json.Marshal([]TokenCache{{Token: token, Expr: expr}})

			return txn.Set([]byte(ownerID), value)
		} else if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []TokenCache
			if err := json.Unmarshal(val, &tokens); err != nil {
				return err
			}
			tokens = append(tokens, TokenCache{Token: token, Expr: expr})
			value, _ := json.Marshal(tokens)

			return txn.Set([]byte(ownerID), value)
		})
	})
}

func DeleteJwtToken(ownerID, token string) error {
	return db.GetBDB().Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []TokenCache
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

func ExistsAndNotExpired(ownerID, token string, now time.Time) error {
	return db.GetBDB().View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(ownerID))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var tokens []TokenCache
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
