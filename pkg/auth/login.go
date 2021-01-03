package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/friendsofgo/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwtTimeout    = 30 * time.Minute
	revokeTimeout = 72 * time.Hour

	AuthorizationHeader = "Authorization"
)

var (
	tokenCache = map[string][]TokenCache{}

	ErrForbidden = errors.New("forbidden")
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

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var jwtToken string
			if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
				api.InternalServerError(w)

				return
			}

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(settings.Get().JwtHMAC), nil
			})

			if err != nil {
				log.Error(err)
				api.Forbidden(w)

				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx, err = chainTokenBySID(w, claims, jwtToken, ctx)
				if err != nil {
					api.Forbidden(w)

					return
				}
			} else {
				log.Errorf("invalid token claim [%#v]", claims)
				api.Forbidden(w)

				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func chainTokenBySID(w http.ResponseWriter, claims jwt.MapClaims, jwtToken string, ctx context.Context) (context.Context, error) {
	if sid, ok := claims["sid"].(string); ok {
		return chainTokenCacheCheck(w, sid, jwtToken, ctx)
	} else {
		log.Errorf("missing sid from claim [%+v]", claims)
		return ctx, ErrForbidden
	}
}

func chainTokenCacheCheck(w http.ResponseWriter, sid string, jwtToken string, ctx context.Context) (context.Context, error) {
	if tokens, ok := tokenCache[sid]; ok {
		now := time.Now()
		var found bool
		for _, t := range tokens {
			if t.Token == jwtToken && now.Before(t.Expr) {
				ctx, found = chainGetUserIfFound(ctx, sid)
				break
			}
		}
		if !found {
			return ctx, ErrForbidden
		}
	} else {
		log.Errorf("missing token to the sid [%s]", sid)
		return ctx, ErrForbidden
	}

	return ctx, nil
}

func chainGetUserIfFound(ctx context.Context, sid string) (context.Context, bool) {
	var found bool
	user, err := models.FindUser(ctx, db.Get(), sid)
	if err == nil && user != nil {
		ctx = context.WithValue(ctx, "user", *user)
		found = true
	} else if err != nil {
		log.Error(err)
	}

	return ctx, found
}
