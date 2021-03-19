package horror

//go:generate go run cmd/generrors/main.go

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

type Error interface {
	http.Handler
	error
}

func WithError(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.ServeHTTP(w, r)

		switch casted := err.(type) {
		case Error:
			casted.ServeHTTP(w, r)
			return
		default:
			fh := getInternalServerErrorHandler(r.Context())
			fh(err, w, r)
			return
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

type statusError struct {
	code int
	body []byte
}

func (s statusError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.code)
	w.Write(s.body)
}

func (s statusError) Error() string {
	msg := string(s.body)
	return fmt.Sprintf(`horror: code=%d, http error body="%s"`, s.code, msg)
}

func Status(code int, body []byte) Error {
	return statusError{code: code, body: body}
}
