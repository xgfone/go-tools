package values

// String is the same as Int64, but string.
func (s Slice) String(i int) (v string, ok bool) {
	if len(s) <= i {
		return
	}
	return ToString(s[i])
}

// MustString is the same as MustInt64, but string.
func (s Slice) MustString(i int) string {
	if v, ok := s.String(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// StringWithDefault is the same as Int64WithDefault, but string.
func (s Slice) StringWithDefault(i int, _default string) string {
	if v, ok := s.String(i); ok {
		return v
	}
	return _default
}

// Bytes is the same as Int64, but []byte.
func (s Slice) Bytes(i int) (v []byte, ok bool) {
	_v, ok := s.String(i)
	return []byte(_v), ok
}

// MustBytes is the same as MustInt64, but []byte.
func (s Slice) MustBytes(i int) []byte {
	if v, ok := s.Bytes(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// BytesWithDefault is the same as Int64WithDefault, but []byte.
func (s Slice) BytesWithDefault(i int, _default []byte) []byte {
	if v, ok := s.Bytes(i); ok {
		return v
	}
	return _default
}
