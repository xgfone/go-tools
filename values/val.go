// Get a value from the slice or the map.
//
// If failed, return (ZERO, false).
//
package values

import (
	"errors"
	"text/template"
)

var (
	FZERO32 float32
	FZERO64 float64

	ErrType = errors.New("The type is error")
)

// IsZero return true if the value is the ZERO value, or false.
//
// For bool, false is ZERO. For the integer, 0 is the ZERO. For the complex,
// it is ZEOR if all the real and the imag are 0.0. For chan, func, map, ptr,
// interface, nil is ZERO. For the slice, it is ZERO if the value is nil or has
// no element. For the array, it is ZERO if the value has no element.
// For string, the empty string is ZERO. For struct, it always is false.
func IsZero(v interface{}) bool {
	ok, _ := template.IsTrue(v)
	return !ok

	// _v := reflect.ValueOf(v)
	// switch _v.Kind() {
	// case reflect.Bool:
	// 	return !_v.Bool()
	// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	// 	return _v.Int() == 0
	// case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	// 	return _v.Uint() == 0
	// case reflect.Complex64, reflect.Complex128:
	// 	vv := _v.Complex()
	// 	if real(vv) == 0.0 && imag(vv) == 0.0 {
	// 		return true
	// 	}
	// 	return false
	// case reflect.Chan, reflect.Func, reflect.Map, reflect.Slice:
	// 	return (_v.IsNil() || _v.Len() == 0)
	// case reflect.Ptr:
	// 	return _v.IsNil()
	// case reflect.Interface:
	// 	return _v.IsNil()
	// case reflect.Array, reflect.String:
	// 	return _v.Len() == 0
	// case reflect.Struct:
	// 	return false
	// case reflect.Uintptr:
	// 	return _v.UnsafeAddr() == 0
	// case reflect.Invalid: // We think it as the interface nil
	// 	return true
	// }
	// return false
}
