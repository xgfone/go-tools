package values

// Float32 is the same as UInt64, but float32.
func (m SMap) Float32(k string) (v float32, ok bool) {
	_v, ok := m.Float64(k)
	return float32(_v), ok
}

// Float64 is the same as UInt64, but float64.
func (m SMap) Float64(k string) (v float64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToFloat64(_v)
}
