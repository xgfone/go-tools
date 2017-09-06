package values

// Complex64 does the best to convert the value whose index is i to complex64.
func (s Slice) Complex64(i int) (complex64, error) {
	_v, err := s.Complex128(i)
	return complex64(_v), err
}

// IsComplex64 returns true when the type of the ith value is complex64; or false.
func (s Slice) IsComplex64(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(complex64)
	return ok
}

// Complex128 does the best to convert the value whose index is i to complex128.
func (s Slice) Complex128(i int) (v complex128, err error) {
	if len(s) <= i {
		err = ErrOutOfLen
		return
	}
	return ToComplex128(s[i])
}

// IsComplex128 returns true when the type of the ith value is complex128; or false.
func (s Slice) IsComplex128(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(complex128)
	return ok
}
