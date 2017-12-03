package types

import (
	"fmt"
)

// CheckMapType checks the type of m[k] is t, and return the value if yes,
// or return an error if not.
//
// If m is nil, return an error.
//
// The function uses VerifyType to verify the type, that's, VerifyType(m[k], t),
// so for t, see VerifyType.
func CheckMapType(m map[string]interface{}, k, t string) (interface{}, error) {
	if len(m) == 0 {
		return nil, fmt.Errorf("the map is nil")
	}
	value := m[k]
	if value == nil {
		return nil, fmt.Errorf("the value of the key[%s] in map is nil", k)
	}
	if !VerifyType(value, t) {
		return nil, fmt.Errorf("the value of the key[%s] in map is not %s",
			k, stype2type[k])
	}
	return value, nil
}

// MapBool check whether m[key] is the bool type, and return the bool value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapBool(m map[string]interface{}, key string) (v bool, err error) {
	_v, err := CheckMapType(m, key, "bool")
	if err == nil {
		v = _v.(bool)
	}
	return
}

// MapString check whether m[key] is the string type, and return the string
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapString(m map[string]interface{}, key string) (v string, err error) {
	_v, err := CheckMapType(m, key, "string")
	if err == nil {
		v = _v.(string)
	}
	return
}

// MapInt check whether m[key] is the int type, and return the int value if yes,
// or return an error if not.
//
// If m is nil, return an error.
func MapInt(m map[string]interface{}, key string) (v int, err error) {
	_v, err := CheckMapType(m, key, "int")
	if err == nil {
		v = _v.(int)
	}
	return
}

// MapInt8 check whether m[key] is the int8 type, and return the int8 value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapInt8(m map[string]interface{}, key string) (v int8, err error) {
	_v, err := CheckMapType(m, key, "int8")
	if err == nil {
		v = _v.(int8)
	}
	return
}

// MapInt16 check whether m[key] is the int16 type, and return the int16 value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapInt16(m map[string]interface{}, key string) (v int16, err error) {
	_v, err := CheckMapType(m, key, "int16")
	if err == nil {
		v = _v.(int16)
	}
	return
}

// MapInt32 check whether m[key] is the int32 type, and return the int32 value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapInt32(m map[string]interface{}, key string) (v int32, err error) {
	_v, err := CheckMapType(m, key, "int32")
	if err == nil {
		v = _v.(int32)
	}
	return
}

// MapInt64 check whether m[key] is the int64 type, and return the int64 value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapInt64(m map[string]interface{}, key string) (v int64, err error) {
	_v, err := CheckMapType(m, key, "int64")
	if err == nil {
		v = _v.(int64)
	}
	return
}

// MapUint check whether m[key] is the uint type, and return the uint value if yes,
// or return an error if not.
//
// If m is nil, return an error.
func MapUint(m map[string]interface{}, key string) (v uint, err error) {
	_v, err := CheckMapType(m, key, "uint")
	if err == nil {
		v = _v.(uint)
	}
	return
}

// MapUint8 check whether m[key] is the uint8 type, and return the uint8 value
// if yes, or return an error if not.
//
// If m is nil, return an error.
func MapUint8(m map[string]interface{}, key string) (v uint8, err error) {
	_v, err := CheckMapType(m, key, "uint8")
	if err == nil {
		v = _v.(uint8)
	}
	return
}

// MapUint16 check whether m[key] is the uint16 type, and return the uint16
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapUint16(m map[string]interface{}, key string) (v uint16, err error) {
	_v, err := CheckMapType(m, key, "uint16")
	if err == nil {
		v = _v.(uint16)
	}
	return
}

// MapUint32 check whether m[key] is the uint32 type, and return the uint32
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapUint32(m map[string]interface{}, key string) (v uint32, err error) {
	_v, err := CheckMapType(m, key, "uint32")
	if err == nil {
		v = _v.(uint32)
	}
	return
}

// MapUint64 check whether m[key] is the uint64 type, and return the uint64
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapUint64(m map[string]interface{}, key string) (v uint64, err error) {
	_v, err := CheckMapType(m, key, "uint64")
	if err == nil {
		v = _v.(uint64)
	}
	return
}

// MapFloat32 check whether m[key] is the Float32 type, and return the Float32
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapFloat32(m map[string]interface{}, key string) (v float32, err error) {
	_v, err := CheckMapType(m, key, "float32")
	if err == nil {
		v = _v.(float32)
	}
	return
}

// MapFloat64 check whether m[key] is the float64 type, and return the float64
// value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapFloat64(m map[string]interface{}, key string) (v float64, err error) {
	_v, err := CheckMapType(m, key, "float64")
	if err == nil {
		v = _v.(float64)
	}
	return
}

// MapComplex64 check whether m[key] is the complex64 type, and return the
// complex64 value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapComplex64(m map[string]interface{}, key string) (v complex64, err error) {
	_v, err := CheckMapType(m, key, "complex64")
	if err == nil {
		v = _v.(complex64)
	}
	return
}

// MapComplex128 check whether m[key] is the complex128 type, and return the
// complex128 value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapComplex128(m map[string]interface{}, key string) (v complex128, err error) {
	_v, err := CheckMapType(m, key, "complex128")
	if err == nil {
		v = _v.(complex128)
	}
	return
}

// MapMap check whether m[key] is the map[string]interface{} type, and return
// the map[string]interface{} value if yes, or return an error if not.
//
// If m is nil, return an error.
func MapMap(m map[string]interface{}, key string) (v map[string]interface{}, err error) {
	_v, err := CheckMapType(m, key, "string2interface")
	if err == nil {
		v = _v.(map[string]interface{})
	}
	return
}
