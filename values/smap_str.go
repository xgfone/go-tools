package values

// String does the best to convert the value whose key is k to string.
func (m SMap) String(k string) (v string, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToString(_v)
}

// Bytes does the best to convert the value whose key is k to []byte.
func (m SMap) Bytes(k string) (v []byte, ok bool) {
	_v, ok := m.String(k)
	return []byte(_v), ok
}
