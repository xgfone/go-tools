package values

func (m SMap) Int64(k string) (v int64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToInt64(_v)
}

func (m SMap) Int64WithDefault(k string, _default int64) int64 {
	if v, ok := m.Int64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt64(k string) int64 {
	if v, ok := m.Int64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Rune(k string) (v rune, ok bool) {
	return m.Int32(k)
}

func (m SMap) RuneWithDefault(k string, _default rune) rune {
	if v, ok := m.Rune(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustRune(k string) rune {
	if v, ok := m.Rune(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int(k string) (v int, ok bool) {
	_v, ok := m.Int64(k)
	return int(_v), ok

}

func (m SMap) IntWithDefault(k string, _default int) int {
	if v, ok := m.Int(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt(k string) int {
	if v, ok := m.Int(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int8(k string) (v int8, ok bool) {
	_v, ok := m.Int64(k)
	return int8(_v), ok
}

func (m SMap) Int8WithDefault(k string, _default int8) int8 {
	if v, ok := m.Int8(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt8(k string) int8 {
	if v, ok := m.Int8(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int16(k string) (v int16, ok bool) {
	_v, ok := m.Int64(k)
	return int16(_v), ok
}

func (m SMap) Int16WithDefault(k string, _default int16) int16 {
	if v, ok := m.Int16(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt16(k string) int16 {
	if v, ok := m.Int16(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int32(k string) (v int32, ok bool) {
	_v, ok := m.Int64(k)
	return int32(_v), ok
}

func (m SMap) Int32WithDefault(k string, _default int32) int32 {
	if v, ok := m.Int32(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt32(k string) int32 {
	if v, ok := m.Int32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
