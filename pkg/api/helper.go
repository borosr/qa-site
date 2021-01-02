package api

import (
	"encoding/json"
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
