// Package slice gets a value from a slice and check whether a value exists in a slice.
//
// If the index is out-of-bounds, return the default value.
//
package slice

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errCannotSet    = errors.New("Can not be set")
	errNotSliceType = errors.New("The type is not slice or array")
	errInvalidIndex = errors.New("The invalid index")
)

func setValue(out interface{}, slice interface{}, index int, _default interface{}, yes bool) error {
	_out := reflect.Indirect(reflect.ValueOf(out))
	if !_out.CanSet() {
		return errCannotSet
	}

	_slice := reflect.ValueOf(slice)
	kind := _slice.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return errNotSliceType
	}

	var value interface{}
	if _slice.Len() > index {
		value = _slice.Index(index).Interface()
	} else if yes {
		value = _default
	} else {
		return errInvalidIndex
	}

	_out.Set(reflect.ValueOf(value))
	return nil
}

// SetValueWithDefault is same as SetValue, but if index >= len(slice),
// set the value of out to _default.
func SetValueWithDefault(out interface{}, slice interface{}, index int, _default interface{}) error {
	return setValue(out, slice, index, _default, true)
}

// SetValue sets the value of 'out' to 'slice[index]' and return nil.
//
// Return an error if the value of out can't be changed, that's, out need to be a pointer.
// Return an error if slice is not a slice type or index >= len(slice).
// Panic if when setting the value but the type is not matching.
func SetValue(out interface{}, slice interface{}, index int) error {
	return setValue(out, slice, index, nil, false)
}

// In returns true if value is in slice. Or false. Also reutrn false if value or
// slice is nil, or the length of slice is 0.
//
// Notice: the type of slice must be either []string or []interface{}, or panic.
// When it's []string, the type of value must be string, or panic.
//
// For []string, it doesn't use the reflect, but use for []interface{}, that's,
// reflect.DeepEqual(x, y).
func In(value interface{}, slice interface{}) bool {
	if value == nil || slice == nil {
		return false
	}

	switch slice.(type) {
	case []string:
		v := value.(string)
		for _, _v := range slice.([]string) {
			if v == _v {
				return true
			}
		}
	case []interface{}:
		for _, _v := range slice.([]interface{}) {
			if reflect.DeepEqual(value, _v) {
				return true
			}
		}
	default:
		panic(fmt.Errorf("only support the slice of []string and []interface"))
	}

	return false
}
