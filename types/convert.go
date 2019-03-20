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
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	// ErrNotSliceOrArray is returned when the value is not a slice.
	ErrNotSliceOrArray = fmt.Errorf("the value is not a slice or array")

	// ErrNotMap is returned when the value is not a map.
	ErrNotMap = fmt.Errorf("the value is not a map")

	// ErrNotString is returned when the type of the key is not string.
	ErrNotString = fmt.Errorf("the type of the key is not string")

	// ErrKindNotExist is returned when a certain kind does not exist.
	ErrKindNotExist = fmt.Errorf("no kind")

	// ErrUnknownType is returned when not to identify a data type.
	ErrUnknownType = fmt.Errorf("unknown type")
)

var converters = make(map[Kind]func(interface{}) (interface{}, error))

// Predefine some kinds.
const (
	Unknown Kind = iota
	Nil
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Bytes
	Time        // For the format "YYYY-MM-DD HH:MM:SS"
	RFC3339Time // For the format time.RFC3339
)

// Kind represents the kind of the converter.
type Kind int

func init() {
	RegisterConverter(Bool, func(v interface{}) (interface{}, error) { return ToBool(v) })
	RegisterConverter(Int, func(v interface{}) (interface{}, error) { return ToInt(v) })
	RegisterConverter(Int32, func(v interface{}) (interface{}, error) { return ToInt32(v) })
	RegisterConverter(Int64, func(v interface{}) (interface{}, error) { return ToInt64(v) })
	RegisterConverter(Uint, func(v interface{}) (interface{}, error) { return ToUint(v) })
	RegisterConverter(Uint32, func(v interface{}) (interface{}, error) { return ToUint32(v) })
	RegisterConverter(Uint64, func(v interface{}) (interface{}, error) { return ToUint64(v) })
	RegisterConverter(Float64, func(v interface{}) (interface{}, error) { return ToFloat64(v) })
	RegisterConverter(String, func(v interface{}) (interface{}, error) { return ToString(v) })
	RegisterConverter(RFC3339Time, func(v interface{}) (interface{}, error) { return ToTime(v, time.RFC3339) })
	RegisterConverter(Time, func(v interface{}) (interface{}, error) {
		if s, _ := v.(string); s == "0000-00-00 00:00:00" {
			return time.Time{}, nil
		}
		return ToTime(v, "2006-01-02 15:04:05")
	})
}

// RegisterConverter registers a converter of the kind k.
//
// By default it has registered the kinds as follow:
//
//     Bool
//     String
//     Float64
//     Int, Int32, Int64
//     Uint, Uint32, Uint64
//     Time, RFC3339Time
func RegisterConverter(k Kind, converter func(interface{}) (interface{}, error)) {
	converters[k] = converter
}

// Convert calls the converter of the kind k to convert the value v.
//
// If the converter of the kind k does not exists, it returns ErrKindNotExist.
func Convert(k Kind, v interface{}) (interface{}, error) {
	if c, ok := converters[k]; ok {
		return c(v)
	}
	return nil, ErrKindNotExist
}

// Converter is used to convert the value by the Scan method. So you can use it
// as the argument of Rows.Scan() in sql.
type Converter struct {
	kind  Kind
	value interface{}
}

// NewConverter returns a Converter to convert a value to the type kind.
func NewConverter(kind Kind) Converter {
	return Converter{kind: kind}
}

// Scan converts the value src.
func (c *Converter) Scan(src interface{}) error {
	value, err := Convert(c.kind, src)
	if err != nil {
		return err
	}
	c.value = value
	return nil
}

// Value returns the inner converted result.
func (c Converter) Value() interface{} {
	return c.value
}

// bool2Int64 converts bool to int64.
func bool2Int64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// ToSlice converts any slice type of []interface{}.
//
// Return nil and an error if v is not a slice type.
//
// For []interface{}, []string and []int, they have already been optimized.
func ToSlice(v interface{}) ([]interface{}, error) {
	switch vs := v.(type) {
	case nil:
		return nil, ErrNotSliceOrArray
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

	_v := reflect.ValueOf(v)
	kind := _v.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrNotSliceOrArray
	}

	_len := _v.Len()
	results := make([]interface{}, _len)
	for i := 0; i < _len; i++ {
		results[i] = _v.Index(i).Interface()
	}
	return results, nil
}

