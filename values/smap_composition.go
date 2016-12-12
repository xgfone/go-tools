package values

func (m SMap) Interface(k string) (v interface{}, ok bool) {
	if v1, ok := m[k]; ok {
		return v1, true
	}
	return nil, false
}

func (m SMap) InterfaceWithDefault(k string, _default interface{}) interface{} {
	if v, ok := m.Interface(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInterface(k string) interface{} {
	if v, ok := m.Interface(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Slice(k string) (v Slice, ok bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(Slice); ok {
			return v2, true
		} else if v3, ok := v1.([]interface{}); ok {
			return Slice(v3), true
		}
	}
	return nil, false
}

func (m SMap) SliceWithDefault(k string, _default Slice) Slice {
	if v, ok := m.Slice(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustSlice(k string) Slice {
	if v, ok := m.Slice(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) SMap(k string) (v SMap, ok bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(SMap); ok {
			return v2, true
		} else if v3, ok := v1.(map[string]interface{}); ok {
			return SMap(v3), true
		}
	}
	return nil, false
}

func (m SMap) SMapWithDefault(k string, _default SMap) SMap {
	if v, ok := m.SMap(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustSMap(k string) SMap {
	if v, ok := m.SMap(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
