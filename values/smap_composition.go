package values

// Interface returns the interface of the value based on the key k.
func (m SMap) Interface(k string) (v interface{}, ok bool) {
	if v1, ok := m[k]; ok {
		return v1, true
	}
	return nil, false
}

// InterfaceWithDefault is the same as Interface, but return the default value,
// not ZERO, when failed.
func (m SMap) InterfaceWithDefault(k string, _default interface{}) interface{} {
	if v, ok := m.Interface(k); ok {
		return v
	}
	return _default
}

// MustInterface is the same as Interface, but panic when failed.
func (m SMap) MustInterface(k string) interface{} {
	if v, ok := m.Interface(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Slice is the same as UInt64, but Slice.
func (m SMap) Slice(k string) (v Slice, ok bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(Slice); ok {
			return v2, true
		} else if v3, ok := v1.([]interface{}); ok {
			return Slice(v3), true
		} else {
			return ConvertToSlice(v1)
		}
	}
	return nil, false
}

// SliceWithDefault is the same as UInt64WithDefault, but Slice.
func (m SMap) SliceWithDefault(k string, _default Slice) Slice {
	if v, ok := m.Slice(k); ok {
		return v
	}
	return _default
}

// MustSlice is the same as MustUInt64, but Slice.
func (m SMap) MustSlice(k string) Slice {
	if v, ok := m.Slice(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// SMap is the same as UInt64, but SMap.
func (m SMap) SMap(k string) (v SMap, ok bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(SMap); ok {
			return v2, true
		} else if v3, ok := v1.(map[string]interface{}); ok {
			return SMap(v3), true
		} else {
			return ConvertToSMap(v1)
		}
	}
	return nil, false
}

// SMapWithDefault is the same as UInt64WithDefault, but SMap.
func (m SMap) SMapWithDefault(k string, _default SMap) SMap {
	if v, ok := m.SMap(k); ok {
		return v
	}
	return _default
}

// MustSMap is the same as MustUInt64, but SMap.
func (m SMap) MustSMap(k string) SMap {
	if v, ok := m.SMap(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
