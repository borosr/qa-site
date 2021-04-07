package healthcheck

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
)

type Controller struct {

}

func NewController() Controller {
	return Controller{}
}

func (c Controller) Route(w http.ResponseWriter, r *http.Request) {
	// TODO check the state in the future
	api.SuccessResponse(w, struct{ Status string }{Status: "ok"})
}
