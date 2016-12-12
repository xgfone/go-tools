package values

func (m SMap) Float32(k string) (v float32, ok bool) {
	_v, ok := m.Float64(k)
	return float32(_v), ok
}

func (m SMap) Float32WithDefault(k string, _default float32) float32 {
	if v, ok := m.Float32(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustFloat32(k string) float32 {
	if v, ok := m.Float32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Float64(k string) (v float64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToFloat64(_v)
}

func (m SMap) Float64WithDefault(k string, _default float64) float64 {
	if v, ok := m.Float64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustFloat64(k string) float64 {
	if v, ok := m.Float64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
