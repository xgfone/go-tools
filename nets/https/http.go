// Package https is the supplement of the standard library `http`,
// not the protocal `https`.
package https

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/xgfone/go-tools/lifecycle"
	"github.com/xgfone/go-tools/log"
)

var (
	// DefaultHealthURL is the url to register the health HTTP handler.
	DefaultHealthURL = "/health"

	// DefaultServeMux is the default serve multiplexer.
	DefaultServeMux = http.DefaultServeMux
)

// SetHealthHandler registers the health handler with DefaultHealthURL into
// DefaultServeMux.
//
// Notice: The handler may be nil, which will use the default handler that only
// returns the status code 200.
func SetHealthHandler(handler interface{}) {
	if handler == nil {
		handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	} else if _, ok := handler.(http.Handler); !ok {
		panic("The handler is not nil or http.Handler")
	}
	DefaultServeMux.Handle(DefaultHealthURL, handler.(http.Handler))
}

// SetHealthHandlerFunc is same as SetHealthHandler, but the handler is
// a function that returns the body content.
//
// Notice: If the handler function panic,it will return the status code 500.
func SetHealthHandlerFunc(handler func() string) {
	SetHealthHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		io.WriteString(w, handler())
	}))
}

// StartHealthServer starts the HTTP server to serve the health check.
//
// The handler may be nil, http.Handler, or a function, the type of which is
// `func() string` that the returned result is the response content.
//
// If you want to start the server on SSL, you can give certFile and keyFile
// as the last two arguments.
func StartHealthServer(addr string, handler interface{}, files ...string) error {
	switch h := handler.(type) {
	case nil, http.Handler:
		SetHealthHandler(h)
	case func() string:
		SetHealthHandlerFunc(h)
	default:
		return fmt.Errorf("Unknown handler type")
	}

	_len := len(files)
	if _len == 0 {
		return ListenAndServe(addr, DefaultServeMux)
	} else if _len == 2 {
		return ListenAndServeTLS(addr, files[0], files[1], DefaultServeMux)
	}
	return fmt.Errorf("the options files is not certFile and keyFile")
}

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
func NewHTTPError(code int, err interface{}) error {
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

// ErrorLogFunc handles the http error log in ErrorHandler and
// ErrorHandlerWithStatusCode.
//
// Notice: The caller doesn't append the new line, so the function should
// append the new line.
//
// DEPRECATED!!! Please ErrorF in the sub-package log instead.
// It is only reserved, and not effect.
var ErrorLogFunc func(format string, args ...interface{})

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
			log.ErrorF("Handling %q: status=%d, err=%v", r.RequestURI, code, err)
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
			log.ErrorF("Failed to handle %q: %s", r.RequestURI, err)
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
			log.ErrorF("Failed to send the response of %q: %s", r.RequestURI, err)
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
	log.ErrorF("The server listening on %s has an error: %s", addr, err)
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
	log.ErrorF("The TLS server listening on %s has an error: %s", addr, err)
	lifecycle.Stop()
	return err
}

// GetBody returns the body of the HTTP request.
func GetBody(r *http.Request) (body []byte, err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.CopyN(buf, r.Body, r.ContentLength); err != nil {
		return
	}
	return buf.Bytes(), nil
}
