package http

import "net/http"

type switchHandler struct {
	handlers []Handler
}

// Switch creates a switching handler with any number of next handlers.
// When it serves an HTTP request, it will find the first next handler that matches and pass the request to that.
// If no next handler is matched, HTTP 500 Internal Server Error is written.
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
		ErrInternalServerError.ServeHTTP(w, req)
	}
}
