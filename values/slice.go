package values

type Slice []interface{}

// ToSlice converts v to the type of Slice.
//
// Return nil if failed.
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
