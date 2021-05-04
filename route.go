package http

import (
	"net/http"
	"regexp"
)

// RouteHandler matches HTTP requests by their method and request URL path.
type RouteHandler struct {
	Method string
	Path   *regexp.Regexp

	Next http.Handler
}

// Route handler.
func Route(method string, path *regexp.Regexp, next http.Handler) RouteHandler {
	return RouteHandler{
		Method: method,
		Path:   path,
		Next:   next,
	}
}

// Match request.
func (h RouteHandler) Match(req *http.Request) bool {
	return req.Method == h.Method && h.Path.Match([]byte(req.URL.Path))
}

func (h RouteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}
