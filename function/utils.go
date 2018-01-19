package function

import (
	"reflect"
)

// IsNil returns true when v is the nil value.
//
// var
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice,
		reflect.Interface:
		if rv.IsNil() {
			return true
		}
	}
	return false
}
