package status

import (
	"fmt"
	"net/http"
)

// Status is standard error used by status package. It satisfies horror.Error
// interface and can be used safely with horror.Handler interface
// implementations.
type Status struct {
	code int
	body []byte
}

func (s *Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.code)
	w.Write(s.body)
}

func (s *Status) Error() string {
	return fmt.Sprintf(`horror: code=%d, http error body="%s"`, s.code, s.body)
}

// New returns pointer to Status with given code and body.
func New(code int, body []byte) *Status {
	return &Status{code: code, body: body}
}
