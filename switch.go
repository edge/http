package http

import (
	"net/http"
)

// SwitchHandler matches HTTP requests arbitrarily, based on any number of 'next' handlers.
// When serving an HTTP request, it tests each 'next' handler in sequence, and resolves through the first that matches the request.
// If no handler matches, the request is not handled.
type SwitchHandler struct {
	Next []Handler
}

// Switch handler.
func Switch(next ...Handler) SwitchHandler {
	return SwitchHandler{next}
}

// Find the first matching 'next' handler.
func (h SwitchHandler) Find(req *http.Request) (Handler, bool) {
	for _, handler := range h.Next {
		if handler.Match(req) {
			return handler, true
		}
	}
	return nil, false
}

// Match request.
func (h SwitchHandler) Match(req *http.Request) bool {
	_, ok := h.Find(req)
	return ok
}

func (h SwitchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := h.Find(req); ok {
		handler.ServeHTTP(w, req)
	}
}
