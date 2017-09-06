package values

// String does the best to convert the value whose index is i to string.
func (s Slice) String(i int) (v string, err error) {
	if len(s) <= i {
		return "", ErrOutOfLen
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
func (s Slice) Bytes(i int) (v []byte, err error) {
	_v, err := s.String(i)
	return []byte(_v), err
}

// IsBytes returns true when the type of the ith value is []byte; or false.
func (s Slice) IsBytes(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].([]byte)
	return ok
}
