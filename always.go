package http

import (
	"net/http"
)

type alwaysHandler struct {
	next http.Handler
}

// Always handler wraps any other http.Handler and causes it to always Match() true.
// This is useful for wrapping http.Handlers that don't also conform to this package's Handler interface.
func Always(next http.Handler) Handler {
	return &alwaysHandler{next}
}

func (h *alwaysHandler) Match(req *http.Request) bool {
	return true
}

func (h *alwaysHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.next.ServeHTTP(w, req)
}
