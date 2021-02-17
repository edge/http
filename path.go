package http

import (
	"net/http"
	"regexp"
)

type pathHandler struct {
	next http.Handler
	path *regexp.Regexp
}

// Path creates a handler which matches on request URL path.
func Path(path *regexp.Regexp, next http.Handler) Handler {
	h := &pathHandler{
		path: path,
		next: next,
	}
	return h
}

func (h *pathHandler) Match(req *http.Request) bool {
	return h.path.Match([]byte(req.URL.Path))
}

func (h *pathHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.next.ServeHTTP(w, req)
}
