// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package io2 is the supplement of the standard library of `io`.
package io2

import (
	"bytes"
	"io"
)

// Peeker is an interface in order for Peek read.
type Peeker interface {
	// Peek returns the next n bytes without advancing the reader.
	Peek(n int) ([]byte, error)
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

// CopyNBuffer is the same as io.CopyN, but uses the given buf as the buffer.
//
// If buf is nil or empty, it will make a new one with 2048.
func CopyNBuffer(dst io.Writer, src io.Reader, n int64, buf []byte) (written int64, err error) {
	if len(buf) == 0 {
		buf = make([]byte, 2048)
	}

	written, err = io.CopyBuffer(dst, io.LimitReader(src, n), buf)
	if written == n {
		return n, nil
	} else if written < n && err == nil {
		// src stopped early; must have been EOF.
		err = io.EOF
	}

	return
}
