package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Simple generic implementation of http.Handler for an error case.
var (
	ErrBadRequest          = Error(errors.New("bad request"), http.StatusBadRequest)
	ErrForbidden           = Error(errors.New("forbidden"), http.StatusForbidden)
	ErrInternalServerError = Error(errors.New("internal server error"), http.StatusInternalServerError)
	ErrMethodNotAllowed    = Error(errors.New("method not allowed"), http.StatusMethodNotAllowed)
	ErrNotFound            = Error(errors.New("not found"), http.StatusNotFound)
	ErrUnauthorized        = Error(errors.New("unauthorized"), http.StatusUnauthorized)
)

type errorHandler struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Error creates a Handler that simply outputs an error message.
func Error(err error, code ...int) Handler {
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

	return &errorHandler{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (h *errorHandler) Match(req *http.Request) bool {
	return true
}

func (h *errorHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(h.StatusCode)
	b, err := json.Marshal(h)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
	w.Write(b)
}
