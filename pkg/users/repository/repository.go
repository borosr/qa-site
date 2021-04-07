package repository

import (
	"context"
	"database/sql"

	"github.com/borosr/qa-site/pkg/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur UserRepository) Insert(ctx context.Context, u models.User) error {
	u.ID = xid.New().String()

	return u.Insert(ctx, ur.db, boil.Infer())
}

func (ur UserRepository) FindByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := models.Users(qm.Where("username=?", username)).One(ctx, ur.db)
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (ur UserRepository) GetAll(ctx context.Context) (models.UserSlice, error) {
	return models.Users(qm.Select("id, username, full_name")).All(ctx, ur.db)
}

func (ur UserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	return models.FindUser(ctx, ur.db, id, "id", "username", "full_name")
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
	return err != nil || exists
}
