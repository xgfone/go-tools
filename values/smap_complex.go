package values

func (m SMap) Complex64(k string) (v complex64, ok bool) {
	_v, ok := m.Complex128(k)
	return complex64(_v), ok
}

func (m SMap) Complex64WithDefault(k string, _default complex64) complex64 {
	if v, ok := m.Complex64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustComplex64(k string) complex64 {
	if v, ok := m.Complex64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Complex128(k string) (v complex128, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToComplex128(_v)
}

func (m SMap) Complex128WithDefault(k string, _default complex128) complex128 {
	if v, ok := m.Complex128(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustComplex128(k string) complex128 {
	if v, ok := m.Complex128(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
