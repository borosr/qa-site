package healthcheck

import (
	"github.com/borosr/qa-site/pkg/settings"
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
	"github.com/borosr/qa-site/pkg/auth/oauth"
)

const (
	pending = "pending"
	failed  = "failed"
	ok      = "ok"
)

type Controller struct{}

func NewController() Controller {
	return Controller{}
}

type State struct {
	Status string `json:"status"`
}

var state *State

func Get() *State {
	if state == nil {
		state = &State{
			Status: pending,
		}
	}

	return state
}

func (s *State) Failed() {
	s.Status = failed
}

func (s *State) Ok() {
	s.Status = ok
}

func (s *State) Healthy() bool {
	return s.Status != failed
}

type InfoResponse struct {
	Visibility string
	Providers map[string]bool
}


func Info() InfoResponse {
	var info InfoResponse
	info.Visibility = settings.Get().Visibility
  	info.Providers = make(map[string]bool)
	for key, value := range oauth.Availability() {
		info.Providers[string(key)] = value
	}
	return info
}

func (c Controller) Route(w http.ResponseWriter, r *http.Request) {
	api.SuccessResponse(w, Get())
}

func (c Controller) Info(w http.ResponseWriter, r *http.Request) {
	api.SuccessResponse(w, Info())
}
