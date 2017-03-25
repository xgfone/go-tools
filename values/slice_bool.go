package values

// Bool is the same as Int64, but bool.
func (s Slice) Bool(i int) (v bool, ok bool) {
	if len(s) <= i {
		return
	}

	return !IsZero(s[i]), true
}

// BoolWithDefault is the same as Int64WithDefault, but bool.
func (s Slice) BoolWithDefault(i int, _default bool) bool {
	if v, ok := s.Bool(i); ok {
		return v
	}
	return _default
}

// MustBool is the same as MustInt64, but bool.
func (s Slice) MustBool(i int) bool {
	if v, ok := s.Bool(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
