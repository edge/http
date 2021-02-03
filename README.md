# Maedan HTTP

This Go library provides a collection of HTTP handlers and middleware that can be combined to form a complex set of request handling rules in a highly expressive, compact way.

## Introduction

A very simple HTTP router with an HTTP 404 fallback might look like this:

```go
package main

import (
	"net/http"
	"regexp"

	mdnhttp "github.com/maedan/http"
)

var howdyRegexp = regexp.MustCompile(`^/howdy/?$`)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Howdy!"))
}

func main() {
	handler := mdnhttp.Switch(
		mdnhttp.Route(http.MethodGet, howdyRegexp, &myHandler{}),
		mdnhttp.ErrNotFound,
	)
	http.ListenAndServe("0.0.0.0:8080", handler)
}
```

As long as your handler implements [http.Handler](https://golang.org/pkg/net/http/#Handler), you can have it serve a page within Maedan HTTP. (If you want to write middleware, you must implement the [Handler](./http.go) superset interface.)

## All Handlers

Except for predefined handlers serving standard HTTP error responses, which you can find in a perfectly readable list in [error.go](./error.go).

| Type | Function | Description |
|:-----|:---------|:------------|
| Page | [Error()](./error.go) | Return an error response |
| Middleware | [Log()](./log.go) | Log connections |
| Middleware | [Route()](./route.go) | Allows request if method and URL path match |
| Middleware | [Switch()](./switch.go) | Passes request to first matching handler in a list |
