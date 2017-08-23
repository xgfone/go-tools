package values

// Bool does the best to convert the value whose key is k to bool.
func (m SMap) Bool(k string) (v bool, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return !IsZero(_v), true
}

// IsBool returns true when the type of the value whose key is k is bool;
// or false.
func (m SMap) IsBool(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(bool)
	return ok
}
