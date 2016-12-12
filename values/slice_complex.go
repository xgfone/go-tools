package values

func (s Slice) Complex64(i int) (complex64, bool) {
	_v, ok := s.Complex128(i)
	return complex64(_v), ok
}

func (s Slice) MustComplex64(i int) complex64 {
	if v, ok := s.Complex64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Complex64WithDefault(i int, _default complex64) complex64 {
	if v, ok := s.Complex64(i); ok {
		return v
	}
	return _default
}

func (s Slice) Complex128(i int) (v complex128, ok bool) {
	if len(s) <= i {
		return
	}
	return ToComplex128(s[i])
}

func (s Slice) MustComplex128(i int) complex128 {
	if v, ok := s.Complex128(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Complex128WithDefault(i int, _default complex128) complex128 {
	if v, ok := s.Complex128(i); ok {
		return v
	}
	return _default
}