// ToMap converts any map type that the key is string to map[string]interface{}.
//
// Return nil and an error if v is not a map type or its key is not the string
// type.
//
// If you ensure that v is a map, and its key is the string type, you can ignore
// the error.
//
// For []interface{}, []string and []int, they have already been optimized.
func ToMap(v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, ErrNotMap
	}

	if _v, ok := v.(map[string]interface{}); ok {
		return _v, nil
	}

	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Map {
		return nil, ErrNotMap
	}

	results := make(map[string]interface{}, _v.Len())
	for _, key := range _v.MapKeys() {
		if key.Kind() != reflect.String {
			return nil, ErrNotString
		}
		results[key.String()] = _v.MapIndex(key).Interface()
	}
	return results, nil
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

// ToTime does the best to convert any certain value to time.Time.
//
// Notice: the layout is time.RFC3339Nano by default.
func ToTime(v interface{}, layout ...string) (time.Time, error) {
	switch _v := v.(type) {
	case nil:
		return time.Time{}, nil
	case time.Time:
		return _v, nil
	case string:
		if len(layout) > 0 && layout[0] != "" {
			return time.Parse(layout[0], _v)
		}
		return time.Parse(time.RFC3339Nano, _v)
	default:
		return time.Time{}, ErrUnknownType
	}
}

// ToLocalTime does the best to convert any certain value to time.Time
// by time.ParseInLocation(layout, v, time.Local).
//
// Notice: the layout is time.RFC3339Nano by default.
func ToLocalTime(v interface{}, layout ...string) (time.Time, error) {
	switch _v := v.(type) {
	case nil:
		return time.Time{}, nil
	case time.Time:
		return _v, nil
	case string:
		if len(layout) > 0 && layout[0] != "" {
			return time.ParseInLocation(layout[0], _v, time.Local)
		}
		return time.ParseInLocation(time.RFC3339Nano, _v, time.Local)
	default:
		return time.Time{}, ErrUnknownType
	}
}

// ToBool does the best to convert any certain value to bool.
//
// For the string, the true value is
//   "t", "T", "1", "on", "On", "ON", "true", "True", "TRUE", "yes", "Yes", "YES"
// the false value is
//   "f", "F", "0", "off", "Off", "OFF", "false", "False", "FALSE", "no", "No", "NO", ""
//
// For other types, if the value is ZERO of the type, it's false. Or it's true.
func ToBool(v interface{}) (bool, error) {
	switch _v := v.(type) {
	case nil:
		return false, nil
	case bool:
		return _v, nil
	case string:
		switch _v {
		case "t", "T", "1", "on", "On", "ON", "true", "True", "TRUE", "yes", "Yes", "YES":
			return true, nil
		case "f", "F", "0", "off", "Off", "OFF", "false", "False", "FALSE", "no", "No", "NO", "":
			return false, nil
		default:
			return false, fmt.Errorf("unrecognized bool string: %s", _v)
		}
	}
	return !IsZero(v), nil
}

// ToInt64 does the best to convert any certain value to int64.
func ToInt64(_v interface{}) (v int64, err error) {
	switch t := _v.(type) {
	case nil:
	case bool:
		v = bool2Int64(t)
	case string:
		v, err = strconv.ParseInt(t, 10, 64)
	case int:
		v = int64(t)
	case int8:
		v = int64(t)
	case int16:
		v = int64(t)
	case int32:
		v = int64(t)
	case int64:
		v = t
	case uint:
		v = int64(t)
	case uint8:
		v = int64(t)
	case uint16:
		v = int64(t)
	case uint32:
		v = int64(t)
	case uint64:
		v = int64(t)
	case float32:
		v = int64(t)
	case float64:
		v = int64(t)
	case complex64:
		v = int64(real(t))
	case complex128:
		v = int64(real(t))
	default:
		err = ErrUnknownType
	}
	return
}

