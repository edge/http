package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

type logHandler struct {
	log  zerolog.Logger
	next http.Handler
}

type logWriter struct {
	log        zerolog.Logger
	req        *http.Request
	statusCode int
	w          http.ResponseWriter
}

// Log creates a logging Handler that logs each HTTP connection.
// It passes the request directly to the next handler without modifying it.
func Log(log zerolog.Logger, next http.Handler) Handler {
	l := &logHandler{
		log:  log,
		next: next,
	}
	return l
}

func (l *logHandler) Match(req *http.Request) bool {
	return true
}

func (l *logHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lw := &logWriter{
		log:        l.log,
		req:        req,
		statusCode: http.StatusOK,
		w:          w,
	}
	l.next.ServeHTTP(lw, req)
}

func (lw *logWriter) Header() http.Header {
	return lw.w.Header()
}

func (lw *logWriter) Write(b []byte) (int, error) {
	if lw.statusCode < http.StatusInternalServerError {
		lw.log.Info().
			Str("method", lw.req.Method).
			Str("path", lw.req.URL.Path).
			Int("size", len(b)).
			Int("status", lw.statusCode).
			Msg("ok")
	} else {
		lw.log.Error().
			Str("method", lw.req.Method).
			Str("path", lw.req.URL.Path).
			Int("size", len(b)).
			Int("status", lw.statusCode).
			Msg(string(b))
	}
	return lw.w.Write(b)
}

func (lw *logWriter) WriteHeader(status int) {
	lw.statusCode = status
	lw.w.WriteHeader(status)
}
