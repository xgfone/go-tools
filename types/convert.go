package types

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	// ErrNotSliceOrArray is returned when the value is not a slice.
	ErrNotSliceOrArray = fmt.Errorf("the value is not a slice or array")

	// ErrNotMap is returned when the value is not a map.
	ErrNotMap = fmt.Errorf("the value is not a map")

	// ErrNotString is returned when the type of the key is not string.
	ErrNotString = fmt.Errorf("the type of the key is not string")
)

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

// MustToSlice must parse the value v to []interface{}, or panic.
//
// Notice: it will do the best to parse v.
func MustToSlice(v interface{}) []interface{} {
	_v, err := ToSlice(v)
	if err != nil {
		panic(err)
	}
	return _v
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

// MustToMap must parse the value v to map[string]interface{}, or panic.
//
// Notice: it will do the best to parse v.
func MustToMap(v interface{}) map[string]interface{} {
	_v, err := ToMap(v)
	if err != nil {
		panic(err)
	}
	return _v
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

// ToBool does the best to convert any certain value to bool.
//
// When the value is string, for "t", "T", "1", "true", "True", "TRUE",
// it's true, for "f", "F", "0", "false", "False", "FALSE", "", it's false.
//
// For other types, if the value is ZERO of the type, it's false. Or it's true.
func ToBool(v interface{}) (bool, error) {
	switch _v := v.(type) {
	case string:
		switch _v {
		case "t", "T", "1", "true", "True", "TRUE":
			return true, nil
		case "f", "F", "0", "false", "False", "FALSE", "":
			return false, nil
		default:
			return false, fmt.Errorf("unrecognized bool string: %s", _v)
		}
	}
	return !IsZero(v), nil
}

// MustToBool must parse the value v to bool, or panic.
//
// Notice: it will do the best to parse v.
func MustToBool(v interface{}) bool {
	_v, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToInt64 does the best to convert any certain value to int64.
func ToInt64(_v interface{}) (v int64, err error) {
	switch _v.(type) {
	case nil:
	case complex64, complex128:
		v = int64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = int64(bool2Int64(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Int()
	case uint, uint8, uint16, uint32, uint64:
		v = int64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = int64(reflect.ValueOf(_v).Float())
	case string:
		return strconv.ParseInt(_v.(string), 10, 64)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToInt64 must parse the value v to int64, or panic.
//
// Notice: it will do the best to parse v.
func MustToInt64(v interface{}) int64 {
	_v, err := ToInt64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToUint64 does the best to convert any certain value to uint64.
func ToUint64(_v interface{}) (v uint64, err error) {
	switch _v.(type) {
	case nil:
	case complex64, complex128:
		v = uint64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = uint64(bool2Int64(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Uint()
	case uint, uint8, uint16, uint32, uint64:
		v = uint64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = uint64(reflect.ValueOf(_v).Float())
	case string:
		return strconv.ParseUint(_v.(string), 10, 64)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToUint64 must parse the value v to uint64, or panic.
//
// Notice: it will do the best to parse v.
func MustToUint64(v interface{}) uint64 {
	_v, err := ToUint64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToString does the best to convert any certain value to string.
func ToString(_v interface{}) (v string, err error) {
	switch _v.(type) {
	case nil:
	case string:
		v = _v.(string)
	case []byte:
		v = string(_v.([]byte))
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		v = fmt.Sprintf("%d", _v)
	case float32, float64:
		v = fmt.Sprintf("%f", _v)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToString must parse the value v to string, or panic.
//
// Notice: it will do the best to parse v.
func MustToString(v interface{}) string {
	_v, err := ToString(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToFloat64 does the best to convert any certain value to float64.
func ToFloat64(_v interface{}) (v float64, err error) {
	switch _v.(type) {
	case nil:
	case complex64, complex128:
		v = float64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = float64(bool2Int64(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = float64(reflect.ValueOf(_v).Int())
	case uint, uint8, uint16, uint32, uint64:
		v = float64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = reflect.ValueOf(_v).Float()
	case string:
		return strconv.ParseFloat(_v.(string), 64)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToFloat64 must parse the value v to float64, or panic.
//
// Notice: it will do the best to parse v.
func MustToFloat64(v interface{}) float64 {
	_v, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToComplex128 does the best to convert any certain value to complex128.
func ToComplex128(_v interface{}) (v complex128, err error) {
	switch _v.(type) {
	case nil:
	case complex64, complex128:
		v = complex128(reflect.ValueOf(_v).Complex())
	case bool:
		v = complex(float64(bool2Int64(_v.(bool))), 0)
	case int, int8, int16, int32, int64:
		v = complex(float64(reflect.ValueOf(_v).Int()), 0)
	case uint, uint8, uint16, uint32, uint64:
		v = complex(float64(reflect.ValueOf(_v).Uint()), 0)
	case float32, float64:
		v = complex(reflect.ValueOf(_v).Float(), 0)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToComplex128 must parse the value v to complex128, or panic.
//
// Notice: it will do the best to parse v.
func MustToComplex128(v interface{}) complex128 {
	_v, err := ToComplex128(v)
	if err != nil {
		panic(err)
	}
	return _v
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
