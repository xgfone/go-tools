package values

// Complex64 does the best to convert the value whose key is k to complex64.
func (m SMap) Complex64(k string) (v complex64, ok bool) {
	_v, ok := m.Complex128(k)
	return complex64(_v), ok
}

// Complex128 does the best to convert the value whose key is k to complex128.
func (m SMap) Complex128(k string) (v complex128, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToComplex128(_v)
}
