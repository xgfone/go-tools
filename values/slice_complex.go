package values

// Complex64 does the best to convert the value whose index is i to complex64.
func (s Slice) Complex64(i int) (complex64, bool) {
	_v, ok := s.Complex128(i)
	return complex64(_v), ok
}

// Complex128 does the best to convert the value whose index is i to complex128.
func (s Slice) Complex128(i int) (v complex128, ok bool) {
	if len(s) <= i {
		return
	}
	return ToComplex128(s[i])
}
