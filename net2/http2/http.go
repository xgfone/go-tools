// Package http2 is the supplement of the standard library `http`,
// not the protocal `http2`.
package http2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xgfone/go-tools/lifecycle"
)

// HTTPError stands for a HTTP error.
type HTTPError struct {
	// The error information
	Err error

	// The status code
	Code int
}

// NewHTTPError returns a new HTTPError.
func NewHTTPError(code int, err interface{}) HTTPError {
	switch err.(type) {
	case error:
	case []byte:
		err = fmt.Errorf("%s", string(err.([]byte)))
	default:
		err = fmt.Errorf("%v", err)
	}
	return HTTPError{Code: code, Err: err.(error)}
}

func (e HTTPError) Error() string {
	return e.Err.Error()
}

// ListenAndServe is equal to http.ListenAndServe, but calling the method
// server.Shutdown(context.TODO()) to shutdown the HTTP server gracefully
// when calling lifecycle.Stop().
//
// If tls exists, it's the CA certificates, that's, certFile and keyFile.
// It will be similar to http.ListenAndServeTLS(addr, tls[0], tls[1], handler).
func ListenAndServe(addr string, handler http.Handler, tls ...string) error {
	server := http.Server{Addr: addr, Handler: handler}
	lifecycle.Register(func() { server.Shutdown(context.TODO()) })
	switch len(tls) {
	case 0:
		return server.ListenAndServe()
	case 2:
		if tls[0] == "" && tls[1] == "" {
			return server.ListenAndServe()
		} else if tls[0] == "" || tls[1] == "" {
			return fmt.Errorf("invalid CA certificate")
		}
		return server.ListenAndServeTLS(tls[0], tls[1])
	default:
		return fmt.Errorf("invalid CA certificate")
	}
}
