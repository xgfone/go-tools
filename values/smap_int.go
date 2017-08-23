package values

// Int64 does the best to convert the value whose key is k to int64.
func (m SMap) Int64(k string) (v int64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToInt64(_v)
}

// Rune does the best to convert the value whose key is k to rune.
func (m SMap) Rune(k string) (v rune, ok bool) {
	return m.Int32(k)
}

// Int does the best to convert the value whose key is k to int.
func (m SMap) Int(k string) (v int, ok bool) {
	_v, ok := m.Int64(k)
	return int(_v), ok

}

// Int8 does the best to convert the value whose key is k to int8.
func (m SMap) Int8(k string) (v int8, ok bool) {
	_v, ok := m.Int64(k)
	return int8(_v), ok
}

// Int16 does the best to convert the value whose key is k to int16.
func (m SMap) Int16(k string) (v int16, ok bool) {
	_v, ok := m.Int64(k)
	return int16(_v), ok
}

// Int32 does the best to convert the value whose key is k to int32.
func (m SMap) Int32(k string) (v int32, ok bool) {
	_v, ok := m.Int64(k)
	return int32(_v), ok
}
