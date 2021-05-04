package http

import (
	"net/http"
	"regexp"
)

// PathHandler matches HTTP requests by their request URL path.
type PathHandler struct {
	Next http.Handler
	Path *regexp.Regexp
}

// Path handler.
func Path(path *regexp.Regexp, next http.Handler) PathHandler {
	return PathHandler{
		Next: next,
		Path: path,
	}
}

// Match request.
func (h PathHandler) Match(req *http.Request) bool {
	return h.Path.Match([]byte(req.URL.Path))
}

func (h PathHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}
