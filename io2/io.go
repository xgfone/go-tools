// Package io2 is the assistant functions of io.
//
// Deprecated!
package io2

import "io"

func readN2Buf(r io.Reader, buf []byte) (m int, err error) {
	total := len(buf)
	var _m int
	for m < total {
		_m, err = r.Read(buf[m:])
		if err != nil {
			return
		}
		m += _m
	}
	return
}

// ReadN2Buf must read and return len(buf) bytes from r to buf.
// Deprecated, see io.ReadFull.
//
// Return nil if reading len(n) bytes successfully, or non-nil.
//
// If you want to get the incomplete bytes read when failed, please use ReadN.
func ReadN2Buf(r io.Reader, buf []byte) error {
	_, err := readN2Buf(r, buf)
	return err
}

// ReadN is same as ReadN2Buf, but return the read buffer.
// Deprecated, see io.ReadFull.
//
// Return the read bytes if failed.
func ReadN(r io.Reader, n int) ([]byte, error) {
	result := make([]byte, n)
	m, err := readN2Buf(r, result)
	return result[:m], err
}

// WriteN must write n bytes to w.
// Deprecated, see io.CopyN.
//
// Return nil if writing len(n) bytes successfully, or non-nil.
func WriteN(w io.Writer, data []byte) error {
	n := len(data)
	for m := 0; m < n; {
		_m, err := w.Write(data[m:])
		if err != nil {
			return err
		}
		m += _m
	}
	return nil
}
