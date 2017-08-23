package values

// Int64 does the best to convert the value whose key is k to int64.
func (m SMap) Int64(k string) (v int64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToInt64(_v)
}

// IsInt64 returns true when the type of the value whose key is k is int64;
// or false.
func (m SMap) IsInt64(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(int64)
	return ok
}

// Rune does the best to convert the value whose key is k to rune.
func (m SMap) Rune(k string) (v rune, ok bool) {
	return m.Int32(k)
}

// IsRune returns true when the type of the value whose key is k is rune;
// or false.
func (m SMap) IsRune(k string) bool {
	return m.IsInt32(k)
}

// Int does the best to convert the value whose key is k to int.
func (m SMap) Int(k string) (v int, ok bool) {
	_v, ok := m.Int64(k)
	return int(_v), ok

}

// IsInt returns true when the type of the value whose key is k is int;
// or false.
func (m SMap) IsInt(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(int)
	return ok
}

// Int8 does the best to convert the value whose key is k to int8.
func (m SMap) Int8(k string) (v int8, ok bool) {
	_v, ok := m.Int64(k)
	return int8(_v), ok
}

// IsInt8 returns true when the type of the value whose key is k is int8;
// or false.
func (m SMap) IsInt8(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(int8)
	return ok
}

// Int16 does the best to convert the value whose key is k to int16.
func (m SMap) Int16(k string) (v int16, ok bool) {
	_v, ok := m.Int64(k)
	return int16(_v), ok
}

// IsInt16 returns true when the type of the value whose key is k is int16;
// or false.
func (m SMap) IsInt16(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(int16)
	return ok
}

// Int32 does the best to convert the value whose key is k to int32.
func (m SMap) Int32(k string) (v int32, ok bool) {
	_v, ok := m.Int64(k)
	return int32(_v), ok
}

// IsInt32 returns true when the type of the value whose key is k is int32;
// or false.
func (m SMap) IsInt32(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(int32)
	return ok
}
