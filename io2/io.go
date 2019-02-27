// Package io2 is the supplement of the standard library of `io`.
package io2

import (
	"bufio"
	"bytes"
	"io"
)

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
	buf := bytes.NewBuffer(nil)
	err = ReadNWriter(buf, r, n)
	return buf.Bytes(), err
}

// ReadNWriter reads n bytes to the writer w from the reader r.
//
// It will return io.EOF if the length of the data from r is less than n.
// But the data has been read into w.
func ReadNWriter(w io.Writer, r io.Reader, n int64) (err error) {
	if n < 1 {
		_, err := io.Copy(w, r)
		return err
	}

	if buf, ok := w.(*bytes.Buffer); ok {
		if n < 32768 { // 32KB
			buf.Grow(int(n))
		} else {
			buf.Grow(32768)
		}
	}

	if m, err := io.Copy(w, io.LimitReader(r, n)); err != nil {
		return err
	} else if m < n {
		return io.EOF
	}
	return nil
}
