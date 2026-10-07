package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/davidporos92/margin-cms/api/internal/api"
)

func NewServerOptions() api.StrictHTTPServerOptions {
	return api.StrictHTTPServerOptions{
		ResponseErrorHandlerFunc: responseErrorHandler,
	}
}

func responseErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, ErrNotImplemented) {
		status = http.StatusNotImplemented
	} else {
		log.Printf("%s %s %s", r.Method, r.URL.Path, err)
	}

	writeProblem(w, status)
}

func writeProblem(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(api.Problem{
		Type:   "about:blank",
		Title:  http.StatusText(status),
		Status: status,
	})
}
