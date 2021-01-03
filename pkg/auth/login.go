package auth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/auth/oauth"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/borosr/qa-site/pkg/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

	if !user.Password.Valid && user.AccessToken.Valid {
		log.Errorf("invalid authentication type for user [%s]", req.Username)
		api.BadRequest(w)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password)); err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	resp, err := loggingIn(ctx, user, time.Now(), DefaultAuthKind)
	if err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	api.SuccessResponse(w, resp)
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
				ctx, err = chainTokenBySID(ctx, claims, jwtToken)
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

func SocialMediaRedirect(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "media")

	if err := oauth.Redirect(w, r, provider); err != nil {
		api.InternalServerError(w)
	}
}

func SocialMediaCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	provider := chi.URLParam(r, "media")

	callbackResp, err := oauth.Callback(w, r, provider)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	user, err := models.Users(qm.Where("username=?", callbackResp.UserDetails.Username())).One(ctx, db.Get())
	if errors.Is(err, sql.ErrNoRows) {
		user = &models.User{
			ID:          xid.New().String(),
			Username:    callbackResp.UserDetails.Username(),
			FullName:    null.StringFrom(callbackResp.UserDetails.FullName()),
			AccessToken: null.StringFrom(callbackResp.Response.AccessToken),
		}
		if err := user.Insert(ctx, db.Get(), boil.Infer()); err != nil {
			api.InternalServerError(w)

			return
		}
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	} else {
		user.AccessToken.SetValid(callbackResp.Response.AccessToken)
		if _, err := user.Update(ctx, db.Get(), boil.Infer()); err != nil {
			log.Error(err)
			api.InternalServerError(w)

			return
		}
	}

	resp, err := loggingIn(ctx, *user, time.Now(), GithubAuthKind)
	if err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	api.SuccessResponse(w, resp)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	loggedInUser := r.Context().Value("user").(models.User)

	var jwtToken string
	if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
		api.InternalServerError(w)

		return
	}

	if err := deleteAccessToken(loggedInUser, jwtToken); err != nil {
		log.Errorf("invalid access token [%s]", jwtToken)
		api.BadRequest(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "Logout success"})
}

func Revoke(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := r.Context().Value("user").(models.User)

	var revoke = struct {
		Token string `json:"revoke_token"`
	}{}
	if err := api.Bind(r, &revoke); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	revokeTokenModel, err := models.RevokeTokens(qm.Where("owner_id=?", loggedInUser.ID), qm.And("token=?", revoke.Token)).One(ctx, db.Get())
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	revokeToken, err := jwt.Parse(revokeTokenModel.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(settings.Get().RevokeHMAC), nil
	})

	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	var jwtToken string
	if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
		api.InternalServerError(w)

		return
	}

	if claims, ok := revokeToken.Claims.(jwt.MapClaims); ok && revokeToken.Valid {
		if exp, ok := claims["exp"].(int64); ok {
			if time.Now().Unix() <= exp {
				resp, err := renewTokens(ctx, loggedInUser, revokeTokenModel, jwtToken)
				if err != nil {
					log.Error(err)
					api.InternalServerError(w)

					return
				}
				api.SuccessResponse(w, resp)

				return
			}
		}
	}

	http.Redirect(w, r, "/api/logout", 301)
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

func chainTokenBySID(ctx context.Context, claims jwt.MapClaims, jwtToken string) (context.Context, error) {
	if sid, ok := claims["sid"].(string); ok {
		return chainTokenCacheCheck(ctx, sid, jwtToken)
	} else {
		log.Errorf("missing sid from claim [%+v]", claims)
		return ctx, ErrForbidden
	}
}

func chainTokenCacheCheck(ctx context.Context, sid string, jwtToken string) (context.Context, error) {
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

func loggingIn(ctx context.Context, user models.User, now time.Time, tokenKind AuthKind) (Response, error) {
	token, err := generateAuthToken(user, now)
	if err != nil {
		return Response{}, err
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
			return Response{}, err
		}
		if err := StoreRevokeToken(ctx, user.ID, revokeToken); err != nil {
			return Response{}, err
		}
	}

	return Response{
		Token:       token,
		RevokeToken: revokeToken,
		AuthKind:    tokenKind,
	}, nil
}

func renewTokens(ctx context.Context, user models.User, model *models.RevokeToken, token string) (Response, error) {
	if err := deleteAccessToken(user, token); err != nil {
		return Response{}, err
	}

	now := time.Now()
	authToken, err := generateAuthToken(user, now)
	if err != nil {
		return Response{}, err
	}

	tokenCache[user.ID] = append(tokenCache[user.ID], TokenCache{
		Token: authToken,
		Expr:  now.Add(jwtTimeout),
	})

	revokeToken, err := generateRevokeToken(user, now)
	if err != nil {
		return Response{}, err
	}
	model.Token = revokeToken
	if _, err := model.Update(ctx, db.Get(), boil.Infer()); err != nil {
		return Response{}, err
	}

	return Response{
		Token:       authToken,
		RevokeToken: revokeToken,
		AuthKind:    DefaultAuthKind,
	}, nil
}

func deleteAccessToken(loggedInUser models.User, jwtToken string) error {
	var tokenIndex = -1
	for i, cache := range tokenCache[loggedInUser.ID] {
		if cache.Token == jwtToken {
			tokenIndex = i

			break
		}
	}

	if tokenIndex == -1 {
		return fmt.Errorf("token not found [%s]", jwtToken)
	}

	tokenCache[loggedInUser.ID][tokenIndex] = tokenCache[loggedInUser.ID][len(tokenCache[loggedInUser.ID])-1]
	tokenCache[loggedInUser.ID] = tokenCache[loggedInUser.ID][:len(tokenCache[loggedInUser.ID])-1]
	return nil
}
