package values

// Float32 does the best to convert the value whose index is i to float32.
func (s Slice) Float32(i int) (float32, bool) {
	_v, ok := s.Float64(i)
	return float32(_v), ok
}

// Float64 does the best to convert the value whose index is i to float64.
func (s Slice) Float64(i int) (float64, bool) {
	if len(s) <= i {
		return FZERO64, false
	}
	return ToFloat64(s[i])
}
