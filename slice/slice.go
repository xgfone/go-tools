// Package slice gets a value from a slice and check whether a value exists in a slice.
//
// If the index is out-of-bounds, return the default value.
//
package slice

import (
	"errors"
	"reflect"

	"github.com/xgfone/go-tools/compare"
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
// The type of value must be consistent with the type of the element of slice.
// Or panic. If the type is the customizable struct, it MUST implement the interface
// Comparer in the package "github.com/xgfone/go-tools/compare".
func In(value interface{}, slice interface{}) bool {
	if value == nil || slice == nil {
		return false
	}

	stype := reflect.ValueOf(slice)
	if stype.Kind() == reflect.Ptr {
		stype = stype.Elem()
	}

	if stype.Kind() != reflect.Array && stype.Kind() != reflect.Slice {
		panic("The second argument is not a slice or an array")
	}

	slen := stype.Len()
	if slen == 0 {
		return false
	}

	vv := reflect.ValueOf(value)
	if stype.Index(0).Kind() != reflect.ValueOf(value).Kind() {
		panic("The type of value must be consistent with the type of the element of slice")
	}

	for i := 0; i < slen; i++ {
		v1 := vv.Interface()
		v2 := stype.Index(i).Interface()
		if compare.EQ(v1, v2) {
			return true
		}
	}

	return false
}
