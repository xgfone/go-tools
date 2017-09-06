package values

import (
	"fmt"
	"reflect"
)

// Slice is a type alias of []interface{}.
type Slice []interface{}

// NewSlice returns a new Slice.
func NewSlice(i int) Slice {
	return make(Slice, i)
}

// ToSlice converts the type of []interface{} or Slice to Slice.
//
// Return nil if the type is not either.
func ToSlice(v interface{}) Slice {
	_v, _ := toSlice(v)
	return _v
}

// MustToSlice must parse the argument v to Slice, or panic.
func MustToSlice(v interface{}) Slice {
	_v, err := toSlice(v)
	if err != nil {
		panic(err)
	}
	return _v
}

func toSlice(v interface{}) (Slice, error) {
	switch v.(type) {
	case []interface{}:
		return Slice(v.([]interface{})), nil
	case Slice:
		return v.(Slice), nil
	default:
		return ConvertToSlice(v)
	}
}

// ConvertToSlice converts any slices to Slice.
//
// Return nil if it's not a slice, or it's nil or has no elements.
//
// Notice: Slice(nil) is not a valid Slice.
func ConvertToSlice(v interface{}) (Slice, error) {
	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the value is not valid or a slice")
	}

	_len := _v.Len()
	results := make(Slice, _len)
	for i := 0; i < _len; i++ {
		results[i] = _v.Index(i).Interface()
	}
	return results, nil
}
