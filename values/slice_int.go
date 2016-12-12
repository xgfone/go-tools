package values

func (s Slice) Int64(i int) (int64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToInt64(s[i])
}

func (s Slice) MustInt64(i int) int64 {
	if v, ok := s.Int64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int64WithDefault(i int, _default int64) int64 {
	if v, ok := s.Int64(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int(i int) (int, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt(i int) int {
	if v, ok := s.Int(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) IntWithDefault(i int, _default int) int {
	if v, ok := s.Int(i); ok {
		return v
	}
	return _default
}

func (s Slice) Rune(i int) (rune, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(rune); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustRune(i int) rune {
	if v, ok := s.Rune(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) RuneWithDefault(i int, _default rune) rune {
	if v, ok := s.Rune(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int8(i int) (int8, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt8(i int) int8 {
	if v, ok := s.Int8(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int8WithDefault(i int, _default int8) int8 {
	if v, ok := s.Int8(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int16(i int) (int16, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt16(i int) int16 {
	if v, ok := s.Int16(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int16WithDefault(i int, _default int16) int16 {
	if v, ok := s.Int16(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int32(i int) (int32, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt32(i int) int32 {
	if v, ok := s.Int32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int32WithDefault(i int, _default int32) int32 {
	if v, ok := s.Int32(i); ok {
		return v
	}
	return _default
}
