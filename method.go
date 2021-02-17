package http

import (
	"net/http"
)

type methodHandler struct {
	method string
	next   http.Handler
}

// Method creates a handler which matches on HTTP method.
func Method(method string, next http.Handler) Handler {
	h := &methodHandler{
		method: method,
		next:   next,
	}
	return h
}

func (h *methodHandler) Match(req *http.Request) bool {
	return req.Method == h.method
}

func (h *methodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.next.ServeHTTP(w, req)
}
