package users

import (
	"context"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := models.User{}
	if err := api.Bind(r, &user); err != nil || !validCreateRequest(ctx, user) {
		api.BadRequest(w)

		return
	}

	if pass, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost); err != nil {
		api.BadRequest(w)

		return
	} else {
		user.Password.SetValid(string(pass))
	}

	if err := Insert(ctx, user); err != nil {
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, struct {
		Msg string
	}{
		Msg: "user successfully created",
	})
}

func validCreateRequest(ctx context.Context, user models.User) bool {
	valid := user.ID == "" && user.Username != "" &&
		(user.FullName.Valid && user.FullName.String != "") && (user.Password.Valid && user.Password.String != "")

	if exists, err := models.Users(qm.Where("username=?", user.Username)).Exists(ctx, db.Get()); err != nil || exists {
		log.Warnf("user with username [%s] already exist", user.Username)
		return false
	}

	return valid
}
