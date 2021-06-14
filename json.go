package http

import "net/http"

// JSONHandler enforces Accept: application/json for requests.
// It also enforces Content-Type: application/json on requests that have a body, e.g. POST.
//
// Note that it does not automatically add a Content-Type header to the response.
// This must be done in user code.
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
	acceptOK := clientAcceptsJSON(req)
	contentOK := req.Body == nil || req.Header.Get("Content-Type") == "application/json"
	return acceptOK && contentOK
}

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}

func clientAcceptsJSON(req *http.Request) bool {
	accept := req.Header.Get("Accept")
	return accept == "application/json" || accept == "*/*"
}
