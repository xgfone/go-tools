package values

// String does the best to convert the value whose index is i to string.
func (s Slice) String(i int) (v string, ok bool) {
	if len(s) <= i {
		return
	}
	return ToString(s[i])
}

// Bytes does the best to convert the value whose index is i to []byte.
func (s Slice) Bytes(i int) (v []byte, ok bool) {
	_v, ok := s.String(i)
	return []byte(_v), ok
}
