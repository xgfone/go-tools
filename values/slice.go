package values

type Slice []interface{}

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
		return nil
	}
}
