package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type code struct {
	Number int
	Name   string
}

type errorsGoData struct {
	Timestamp time.Time
	Codes     []code
}

var data errorsGoData

func init() {
	data = errorsGoData{
		Timestamp: time.Now(),
		Codes: []code{
			{http.StatusBadRequest, "BadRequest"},
			{http.StatusUnauthorized, "Unauthorized"},
			{http.StatusPaymentRequired, "PaymentRequired"},
			{http.StatusForbidden, "Forbidden"},
			{http.StatusNotFound, "NotFound"},
			{http.StatusMethodNotAllowed, "MethodNotAllowed"},
			{http.StatusNotAcceptable, "NotAcceptable"},
			{http.StatusProxyAuthRequired, "ProxyAuthRequired"},
			{http.StatusRequestTimeout, "RequestTimeout"},
			{http.StatusConflict, "Conflict"},
			{http.StatusGone, "Gone"},
			{http.StatusLengthRequired, "LengthRequired"},
			{http.StatusPreconditionFailed, "PreconditionFailed"},
			{http.StatusRequestEntityTooLarge, "RequestEntityTooLarge"},
			{http.StatusRequestURITooLong, "RequestURITooLong"},
			{http.StatusUnsupportedMediaType, "UnsupportedMediaType"},
			{http.StatusRequestedRangeNotSatisfiable, "RequestedRangeNotSatisfiable"},
			{http.StatusExpectationFailed, "ExpectationFailed"},
			{http.StatusTeapot, "Teapot"},
			{http.StatusMisdirectedRequest, "MisdirectedRequest"},
			{http.StatusUnprocessableEntity, "UnprocessableEntity"},
			{http.StatusLocked, "Locked"},
			{http.StatusFailedDependency, "FailedDependency"},
			{http.StatusTooEarly, "TooEarly"},
			{http.StatusUpgradeRequired, "UpgradeRequired"},
			{http.StatusPreconditionRequired, "PreconditionRequired"},
			{http.StatusTooManyRequests, "TooManyRequests"},
			{http.StatusRequestHeaderFieldsTooLarge, "RequestHeaderFieldsTooLarge"},
			{http.StatusUnavailableForLegalReasons, "UnavailableForLegalReasons"},
			{http.StatusInternalServerError, "InternalServerError"},
			{http.StatusNotImplemented, "NotImplemented"},
			{http.StatusBadGateway, "BadGateway"},
			{http.StatusServiceUnavailable, "ServiceUnavailable"},
			{http.StatusGatewayTimeout, "GatewayTimeout"},
			{http.StatusHTTPVersionNotSupported, "HTTPVersionNotSupported"},
			{http.StatusVariantAlsoNegotiates, "VariantAlsoNegotiates"},
			{http.StatusInsufficientStorage, "InsufficientStorage"},
			{http.StatusLoopDetected, "LoopDetected"},
			{http.StatusNotExtended, "NotExtended"},
			{http.StatusNetworkAuthenticationRequired, "NetworkAuthenticationRequired"},
		},
	}

}

func main() {
	errorsGo, err := os.Create("status/errors.go")
	die(err)
	defer errorsGo.Close()
	errorsGoTemplate.Execute(errorsGo, data)

	tests, err := os.Create("horror_test.go")
	defer tests.Close()
	testsGoTemplate.Execute(tests, data)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var errorsGoTemplate = template.Must(
	template.New("").Parse(`package status
// Code generated by go generate; DO NOT EDIT.
// This file was generated at:
// {{ .Timestamp }}

import "net/http"
{{ range $index, $element := .Codes }}
// {{ $element.Name }} return Error with given body and {{ $element.Number }} http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func {{ $element.Name }}(body []byte) *Status {
	return New(http.Status{{ $element.Name }}, body)
}
{{- end}}
`))

var testsGoTemplate = template.Must(
	template.New("").Parse(`package horror
// Code generated by go generate; DO NOT EDIT.
// This file was generated at:
// {{ .Timestamp }}

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thinkofher/horror/status"
)

func statusErrorFunc(body []byte, f func([]byte) *status.Status) Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return f(body)
	})
}

func TestHTTPErrorFuncs(t *testing.T) {
	tests := []struct {
		name     string
		factory  func([]byte) *status.Status
		wantCode int
	}{
{{- range  .Codes }}
		{
			name: "{{ .Name }}",
			factory: status.{{ .Name }},
			wantCode: http.Status{{ .Name }},
		},
{{- end}}
	}

	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("Requesting handler which returns only %s status.", tt.name),
			func(t *testing.T) {
				req, err := http.NewRequest(http.MethodGet, "/", nil)
				if err != nil {
					t.Errorf("got unwanted error: %v", err)
				}

				body := []byte(tt.name + " test")
				mux := http.NewServeMux()
				mux.Handle("/", WithError(statusErrorFunc(body, tt.factory)))
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)

				if tt.wantCode != rr.Code {
					t.Errorf("got: %d, want: %d", rr.Code, tt.wantCode)
				}

				if !bytes.Equal(body, rr.Body.Bytes()) {
					t.Errorf("got: b\"%s\", want: b\"%s\"", body, rr.Body.Bytes())
				}
			},
		)
		t.Run(
			fmt.Sprintf("Requesting handler which returns only %s status with adapter.",
				tt.name),
			func(t *testing.T) {
				req, err := http.NewRequest(http.MethodGet, "/", nil)
				if err != nil {
					t.Errorf("got unwanted error: %v", err)
				}

				adapter := NewAdapter(&AdapterBuilder{})

				body := []byte(tt.name + " test")
				mux := http.NewServeMux()
				mux.Handle("/", adapter.WithError(statusErrorFunc(body, tt.factory)))
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)

				if tt.wantCode != rr.Code {
					t.Errorf("got: %d, want: %d", rr.Code, tt.wantCode)
				}

				if !bytes.Equal(body, rr.Body.Bytes()) {
					t.Errorf("got: b\"%s\", want: b\"%s\"", body, rr.Body.Bytes())
				}
			},
		)
	}
}
`))
