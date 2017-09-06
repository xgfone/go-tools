package values

// String does the best to convert the value whose key is k to string.
func (m SMap) String(k string) (v string, err error) {
	_v, ok := m[k]
	if !ok {
		err = ErrNoKey
		return
	}
	return ToString(_v)
}

// IsString returns true when the type of the value whose key is k is string;
// or false.
func (m SMap) IsString(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(string)
	return ok
}

// Bytes does the best to convert the value whose key is k to []byte.
func (m SMap) Bytes(k string) (v []byte, err error) {
	_v, ok := m.String(k)
	return []byte(_v), ok
}

// IsBytes returns true when the type of the value whose key is k is []bytes;
// or false.
func (m SMap) IsBytes(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.([]byte)
	return ok
}
