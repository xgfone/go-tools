package values

// Float32 is the same as UInt64, but float32.
func (m SMap) Float32(k string) (v float32, err error) {
	_v, err := m.Float64(k)
	return float32(_v), err
}

// IsFloat32 returns true when the type of the value whose key is k is float32;
// or false.
func (m SMap) IsFloat32(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(float32)
	return ok
}

// Float64 is the same as UInt64, but float64.
func (m SMap) Float64(k string) (v float64, err error) {
	_v, ok := m[k]
	if !ok {
		err = ErrNoKey
		return
	}
	return ToFloat64(_v)
}

// IsFloat64 returns true when the type of the value whose key is k is float64;
// or false.
func (m SMap) IsFloat64(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(float64)
	return ok
}
