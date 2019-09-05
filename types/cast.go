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

package types

import (
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errNegativeNotAllowed = fmt.Errorf("unable to cast negative value")

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// StringToTime does the best to parse a string into a time.Time.
//
// If giving layout, it will use it, or attempt to guess it by using
// a predefined list of formats.
func StringToTime(s string, layout ...string) (t time.Time, err error) {
	if s == "" || s == "0000-00-00 00:00:00" {
		return
	} else if len(layout) > 0 && layout[0] != "" {
		return time.Parse(layout[0], s)
	}

	layouts := []string{
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		if t, err = time.Parse(layout, s); err == nil {
			return
		}
	}
	return t, fmt.Errorf("unable to parse time: '%s'", s)
}

// ToTime does the best to convert any certain value to time.Time.
//
// If value is string or []byte, it will use StringToTime to convert it.
func ToTime(value interface{}, layout ...string) (time.Time, error) {
	value = indirect(value)

	switch v := value.(type) {
	case nil:
		return time.Time{}, nil
	case time.Time:
		return v, nil
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	case string:
		return StringToTime(v, layout...)
	case []byte:
		if len(v) == 0 {
			return time.Time{}, nil
		}
		return StringToTime(string(v), layout...)
	case fmt.Stringer:
		return StringToTime(v.String(), layout...)
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to time.Time", v, v)
	}
}

// ToLocalTime does the best to convert any certain value to time.Time
//
// DEPRECATED!!! It is equal to ToTime.
func ToLocalTime(value interface{}, layout ...string) (time.Time, error) {
	return ToTime(value, layout...)
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(value interface{}) (time.Duration, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case time.Duration:
		return v, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float32, float64:
		i, _ := ToInt64(value)
		return time.Duration(i), nil
	case string:
		s = v
	case []byte:
		s = string(v)
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to Duration", v, v)
	}

	if strings.ContainsAny(s, "nsuµmh") {
		return time.ParseDuration(s)
	}
	return time.ParseDuration(s + "ns")
}

// ToBool does the best to convert any certain value to bool.
//
// For the string, the true value is
//   "t", "T", "1", "on", "On", "ON", "true", "True", "TRUE", "yes", "Yes", "YES"
// the false value is
//   "f", "F", "0", "off", "Off", "OFF", "false", "False", "FALSE", "no", "No", "NO", ""
//
// For other types, if the value is ZERO of the type, it's false. Or it's true.
func ToBool(value interface{}) (bool, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return false, nil
	case bool:
		return v, nil
	case []byte:
		if len(v) == 0 {
			return false, nil
		}
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return !IsZero(v), nil
	}

	switch s {
	case "t", "T", "1", "on", "On", "ON", "true", "True", "TRUE", "yes", "Yes", "YES":
		return true, nil
	case "f", "F", "0", "off", "Off", "OFF", "false", "False", "FALSE", "no", "No", "NO", "":
		return false, nil
	default:
		return false, fmt.Errorf("unrecognized bool string: %s", s)
	}
}

// ToFloat64 does the best to convert any certain value to float64.
func ToFloat64(value interface{}) (float64, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case complex64:
		return float64(real(v)), nil
	case complex128:
		return real(v), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", v, v)
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to float64", value, value)
}

// ToFloat32 does the best to convert any certain value to float32.
func ToFloat32(value interface{}) (float32, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case complex64:
		return float32(real(v)), nil
	case complex128:
		return float32(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", v, v)
	}

	if f, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(f), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to float32", value, value)
}

