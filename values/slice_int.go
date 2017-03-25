package values

// Int64 converts the value whose key is k to int64.
func (s Slice) Int64(i int) (int64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToInt64(s[i])
}

// MustInt64 is the same as Int64, but panic when failed.
func (s Slice) MustInt64(i int) int64 {
	if v, ok := s.Int64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Int64WithDefault is the same as Int64, but return the default value,
// not ZERO, when failed.
func (s Slice) Int64WithDefault(i int, _default int64) int64 {
	if v, ok := s.Int64(i); ok {
		return v
	}
	return _default
}

// Int is the same as Int64, but int.
func (s Slice) Int(i int) (int, bool) {
	v, ok := s.Int64(i)
	return int(v), ok
}

// MustInt is the same as MustInt64, but int.
func (s Slice) MustInt(i int) int {
	if v, ok := s.Int(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// IntWithDefault is the Int64WithDefault, but int.
func (s Slice) IntWithDefault(i int, _default int) int {
	if v, ok := s.Int(i); ok {
		return v
	}
	return _default
}

// Rune is the same as Int64, but rune.
func (s Slice) Rune(i int) (rune, bool) {
	return s.Int32(i)
}

// MustRune is the same as MustInt64, but rune.
func (s Slice) MustRune(i int) rune {
	if v, ok := s.Rune(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// RuneWithDefault is the same as Int64WithDefault, but rune.
func (s Slice) RuneWithDefault(i int, _default rune) rune {
	if v, ok := s.Rune(i); ok {
		return v
	}
	return _default
}

// Int8 is the same as Int64, but int8.
func (s Slice) Int8(i int) (int8, bool) {
	v, ok := s.Int64(i)
	return int8(v), ok
}

// MustInt8 is the same as MustInt64, but int8.
func (s Slice) MustInt8(i int) int8 {
	if v, ok := s.Int8(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Int8WithDefault is the same as Int64WithDefault, but int8.
func (s Slice) Int8WithDefault(i int, _default int8) int8 {
	if v, ok := s.Int8(i); ok {
		return v
	}
	return _default
}

// Int16 is the same as Int64, but int16.
func (s Slice) Int16(i int) (int16, bool) {
	v, ok := s.Int64(i)
	return int16(v), ok
}

// MustInt16 is the same as MustInt64, but int16.
func (s Slice) MustInt16(i int) int16 {
	if v, ok := s.Int16(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Int16WithDefault is the same as Int64WithDefault, but int16.
func (s Slice) Int16WithDefault(i int, _default int16) int16 {
	if v, ok := s.Int16(i); ok {
		return v
	}
	return _default
}

// Int32 is the same as Int64, but int32.
func (s Slice) Int32(i int) (int32, bool) {
	v, ok := s.Int64(i)
	return int32(v), ok
}

// MustInt32 is the same as MustInt64, but int32.
func (s Slice) MustInt32(i int) int32 {
	if v, ok := s.Int32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Int32WithDefault is the same as Int64WithDefault, but int32.
func (s Slice) Int32WithDefault(i int, _default int32) int32 {
	if v, ok := s.Int32(i); ok {
		return v
	}
	return _default
}
