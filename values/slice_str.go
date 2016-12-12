package values

func (s Slice) String(i int) (v string, ok bool) {
	if len(s) <= i {
		return
	}
	return ToString(s[i])
}

func (s Slice) MustString(i int) string {
	if v, ok := s.String(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) StringWithDefault(i int, _default string) string {
	if v, ok := s.String(i); ok {
		return v
	}
	return _default
}

func (s Slice) Bytes(i int) (v []byte, ok bool) {
	_v, ok := s.String(i)
	return []byte(_v), ok
}

func (s Slice) MustBytes(i int) []byte {
	if v, ok := s.Bytes(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) BytesWithDefault(i int, _default []byte) []byte {
	if v, ok := s.Bytes(i); ok {
		return v
	}
	return _default
}
