package healthcheck

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
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

var state *State

func Get() *State {
	if state == nil {
		state = &State{
			Status: pending,
		}
	}

	return state
}

type State struct {
	Status string `json:"status"`
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

func (c Controller) Route(w http.ResponseWriter, r *http.Request) {
	api.SuccessResponse(w, Get())
}
