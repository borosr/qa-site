package users

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/users/repository"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewController(userRepository repository.UserRepository) UserController {
	return UserController{userRepository: userRepository}
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := models.User{}
	if err := api.Bind(r, &user); err != nil || !c.validCreateRequest(ctx, user) {
		api.BadRequest(w)

		return
	}

	if pass, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost); err != nil {
		api.BadRequest(w)

		return
	} else {
		user.Password.SetValid(string(pass))
	}

	if err := c.userRepository.Insert(ctx, user); err != nil {
		api.InternalServerError(w)

		return
	}

	// TODO use response type here
	api.SuccessResponse(w, struct {
		Msg string
	}{
		Msg: "user successfully created",
	})
}

func (c UserController) validCreateRequest(ctx context.Context, user models.User) bool {
	valid := user.ID == "" && user.Username != "" &&
		(user.FullName.Valid && user.FullName.String != "") && (user.Password.Valid && user.Password.String != "")

	if exists, err := models.Users(qm.Where("username=?", user.Username)).Exists(ctx, db.Get()); err != nil || exists {
		log.Warnf("user with username [%s] already exist", user.Username)
		return false
	}

	return valid
}

func (c UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := c.userRepository.GetAll(ctx)
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

func (c UserController) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	user, err := c.userRepository.Get(r.Context(), id)
	if err != nil {
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, user)
}

func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	if err := c.userRepository.Delete(r.Context(), loggedInUser); err != nil {
		api.InternalServerError(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "success"})
}

func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	loggedInUser := r.Context().Value("user").(models.User)
	if loggedInUser.ID != id {
		api.Forbidden(w)

		return
	}

	var u models.User
	if err := api.Bind(r, &u); err != nil || !c.validUpdateRequest(u) {
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

	if _, err := c.userRepository.Update(ctx, u); err != nil {
		api.InternalServerError(w)

		return
	}

	u.Password = null.String{}
	api.SuccessResponse(w, u)
}

func (c UserController) validUpdateRequest(u models.User) bool {
	return u.Username != "" && u.FullName.String != ""
}
