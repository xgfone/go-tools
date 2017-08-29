package values

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt64 does the best to convert a certain value to int64.
func ToInt64(_v interface{}) (v int64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = int64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = int64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Int()
	case uint, uint8, uint16, uint32, uint64:
		v = int64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = int64(reflect.ValueOf(_v).Float())
	case string:
		if vv, err := strconv.ParseInt(_v.(string), 10, 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

// MustToInt64 must parse the value v to int64, or panic.
//
// Notice: it will do the best to parse v.
func MustToInt64(v interface{}) int64 {
	if _v, ok := ToInt64(v); ok {
		return _v
	}
	panic(fmt.Errorf("can't parse the value to int64"))
}

// ToUInt64 does the best to convert a certain value to uint64.
func ToUInt64(_v interface{}) (v uint64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = uint64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = uint64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Uint()
	case uint, uint8, uint16, uint32, uint64:
		v = uint64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = uint64(reflect.ValueOf(_v).Float())
	case string:
		if vv, err := strconv.ParseUint(_v.(string), 10, 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

// MustToUInt64 must parse the value v to uint64, or panic.
//
// Notice: it will do the best to parse v.
func MustToUInt64(v interface{}) uint64 {
	if _v, ok := ToUInt64(v); ok {
		return _v
	}
	panic(fmt.Errorf("can't parse the value to uint64"))
}

// ToString does the best to convert a certain value to string.
func ToString(_v interface{}) (v string, ok bool) {
	ok = true
	switch _v.(type) {
	case string:
		v = _v.(string)
	case []byte:
		v = string(_v.([]byte))
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		v = fmt.Sprintf("%v", _v)
	default:
		ok = false
	}
	return
}

// MustToString must parse the value v to string, or panic.
//
// Notice: it will do the best to parse v.
func MustToString(v interface{}) string {
	if _v, ok := ToString(v); ok {
		return _v
	}
	panic(fmt.Errorf("can't parse the value to string"))
}

// ToFloat64 does the best to convert a certain value to float64.
func ToFloat64(_v interface{}) (v float64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = float64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = float64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = float64(reflect.ValueOf(_v).Int())
	case uint, uint8, uint16, uint32, uint64:
		v = float64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = reflect.ValueOf(_v).Float()
	case string:
		if vv, err := strconv.ParseFloat(_v.(string), 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

// MustToFloat64 must parse the value v to float64, or panic.
//
// Notice: it will do the best to parse v.
func MustToFloat64(v interface{}) float64 {
	if _v, ok := ToFloat64(v); ok {
		return _v
	}
	panic(fmt.Errorf("can't parse the value to float64"))
}

// ToComplex128 does the best to convert a certain value to complex128.
func ToComplex128(_v interface{}) (v complex128, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = complex128(reflect.ValueOf(_v).Complex())
	case bool:
		v = complex(float64(Bool2Int(_v.(bool))), 0)
	case int, int8, int16, int32, int64:
		v = complex(float64(reflect.ValueOf(_v).Int()), 0)
	case uint, uint8, uint16, uint32, uint64:
		v = complex(float64(reflect.ValueOf(_v).Uint()), 0)
	case float32, float64:
		v = complex(reflect.ValueOf(_v).Float(), 0)
	default:
		ok = false
	}
	return
}

// MustToComplex128 must parse the value v to complex128, or panic.
//
// Notice: it will do the best to parse v.
func MustToComplex128(v interface{}) complex128 {
	if _v, ok := ToComplex128(v); ok {
		return _v
	}
	panic(fmt.Errorf("can't parse the value to complex128"))
}
