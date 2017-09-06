package values

// Float32 does the best to convert the value whose index is i to float32.
func (s Slice) Float32(i int) (float32, error) {
	_v, err := s.Float64(i)
	return float32(_v), err
}

// IsFloat32 returns true when the type of the ith value is float32; or false.
func (s Slice) IsFloat32(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(float32)
	return ok
}

// Float64 does the best to convert the value whose index is i to float64.
func (s Slice) Float64(i int) (float64, error) {
	if len(s) <= i {
		return 0, ErrOutOfLen
	}
	return ToFloat64(s[i])
}

// IsFloat64 returns true when the type of the ith value is float64; or false.
func (s Slice) IsFloat64(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(float64)
	return ok
}