// ToUint64 does the best to convert any certain value to uint64.
func ToUint64(_v interface{}) (v uint64, err error) {
	switch t := _v.(type) {
	case nil:
	case bool:
		v = uint64(bool2Int64(t))
	case string:
		v, err = strconv.ParseUint(t, 10, 64)
	case int:
		v = uint64(t)
	case int8:
		v = uint64(t)
	case int16:
		v = uint64(t)
	case int32:
		v = uint64(t)
	case int64:
		v = uint64(t)
	case uint:
		v = uint64(t)
	case uint8:
		v = uint64(t)
	case uint16:
		v = uint64(t)
	case uint32:
		v = uint64(t)
	case uint64:
		v = t
	case float32:
		v = uint64(t)
	case float64:
		v = uint64(t)
	case complex64:
		v = uint64(real(t))
	case complex128:
		v = uint64(real(t))
	default:
		err = ErrUnknownType
	}
	return
}

// ToString does the best to convert any certain value to string.
//
// For time.Time, it will use time.RFC3339Nano to format it.
func ToString(_v interface{}) (v string, err error) {
	switch t := _v.(type) {
	case nil:
	case string:
		v = t
	case []byte:
		v = string(t)
	case bool:
		if t {
			v = "true"
		} else {
			v = "false"
		}
	case int:
		v = strconv.FormatInt(int64(t), 10)
	case int8:
		v = strconv.FormatInt(int64(t), 10)
	case int16:
		v = strconv.FormatInt(int64(t), 10)
	case int32:
		v = strconv.FormatInt(int64(t), 10)
	case int64:
		v = strconv.FormatInt(t, 10)
	case uint:
		v = strconv.FormatUint(uint64(t), 10)
	case uint8:
		v = strconv.FormatUint(uint64(t), 10)
	case uint16:
		v = strconv.FormatUint(uint64(t), 10)
	case uint32:
		v = strconv.FormatUint(uint64(t), 10)
	case uint64:
		v = strconv.FormatUint(t, 10)
	case float32:
		v = strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		v = strconv.FormatFloat(t, 'f', -1, 64)
	case error:
		v = t.Error()
	case time.Time:
		v = t.Format(time.RFC3339Nano)
	case fmt.Stringer:
		v = t.String()
	default:
		err = ErrUnknownType
	}
	return
}

// ToFloat64 does the best to convert any certain value to float64.
func ToFloat64(_v interface{}) (v float64, err error) {
	switch t := _v.(type) {
	case nil:
	case bool:
		v = float64(bool2Int64(t))
	case string:
		v, err = strconv.ParseFloat(t, 64)
	case int:
		v = float64(t)
	case int8:
		v = float64(t)
	case int16:
		v = float64(t)
	case int32:
		v = float64(t)
	case int64:
		v = float64(t)
	case uint:
		v = float64(t)
	case uint8:
		v = float64(t)
	case uint16:
		v = float64(t)
	case uint32:
		v = float64(t)
	case uint64:
		v = float64(t)
	case float32:
		v = float64(t)
	case float64:
		v = t
	case complex64:
		v = float64(real(t))
	case complex128:
		v = real(t)
	default:
		err = ErrUnknownType
	}
	return
}

