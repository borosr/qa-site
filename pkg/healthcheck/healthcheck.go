package healthcheck

import (
	"net/http"
	"strings"
	"sync"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/auth/oauth"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/samber/do"
)

const (
	pending = "pending"
	failed  = "failed"
	ok      = "ok"
)

type Controller struct{}

func NewController(_ *do.Injector) Controller {
	return Controller{}
}

var state *State
var stateMutex sync.Mutex

func Instance() *State {
	if state == nil {
		stateMutex.Lock()
		state = &State{
			Status: pending,
		}
		stateMutex.Unlock()
	}

	return state
}

type State struct {
	Status string `json:"status"`
	sync.RWMutex
}

func (s *State) Failed() {
	s.Lock()
	defer s.Unlock()
	s.Status = failed
}

func (s *State) Ok() {
	s.Lock()
	defer s.Unlock()
	s.Status = ok
}

func (s *State) Healthy() bool {
	s.RLock()
	defer s.RUnlock()
	return s.Status != failed
}

type InfoResponse struct {
	Visibility     string          `json:"visibility"`
	OauthProviders map[string]bool `json:"oauth_providers"`
}

func (c Controller) Route(w http.ResponseWriter, r *http.Request) {
	api.SuccessResponse(w, Instance())
}

func (c Controller) Info(w http.ResponseWriter, r *http.Request) {
	api.SuccessResponse(w, InfoResponse{
		Visibility:     settings.Get().Visibility,
		OauthProviders: oauth.Availability(),
	})
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.RequestURI, "/api/status") && !Instance().Healthy() {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}
		next.ServeHTTP(w, r)
	})
}
