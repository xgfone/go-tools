package values

// Bool is the same as UInt64, but bool.
func (m SMap) Bool(k string) (v bool, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return !IsZero(_v), true
}

// BoolWithDefault is the same as UInt64WithDefault, but bool.
func (m SMap) BoolWithDefault(k string, _default bool) bool {
	if v, ok := m.Bool(k); ok {
		return v
	}
	return _default
}

// MustBool is the same as MustUInt64, but bool.
func (m SMap) MustBool(k string) bool {
	if v, ok := m.Bool(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
