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

package strings2

import (
	"io"
	"strconv"
	"time"
	"unicode/utf8"
)

// Builder is a thin wrapper around a byte slice. It's intended to be pooled, so
// the only way to construct one is via a Pool.
type Builder struct {
	buf []byte
}

// NewBuilder returns a new Builder with a initial capacity n.
func NewBuilder(n int) *Builder {
	return NewBuilderBytes(make([]byte, 0, n))
}

// NewBuilderBytes returns a new Builder with a initial data.
func NewBuilderBytes(buf []byte) *Builder {
	return &Builder{buf: buf}
}

// NewBuilderString returns a new Builder with a initial string.
func NewBuilderString(s string) *Builder {
	b := NewBuilderBytes(make([]byte, 0, len(s)*2))
	b.WriteString(s)
	return b
}

// Len returns the length of the underlying byte slice.
func (b *Builder) Len() int {
	return len(b.buf)
}

// Cap returns the capacity of the underlying byte slice.
func (b *Builder) Cap() int {
	return cap(b.buf)
}

// Bytes returns a mutable reference to the underlying byte slice.
func (b *Builder) Bytes() []byte {
	return b.buf
}

// String returns a string copy of the underlying byte slice.
func (b *Builder) String() string {
	return string(b.buf)
}

// Reset resets the underlying byte slice.
//
// Subsequent writes will re-use the slice's backing array.
func (b *Builder) Reset() {
	b.buf = b.buf[:0]
}

// ResetBytes resets the underlying byte slice to bs.
func (b *Builder) ResetBytes(bs []byte) {
	b.buf = bs
}

// TruncateBefore truncates and discards first n bytes.
//
// It will is equal to reset if n is greater than the length of the underlying
// byte slice,
func (b *Builder) TruncateBefore(n int) {
	if n = len(b.buf) - n; n > 0 {
		copy(b.buf, b.buf[n-1:])
		b.buf = b.buf[:n]
	} else {
		b.buf = b.buf[:0]
	}
}

// TruncateAfter truncates and discards last n bytes.
//
// It will is equal to reset if n is greater than the length of the underlying
// byte slice,
func (b *Builder) TruncateAfter(n int) {
	if n = len(b.buf) - n; n < 0 {
		n = 0
	}
	b.buf = b.buf[:n]
}

// AppendByte is the same as WriteByte, but no return.
func (b *Builder) AppendByte(c byte) {
	b.WriteByte(c)
}

// AppendString is the same as WriteString, but no return.
func (b *Builder) AppendString(s string) {
	b.WriteString(s)
}

// AppendInt appends an integer to the underlying buffer (assuming base 10).
func (b *Builder) AppendInt(i int64) {
	b.buf = strconv.AppendInt(b.buf, i, 10)
}

// AppendUint appends an unsigned integer to the underlying buffer (assuming
// base 10).
func (b *Builder) AppendUint(i uint64) {
	b.buf = strconv.AppendUint(b.buf, i, 10)
}

// AppendBool appends a bool to the underlying buffer.
func (b *Builder) AppendBool(v bool) {
	b.buf = strconv.AppendBool(b.buf, v)
}

// AppendFloat appends a float to the underlying buffer. It doesn't quote NaN
// or +/- Inf.
func (b *Builder) AppendFloat(f float64, bitSize int) {
	b.buf = strconv.AppendFloat(b.buf, f, 'f', -1, bitSize)
}

// AppendTime appends a time to the underlying buffer.
func (b *Builder) AppendTime(t time.Time, layout string) {
	b.buf = t.AppendFormat(b.buf, layout)
}

// Write implements io.Writer.
func (b *Builder) Write(bs []byte) (int, error) {
	b.buf = append(b.buf, bs...)
	return len(bs), nil
}

// WriteByte writes a byte into the builder.
func (b *Builder) WriteByte(c byte) error {
	b.buf = append(b.buf, c)
	return nil
}

// WriteRune writes a rune into the builder.
func (b *Builder) WriteRune(r rune) (int, error) {
	if r < utf8.RuneSelf {
		b.WriteByte(byte(r))
		return 1, nil
	}

	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	b.buf = append(b.buf, buf[:n]...)
	return n, nil
}

// WriteString writes a string into the builder.
func (b *Builder) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}

// WriteTo implements io.WriterTo.
func (b *Builder) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.buf)
	return int64(n), err
}

// TrimNewline trims any final "\n" byte from the end of the buffer.
func (b *Builder) TrimNewline() {
	maxIndex := len(b.buf) - 1
	for i := maxIndex; i >= 0; i-- {
		if b.buf[i] != '\n' {
			if i < maxIndex {
				b.buf = b.buf[:i+1]
			}
			return
		}
	}
}
