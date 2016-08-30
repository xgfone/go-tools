// The assistant functions of io.
package io2

import "io"

// Must read and return n bytes from r. err is not nil, if failed.
func ReadN(r io.Reader, n int) (result []byte, err error) {
	result = make([]byte, n)
	for m := 0; m < n; {
		m, err = r.Read(result[m:])
		if err != nil {
			return
		}
	}
	return
}

// Must write n bytes to w. err is not nil, if failed.
func WriteN(w io.Writer, data []byte) (err error) {
	n := len(data)
	for m := 0; m < n; {
		m, err = w.Write(data[m:])
		if err != nil {
			return err
		}
	}
	return nil
}
