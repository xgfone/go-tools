package values

import "reflect"

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
	switch v.(type) {
	case []interface{}:
		return Slice(v.([]interface{}))
	case Slice:
		return v.(Slice)
	default:
		_v, _ := ConvertToSlice(v)
		return _v
	}
}

// ConvertToSlice converts any slices to Slice.
//
// Return nil if it's not a slice, or it's nil or has no elements.
func ConvertToSlice(v interface{}) (Slice, bool) {
	_v := reflect.ValueOf(v)
	if !_v.IsValid() || _v.Kind() != reflect.Slice {
		return nil, false
	}

	_len := _v.Len()
	results := make(Slice, _len)
	for i := 0; i < _len; i++ {
		results[i] = _v.Index(i).Interface()
	}
	return results, true
}
