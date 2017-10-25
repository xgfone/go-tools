// Package https is the supplement of the standard library `http`,
// not the protocal `https`.
package https

import (
	"fmt"
	"net/http"
	"os"
)

// HTTPError stands for a HTTP error.
type HTTPError struct {
	// The error information
	Err error

	// You can assign it any what you want.
	Flag int

	// You can place data into it to carry in an error.
	Data map[string]interface{}
}

// NewHTTPError returns a new HTTPError.
func NewHTTPError(flag int, err interface{}) error {
	switch err.(type) {
	case error:
	case []byte:
		err = fmt.Errorf("%s", string(err.([]byte)))
	default:
		err = fmt.Errorf("%v", err)
	}
	return HTTPError{Flag: flag, Err: err.(error)}
}

func (e HTTPError) Error() string {
	return e.Err.Error()
}

// ErrorLogFunc handles the http error log in ErrorHandler and
// ErrorHandlerWithStatusCode.
//
// Notice: The caller doesn't append the new line, so the function should
// append the new line.
var ErrorLogFunc func(format string, args ...interface{})

func init() {
	ErrorLogFunc = func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}
}

// ErrorHandler handles the error and responds it the client.
func ErrorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return ErrorHandlerWithStatusCode(func(w http.ResponseWriter,
		r *http.Request) (int, error) {
		return http.StatusInternalServerError, f(w, r)
	})
}

// ErrorHandlerWithStatusCode handles the error and responds it the client
// with the status code.
func ErrorHandlerWithStatusCode(f func(http.ResponseWriter, *http.Request) (
	int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if code, err := f(w, r); err != nil {
			if code == 0 {
				code = http.StatusInternalServerError
			}
			http.Error(w, err.Error(), code)
			ErrorLogFunc("Handling %q: status=%d, err=%v", r.RequestURI, code, err)
		}
	}
}
