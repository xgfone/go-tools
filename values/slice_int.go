package values

// Int64 does the best to convert the value whose index is i to int64.
func (s Slice) Int64(i int) (int64, error) {
	if len(s) <= i {
		return 0, ErrOutOfLen
	}
	return ToInt64(s[i])
}

// IsInt64 returns true when the type of the ith value is int64; or false.
func (s Slice) IsInt64(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(int64)
	return ok
}

// Int does the best to convert the value whose index is i to int.
func (s Slice) Int(i int) (int, error) {
	v, err := s.Int64(i)
	return int(v), err
}

// IsInt returns true when the type of the ith value is int; or false.
func (s Slice) IsInt(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(int)
	return ok
}

// Rune does the best to convert the value whose index is i to rune.
func (s Slice) Rune(i int) (rune, error) {
	return s.Int32(i)
}

// IsRune returns true when the type of the ith value is rune; or false.
func (s Slice) IsRune(i int) bool {
	return s.IsInt32(i)
}

// Int8 does the best to convert the value whose index is i to int8.
func (s Slice) Int8(i int) (int8, error) {
	v, err := s.Int64(i)
	return int8(v), err
}

// IsInt8 returns true when the type of the ith value is int8; or false.
func (s Slice) IsInt8(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(int8)
	return ok
}

// Int16 does the best to convert the value whose index is i to int16.
func (s Slice) Int16(i int) (int16, error) {
	v, err := s.Int64(i)
	return int16(v), err
}

// IsInt16 returns true when the type of the ith value is int16; or false.
func (s Slice) IsInt16(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(int16)
	return ok
}

// Int32 does the best to convert the value whose index is i to int32.
func (s Slice) Int32(i int) (int32, error) {
	v, err := s.Int64(i)
	return int32(v), err
}

// IsInt32 returns true when the type of the ith value is int32; or false.
func (s Slice) IsInt32(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(int32)
	return ok
}
