package values

// Complex64 does the best to convert the value whose key is k to complex64.
func (m SMap) Complex64(k string) (v complex64, err error) {
	_v, err := m.Complex128(k)
	return complex64(_v), err
}

// IsComplex64 returns true when the type of the value whose key is k is complex64;
// or false.
func (m SMap) IsComplex64(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(complex64)
	return ok
}

// Complex128 does the best to convert the value whose key is k to complex128.
func (m SMap) Complex128(k string) (v complex128, err error) {
	_v, ok := m[k]
	if !ok {
		err = ErrNoKey
		return
	}
	return ToComplex128(_v)
}

// IsComplex128 returns true when the type of the value whose key is k is complex128;
// or false.
func (m SMap) IsComplex128(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(complex128)
	return ok
}
