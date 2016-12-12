package values

func (m SMap) Bool(k string) (v bool, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return !IsZero(_v), true
}

func (m SMap) BoolWithDefault(k string, _default bool) bool {
	if v, ok := m.Bool(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustBool(k string) bool {
	if v, ok := m.Bool(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
