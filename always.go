package http

import (
	"net/http"
)

// AlwaysHandler wraps any other http.Handler and causes it to always match a request.
type AlwaysHandler struct {
	Next http.Handler
}

// Always handler.
func Always(next http.Handler) Handler {
	return AlwaysHandler{next}
}

// Match request.
func (h AlwaysHandler) Match(req *http.Request) bool {
	return true
}

func (h AlwaysHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}
