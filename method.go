package http

import (
	"net/http"
)

// MethodHandler matches HTTP requests by their method.
type MethodHandler struct {
	Method string
	Next   http.Handler
}

// Method handler.
func Method(method string, next http.Handler) MethodHandler {
	return MethodHandler{
		Method: method,
		Next:   next,
	}
}

// Match request.
func (h *MethodHandler) Match(req *http.Request) bool {
	return req.Method == h.Method
}

func (h *MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}
