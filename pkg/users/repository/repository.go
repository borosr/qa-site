package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/samber/do"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(i *do.Injector) (UserRepository, error) {
	db, err := do.Invoke[*sql.DB](i)
	if err != nil {
		return UserRepository{}, err
	}
	return UserRepository{
		db: db,
	}, nil
}

func (ur UserRepository) Insert(ctx context.Context, u models.User) error {
	u.ID = xid.New().String()

	return u.Insert(ctx, ur.db, boil.Infer())
}

func (ur UserRepository) FindByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := models.Users(qm.Where("username=?", username)).One(ctx, ur.db)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			healthcheck.Instance().Failed()
		}
		return models.User{}, err
	}

	return *user, nil
}

func (ur UserRepository) GetAll(ctx context.Context) (models.UserSlice, error) {
	users, err := models.Users(qm.Select("id, username, full_name")).All(ctx, ur.db)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			healthcheck.Instance().Failed()
		}
		return models.UserSlice{}, err
	}
	return users, nil
}

func (ur UserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	user, err := models.FindUser(ctx, ur.db, id, "id", "username", "full_name")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			healthcheck.Instance().Failed()
		}
		return &models.User{}, err
	}
	return user, nil
}

func (ur UserRepository) Delete(ctx context.Context, m models.User) error {
	_, err := m.Delete(ctx, ur.db)
	return err
}

func (ur UserRepository) Update(ctx context.Context, u models.User) (models.User, error) {
	_, err := u.Update(ctx, ur.db, boil.Infer())
	return u, err
}

func (ur UserRepository) ExistsByUsername(ctx context.Context, username string) bool {
	exists, err := models.Users(qm.Where("username=?", username)).Exists(ctx, ur.db)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Error(err, "error getting user by username")
			return false
		}
	}
	return exists
}
