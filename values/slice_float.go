package values

// Float32 is the same as Int64, but float32.
func (s Slice) Float32(i int) (float32, bool) {
	_v, ok := s.Float64(i)
	return float32(_v), ok
}

// MustFloat32 is the same as MustInt64, but float32.
func (s Slice) MustFloat32(i int) float32 {
	if v, ok := s.Float32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Float32WithDefault is the same as Int64WithDefault, but float32.
func (s Slice) Float32WithDefault(i int, _default float32) float32 {
	if v, ok := s.Float32(i); ok {
		return v
	}
	return _default
}

// Float64 is the same as Int64, but float64.
func (s Slice) Float64(i int) (float64, bool) {
	if len(s) <= i {
		return FZERO64, false
	}
	return ToFloat64(s[i])
}

// MustFloat64 is the same as MustInt64, but float64.
func (s Slice) MustFloat64(i int) float64 {
	if v, ok := s.Float64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Float64WithDefault is the same as Int64WithDefault, but float64.
func (s Slice) Float64WithDefault(i int, _default float64) float64 {
	if v, ok := s.Float64(i); ok {
		return v
	}
	return _default
}
