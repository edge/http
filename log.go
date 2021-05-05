package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

var (
	okStatusCode = []int{
		http.StatusOK,
		http.StatusPermanentRedirect,
		http.StatusTemporaryRedirect,
	}
)

// LogHandler logs each HTTP connection passing through it.
type LogHandler struct {
	Level zerolog.Level
	Log   zerolog.Logger
	Next  http.Handler
}

// Log handler.
func Log(log zerolog.Logger, level zerolog.Level, next http.Handler) LogHandler {
	return LogHandler{
		Level: level,
		Log:   log,
		Next:  next,
	}
}

// Match request.
func (h LogHandler) Match(*http.Request) bool {
	return true
}

func (h LogHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lw := &LogWriter{
		Level:      h.Level,
		Log:        h.Log,
		Req:        req,
		StatusCode: http.StatusOK,
		W:          w,
	}
	h.Next.ServeHTTP(lw, req)
}

// LogWriter wraps http.ResponseWriter to log a request as the response is served.
//
// LogHandler uses this type internally to enable per-request logging.
// This type is exported for documentary reasons, and should not normally be used directly.
type LogWriter struct {
	Level      zerolog.Level
	Log        zerolog.Logger
	Req        *http.Request
	StatusCode int
	W          http.ResponseWriter
}

// Header returns the header map that will be returned by WriteHeader.
// See http.ResponseWriter.
func (lw *LogWriter) Header() http.Header {
	return lw.W.Header()
}

func (lw *LogWriter) Write(b []byte) (int, error) {
	size, err := lw.W.Write(b)
	var evt *zerolog.Event
	if lw.isError() {
		evt = lw.Log.Error()
	} else {
		evt = lw.Log.WithLevel(lw.Level)
	}
	evt.Str("method", lw.Req.Method).
		Str("path", lw.Req.URL.Path).
		Int("status", lw.StatusCode).
		Msgf("%dB", size)
	return size, err
}

// WriteHeader sends an HTTP response header with the provided status code.
// See http.ResponseWriter.
func (lw *LogWriter) WriteHeader(status int) {
	lw.StatusCode = status
	lw.W.WriteHeader(status)
}

func (lw *LogWriter) isError() bool {
	for _, okcode := range okStatusCode {
		if lw.StatusCode == okcode {
			return false
		}
	}
	return true
}
