package values

func (s Slice) Float32(i int) (float32, bool) {
	_v, ok := s.Float64(i)
	return float32(_v), ok
}

func (s Slice) MustFloat32(i int) float32 {
	if v, ok := s.Float32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Float32WithDefault(i int, _default float32) float32 {
	if v, ok := s.Float32(i); ok {
		return v
	}
	return _default
}

func (s Slice) Float64(i int) (float64, bool) {
	if len(s) <= i {
		return FZERO64, false
	}
	return ToFloat64(s[i])
}

func (s Slice) MustFloat64(i int) float64 {
	if v, ok := s.Float64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Float64WithDefault(i int, _default float64) float64 {
	if v, ok := s.Float64(i); ok {
		return v
	}
	return _default
}
