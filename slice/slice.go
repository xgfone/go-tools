package slice

import (
	"reflect"
)

func setValue(out interface{}, slice interface{}, index int, _default interface{}, yes bool) bool {
	_out := reflect.Indirect(reflect.ValueOf(out))
	if !_out.CanSet() {
		return false
	}

	_slice := reflect.ValueOf(slice)
	kind := _slice.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return false
	}

	var value interface{}
	if _slice.Len() > index {
		value = _slice.Index(index).Interface()
	} else if yes {
		value = _default
	} else {
		return false
	}

	_out.Set(reflect.ValueOf(value))
	return true
}

// Same as SetValue, but if index >= len(slice), set the value of out to _default,
// and return true.
func SetValueWithDefault(out interface{}, slice interface{}, index int, _default interface{}) bool {
	return setValue(out, slice, index, _default, true)
}

// Set the value of 'out' to 'slice[index]' and return true.
//
// Return false if the value of out can't be changed, that's, out need to be a pointer.
// Return false if slice is not a slice type or index >= len(slice).
// Panic for other cases.
func SetValue(out interface{}, slice interface{}, index int) bool {
	return setValue(out, slice, index, nil, false)
}
