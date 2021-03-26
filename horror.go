package horror

//go:generate go run cmd/generrors/main.go
//go:generate go fmt ./...

import (
	"errors"
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

type Error interface {
	http.Handler
	error
}

// WithError wraps horror.Handler, adapts it and returns http.Handler that
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

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (h handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return h(w, r)
}

func HandlerFunc(f func(w http.ResponseWriter, r *http.Request) error) Handler {
	return handlerFunc(f)
}
