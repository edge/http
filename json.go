package http

import "net/http"

// JSONHandler enforces Content-Type: application/json for requests and adds it to all responses.
type JSONHandler struct {
	Next http.Handler
}

// JSON handler.
func JSON(next http.Handler) JSONHandler {
	return JSONHandler{
		Next: next,
	}
}

// Match request.
func (h JSONHandler) Match(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	h.Next.ServeHTTP(w, req)
}
