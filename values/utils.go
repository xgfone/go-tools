package values

import (
	"fmt"
	"reflect"
	"strconv"
)

// In returns true if the key exists.
func In(m map[string]interface{}, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// Ins returns true if the key exists.
func Ins(m map[string]string, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// GetString returns the value of the key.
//
// The key must exist, or panic.
func GetString(m map[string]string, key string) string {
	if v, ok := m[key]; ok {
		return v
	}
	panic(fmt.Errorf("the key '%s' is missing", key))
}

// GetStringWithDefault returns the value of the key.
//
// Return the default if the key does not exist.
func GetStringWithDefault(m map[string]string, key string, _default string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return _default
}

// GetInterface returns the value of the key.
//
// The key must exist, or panic.
func GetInterface(m map[string]interface{}, key string) interface{} {
	if v, ok := m[key]; ok {
		return v
	}
	panic(fmt.Errorf("the key '%s' is missing", key))
}

// GetInterfaceWithDefault returns the value of the key.
//
// Return the default if the key does not exist.
func GetInterfaceWithDefault(m map[string]interface{}, key string, _default interface{}) interface{} {
	if v, ok := m[key]; ok {
		return v
	}
	return _default
}

// ToInt64 does the best to convert a certain value to int64.
func ToInt64(_v interface{}) (v int64, err error) {
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

// ToUInt64 does the best to convert a certain value to uint64.
func ToUInt64(_v interface{}) (v uint64, err error) {
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
		return strconv.ParseUint(_v.(string), 10, 64)
	default:
		err = fmt.Errorf("unknown type of %t", _v)
	}
	return
}

// MustToUInt64 must parse the value v to uint64, or panic.
//
// Notice: it will do the best to parse v.
func MustToUInt64(v interface{}) uint64 {
	_v, err := ToUInt64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToString does the best to convert a certain value to string.
func ToString(_v interface{}) (v string, err error) {
	switch _v.(type) {
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

// ToFloat64 does the best to convert a certain value to float64.
func ToFloat64(_v interface{}) (v float64, err error) {
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

// ToComplex128 does the best to convert a certain value to complex128.
func ToComplex128(_v interface{}) (v complex128, err error) {
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

// ToInt does the best to convert a certain value to int.
func ToInt(v interface{}) (int, error) {
	_v, err := ToInt64(v)
	return int(_v), err
}

// ToUInt does the best to convert a certain value to uint.
func ToUInt(v interface{}) (uint, error) {
	_v, err := ToUInt64(v)
	return uint(_v), err
}

// ToInt32 does the best to convert a certain value to int32.
func ToInt32(v interface{}) (int32, error) {
	_v, err := ToInt64(v)
	return int32(_v), err
}

// ToUInt32 does the best to convert a certain value to uint32.
func ToUInt32(v interface{}) (uint32, error) {
	_v, err := ToUInt64(v)
	return uint32(_v), err
}
