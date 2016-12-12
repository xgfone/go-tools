package values

func (m SMap) String(k string) (v string, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToString(_v)
}

func (m SMap) StringWithDefault(k string, _default string) string {
	if v, ok := m.String(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustString(k string) string {
	if v, ok := m.String(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Bytes(k string) (v []byte, ok bool) {
	_v, ok := m.String(k)
	return []byte(_v), ok
}

func (m SMap) BytesWithDefault(k string, _default []byte) []byte {
	if v, ok := m.Bytes(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustBytes(k string) []byte {
	if v, ok := m.Bytes(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
