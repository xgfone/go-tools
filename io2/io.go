// Package io2 is the supplement of the standard library of `io`,
// such as `Close`.
package io2

import (
	"bufio"
	"io"

	"github.com/xgfone/go-tools/pools"
)

// Close implements the interface with the method Close(), which does not return
// an error.
//
// DEPRECATED!!!
type Close struct {
	Value io.Closer
}

// Close implements the method Close().
func (c Close) Close() {
	c.Value.Close()
}

// NewClose returns an new Close.
//
// DEPRECATED!!!
func NewClose(v io.Closer) Close {
	return Close{Value: v}
}

// ReadLine reads the content in the buffer by line.
func ReadLine(r *bufio.Reader) (lines [][]byte, err error) {
	var line []byte
	isPrefix := true
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines, err
}

// ReadN reads the data from io.Reader until n bytes or no incoming data
// if n is equal to or less than 0.
func ReadN(r io.Reader, n int64) (v []byte, err error) {
	w := pools.DefaultBufferPool.Get()
	err = ReadNWriter(w, r, n)
	v = w.Bytes()
	pools.DefaultBufferPool.Put(w)
	return v, err
}

// ReadNWriter is the same as ReadN, but writes the data to the writer
// from the reader.
func ReadNWriter(w io.Writer, r io.Reader, n int64) (err error) {
	buf := pools.BytesPool2K.Get()

	if n > 0 {
		var m int64
		m, err = io.CopyBuffer(w, io.LimitReader(r, n), buf)
		if m < n && err == nil {
			err = io.EOF
		}
	} else {
		_, err = io.CopyBuffer(w, r, buf)
	}

	pools.BytesPool2K.Put(buf)
	return
}
