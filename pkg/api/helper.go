package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	applicationJSON   = "application/json"
	contentTypeHeader = "Content-Type"
)

func SuccessResponse(w http.ResponseWriter, v interface{}) {
	result, err := json.Marshal(v)
	if err != nil {
		InternalServerError(w)

		return
	}
	w.Header().Set(contentTypeHeader, applicationJSON)
	if _, err := w.Write(result); err != nil {
		InternalServerError(w)

		return
	}
	w.WriteHeader(http.StatusOK)
}

func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter) {
	http.Error(w, "bad request", http.StatusBadRequest)
}

func NotFound(w http.ResponseWriter) {
	http.Error(w, "not found", http.StatusNotFound)
}

func Forbidden(w http.ResponseWriter) {
	http.Error(w, "forbidden", http.StatusForbidden)
}

// Bind unmarshall the http.Request's body to v
// v should be a pointer
func Bind(r *http.Request, v interface{}) error {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(rawBody, v); err != nil {
		return err
	}

	return err
}
