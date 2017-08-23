package values

// Interface returns the interface of the value based on the key k.
func (m SMap) Interface(k string) (v interface{}, ok bool) {
	if v1, ok := m[k]; ok {
		return v1, true
	}
	return nil, false
}

// Slice does the best to convert the value whose key is k to Slice.
func (m SMap) Slice(k string) (v Slice, ok bool) {
	if v1, ok := m[k]; ok {
		return toSlice(v1)
	}
	return nil, false
}

// SMap does the best to convert the value whose key is k to SMap.
func (m SMap) SMap(k string) (v SMap, ok bool) {
	if v1, ok := m[k]; ok {
		return toSMap(v1)
	}
	return nil, false
}

// MapString does the best to convert the value whose key is k to map[string]string.
func (m SMap) MapString(k string) (v map[string]string, ok bool) {
	if _v, _ok := m[k]; _ok {
		switch _v.(type) {
		case map[string]string:
			return _v.(map[string]string), true
		case map[string]interface{}:
			sm := ToSMap(_v)
			v = make(map[string]string, len(sm))
			for k := range sm {
				if v[k], ok = sm.String(k); !ok {
					return nil, false
				}
			}
		}
	}
	return
}
