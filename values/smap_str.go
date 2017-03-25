package values

// String is the same as UInt64, but string.
func (m SMap) String(k string) (v string, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToString(_v)
}

// StringWithDefault is the same as UInt64WithDefault, but string.
func (m SMap) StringWithDefault(k string, _default string) string {
	if v, ok := m.String(k); ok {
		return v
	}
	return _default
}

// MustString is the same as MustUInt64, but string.
func (m SMap) MustString(k string) string {
	if v, ok := m.String(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Bytes is the same as UInt64, but []byte.
func (m SMap) Bytes(k string) (v []byte, ok bool) {
	_v, ok := m.String(k)
	return []byte(_v), ok
}

// BytesWithDefault is the same as UInt64WithDefault, but []byte.
func (m SMap) BytesWithDefault(k string, _default []byte) []byte {
	if v, ok := m.Bytes(k); ok {
		return v
	}
	return _default
}

// MustBytes is the same as MustUInt64, but []byte.
func (m SMap) MustBytes(k string) []byte {
	if v, ok := m.Bytes(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