// ToInt does the best to convert any certain value to int.
func ToInt(value interface{}) (int, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case complex64:
		return int(real(v)), nil
	case complex128:
		return int(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", v, v)
	}

	if i, err := strconv.ParseInt(s, 0, 64); err == nil {
		return int(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int", value, value)
}

// ToInt8 does the best to convert any certain value to int8.
func ToInt8(value interface{}) (int8, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return int8(v), nil
	case int8:
		return v, nil
	case int16:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case uint:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	case float32:
		return int8(v), nil
	case float64:
		return int8(v), nil
	case complex64:
		return int8(real(v)), nil
	case complex128:
		return int8(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", v, v)
	}

	if i, err := strconv.ParseInt(s, 0, 8); err == nil {
		return int8(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int8", value, value)
}

// ToInt16 does the best to convert any certain value to int16.
func ToInt16(value interface{}) (int16, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case float32:
		return int16(v), nil
	case float64:
		return int16(v), nil
	case complex64:
		return int16(real(v)), nil
	case complex128:
		return int16(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", v, v)
	}

	if i, err := strconv.ParseInt(s, 0, 16); err == nil {
		return int16(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int16", value, value)
}

// ToInt32 does the best to convert any certain value to int32.
func ToInt32(value interface{}) (int32, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case complex64:
		return int32(real(v)), nil
	case complex128:
		return int32(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", v, v)
	}

	if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return int32(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int32", value, value)
}

// ToInt64 does the best to convert any certain value to int64.
func ToInt64(value interface{}) (int64, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case complex64:
		return int64(real(v)), nil
	case complex128:
		return int64(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", v, v)
	}

	if i, err := strconv.ParseInt(s, 0, 64); err == nil {
		return i, nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int64", value, value)
}

// ToUint does the best to convert any certain value to uint.
func ToUint(value interface{}) (uint, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		return uint(v), nil
	case float32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), nil
	case complex64:
		return uint(real(v)), nil
	case complex128:
		return uint(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", v, v)
	}

	if i, err := strconv.ParseUint(s, 0, 64); err == nil {
		return uint(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to uint", value, value)
}

// ToUint8 does the best to convert any certain value to uint8.
func ToUint8(value interface{}) (uint8, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case uint:
		return uint8(v), nil
	case uint8:
		return v, nil
	case uint16:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case float32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case float64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), nil
	case complex64:
		return uint8(real(v)), nil
	case complex128:
		return uint8(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", v, v)
	}

	if i, err := strconv.ParseUint(s, 0, 8); err == nil {
		return uint8(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", value, value)
}

// ToUint16 does the best to convert any certain value to uint16.
func ToUint16(value interface{}) (uint16, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case uint:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case float32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case float64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), nil
	case complex64:
		return uint16(real(v)), nil
	case complex128:
		return uint16(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", v, v)
	}

	if i, err := strconv.ParseUint(s, 0, 16); err == nil {
		return uint16(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", value, value)
}

// ToUint32 does the best to convert any certain value to uint32.
func ToUint32(value interface{}) (uint32, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case uint:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		return uint32(v), nil
	case float32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), nil
	case complex64:
		return uint32(real(v)), nil
	case complex128:
		return uint32(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", v, v)
	}

	if i, err := strconv.ParseUint(s, 0, 32); err == nil {
		return uint32(i), nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", value, value)
}

// ToUint64 does the best to convert any certain value to uint64.
func ToUint64(value interface{}) (uint64, error) {
	value = indirect(value)

	var s string
	switch v := value.(type) {
	case nil:
		return 0, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), nil
	case complex64:
		return uint64(real(v)), nil
	case complex128:
		return uint64(real(v)), nil
	case []byte:
		s = string(v)
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", v, v)
	}

	if i, err := strconv.ParseUint(s, 0, 64); err == nil {
		return i, nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", value, value)
}

// ToString does the best to convert any certain value to string.
//
// For time.Time, it will use time.RFC3339Nano to format it.
func ToString(value interface{}) (string, error) {
	value = indirect(value)

	switch v := value.(type) {
	case nil:
		return "", nil
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case error:
		return v.Error(), nil
	case time.Time:
		return v.Format(time.RFC3339Nano), nil
	case template.HTML:
		return string(v), nil
	case template.URL:
		return string(v), nil
	case template.JS:
		return string(v), nil
	case template.CSS:
		return string(v), nil
	case template.HTMLAttr:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", v, v)
	}
}

// ToStringMapBool casts an interface to a map[string]bool type.
func ToStringMapBool(value interface{}) (m map[string]bool, err error) {
	m = make(map[string]bool)

	switch v := value.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			key, e := ToString(k)
			if e != nil {
				return m, e
			}
			if m[key], err = ToBool(val); err != nil {
				return
			}
		}
		return
	case map[string]interface{}:
		for k, val := range v {
			if m[k], err = ToBool(val); err != nil {
				return
			}
		}
		return
	case map[string]bool:
		return v, nil
	case []byte:
		err = json.Unmarshal(v, &m)
		return
	case string:
		err = json.Unmarshal([]byte(v), &m)
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]bool", v, v)
		return
	}
}

// ToStringMapString casts an interface to a map[string]string type.
func ToStringMapString(value interface{}) (m map[string]string, err error) {
	m = make(map[string]string)

	switch v := value.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		for k, val := range v {
			if m[k], err = ToString(val); err != nil {
				return
			}
		}
		return
	case map[interface{}]string:
		for k, val := range v {
			key, e := ToString(k)
			if e != nil {
				return m, e
			}
			m[key] = val
		}
		return
	case map[interface{}]interface{}:
		for k, val := range v {
			key, e := ToString(k)
			if e != nil {
				return m, e
			}
			if m[key], err = ToString(val); err != nil {
				return
			}
		}
		return
	case []byte:
		err = json.Unmarshal(v, &m)
		return
	case string:
		err = json.Unmarshal([]byte(v), &m)
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]string", v, v)
		return
	}
}

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(value interface{}) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})

	switch v := value.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			key, e := ToString(k)
			if e != nil {
				return m, e
			}
			m[key] = val
		}
		return
	case map[string]interface{}:
		return v, nil
	case map[string]string:
		for k, val := range v {
			m[k] = val
		}
		return
	case []byte:
		err = json.Unmarshal(v, &m)
		return
	case string:
		err = json.Unmarshal([]byte(v), &m)
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", v, v)
		return
	}
}

// ToMapKeys returns all the keys of a map.
//
// If the value is not a map or the key is not string, it returns an error.
// But if the value is nil, it will return a empty slice, not an error instead.
//
// If you ensure that v is a map, and its key is the string type, you can ignore
// the error.
//
// For map[string]interface{}, map[string]string and map[string]int, they have
// already been optimized.
func ToMapKeys(v interface{}) ([]string, error) {
	switch _v := v.(type) {
	case nil:
		return []string{}, nil
	case map[string]interface{}:
		results := make([]string, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	case map[string]string:
		results := make([]string, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	case map[string]int:
		results := make([]string, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	}

	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Map {
		return nil, ErrNotMap
	}

	results := make([]string, _v.Len())
	for i, key := range _v.MapKeys() {
		if key.Kind() != reflect.String {
			return nil, ErrNotString
		}
		results[i] = key.String()
	}
	return results, nil
}

// ToMapValues returns all the values of a map.
//
// If the value is not a map, it returns an error.
// But if the value is nil, it will return a empty slice, not an error instead.
//
// If you ensure that v is a map, you can ignore the error.
//
// For map[string]interface{}, map[string]string and map[string]int, they have
// already been optimized.
func ToMapValues(v interface{}) ([]interface{}, error) {
	switch _v := v.(type) {
	case nil:
		return []interface{}{}, nil
	case map[string]interface{}:
		results := make([]interface{}, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	case map[string]string:
		results := make([]interface{}, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	case map[string]int:
		results := make([]interface{}, len(_v))
		for k := range _v {
			results = append(results, k)
		}
		return results, nil
	}

	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Map {
		return nil, ErrNotMap
	}

	results := make([]interface{}, _v.Len())
	for i, key := range _v.MapKeys() {
		results[i] = _v.MapIndex(key).Interface()
	}
	return results, nil
}

// ToSlice converts any slice type of []interface{}.
func ToSlice(value interface{}) ([]interface{}, error) {
	switch vs := value.(type) {
	case nil:
		return []interface{}{}, nil
	case []interface{}:
		return vs, nil
	case []string:
		results := make([]interface{}, len(vs))
		for i, v := range vs {
			results[i] = v
		}
		return results, nil
	case []int:
		results := make([]interface{}, len(vs))
		for i, v := range vs {
			results[i] = v
		}
		return results, nil
	}

	vf := reflect.ValueOf(value)
	if kind := vf.Kind(); kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []interface{}", value, value)
	}

	results := make([]interface{}, vf.Len())
	for i, _len := 0, vf.Len(); i < _len; i++ {
		results[i] = vf.Index(i).Interface()
	}
	return results, nil
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice(value interface{}) ([]bool, error) {
	switch v := value.(type) {
	case nil:
		return []bool{}, nil
	case []bool:
		return v, nil
	case []int:
		vs := make([]bool, len(v))
		for i, _v := range v {
			if _v == 0 {
				vs[i] = false
			} else {
				vs[i] = true
			}
		}
		return vs, nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]bool, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToBool(vf.Index(i).Interface()); err != nil {
				return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", value, value)
			}
		}
		return vs, nil
	default:
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", value, value)
	}
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(value interface{}) ([]int, error) {
	switch v := value.(type) {
	case nil:
		return []int{}, nil
	case []int:
		return v, nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]int, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToInt(vf.Index(i).Interface()); err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", value, value)
			}
		}
		return vs, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", value, value)
	}
}

// ToUintSlice casts an interface to a []uint type.
func ToUintSlice(value interface{}) ([]uint, error) {
	switch v := value.(type) {
	case nil:
		return []uint{}, nil
	case []uint:
		return v, nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]uint, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToUint(vf.Index(i).Interface()); err != nil {
				return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []uint", value, value)
			}
		}
		return vs, nil
	default:
		return []uint{}, fmt.Errorf("unable to cast %#v of type %T to []uint", value, value)
	}
}

