package ratings

import (
	"context"
	"time"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Exists(ctx context.Context, k kind, userID, recordID string) (bool, error) {
	return models.Ratings(buildRateFilter(k, userID, recordID)...).Exists(ctx, db.Get())
}

func Rating(ctx context.Context, k kind, userID, id string) (int64, error) {
	if _, err := store(ctx, k, userID, id, defaultRateValue); err != nil {
		return 0, err
	}

	return getCurrentValue(ctx, k, userID, id)
}

func Unrating(ctx context.Context, k kind, userID, id string) (int64, error) {
	if _, err := store(ctx, k, userID, id, defaultUnrateValue); err != nil {
		return 0, err
	}

	return getCurrentValue(ctx, k, userID, id)
}

func store(ctx context.Context, k kind, userID string, id string, value int16) (models.Rating, error) {
	exists, err := Exists(ctx, k, userID, id)
	if err != nil {
		return models.Rating{}, err
	}

	if exists {
		r, err := models.Ratings(buildRateFilter(k, userID, id)...).One(ctx, db.Get())
		if err != nil {
			return models.Rating{}, err
		}

		r.RatedAt = time.Now()
		r.Value = value
		_, err = r.Update(ctx, db.Get(), boil.Infer())

		return *r, err
	}

	r := models.Rating{
		ID:       xid.New().String(),
		Kind:     string(k),
		RecordID: id,
		RatedBy:  userID,
		RatedAt:  time.Now(),
		Value:    value,
	}

	return r, r.Insert(ctx, db.Get(), boil.Infer())
}

func getCurrentValue(ctx context.Context, k kind, userID, recordID string) (int64, error) {
	var data struct {
		Value int64 `boil:"value"`
	}

	if err := models.Ratings(append(buildRateFilter(k, userID, recordID),
		qm.Select("COALESCE(SUM(value), 0) as value"))...).Bind(ctx, db.Get(), &data); err != nil {
		return 0, err
	}
	return data.Value, nil
}

func buildRateFilter(k kind, userID, recordID string) []qm.QueryMod {
	return []qm.QueryMod{
		qm.Where("kind=?", k),
		qm.And("rated_by=?", userID),
		qm.And("record_id=?", recordID),
	}
}
