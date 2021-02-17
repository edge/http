package http

import (
	"net/http"
	"regexp"
)

type routeHandler struct {
	method string
	path   *regexp.Regexp

	next http.Handler
}

// Route creates a handler which matches on HTTP method and request URL path.
func Route(method string, path *regexp.Regexp, next http.Handler) Handler {
	h := &routeHandler{
		method: method,
		path:   path,
		next:   next,
	}
	return h
}

func (h *routeHandler) Match(req *http.Request) bool {
	return req.Method == h.method && h.path.Match([]byte(req.URL.Path))
}

func (h *routeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.next.ServeHTTP(w, req)
}
