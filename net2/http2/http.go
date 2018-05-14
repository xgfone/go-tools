// Package http2 is the supplement of the standard library `http`,
// not the protocal `http2`.
package http2

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/xgfone/go-tools/lifecycle"
	"github.com/xgfone/go-tools/log2"
)

// HTTPError stands for a HTTP error.
type HTTPError struct {
	// The error information
	Err error

	// The status code
	Code int

	// You can place data into it to carry in an error.
	Data map[string]interface{}
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

// ErrorWrapper converts the http handler function with an error return value
// to http.Handler.
//
// Notice: the wrapper only wraps the panic and writes the error into log.
func ErrorWrapper(f func(http.ResponseWriter, *http.Request) error) http.Handler {
	return ErrorWrapperFunc(f)
}

// ErrorWrapperFunc converts the http handler function with an error return
// value to http.HandlerFunc.
//
// Notice: the wrapper only wraps the panic and writes the error into log.
func ErrorWrapperFunc(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if e := recover(); e != nil {
				log2.ErrorF("Get a panic when handling %q: %s", r.RequestURI, e)
			} else if err != nil {
				log2.ErrorF("Failed to handle %q: %s", r.RequestURI, err)
			}
		}()
		err = f(w, r)
	}
}

// ErrorHandler handles the error and responds it the client.
func ErrorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return ErrorHandlerWithStatusCode(func(w http.ResponseWriter,
		r *http.Request) (int, error) {
		return 0, f(w, r)
	})
}

// ErrorHandlerWithStatusCode handles the error and responds it the client
// with the status code.
func ErrorHandlerWithStatusCode(f func(http.ResponseWriter, *http.Request) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if code, err := f(w, r); err != nil {
			if code == 0 {
				if _err, ok := err.(HTTPError); ok {
					code = _err.Code
				} else {
					code = http.StatusInternalServerError
				}
			}
			http.Error(w, err.Error(), code)
			log2.ErrorF("Handling %q: status=%d, err=%v", r.RequestURI, code, err)
		}
	}
}

// HandlerWrapper handles the response result.
func HandlerWrapper(f func(http.ResponseWriter, *http.Request) (int, []byte, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, resp, err := f(w, r)

		// Handle the error.
		if err != nil {
			if code == 0 {
				if _err, ok := err.(HTTPError); ok {
					code = _err.Code
				} else {
					code = http.StatusInternalServerError
				}
			}
			http.Error(w, err.Error(), code)
			log2.ErrorF("Failed to handle %q: %s", r.RequestURI, err)
			return
		}

		// Determine the status code.
		if code == 0 {
			if len(resp) == 0 {
				code = http.StatusNoContent
			} else {
				code = http.StatusOK
			}
		}

		// Send the response result.
		w.WriteHeader(code)
		if _, err = io.CopyN(w, bytes.NewBuffer(resp), int64(len(resp))); err != nil {
			log2.ErrorF("Failed to send the response of %q: %s", r.RequestURI, err)
		}
	}
}

// ListenAndServe is equal to http.ListenAndServe, but calling the method
// server.Shutdown(context.TODO()) to shutdown the HTTP server gracefully
// when calling lifecycle.Stop().
//
// Notice: It will call lifecycle.Stop() when the server exits.
func ListenAndServe(addr string, handler http.Handler) error {
	server := http.Server{Addr: addr, Handler: handler}
	lifecycle.Register(func() { server.Shutdown(context.TODO()) })
	err := server.ListenAndServe()
	lifecycle.Stop()
	return err
}

// ListenAndServeTLS is equal to http.ListenAndServeTLS, but calling the method
// server.Shutdown(context.TODO()) to shutdown the HTTP server gracefully
// when calling lifecycle.Stop().
//
// Notice: It will call lifecycle.Stop() when the server exits.
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	server := http.Server{Addr: addr, Handler: handler}
	lifecycle.Register(func() { server.Shutdown(context.TODO()) })
	err := server.ListenAndServeTLS(certFile, keyFile)
	lifecycle.Stop()
	return err
}
