package users

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := models.Users(qm.Select("id, username, full_name")).All(ctx, db.Get())
	if errors.Is(err, sql.ErrNoRows) {
		api.NotFound(w)

		return
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, users)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	user, err := models.FindUser(r.Context(), db.Get(), id, "id", "username", "full_name")
	if err != nil {
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	if _, err := loggedInUser.Delete(r.Context(), db.Get()); err != nil {
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "success"})
}

func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	var u models.User
	if err := api.Bind(r, &u); err != nil || !validUpdateRequest(u) {
		api.BadRequest(w)

		return
	}
	u.ID = id
	if u.Password.Valid && u.Password.String != "" {
		if pass, err := bcrypt.GenerateFromPassword([]byte(u.Password.String), bcrypt.DefaultCost); err != nil {
			api.BadRequest(w)

			return
		} else {
			u.Password.SetValid(string(pass))
		}
	} else {
		u.Password = loggedInUser.Password
	}

	if _, err := u.Update(ctx, db.Get(), boil.Infer()); err != nil {
		api.InternalServerError(w)

		return
	}

	u.Password = null.String{}
	api.SuccessResponse(w, u)
}

func validUpdateRequest(u models.User) bool {
	return u.Username != "" && u.FullName.String != ""
}
