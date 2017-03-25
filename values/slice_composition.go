package values

// Interface returns the interface of the ith value.
func (s Slice) Interface(i int) (v interface{}, ok bool) {
	if len(s) <= i {
		return nil, false
	}

	return s[i], true
}

// MustInterface is the same as Interface, but panic when failed.
func (s Slice) MustInterface(i int) interface{} {
	if v, ok := s.Interface(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// InterfaceWithDefault is the same as Interface, but return the default value,
// not ZERO, when failed.
func (s Slice) InterfaceWithDefault(i int, _default interface{}) interface{} {
	if v, ok := s.Interface(i); ok {
		return v
	}
	return _default
}

// Slice is the same as Int64, but Slice.
func (s Slice) Slice(i int) (v Slice, ok bool) {
	if len(s) <= i {
		return nil, false
	}

	if v, ok := s[i].(Slice); ok {
		return v, true
	} else if v, ok := s[i].([]interface{}); ok {
		return Slice(v), true
	} else {
		return nil, false
	}
}

// MustSlice is the same as MustInt64, but Slice.
func (s Slice) MustSlice(i int) Slice {
	if v, ok := s.Slice(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// SliceWithDefault is the same as Int64WithDefault, but Slice.
func (s Slice) SliceWithDefault(i int, _default Slice) Slice {
	if v, ok := s.Slice(i); ok {
		return v
	}
	return _default
}

// SMap is the same as Int64, but SMap.
func (s Slice) SMap(i int) (v SMap, ok bool) {
	if len(s) <= i {
		return nil, false
	}

	if v, ok := s[i].(SMap); ok {
		return v, true
	} else if v, ok := s[i].(map[string]interface{}); ok {
		return SMap(v), true
	} else {
		return nil, false
	}
}

// MustSMap is the same as MustInt64, but SMap.
func (s Slice) MustSMap(i int) SMap {
	if v, ok := s.SMap(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// SMapWithDefault is the same as Int64WithDefault, but SMap.
func (s Slice) SMapWithDefault(i int, _default SMap) SMap {
	if v, ok := s.SMap(i); ok {
		return v
	}
	return _default
}
