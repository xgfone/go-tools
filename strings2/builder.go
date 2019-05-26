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
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
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

// AppendBool appends a bool to the underlying buffer.
func (b *Builder) AppendBool(v bool) {
	b.buf = strconv.AppendBool(b.buf, v)
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

// AppendFloat appends a float to the underlying buffer. It doesn't quote NaN
// or +/- Inf.
func (b *Builder) AppendFloat(f float64, bitSize int) {
	b.buf = strconv.AppendFloat(b.buf, f, 'f', -1, bitSize)
}

// AppendTime appends a time to the underlying buffer.
func (b *Builder) AppendTime(t time.Time, layout string) {
	b.buf = t.AppendFormat(b.buf, layout)
}

// AppendAny appends any value to the underlying buffer.
//
// It supports the type:
//    nil     ==> "<nil>"
//    bool    ==> "true" or "false"
//    []byte
//    string
//    float32
//    float64
//    int
//    int8
//    int16
//    int32
//    int64
//    uint
//    uint8
//    uint16
//    uint32
//    uint64
//    time.Time ==> time.RFC3339Nano
//    Slice
//    Map
//    interface error
//    interface fmt.Stringer
//    interface encoding.TextMarshaler
//
// For the unknown type, it does not append it and return false, or return true.
func (b *Builder) AppendAny(any interface{}) (ok bool, err error) {
	switch v := any.(type) {
	case nil:
		b.WriteString("<nil>")
	case bool:
		b.AppendBool(v)
	case []byte:
		b.Write(v)
	case string:
		b.WriteString(v)
	case float32:
		b.AppendFloat(float64(v), 32)
	case float64:
		b.AppendFloat(v, 64)
	case int:
		b.AppendInt(int64(v))
	case int8:
		b.AppendInt(int64(v))
	case int16:
		b.AppendInt(int64(v))
	case int32:
		b.AppendInt(int64(v))
	case int64:
		b.AppendInt(v)
	case uint:
		b.AppendUint(uint64(v))
	case uint8:
		b.AppendUint(uint64(v))
	case uint16:
		b.AppendUint(uint64(v))
	case uint32:
		b.AppendUint(uint64(v))
	case uint64:
		b.AppendUint(v)
	case time.Time:
		b.AppendTime(v, time.RFC3339Nano)
	case fmt.Stringer:
		b.WriteString(v.String())
	case error:
		b.WriteString(v.Error())
	case encoding.TextMarshaler:
		data, err := v.MarshalText()
		if err != nil {
			return true, err
		}
		b.Write(data)
	case []interface{}:
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(' ')
			}
			if ok, err = b.AppendAny(_v); !ok || err != nil {
				return
			}
		}
		b.WriteByte(']')
	case []string:
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(_v)
		}
		b.WriteByte(']')
	case []int:
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(' ')
			}
			b.AppendInt(int64(_v))
		}
		b.WriteByte(']')
	case map[string]interface{}:
		b.WriteString("map[")
		var i int
		for key, value := range v {
			if i > 0 {
				b.WriteByte(' ')
			}
			i++
			b.WriteString(key)
			b.WriteByte(':')
			b.AppendAny(value)
		}
		b.WriteByte(']')
	case map[string]string:
		b.WriteString("map[")
		var i int
		for key, value := range v {
			if i > 0 {
				b.WriteByte(' ')
			}
			i++
			b.WriteString(key)
			b.WriteByte(':')
			b.AppendString(value)
		}
		b.WriteByte(']')
	default:
		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Map && kind != reflect.Slice && kind != reflect.Array {
			return false, nil
		}
		fmt.Fprintf(b, "%v", v)
	}
	return true, nil
}

// AppendAnyFmt is the same as AppendAny(any), but it will use
// `fmt.Sprintf("%+v", any)` to format the unknown type `any`.
func (b *Builder) AppendAnyFmt(any interface{}) error {
	switch any.(type) {
	case nil:
	case bool:
	case []byte:
	case string:
	case float32:
	case float64:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case time.Time:
	case fmt.Stringer:
	case error:
	case encoding.TextMarshaler:
	case []interface{}:
	case []string:
	case []int:
	case map[string]interface{}:
	case map[string]string:
	default:
		fmt.Fprintf(b, "%+v", any)
		return nil
	}
	_, err := b.AppendAny(any)
	return err
}

// AppendJSONString appends a string as JSON string, which will escape
// the double quotation(") and enclose it with a pair of the double quotation.
func (b *Builder) AppendJSONString(s string) {
	if strings.IndexByte(s, '"') > -1 {
		b.buf = strconv.AppendQuote(b.buf, s)
	} else {
		b.WriteByte('"')
		b.WriteString(s)
		b.WriteByte('"')
	}
}

// AppendJSON appends the value as the JSON value, that's, the value will
// be encoded to JSON and appended into the underlying byte slice.
func (b *Builder) AppendJSON(value interface{}) error {
	switch v := value.(type) {
	case nil:
		b.WriteString(`null`)
	case bool:
		if v {
			b.WriteString(`true`)
		} else {
			b.WriteString(`false`)
		}
	case int:
		b.AppendInt(int64(v))
	case int8:
		b.AppendInt(int64(v))
	case int16:
		b.AppendInt(int64(v))
	case int32:
		b.AppendInt(int64(v))
	case int64:
		b.AppendInt(v)
	case uint:
		b.AppendUint(uint64(v))
	case uint8:
		b.AppendUint(uint64(v))
	case uint16:
		b.AppendUint(uint64(v))
	case uint32:
		b.AppendUint(uint64(v))
	case uint64:
		b.AppendUint(v)
	case float32:
		b.AppendFloat(float64(v), 32)
	case float64:
		b.AppendFloat(v, 64)
	case time.Time:
		b.WriteByte('"')
		b.AppendTime(v, time.RFC3339Nano)
		b.WriteByte('"')
	case error:
		b.AppendJSONString(v.Error())
	case string:
		b.AppendJSONString(v)
	case fmt.Stringer:
		b.AppendJSONString(v.String())
	case json.Marshaler:
		data, err := v.MarshalJSON()
		if err != nil {
			return err
		}
		b.Write(data)
	case []interface{}: // Optimize []interface{}
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			if err := b.AppendJSON(_v); err != nil {
				return err
			}
		}
		b.WriteByte(']')
	case []string: // Optimize []string
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			b.AppendJSONString(_v)
		}
		b.WriteByte(']')
	case []int: // Optimize []int
		b.WriteByte('[')
		for i, _v := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			b.AppendInt(int64(_v))
		}
		b.WriteByte(']')
	case map[string]interface{}: // Optimize map[string]interface{}
		b.WriteByte('{')
		var i int
		for key, value := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			i++
			b.AppendJSONString(key)
			b.WriteByte(':')
			if err := b.AppendJSON(value); err != nil {
				return err
			}
		}
		b.WriteByte('}')
	case map[string]string: // Optimize map[string]string
		b.WriteByte('{')
		var i int
		for key, value := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			i++
			b.AppendJSONString(key)
			b.WriteByte(':')
			b.AppendJSONString(value)
		}
		b.WriteByte('}')
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		b.Write(data)
	}
	return nil
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
