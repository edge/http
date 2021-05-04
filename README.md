# Edge HTTP

This Go library provides a collection of HTTP handlers and middleware that can be combined to form a complex set of request handling rules in a highly expressive, compact way.

## Example

A very simple HTTP router with an HTTP 404 fallback might look like this:

```go
package main

import (
	"net/http"
	"regexp"

	edgehttp "github.com/edge/http"
)

var howdyRegexp = regexp.MustCompile(`^/howdy/?$`)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Howdy!"))
}

func main() {
	handler := edgehttp.Switch(
		edgehttp.Route(http.MethodGet, howdyRegexp, &myHandler{}),
		edgehttp.ErrNotFound,
	)
	http.ListenAndServe("0.0.0.0:8080", handler)
}
```

As long as your handler implements [http.Handler](https://golang.org/pkg/net/http/#Handler), you can have it serve a page within Edge HTTP. (If you want to write middleware, you must implement the [Handler](./http.go) superset interface.)

## Further Reading

Find out more with `go doc` - this package is fully documented.
