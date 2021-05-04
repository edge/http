package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Simple handler for a generic HTTP error.
var (
	ErrBadRequest          = Error(errors.New("bad request"), http.StatusBadRequest)
	ErrForbidden           = Error(errors.New("forbidden"), http.StatusForbidden)
	ErrInternalServerError = Error(errors.New("internal server error"), http.StatusInternalServerError)
	ErrMethodNotAllowed    = Error(errors.New("method not allowed"), http.StatusMethodNotAllowed)
	ErrNotFound            = Error(errors.New("not found"), http.StatusNotFound)
	ErrUnauthorized        = Error(errors.New("unauthorized"), http.StatusUnauthorized)
)

// ErrorHandler is a simple handler that responds to HTTP requests with a JSON error.
type ErrorHandler struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Error handler.
func Error(err error, code ...int) ErrorHandler {
	var statusCode int
	var message string

	if err != nil {
		message = fmt.Sprintf("%s", err)
	} else {
		message = "internal server error"
	}

	if len(code) > 0 {
		statusCode = code[0]
	} else {
		statusCode = http.StatusInternalServerError
	}

	return ErrorHandler{
		Message:    message,
		StatusCode: statusCode,
	}
}

// Match request.
func (h ErrorHandler) Match(*http.Request) bool {
	return true
}

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(h.StatusCode)
	b, err := json.Marshal(h)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
	w.Write(b)
}
