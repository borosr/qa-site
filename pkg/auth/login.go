package auth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/auth/oauth"
	authRepo "github.com/borosr/qa-site/pkg/auth/repository"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/models"
	"github.com/borosr/qa-site/pkg/settings"
	userRepo "github.com/borosr/qa-site/pkg/users/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwtTimeout    = 30 * time.Minute
	revokeTimeout = 72 * time.Hour

	AuthorizationHeader = "Authorization"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type TokenCache struct {
	Token string
	Expr  time.Time
}

type Controller struct {
	userRepository userRepo.UserRepository
	authRepository authRepo.AuthRepository
}

func NewController(userRepository userRepo.UserRepository,
	authRepository authRepo.AuthRepository) Controller {
	return Controller{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (ac Controller) DefaultLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := Request{}
	if err := api.Bind(r, &req); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	var user models.User
	var err error
	if user, err = ac.userRepository.FindByUsername(ctx, req.Username); err != nil {
		log.Errorf("missing username from database: %s, error: %v", req.Username, err)
		api.Forbidden(w)

		return
	}

	if !user.Password.Valid && user.AccessToken.Valid {
		log.Errorf("invalid authentication type for user [%s]", req.Username)
		api.BadRequest(w)

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password)); err != nil {
		log.Errorf("comparing passwords: %v", err)
		api.Forbidden(w)

		return
	}

	resp, err := ac.loggingIn(ctx, user, time.Now(), DefaultAuthKind)
	if err != nil {
		log.Errorf("logging in: %v", err)
		api.Forbidden(w)

		return
	}

	api.SuccessResponse(w, resp)
}

func (ac Controller) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var jwtToken string
			if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
				log.Error("missing token from header")
				api.Forbidden(w)

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
				log.Errorf("error parsing jwt token [%s]: %v", jwtToken, err)
				api.Forbidden(w)

				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx, err = ac.chainTokenBySID(ctx, claims, jwtToken)
				if err != nil {
					log.Errorf("error building jwt token: %v", err)
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

func (ac Controller) SocialMediaRedirect(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "media")

	if err := oauth.Redirect(w, r, provider); err != nil {
		api.InternalServerError(w)
	}
}

func (ac Controller) SocialMediaCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	provider := chi.URLParam(r, "media")

	callbackResp, err := oauth.Callback(w, r, provider)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	}

	user, err := ac.userRepository.FindByUsername(ctx, callbackResp.UserDetails.Username())
	if errors.Is(err, sql.ErrNoRows) {
		user = models.User{
			ID:          xid.New().String(),
			Username:    callbackResp.UserDetails.Username(),
			FullName:    null.StringFrom(callbackResp.UserDetails.FullName()),
			AccessToken: null.StringFrom(callbackResp.Response.AccessToken),
		}
		if err := ac.userRepository.Insert(ctx, user); err != nil {
			api.InternalServerError(w)

			return
		}
	} else if err != nil {
		log.Error(err)
		api.InternalServerError(w)

		return
	} else {
		user.AccessToken.SetValid(callbackResp.Response.AccessToken)
		if _, err := ac.userRepository.Update(ctx, user); err != nil {
			log.Error(err)
			api.InternalServerError(w)

			return
		}
	}

	resp, err := ac.loggingIn(ctx, user, time.Now(), GithubAuthKind)
	if err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	api.SuccessResponse(w, resp)
}

func (ac Controller) Logout(w http.ResponseWriter, r *http.Request) {
	loggedInUser := r.Context().Value("user").(models.User)

	var jwtToken string
	if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
		api.InternalServerError(w)

		return
	}

	if err := ac.authRepository.DeleteJwtToken(loggedInUser.ID, jwtToken); err != nil {
		log.Errorf("invalid access token [%s]", jwtToken)
		api.BadRequest(w)

		return
	}

	api.SuccessResponse(w, struct{ Msg string }{Msg: "Logout success"})
}

