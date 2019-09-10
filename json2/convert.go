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

package json2

import (
	"encoding"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/xgfone/go-tools/types"
)

// ToBytesErr encodes a value to []byte.
//
// For the time.Time, it uses time.RFC3339Nano to format it.
//
// Support the types:
//   nil
//   bool
//   []byte
//   string
//   float32
//   float64
//   int
//   int8
//   int16
//   int32
//   int64
//   uint
//   uint8
//   uint16
//   uint32
//   uint64
//   time.Time
//   interface error
//   interface fmt.Stringer
//   interface Byter
//   interface MarshalByter
//   interface encoding.TextMarshaler
//
// For other types, use fmt.Sprintf("%v") to format it if fmtSprintf is true,
// or return the error types.ErrUnknownType.
func ToBytesErr(i interface{}, fmtSprintf ...bool) ([]byte, error) {
	switch v := i.(type) {
	case nil:
		return nilBytes, nil
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case bool:
		if v {
			return trueBytes, nil
		}
		return falseBytes, nil
	case float32:
		return strconv.AppendFloat(make([]byte, 0, 24), float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.AppendFloat(make([]byte, 0, 24), v, 'f', -1, 64), nil
	case int:
		return strconv.AppendInt(make([]byte, 0, 20), int64(v), 10), nil
	case int8:
		return strconv.AppendInt(make([]byte, 0, 20), int64(v), 10), nil
	case int16:
		return strconv.AppendInt(make([]byte, 0, 20), int64(v), 10), nil
	case int32:
		return strconv.AppendInt(make([]byte, 0, 20), int64(v), 10), nil
	case int64:
		return strconv.AppendInt(make([]byte, 0, 20), v, 10), nil
	case uint:
		return strconv.AppendUint(make([]byte, 0, 20), uint64(v), 10), nil
	case uint8:
		return strconv.AppendUint(make([]byte, 0, 20), uint64(v), 10), nil
	case uint16:
		return strconv.AppendUint(make([]byte, 0, 20), uint64(v), 10), nil
	case uint32:
		return strconv.AppendUint(make([]byte, 0, 20), uint64(v), 10), nil
	case uint64:
		return strconv.AppendUint(make([]byte, 0, 20), v, 10), nil
	case time.Time:
		return EncodeTime(v, time.RFC3339Nano), nil
	case encoding.TextMarshaler:
		return v.MarshalText()
	case error:
		return []byte(v.Error()), nil
	case fmt.Stringer:
		return []byte(v.String()), nil
	default:
		if len(fmtSprintf) > 0 && fmtSprintf[0] {
			return []byte(fmt.Sprintf("%v", v)), nil
		}
		return nil, types.ErrUnknownType
	}
}

// ToBytes is the same as ToBytesErr, but ignoring the error.
func ToBytes(i interface{}, fmtSprintf ...bool) []byte {
	bs, _ := ToBytesErr(i, fmtSprintf...)
	return bs
}

// ToStringErr is the same as ToBytesErr, but returns string.
func ToStringErr(i interface{}, fmtSprintf ...bool) (string, error) {
	switch v := i.(type) {
	case nil:
		return "nil", nil
	case string:
		return v, nil
	case error:
		return v.Error(), nil
	case fmt.Stringer:
		return v.String(), nil
	default:
		bs, err := ToBytesErr(i, fmtSprintf...)
		return string(bs), err
	}
}

// ToString is the same as ToBytesErr, but returns string and ignores the error.
func ToString(i interface{}, fmtSprintf ...bool) string {
	s, _ := ToStringErr(i, fmtSprintf...)
	return s
}

// EncodeNowTime is the same as EncodeTime, but encodes the now time.
func EncodeNowTime(layout string, utc ...bool) []byte {
	return EncodeTime(time.Now(), layout, utc...)
}

// EncodeTime encodes the time t to []byte, which will convrt it to UTC
// if utc is true.
func EncodeTime(t time.Time, layout string, utc ...bool) []byte {
	if len(utc) > 0 && utc[0] {
		t = t.UTC()
	}
	return t.AppendFormat(make([]byte, 0, 36), layout)
}

// Write is the same as ToBytesErr, but writes the result into w,
// and do some optimizations.
func Write(w io.Writer, i interface{}, fmtSprintf ...bool) error {
	switch v := i.(type) {
	case nil:
		io.WriteString(w, "nil")
	case []byte:
		w.Write(v)
	case string:
		io.WriteString(w, v)
	case error:
		io.WriteString(w, v.Error())
	case fmt.Stringer:
		io.WriteString(w, v.String())
	default:
		bs, err := ToBytesErr(i, fmtSprintf...)
		if err != nil {
			return err
		}
		w.Write(bs)
	}
	return nil
}
