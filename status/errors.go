package status

// Code generated by go generate; DO NOT EDIT.
// This file was generated at:
// 2021-04-30 13:16:44.056500255 &#43;0200 CEST m=&#43;0.000794951

import "net/http"

// BadRequest return Error with given body and 400 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func BadRequest(body []byte) *Status {
	return New(http.StatusBadRequest, body)
}

// Unauthorized return Error with given body and 401 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Unauthorized(body []byte) *Status {
	return New(http.StatusUnauthorized, body)
}

// PaymentRequired return Error with given body and 402 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func PaymentRequired(body []byte) *Status {
	return New(http.StatusPaymentRequired, body)
}

// Forbidden return Error with given body and 403 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Forbidden(body []byte) *Status {
	return New(http.StatusForbidden, body)
}

// NotFound return Error with given body and 404 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func NotFound(body []byte) *Status {
	return New(http.StatusNotFound, body)
}

// MethodNotAllowed return Error with given body and 405 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func MethodNotAllowed(body []byte) *Status {
	return New(http.StatusMethodNotAllowed, body)
}

// NotAcceptable return Error with given body and 406 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func NotAcceptable(body []byte) *Status {
	return New(http.StatusNotAcceptable, body)
}

// ProxyAuthRequired return Error with given body and 407 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func ProxyAuthRequired(body []byte) *Status {
	return New(http.StatusProxyAuthRequired, body)
}

// RequestTimeout return Error with given body and 408 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func RequestTimeout(body []byte) *Status {
	return New(http.StatusRequestTimeout, body)
}

// Conflict return Error with given body and 409 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Conflict(body []byte) *Status {
	return New(http.StatusConflict, body)
}

// Gone return Error with given body and 410 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Gone(body []byte) *Status {
	return New(http.StatusGone, body)
}

// LengthRequired return Error with given body and 411 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func LengthRequired(body []byte) *Status {
	return New(http.StatusLengthRequired, body)
}

// PreconditionFailed return Error with given body and 412 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func PreconditionFailed(body []byte) *Status {
	return New(http.StatusPreconditionFailed, body)
}

// RequestEntityTooLarge return Error with given body and 413 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func RequestEntityTooLarge(body []byte) *Status {
	return New(http.StatusRequestEntityTooLarge, body)
}

// RequestURITooLong return Error with given body and 414 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func RequestURITooLong(body []byte) *Status {
	return New(http.StatusRequestURITooLong, body)
}

// UnsupportedMediaType return Error with given body and 415 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func UnsupportedMediaType(body []byte) *Status {
	return New(http.StatusUnsupportedMediaType, body)
}

// RequestedRangeNotSatisfiable return Error with given body and 416 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func RequestedRangeNotSatisfiable(body []byte) *Status {
	return New(http.StatusRequestedRangeNotSatisfiable, body)
}

// ExpectationFailed return Error with given body and 417 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func ExpectationFailed(body []byte) *Status {
	return New(http.StatusExpectationFailed, body)
}

// Teapot return Error with given body and 418 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Teapot(body []byte) *Status {
	return New(http.StatusTeapot, body)
}

// MisdirectedRequest return Error with given body and 421 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func MisdirectedRequest(body []byte) *Status {
	return New(http.StatusMisdirectedRequest, body)
}

// UnprocessableEntity return Error with given body and 422 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func UnprocessableEntity(body []byte) *Status {
	return New(http.StatusUnprocessableEntity, body)
}

// Locked return Error with given body and 423 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func Locked(body []byte) *Status {
	return New(http.StatusLocked, body)
}

// FailedDependency return Error with given body and 424 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func FailedDependency(body []byte) *Status {
	return New(http.StatusFailedDependency, body)
}

// TooEarly return Error with given body and 425 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func TooEarly(body []byte) *Status {
	return New(http.StatusTooEarly, body)
}

// UpgradeRequired return Error with given body and 426 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func UpgradeRequired(body []byte) *Status {
	return New(http.StatusUpgradeRequired, body)
}

// PreconditionRequired return Error with given body and 428 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func PreconditionRequired(body []byte) *Status {
	return New(http.StatusPreconditionRequired, body)
}

// TooManyRequests return Error with given body and 429 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func TooManyRequests(body []byte) *Status {
	return New(http.StatusTooManyRequests, body)
}

// RequestHeaderFieldsTooLarge return Error with given body and 431 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func RequestHeaderFieldsTooLarge(body []byte) *Status {
	return New(http.StatusRequestHeaderFieldsTooLarge, body)
}

// UnavailableForLegalReasons return Error with given body and 451 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func UnavailableForLegalReasons(body []byte) *Status {
	return New(http.StatusUnavailableForLegalReasons, body)
}

// InternalServerError return Error with given body and 500 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func InternalServerError(body []byte) *Status {
	return New(http.StatusInternalServerError, body)
}

// NotImplemented return Error with given body and 501 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func NotImplemented(body []byte) *Status {
	return New(http.StatusNotImplemented, body)
}

// BadGateway return Error with given body and 502 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func BadGateway(body []byte) *Status {
	return New(http.StatusBadGateway, body)
}

// ServiceUnavailable return Error with given body and 503 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func ServiceUnavailable(body []byte) *Status {
	return New(http.StatusServiceUnavailable, body)
}

// GatewayTimeout return Error with given body and 504 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func GatewayTimeout(body []byte) *Status {
	return New(http.StatusGatewayTimeout, body)
}

// HTTPVersionNotSupported return Error with given body and 505 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func HTTPVersionNotSupported(body []byte) *Status {
	return New(http.StatusHTTPVersionNotSupported, body)
}

// VariantAlsoNegotiates return Error with given body and 506 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func VariantAlsoNegotiates(body []byte) *Status {
	return New(http.StatusVariantAlsoNegotiates, body)
}

// InsufficientStorage return Error with given body and 507 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func InsufficientStorage(body []byte) *Status {
	return New(http.StatusInsufficientStorage, body)
}

// LoopDetected return Error with given body and 508 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func LoopDetected(body []byte) *Status {
	return New(http.StatusLoopDetected, body)
}

// NotExtended return Error with given body and 510 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func NotExtended(body []byte) *Status {
	return New(http.StatusNotExtended, body)
}

// NetworkAuthenticationRequired return Error with given body and 511 http status code.
//
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func NetworkAuthenticationRequired(body []byte) *Status {
	return New(http.StatusNetworkAuthenticationRequired, body)
}
