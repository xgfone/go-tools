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

// Package json is the supplement of the standard library of `json`.
package json

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"time"
)

// Predefine some json mark
var (
	nullBytes  = []byte("null")
	trueBytes  = []byte("true")
	falseBytes = []byte("false")

	commaBytes        = []byte{','}
	colonBytes        = []byte{':'}
	leftBracketBytes  = []byte{'['}
	rightBracketBytes = []byte{']'}
	leftBraceBytes    = []byte{'{'}
	rightBraceBytes   = []byte{'}'}
)

// MarshalJSON marshals a value v as JSON into w.
//
// Support the types:
//   nil
//   bool
//   string | error
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
//   time.Time  // The layout is time.RFC3339Nano.
//   map[string]interface{} or map[string]string for json object
//   json.Marshaler
//   fmt.Stringer
//   Array or Slice of the type above
//
// For other types, it will use json.Marshal() to marshal it.
func MarshalJSON(w io.Writer, v interface{}) (n int, err error) {
	switch _v := v.(type) {
	case nil:
		return w.Write(nullBytes)
	case bool:
		if _v {
			return w.Write(trueBytes)
		}
		return w.Write(falseBytes)
	case string:
		return marshalString(w, _v)
	case error:
		return marshalString(w, _v.Error())
	case time.Time:
		return marshalString(w, _v.Format(time.RFC3339Nano))
	case fmt.Stringer:
		return marshalString(w, _v.String())
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return marshalNumber(w, v)
	case map[string]interface{}:
		// Write {
		if n, err = w.Write(leftBraceBytes); err != nil {
			return n, err
		}
		total := n

		count := 0
		for key, value := range _v {
			if count > 0 {
				// Write comma
				if n, err = w.Write(commaBytes); err != nil {
					return total, err
				}
				total += n
			}

			// Write key
			if n, err = marshalString(w, key); err != nil {
				return total, err
			}
			total += n

			// Write :
			if n, err = w.Write(colonBytes); err != nil {
				return total, err
			}
			total += n

			// Write value
			if n, err = MarshalJSON(w, value); err != nil {
				return total, err
			}
			total += n

			count++
		}

		// Write }
		n, err = w.Write(rightBraceBytes)
		return total + n, err
	case map[string]string: // Optimize for map[string]string
		// Write {
		if n, err = w.Write(leftBraceBytes); err != nil {
			return n, err
		}
		total := n

		count := 0
		for key, value := range _v {
			if count > 0 {
				// Write comma
				if n, err = w.Write(commaBytes); err != nil {
					return total, err
				}
				total += n
			}

			// Write key
			if n, err = marshalString(w, key); err != nil {
				return total, err
			}
			total += n

			// Write :
			if n, err = w.Write(colonBytes); err != nil {
				return total, err
			}
			total += n

			// Write value
			if n, err = marshalString(w, value); err != nil {
				return total, err
			}
			total += n

			count++
		}

		// Write }
		n, err = w.Write(rightBraceBytes)
		return total + n, err
	case json.Marshaler:
		bs, err := _v.MarshalJSON()
		if err != nil {
			return 0, err
		}
		return w.Write(bs)
	case []string: // Optimzie []string
		if n, err = w.Write(leftBracketBytes); err != nil {
			return n, err
		}

		total := n
		for i, _len := 0, len(_v); i < _len; i++ {
			if i > 0 {
				if n, err = w.Write(commaBytes); err != nil {
					return total, err
				}
				total += n
			}

			if n, err = marshalString(w, _v[i]); err != nil {
				return total, err
			}
			total += n
		}

		n, err = w.Write(rightBracketBytes)
		return total + n, err
	case []interface{}: // Optimzie []interface{}
		if n, err = w.Write(leftBracketBytes); err != nil {
			return n, err
		}

		total := n
		for i, _len := 0, len(_v); i < _len; i++ {
			if i > 0 {
				if n, err = w.Write(commaBytes); err != nil {
					return total, err
				}
				total += n
			}

			if n, err = MarshalJSON(w, _v[i]); err != nil {
				return total, err
			}
			total += n
		}

		n, err = w.Write(rightBracketBytes)
		return total + n, err
	default:
		// Check whether it's an array or slice.
		value := reflect.ValueOf(v)
		kind := value.Kind()
		if kind != reflect.Array && kind != reflect.Slice {
			data, err := json.Marshal(v)
			if err != nil {
				return 0, err
			}
			return w.Write(data)
		}

		if n, err = w.Write(leftBracketBytes); err != nil {
			return n, err
		}

		total := n
		_len := value.Len()
		for i := 0; i < _len; i++ {
			if i > 0 {
				if n, err = w.Write(commaBytes); err != nil {
					return total, err
				}
				total += n
			}

			if n, err = MarshalJSON(w, value.Index(i).Interface()); err != nil {
				return total, err
			}
			total += n
		}

		n, err = w.Write(rightBracketBytes)
		return total + n, err
	}
}

