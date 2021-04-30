/*
Package horror implements interfaces for http error handling.

Horror was created to simplify error handling with standard go http.Handlers
and give developers tools to encapsulate common error scenarios.

Thanks to WithError function and Adapter type you can use horror.Handlers
with existing go code based on http.Handler interface.

If you're interested how to use this module, please see examples that are
placed in `examples` directory.

See "github.com/thinkofher/horror/status" for convenient http errors that
satisfy Error interface from this module.
*/
package horror

//go:generate go run cmd/generrors/main.go
//go:generate go fmt ./...

import (
	"errors"
	"net/http"
)

// Handler is an alternative to http.Handler. It expands ServerHTTP method by
// adding error as return value.
//
// You can propagate error value further and eventually wrap it with
// fmt.Errorf when implementing ServeHTTP method. The logic for handling
// specific erorrs can be encapsulated in the Error interface.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

// Error interface is the resultant of error and http.Handler interfaces.
// You can use it to encapsulate error handling logic within error itself.
//
// See github.com/thinkofher/horror/status module for standard http status
// implementations of Error interface.
type Error interface {
	http.Handler
	error
}

// WithError wraps Handler, adapts it and returns http.Handler that
// can be used with other APIs that relies on go standard library.
func WithError(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h.ServeHTTP(w, r); err != nil {
			var horrorError Error
			if errors.As(err, &horrorError) {
				horrorError.ServeHTTP(w, r)
				return
			} else {
				fh := getInternalServerErrorHandler(r.Context())
				fh(err, w, r)
				return
			}
		}
	})
}

// HandlerFunc is an adapter that will use given f as ServeHTTP method for
// returned Handler.
func HandlerFunc(f func(w http.ResponseWriter, r *http.Request) error) Handler {
	return handlerFunc(f)
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (h handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return h(w, r)
}
