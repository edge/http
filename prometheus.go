package http

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusHandler counts each HTTP request passing through it.
type PrometheusHandler struct {
	RequestCounter *prometheus.CounterVec
	Next           http.Handler
}

// Prometheus handler.
func Prometheus(namePrefix string, next Handler) PrometheusHandler {
	rcName := "http_request"
	if len(namePrefix) > 0 {
		rcName = fmt.Sprintf("%s_%s", namePrefix, rcName)
	}
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: rcName,
			Help: "Number of HTTP requests handled during current process.",
		},
		[]string{"method", "path", "result"},
	)

	return PrometheusHandler{
		RequestCounter: requestCounter,
		Next:           next,
	}
}

// Match request.
func (h PrometheusHandler) Match(*http.Request) bool {
	return true
}

func (h PrometheusHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pw := &PrometheusWriter{
		h:          &h,
		w:          w,
		method:     req.Method,
		path:       req.URL.Path,
		statusCode: http.StatusOK,
	}
	h.Next.ServeHTTP(pw, req)
}

// PrometheusWriter wraps http.ResponseWriter to count a request as the response is served.
//
// PrometheusHandler uses this type internally.
// This type is exported for documentary reasons, and should not normally be used directly.
type PrometheusWriter struct {
	h *PrometheusHandler
	w http.ResponseWriter

	method     string
	path       string
	statusCode int
}

// Header implements http.ResponseWriter.
func (pw *PrometheusWriter) Header() http.Header {
	return pw.w.Header()
}

func (pw *PrometheusWriter) Write(b []byte) (int, error) {
	pw.h.RequestCounter.WithLabelValues(pw.method, pw.path, fmt.Sprint(pw.statusCode)).Inc()
	return pw.w.Write(b)
}

// WriteHeader implements http.ResponseWriter.
func (pw *PrometheusWriter) WriteHeader(statusCode int) {
	pw.statusCode = statusCode
	pw.w.WriteHeader(statusCode)
}
