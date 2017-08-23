package values

// Int64 does the best to convert the value whose index is i to int64.
func (s Slice) Int64(i int) (int64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToInt64(s[i])
}

// Int does the best to convert the value whose index is i to int.
func (s Slice) Int(i int) (int, bool) {
	v, ok := s.Int64(i)
	return int(v), ok
}

// Rune does the best to convert the value whose index is i to rune.
func (s Slice) Rune(i int) (rune, bool) {
	return s.Int32(i)
}

// Int8 does the best to convert the value whose index is i to int8.
func (s Slice) Int8(i int) (int8, bool) {
	v, ok := s.Int64(i)
	return int8(v), ok
}

// Int16 does the best to convert the value whose index is i to int16.
func (s Slice) Int16(i int) (int16, bool) {
	v, ok := s.Int64(i)
	return int16(v), ok
}

// Int32 does the best to convert the value whose index is i to int32.
func (s Slice) Int32(i int) (int32, bool) {
	v, ok := s.Int64(i)
	return int32(v), ok
}
