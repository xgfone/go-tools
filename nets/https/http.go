// Package https is the supplement of the standard library `http`,
// not the protocal `https`.
package https

import (
	"fmt"
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
