package http

import (
	"crypto/subtle"
	"net/http"
	"strings"
)

// BearerAuthHandler matches HTTP requests by a bearer token in their Authorization header.
type BearerAuthHandler struct {
	Next http.Handler

	// Token can be empty, in which case all requests are implicitly authorized and not checked.
	Token string
}

// BearerAuth handler.
func BearerAuth(token string, next http.Handler) BearerAuthHandler {
	return BearerAuthHandler{next, token}
}

// Authorize request.
func (h BearerAuthHandler) Authorize(req *http.Request) bool {
	if h.Token == "" {
		return true
	}
	header := req.Header.Get("Authorization")
	splitToken := strings.SplitN(header, " ", 2)
	if len(splitToken) != 2 {
		return false
	}
	reqToken := strings.TrimSpace(splitToken[1])
	if subtle.ConstantTimeCompare([]byte(h.Token), []byte(reqToken)) != 1 {
		return false
	}
	return true
}

// Match request.
func (h BearerAuthHandler) Match(req *http.Request) bool {
	return h.Authorize(req)
}

func (h BearerAuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.Next.ServeHTTP(w, req)
}