func (ac Controller) Revoke(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var revoke = struct {
		Token string `json:"revoke_token"`
	}{}
	if err := api.Bind(r, &revoke); err != nil {
		log.Error(err)
		api.BadRequest(w)

		return
	}

	revokeToken, err := jwt.Parse(revoke.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(settings.Get().RevokeHMAC), nil
	})

	if err != nil {
		log.Error(err)
		api.Forbidden(w)

		return
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = revokeToken.Claims.(jwt.MapClaims); !ok {
		log.Errorf("invalid revoke token, unable to parse jwt claims: [%#v]", revokeToken.Claims)
		api.Forbidden(w)

		return
	}

	userID := claims["sid"].(string)
	revokeTokenModel, err := ac.authRepository.FindRevokeTokenBy(ctx, userID, revoke.Token)
	if err != nil {
		log.Errorf("error getting revoke token from database: %v", err)
		api.Forbidden(w)

		return
	}

	if exp, ok := claims["exp"].(int64); ok && time.Now().Unix() > exp {
		log.Errorf("revoke token [%s] expired %d", revoke.Token, exp)
		api.Forbidden(w)

		return
	}
	var jwtToken string
	if jwtToken = r.Header.Get(AuthorizationHeader); jwtToken == "" {
		log.Error("missing jwt token from header")
		api.Forbidden(w)

		return
	}

	user, err := models.FindUser(ctx, db.Get(), userID)
	if err != nil {
		log.Errorf("error getting owner of the token: %v", err)
		api.Forbidden(w)

		return
	}

	resp, err := ac.renewTokens(ctx, *user, revokeTokenModel, jwtToken)
	if err != nil {
		log.Errorf("error renew the tokens token: %v", err)
		api.Forbidden(w)

		return
	}
	api.SuccessResponse(w, resp)

}

func (ac Controller) generateAuthToken(user models.User, now time.Time) (string, error) {
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

func (ac Controller) generateRevokeToken(user models.User, now time.Time) (string, error) {
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

func (ac Controller) chainTokenBySID(ctx context.Context, claims jwt.MapClaims, jwtToken string) (context.Context, error) {
	if sid, ok := claims["sid"].(string); ok {
		return ac.chainTokenCacheCheck(ctx, sid, jwtToken)
	}
	log.Errorf("missing sid from claim [%+v]", claims)

	return ctx, ErrForbidden
}

func (ac Controller) chainTokenCacheCheck(ctx context.Context, sid string, jwtToken string) (context.Context, error) {
	if err := ac.authRepository.ExistsAndNotExpired(sid, jwtToken, time.Now()); err != nil {
		log.Error(err)

		return ctx, err
	}

	var found bool
	ctx, found = ac.chainGetUserIfFound(ctx, sid)
	if !found {
		return ctx, ErrForbidden
	}

	return ctx, nil
}

func (ac Controller) chainGetUserIfFound(ctx context.Context, sid string) (context.Context, bool) {
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

func (ac Controller) loggingIn(ctx context.Context, user models.User, now time.Time, tokenKind Kind) (Response, error) {
	token, err := ac.generateAuthToken(user, now)
	if err != nil {
		return Response{}, err
	}

	if err := ac.authRepository.StoreJwtToken(user.ID, token, now.Add(jwtTimeout)); err != nil {
		log.Error(err)

		return Response{}, err
	}

	var revokeToken string
	if revokeToken = ac.authRepository.GetRevokeToken(ctx, user.ID); revokeToken == "" {
		var err error
		revokeToken, err = ac.generateRevokeToken(user, now)
		if err != nil {
			return Response{}, err
		}
		if err := ac.authRepository.StoreRevokeToken(ctx, user.ID, revokeToken); err != nil {
			return Response{}, err
		}
	}

	return Response{
		Token:       token,
		RevokeToken: revokeToken,
		AuthKind:    tokenKind,
	}, nil
}

func (ac Controller) renewTokens(ctx context.Context, user models.User, model *models.RevokeToken, token string) (Response, error) {
	if err := ac.authRepository.DeleteJwtToken(user.ID, token); err != nil {
		log.Error(err)

		return Response{}, err
	}

	now := time.Now()
	authToken, err := ac.generateAuthToken(user, now)
	if err != nil {
		return Response{}, err
	}

	if err := ac.authRepository.StoreJwtToken(user.ID, authToken, now.Add(jwtTimeout)); err != nil {
		log.Error(err)

		return Response{}, err
	}

	revokeToken, err := ac.generateRevokeToken(user, now)
	if err != nil {
		return Response{}, err
	}
	model.Token = revokeToken
	if _, err := ac.authRepository.Update(ctx, model); err != nil {
		return Response{}, err
	}

	return Response{
		Token:       authToken,
		RevokeToken: revokeToken,
		AuthKind:    DefaultAuthKind,
	}, nil
}
