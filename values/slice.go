package values

type Slice []interface{}

func ToSlice(v interface{}) (Slice, error) {
	switch v.(type) {
	case []interface{}:
		return Slice(v.([]interface{})), nil
	case Slice:
		return v.(Slice), nil
	default:
		return nil, ErrTypeOrIndex
	}
}
