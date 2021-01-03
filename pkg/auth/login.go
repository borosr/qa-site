package auth

import (
	"net/http"
	"time"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwtTimeout    = 30 * time.Minute
	revokeTimeout = 72 * time.Hour
)

var (
	tokenCache = map[string][]TokenCache{}
)

type TokenCache struct {
	Token string
	Expr  time.Time
}

func DefaultLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := Request{}
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	var user models.User
	var err error
	if user, err = users.FindByUsername(ctx, req.Username); err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password)); err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	now := time.Now()

	token, err := generateAuthToken(user, now)
	if err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	tokenCache[user.ID] = append(tokenCache[user.ID], TokenCache{
		Token: token,
		Expr:  now.Add(jwtTimeout),
	})

	var revokeToken string
	if revokeToken = GetRevokeToken(ctx, user.ID); revokeToken == "" {
		var err error
		revokeToken, err = generateRevokeToken(user, now)
		if err != nil {
			log.Error(err)
			api.Forbidden(w)

			return
		}
		if err := StoreRevokeToken(ctx, user.ID, revokeToken); err != nil {
			log.Error(err)
			api.Forbidden(w)

			return
		}
	}

	api.SuccessResponse(w, Response{
		Token:       token,
		RevokeToken: revokeToken,
		AuthKind:    DefaultAuthKind,
	})
}

func generateAuthToken(user models.User, now time.Time) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"sid": user.ID,
		"exp": now.Add(jwtTimeout).Unix(),
	}).SignedString([]byte(settings.Get().JwtHMAC))
	if err != nil {
		return "", err
	}

	return token, err
}

func generateRevokeToken(user models.User, now time.Time) (string, error) {
	revokeToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"sid": user.ID,
		"exp": now.Add(revokeTimeout).Unix(),
	}).SignedString([]byte(settings.Get().RevokeHMAC))

	if err != nil {
		return "", err
	}

	return revokeToken, err
}
