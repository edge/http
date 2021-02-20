package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

type logHandler struct {
	level zerolog.Level
	log   zerolog.Logger
	next  http.Handler
}

type logWriter struct {
	level      zerolog.Level
	log        zerolog.Logger
	req        *http.Request
	statusCode int
	w          http.ResponseWriter
}

// Log creates a handler middleware that logs each HTTP connection passing through it.
func Log(log zerolog.Logger, level zerolog.Level, next http.Handler) Handler {
	h := &logHandler{
		level: level,
		log:   log,
		next:  next,
	}
	return Always(h)
}

func (h *logHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lw := &logWriter{
		level:      h.level,
		log:        h.log,
		req:        req,
		statusCode: http.StatusOK,
		w:          w,
	}
	h.next.ServeHTTP(lw, req)
}

func (lw *logWriter) Header() http.Header {
	return lw.w.Header()
}

func (lw *logWriter) Write(b []byte) (int, error) {
	var evt *zerolog.Event
	if lw.isError() {
		evt = lw.log.Error()
	} else {
		evt = lw.log.WithLevel(lw.level)
	}
	size := len(b)
	evt.Str("method", lw.req.Method).
		Str("path", lw.req.URL.Path).
		Int("status", lw.statusCode).
		Msgf("(%dB)", size)

	return lw.w.Write(b)
}

func (lw *logWriter) WriteHeader(status int) {
	lw.statusCode = status
	lw.w.WriteHeader(status)
}

func (lw *logWriter) isError() bool {
	return lw.statusCode >= http.StatusInternalServerError
}