// MarshalKvJSON marshals some key-value pairs as JSON into w.
//
// Notice: the key must be string, and the value may be one of the following:
//   nil
//   bool
//   string | error
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
//   time.Time  // The layout is time.RFC3339Nano.
//   map[string]interface{} or map[string]string for json object
//   json.Marshaler
//   fmt.Stringer
//   Array or Slice of the type above
func MarshalKvJSON(w io.Writer, args ...interface{}) (n int, err error) {
	_len := len(args)
	if _len == 0 {
		return
	} else if _len%2 == 1 {
		return 0, fmt.Errorf("args must be even")
	}

	var m int

	// Write {
	if m, err = w.Write(leftBraceBytes); err != nil {
		return
	}
	n += m

	for i := 0; i < _len; i += 2 {
		// Write comma
		if i > 0 {
			if m, err = w.Write(commaBytes); err != nil {
				return
			}
			n += m
		}

		// Write Key
		key, ok := args[i].(string)
		if !ok {
			return 0, fmt.Errorf("the %dth key is not string", i/2)
		}
		if m, err = marshalString(w, key); err != nil {
			return
		}
		n += m

		// Write :
		if m, err = io.WriteString(w, ":"); err != nil {
			return
		}
		n += m

		// Write Value
		switch v := args[i+1].(type) {
		case nil:
			if m, err = w.Write(nullBytes); err != nil {
				return
			}
			n += m
		case bool:
			if v {
				m, err = w.Write(trueBytes)
			} else {
				m, err = w.Write(falseBytes)
			}
			if err != nil {
				return
			}
			n += m
		case string:
			if m, err = marshalString(w, v); err != nil {
				return
			}
			n += m
		case error:
			if m, err = marshalString(w, v.Error()); err != nil {
				return
			}
			n += m
		case time.Time:
			if m, err = marshalString(w, v.Format(time.RFC3339Nano)); err != nil {
				return
			}
			n += m
		case fmt.Stringer:
			if m, err = marshalString(w, v.String()); err != nil {
				return
			}
			n += m
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64:
			if m, err = marshalNumber(w, v); err != nil {
				return
			}
			n += m
		case json.Marshaler:
			var bs []byte
			if bs, err = v.MarshalJSON(); err != nil {
				return
			} else if m, err = w.Write(bs); err != nil {
				return
			}
			n += m
		default: // For array, slice or map
			if m, err = MarshalJSON(w, v); err != nil {
				return
			}
			n += m
		}
	}

	// Write }
	if m, err = w.Write(rightBraceBytes); err != nil {
		return
	}
	n += m
	return
}

func marshalNumber(w io.Writer, n interface{}) (int, error) {
	var buf [8]byte
	switch v := n.(type) {
	case int:
		return w.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	case int8:
		return w.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	case int16:
		return w.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	case int32:
		return w.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	case int64:
		return w.Write(strconv.AppendInt(buf[:0], v, 10))
	case uint:
		return w.Write(strconv.AppendUint(buf[:0], uint64(v), 10))
	case uint8:
		return w.Write(strconv.AppendUint(buf[:0], uint64(v), 10))
	case uint16:
		return w.Write(strconv.AppendUint(buf[:0], uint64(v), 10))
	case uint32:
		return w.Write(strconv.AppendUint(buf[:0], uint64(v), 10))
	case uint64:
		return w.Write(strconv.AppendUint(buf[:0], v, 10))
	case float32:
		return w.Write(strconv.AppendFloat(buf[:0], float64(v), 'f', -1, 64))
	case float64:
		return w.Write(strconv.AppendFloat(buf[:0], v, 'f', -1, 64))
	default:
		return 0, nil
	}
}

var doubleQuotationByte = []byte{'"'}

func marshalString(w io.Writer, s string) (n int, err error) {
	// Check whether it needs to be escaped.
	var escape bool
	for _, c := range s {
		if c == '"' {
			escape = true
		}
	}
	if escape {
		return io.WriteString(w, strconv.Quote(s))
	}

	var m int

	m, err = w.Write(doubleQuotationByte)
	n += m
	if err != nil {
		return
	}

	m, err = io.WriteString(w, s)
	n += m
	if err != nil {
		return
	}

	m, err = w.Write(doubleQuotationByte)
	n += m

	return
}
