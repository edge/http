package http

import "net/http"

type switchHandler struct {
	handlers []Handler
}

// Switch creates a switching http.Handler with any number of sub-handlers.
// When it serves an HTTP request, it will find the first sub-handler that matches and pass the request to that.
// If no sub-handler is matched, HTTP 400 Bad Request is written.
func Switch(handlers ...Handler) Handler {
	return &switchHandler{handlers}
}

func (h *switchHandler) Find(req *http.Request) (Handler, bool) {
	for _, handler := range h.handlers {
		if handler.Match(req) {
			return handler, true
		}
	}
	return nil, false
}

func (h *switchHandler) Match(req *http.Request) bool {
	_, ok := h.Find(req)
	return ok
}

func (h *switchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := h.Find(req); ok {
		handler.ServeHTTP(w, req)
	} else {
		ErrBadRequest.ServeHTTP(w, req)
	}
}
