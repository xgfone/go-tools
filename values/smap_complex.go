package values

// Complex64 is the same as UInt64, but complex64.
func (m SMap) Complex64(k string) (v complex64, ok bool) {
	_v, ok := m.Complex128(k)
	return complex64(_v), ok
}

// Complex64WithDefault is the same as UInt64WithDefault, but complex64.
func (m SMap) Complex64WithDefault(k string, _default complex64) complex64 {
	if v, ok := m.Complex64(k); ok {
		return v
	}
	return _default
}

// MustComplex64 is the same as MustUInt64, but complex64.
func (m SMap) MustComplex64(k string) complex64 {
	if v, ok := m.Complex64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Complex128 is the same as UInt64, but complex128.
func (m SMap) Complex128(k string) (v complex128, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToComplex128(_v)
}

// Complex128WithDefault is the same as UInt64WithDefault, but complex128.
func (m SMap) Complex128WithDefault(k string, _default complex128) complex128 {
	if v, ok := m.Complex128(k); ok {
		return v
	}
	return _default
}

// MustComplex128 is the same as MustUInt64, but complex128.
func (m SMap) MustComplex128(k string) complex128 {
	if v, ok := m.Complex128(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
