// Package io2 is the supplement of the standard library of `io`,
// such as `Close`.
package io2

import "io"

// Close implements the interface with the method Close(), which does not return
// an error.
type Close struct {
	Value io.Closer
}

// Close implements the method Close().
func (c Close) Close() {
	c.Value.Close()
}

// NewClose returns an new Close.
func NewClose(v io.Closer) Close {
	return Close{Value: v}
}
