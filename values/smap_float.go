package values

// Float32 is the same as UInt64, but float32.
func (m SMap) Float32(k string) (v float32, ok bool) {
	_v, ok := m.Float64(k)
	return float32(_v), ok
}

// Float32WithDefault is the same as UInt64WithDefault, but float32.
func (m SMap) Float32WithDefault(k string, _default float32) float32 {
	if v, ok := m.Float32(k); ok {
		return v
	}
	return _default
}

// MustFloat32 is the same as MustUInt64, but float32.
func (m SMap) MustFloat32(k string) float32 {
	if v, ok := m.Float32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Float64 is the same as UInt64, but float64.
func (m SMap) Float64(k string) (v float64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToFloat64(_v)
}

// Float64WithDefault is the same as UInt64WithDefault, but float64.
func (m SMap) Float64WithDefault(k string, _default float64) float64 {
	if v, ok := m.Float64(k); ok {
		return v
	}
	return _default
}

// MustFloat64 is the same as MustUInt64, but float64.
func (m SMap) MustFloat64(k string) float64 {
	if v, ok := m.Float64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
