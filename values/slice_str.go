package values

// String does the best to convert the value whose index is i to string.
func (s Slice) String(i int) (v string, ok bool) {
	if len(s) <= i {
		return
	}
	return ToString(s[i])
}

// IsString returns true when the type of the ith value is string; or false.
func (s Slice) IsString(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(string)
	return ok
}

// Bytes does the best to convert the value whose index is i to []byte.
func (s Slice) Bytes(i int) (v []byte, ok bool) {
	_v, ok := s.String(i)
	return []byte(_v), ok
}

// IsBytes returns true when the type of the ith value is []byte; or false.
func (s Slice) IsBytes(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].([]byte)
	return ok
}
