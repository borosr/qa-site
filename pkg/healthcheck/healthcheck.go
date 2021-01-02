package healthcheck

import (
	"net/http"

	"github.com/borosr/qa-site/pkg/api"
)

func Route(w http.ResponseWriter, r *http.Request) {
	// TODO check the state in the future
	api.SuccessResponse(w, struct{ Status string }{Status: "ok"})
}
