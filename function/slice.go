package function

import (
	"fmt"
	"reflect"
)

var (
	// ErrNotPointer is returned when the value is not a pointer.
	ErrNotPointer = fmt.Errorf("the value is not a pointer")

	// ErrNotSliceOrArray is returned when the value is not a slice or array.
	ErrNotSliceOrArray = fmt.Errorf("The value is not slice or array")

	// ErrInvalidIndex is returned when the index exceeds over the length of
	// the slice.
	ErrInvalidIndex = fmt.Errorf("the index is exceeds the length of the slice")

	// ErrTypeNotCompatible is returned when the type is not compatible.
	ErrTypeNotCompatible = fmt.Errorf("the type is not compatible")
)

// GetSliceValue returns the ith element of slice.
//
// If slice is not a slice or array type, it will return ErrNotSliceOrArray.
// If the index i exceeds over the length of slice, that's, i>=len(slice),
// it will return ErrInvalidIndex.
//
// For []interface{}, []string and []int, they have already been optimized.
func GetSliceValue(slice interface{}, i int) (interface{}, error) {
	switch s := slice.(type) {
	case nil:
		return nil, ErrNotSliceOrArray
	case []string:
		if len(s) > i {
			return s[i], nil
		}
		return nil, ErrInvalidIndex
	case []int:
		if len(s) > i {
			return s[i], nil
		}
		return nil, ErrInvalidIndex
	case []interface{}:
		if len(s) > i {
			return s[i], nil
		}
		return nil, ErrInvalidIndex
	}

	s := reflect.ValueOf(slice)
	kind := s.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrNotSliceOrArray
	}

	if s.Len() > i {
		return s.Index(i).Interface(), nil
	}

	return nil, ErrInvalidIndex
}

func pullSliceValue(out interface{}, slice interface{}, index int,
	_default interface{}, yes bool) error {

	if out == nil {
		return ErrNotPointer
	}

	_out := reflect.ValueOf(out)
	kind := _out.Kind()
	if kind != reflect.Ptr && kind != reflect.UnsafePointer {
		return ErrNotPointer
	}
	_out = reflect.Indirect(_out)

	_slice := reflect.ValueOf(slice)
	kind = _slice.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return ErrNotSliceOrArray
	}

	var value interface{}
	if _slice.Len() > index {
		value = _slice.Index(index).Interface()
	} else if yes {
		value = _default
	} else {
		return ErrInvalidIndex
	}

	v := reflect.ValueOf(value)
	if v.Kind() != _out.Kind() {
		return ErrTypeNotCompatible
	}

	_out.Set(v)
	return nil
}

// PullSliceValueWithDefault is the same as PullSliceValue,
// but if index >= len(slice), set the value of out to _default.
func PullSliceValueWithDefault(out interface{}, slice interface{}, i int,
	_default interface{}) error {
	return pullSliceValue(out, slice, i, _default, true)
}

// PullSliceValue gets the ith element from the slice and puts it into the
// variable out, then returns nil if successfully.
//
// Return an error If out isn't a pointer, or slice isn't a slice or array type,
// or the length of slice is less than or equal to the index i, or the type of
// slice[i] and the underlying type of out are not compatible.
func PullSliceValue(out interface{}, slice interface{}, i int) error {
	return pullSliceValue(out, slice, i, nil, false)
}

// InSlice returns true if v is in slice, or returns false.
//
// It returns false if slice is not a slice type or the type is not compatible.
//
// For []interface{}, []string and []int, they have already been optimized.
func InSlice(v interface{}, slice interface{}) bool {
	// Optimize the types of []interface{}, []string and []int.
	switch s := slice.(type) {
	case nil:
		return false
	case []interface{}:
		for _, _v := range s {
			if reflect.DeepEqual(v, _v) {
				return true
			}
		}
		return false
	case []string:
		if _v, ok := v.(string); ok {
			for _, _s := range s {
				if _v == _s {
					return true
				}
			}
		}
		return false
	case []int:
		if _v, ok := v.(int); ok {
			for _, _s := range s {
				if _v == _s {
					return true
				}
			}
		}
		return false
	}

	s := reflect.ValueOf(slice)
	kind := s.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return false
	}

	slen := s.Len()
	for i := 0; i < slen; i++ {
		if reflect.DeepEqual(v, s.Index(i).Interface()) {
			return true
		}
	}
	return false
}