// ToComplex128 does the best to convert any certain value to complex128.
func ToComplex128(_v interface{}) (v complex128, err error) {
	switch t := _v.(type) {
	case nil:
	case complex64:
		v = complex128(t)
	case complex128:
		v = t
	case float32:
		v = complex(float64(t), 0)
	case float64:
		v = complex(t, 0)
	case bool:
		if t {
			v = complex(1, 0)
		} else {
			v = complex(0, 0)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		f, _ := ToFloat64(_v)
		v = complex(f, 0)
	default:
		err = ErrUnknownType
	}
	return
}

// ToInt does the best to convert any certain value to int.
func ToInt(v interface{}) (int, error) {
	_v, err := ToInt64(v)
	return int(_v), err
}

// ToUint does the best to convert any certain value to uint.
func ToUint(v interface{}) (uint, error) {
	_v, err := ToUint64(v)
	return uint(_v), err
}

// ToInt8 does the best to convert any certain value to int8.
func ToInt8(v interface{}) (int8, error) {
	_v, err := ToInt64(v)
	return int8(_v), err
}

// ToUint8 does the best to convert any certain value to uint8.
func ToUint8(v interface{}) (uint8, error) {
	_v, err := ToUint64(v)
	return uint8(_v), err
}

// ToInt16 does the best to convert any certain value to int16.
func ToInt16(v interface{}) (int16, error) {
	_v, err := ToInt64(v)
	return int16(_v), err
}

// ToUint16 does the best to convert any certain value to uint16.
func ToUint16(v interface{}) (uint16, error) {
	_v, err := ToUint64(v)
	return uint16(_v), err
}

// ToInt32 does the best to convert any certain value to int32.
func ToInt32(v interface{}) (int32, error) {
	_v, err := ToInt64(v)
	return int32(_v), err
}

// ToUint32 does the best to convert any certain value to uint32.
func ToUint32(v interface{}) (uint32, error) {
	_v, err := ToUint64(v)
	return uint32(_v), err
}

// MustToSlice is equal to ToSlice, but panic if there is an error.
func MustToSlice(v interface{}) []interface{} {
	_v, err := ToSlice(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToMap is equal to ToMap, but panic if there is an error.
func MustToMap(v interface{}) map[string]interface{} {
	_v, err := ToMap(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToMapKeys is equal to ToMapKeys, but panic if there is an error.
func MustToMapKeys(v interface{}) []string {
	_v, err := ToMapKeys(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToMapValues is equal to ToMapValues, but panic if there is an error.
func MustToMapValues(v interface{}) []interface{} {
	_v, err := ToMapValues(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToTime is equal to ToTime, but panic if there is an error.
func MustToTime(v interface{}, layout ...string) time.Time {
	_v, err := ToTime(v, layout...)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToBool is equal to ToBool, but panic if there is an error.
func MustToBool(v interface{}) bool {
	_v, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt64 is equal to ToInt64, but panic if there is an error.
func MustToInt64(v interface{}) int64 {
	_v, err := ToInt64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint64 is equal to ToUint64, but panic if there is an error.
func MustToUint64(v interface{}) uint64 {
	_v, err := ToUint64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToString is equal to ToString, but panic if there is an error.
func MustToString(v interface{}) string {
	_v, err := ToString(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToFloat64 is equal to ToFloat64, but panic if there is an error.
func MustToFloat64(v interface{}) float64 {
	_v, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToComplex128 is equal to ToComplex128, but panic if there is an error.
func MustToComplex128(v interface{}) complex128 {
	_v, err := ToComplex128(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt is equal to ToInt, but panic if there is an error.
func MustToInt(v interface{}) int {
	_v, err := ToInt(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint is equal to ToUint, but panic if there is an error.
func MustToUint(v interface{}) uint {
	_v, err := ToUint(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt8 is equal to ToInt8, but panic if there is an error.
func MustToInt8(v interface{}) int8 {
	_v, err := ToInt8(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint8 is equal to ToUint8, but panic if there is an error.
func MustToUint8(v interface{}) uint8 {
	_v, err := ToUint8(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt16 is equal to ToInt16, but panic if there is an error.
func MustToInt16(v interface{}) int16 {
	_v, err := ToInt16(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint16 is equal to ToUint16, but panic if there is an error.
func MustToUint16(v interface{}) uint16 {
	_v, err := ToUint16(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt32 is equal to ToInt32, but panic if there is an error.
func MustToInt32(v interface{}) int32 {
	_v, err := ToInt32(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint32 is equal to ToUint32, but panic if there is an error.
func MustToUint32(v interface{}) uint32 {
	_v, err := ToUint32(v)
	if err != nil {
		panic(err)
	}
	return _v
}
