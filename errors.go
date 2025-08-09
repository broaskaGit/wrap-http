package wrap

import (
	"errors"
	"net/http"
)

var (
	ErrConn = errors.New("connection") // Represents an error when a connection cannot be established or maintained.

	ErrStatus = errors.New("status") // Represents an error when a failed status code is encountered.

	ErrUnknownCode = errors.New("unknown status code")

	ErrUnmarshal = errors.New("unmarshal")

	ErrInvalidProxy = errors.New("invalid proxy")

	ErrInvalidProxyURL = errors.New("invalid proxy url")

	ErrInvalidProxyCheckURL = errors.New("invalid proxy check url")
)

var (
	ErrInvalidHeadersType = errors.New("invalid headers type")
	ErrInvalidCookiesType = errors.New("invalid cookies type")
)

var (
	// 4xx Client Errors
	ErrBadRequest                   = errors.New("bad request")
	ErrUnauthorized                 = errors.New("unauthorized")
	ErrPaymentRequired              = errors.New("payment required")
	ErrForbidden                    = errors.New("forbidden")
	ErrNotFound                     = errors.New("not found")
	ErrMethodNotAllowed             = errors.New("method not allowed")
	ErrNotAcceptable                = errors.New("not acceptable")
	ErrProxyAuthRequired            = errors.New("proxy authentication required")
	ErrRequestTimeout               = errors.New("request timeout")
	ErrConflict                     = errors.New("conflict")
	ErrGone                         = errors.New("gone")
	ErrLengthRequired               = errors.New("length required")
	ErrPreconditionFailed           = errors.New("precondition failed")
	ErrRequestEntityTooLarge        = errors.New("request entity too large")
	ErrRequestURITooLong            = errors.New("request uri too long")
	ErrUnsupportedMediaType         = errors.New("unsupported media type")
	ErrRequestedRangeNotSatisfiable = errors.New("requested range not satisfiable")
	ErrExpectationFailed            = errors.New("expectation failed")
	ErrTeapot                       = errors.New("i'm a teapot")
	ErrMisdirectedRequest           = errors.New("misdirected request")
	ErrUnprocessableEntity          = errors.New("unprocessable entity")
	ErrLocked                       = errors.New("locked")
	ErrFailedDependency             = errors.New("failed dependency")
	ErrTooEarly                     = errors.New("too early")
	ErrUpgradeRequired              = errors.New("upgrade required")
	ErrPreconditionRequired         = errors.New("precondition required")
	ErrTooManyRequests              = errors.New("too many requests")
	ErrRequestHeaderFieldsTooLarge  = errors.New("request header fields too large")
	ErrUnavailableForLegalReasons   = errors.New("unavailable for legal reasons")

	// 5xx Server Errors
	ErrInternalServerError           = errors.New("internal server error")
	ErrNotImplemented                = errors.New("not implemented")
	ErrBadGateway                    = errors.New("bad gateway")
	ErrServiceUnavailable            = errors.New("service unavailable")
	ErrGatewayTimeout                = errors.New("gateway timeout")
	ErrHTTPVersionNotSupported       = errors.New("http version not supported")
	ErrVariantAlsoNegotiates         = errors.New("variant also negotiates")
	ErrInsufficientStorage           = errors.New("insufficient storage")
	ErrLoopDetected                  = errors.New("loop detected")
	ErrNotExtended                   = errors.New("not extended")
	ErrNetworkAuthenticationRequired = errors.New("network authentication required")
)

// httpstatusErrors maps HTTP status codes to their corresponding errors
var httpstatusErrors = map[int]error{
	// 4xx Client Errors
	http.StatusBadRequest:                   ErrBadRequest,
	http.StatusUnauthorized:                 ErrUnauthorized,
	http.StatusPaymentRequired:              ErrPaymentRequired,
	http.StatusForbidden:                    ErrForbidden,
	http.StatusNotFound:                     ErrNotFound,
	http.StatusMethodNotAllowed:             ErrMethodNotAllowed,
	http.StatusNotAcceptable:                ErrNotAcceptable,
	http.StatusProxyAuthRequired:            ErrProxyAuthRequired,
	http.StatusRequestTimeout:               ErrRequestTimeout,
	http.StatusConflict:                     ErrConflict,
	http.StatusGone:                         ErrGone,
	http.StatusLengthRequired:               ErrLengthRequired,
	http.StatusPreconditionFailed:           ErrPreconditionFailed,
	http.StatusRequestEntityTooLarge:        ErrRequestEntityTooLarge,
	http.StatusRequestURITooLong:            ErrRequestURITooLong,
	http.StatusUnsupportedMediaType:         ErrUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable: ErrRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:            ErrExpectationFailed,
	http.StatusTeapot:                       ErrTeapot,
	http.StatusMisdirectedRequest:           ErrMisdirectedRequest,
	http.StatusUnprocessableEntity:          ErrUnprocessableEntity,
	http.StatusLocked:                       ErrLocked,
	http.StatusFailedDependency:             ErrFailedDependency,
	http.StatusTooEarly:                     ErrTooEarly,
	http.StatusUpgradeRequired:              ErrUpgradeRequired,
	http.StatusPreconditionRequired:         ErrPreconditionRequired,
	http.StatusTooManyRequests:              ErrTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge:  ErrRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:   ErrUnavailableForLegalReasons,

	// 5xx Server Errors

	http.StatusInternalServerError:           ErrInternalServerError,
	http.StatusNotImplemented:                ErrNotImplemented,
	http.StatusBadGateway:                    ErrBadGateway,
	http.StatusServiceUnavailable:            ErrServiceUnavailable,
	http.StatusGatewayTimeout:                ErrGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrInsufficientStorage,
	http.StatusLoopDetected:                  ErrLoopDetected,
	http.StatusNotExtended:                   ErrNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrNetworkAuthenticationRequired,
}

func CodeToErr(code int) error {
	if err, ok := httpstatusErrors[code]; ok {
		return err
	}
	return ErrUnknownCode
}
