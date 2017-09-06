package values

// Interface returns the interface of the value based on the key k.
func (m SMap) Interface(k string) (v interface{}, err error) {
	if v1, ok := m[k]; ok {
		return v1, nil
	}
	return nil, ErrNoKey
}

// Slice does the best to convert the value whose key is k to Slice.
func (m SMap) Slice(k string) (v Slice, err error) {
	if v1, ok := m[k]; ok {
		return toSlice(v1)
	}
	return nil, ErrNoKey
}

// IsSlice returns true when the type of the value whose key is k is Slice or
// []interface{}; or false.
func (m SMap) IsSlice(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	return isSlice(_v)
}

// SMap does the best to convert the value whose key is k to SMap.
func (m SMap) SMap(k string) (v SMap, err error) {
	if v1, ok := m[k]; ok {
		return toSMap(v1)
	}
	return nil, ErrNoKey
}

// IsSMap returns true when the type of the value whose key is k is SMap or
// map[string]interface{}; or false.
func (m SMap) IsSMap(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	return isSMap(_v)
}

// MapString does the best to convert the value whose key is k to map[string]string.
func (m SMap) MapString(k string) (v map[string]string, err error) {
	if _v, _ok := m[k]; _ok {
		switch _v.(type) {
		case map[string]string:
			return _v.(map[string]string), nil
		case map[string]interface{}:
			sm := ToSMap(_v)
			v = make(map[string]string, len(sm))
			for k := range sm {
				if v[k], err = sm.String(k); err != nil {
					return nil, err
				}
			}
		}
	}
	return
}

// IsMapString returns true when the type of the value whose key is k is map[string]string;
// or false.
func (m SMap) IsMapString(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(map[string]string)
	return ok
}
