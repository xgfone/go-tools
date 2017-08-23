package values

// Bool does the best to convert the value whose key is k to bool.
func (m SMap) Bool(k string) (v bool, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return !IsZero(_v), true
}
