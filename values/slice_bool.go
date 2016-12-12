package values

func (s Slice) Bool(i int) (v bool, ok bool) {
	if len(s) <= i {
		return
	}

	return !IsZero(s[i]), true
}

func (s Slice) BoolWithDefault(i int, _default bool) bool {
	if v, ok := s.Bool(i); ok {
		return v
	}
	return _default
}

func (s Slice) MustBool(i int) bool {
	if v, ok := s.Bool(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
