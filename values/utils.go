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

// ToBool does the best to convert a certain value to bool.
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

// ToUInt64 is the alias of ToUint64. This function is DEPRECATED, and please
// use ToUint64.
func ToUInt64(_v interface{}) (v uint64, err error) {
	return ToUint64(_v)
}

// ToUint64 does the best to convert a certain value to uint64.
func ToUint64(_v interface{}) (v uint64, err error) {
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

// MustToUInt64 is the alias of MustToUint64. This function is DEPRECATED,
// and please use MustToUint64.
func MustToUInt64(v interface{}) uint64 {
	return MustToUint64(v)
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

// ToUInt is the alias of ToUint. This function is DEPRECATED, and please use
// ToUint.
func ToUInt(v interface{}) (uint, error) {
	return ToUint(v)
}

// ToUint does the best to convert a certain value to uint.
func ToUint(v interface{}) (uint, error) {
	_v, err := ToUint64(v)
	return uint(_v), err
}

// ToInt32 does the best to convert a certain value to int32.
func ToInt32(v interface{}) (int32, error) {
	_v, err := ToInt64(v)
	return int32(_v), err
}

// ToUInt32 is the alias of ToUint32. This function is DEPRECATED, and please
// use ToUint32.
func ToUInt32(v interface{}) (uint32, error) {
	return ToUint32(v)
}

// ToUint32 does the best to convert a certain value to uint32.
func ToUint32(v interface{}) (uint32, error) {
	_v, err := ToUint64(v)
	return uint32(_v), err
}