// ToFloat64Slice casts an interface to a []float64 type.
func ToFloat64Slice(value interface{}) ([]float64, error) {
	switch v := value.(type) {
	case nil:
		return []float64{}, nil
	case []float64:
		return v, nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]float64, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToFloat64(vf.Index(i).Interface()); err != nil {
				return []float64{}, fmt.Errorf("unable to cast %#v of type %T to []float64", value, value)
			}
		}
		return vs, nil
	default:
		return []float64{}, fmt.Errorf("unable to cast %#v of type %T to []float64", value, value)
	}
}

// StringSeparator is the separator of the string slice to split the string.
var StringSeparator = ""

func isStringSeparator(r rune) bool {
	for _, c := range StringSeparator {
		if c == r {
			return true
		}
	}
	return false
}

// ToStringSlice casts an interface to a []string type.
//
// If value is string and the global variable StringSeparator is not "",
// the value will be split into []string by the string separator.
func ToStringSlice(value interface{}) ([]string, error) {
	switch v := value.(type) {
	case nil:
		return []string{}, nil
	case []interface{}:
		ss := make([]string, len(v))
		for i, _v := range v {
			ss[i], _ = ToString(_v)
		}
		return ss, nil
	case []string:
		return v, nil
	case string:
		if StringSeparator == "" {
			return strings.Fields(v), nil
		}
		return strings.FieldsFunc(v, isStringSeparator), nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]string, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToString(vf.Index(i).Interface()); err != nil {
				return []string{}, fmt.Errorf("unable to cast %#v of type %T to []string", value, value)
			}
		}
		return vs, nil
	default:
		if s, err := ToString(value); err == nil {
			return []string{s}, nil
		}
		return []string{}, fmt.Errorf("unable to cast %#v of type %T to []string", value, value)
	}
}

// ToDurationSlice casts an interface to a []time.Duration type.
func ToDurationSlice(value interface{}) ([]time.Duration, error) {
	switch v := value.(type) {
	case nil:
		return []time.Duration{}, nil
	case []time.Duration:
		return v, nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		var err error
		vf := reflect.ValueOf(value)
		vs := make([]time.Duration, vf.Len())
		for i, _len := 0, vf.Len(); i < _len; i++ {
			if vs[i], err = ToDuration(vf.Index(i).Interface()); err != nil {
				return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", value, value)
			}
		}
		return vs, nil
	default:
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", value, value)
	}
}
