package http

import "net/http"

// Handler extends http.Handler and can Match() an HTTP request based on arbitrary criteria.
//
// For example, Route().Match tests the HTTP request's method and URL path against its configuration to determine whether to serve it.
// Conversely, Log().Match always returns true and will always serve the request, effectively acting as middleware.
//
// All handlers in this package implement the extended Handler interface.
// In turn, they may require other Handler[s] or just http.Handler[s] depending on internal requirements.
// Most functions, such as Log() and Route(), only require an http.Handler to pass the request along to.
// However, Switch() requires other Handler[s] as it depends on each handler implementing Handler.Match.
type Handler interface {
	http.Handler
	Match(req *http.Request) bool
}
