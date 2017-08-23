package values

// UInt64 does the best to convert the value whose key is k to uint64.
func (m SMap) UInt64(k string) (v uint64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToUInt64(_v)
}

// Byte does the best to convert the value whose key is k to byte.
func (m SMap) Byte(k string) (v byte, ok bool) {
	return m.UInt8(k)
}

// Uintptr does the best to convert the value whose key is k to uintptr.
func (m SMap) Uintptr(k string) (v uintptr, ok bool) {
	_v, ok := m.UInt(k)
	return uintptr(_v), ok
}

// UInt does the best to convert the value whose key is k to uint.
func (m SMap) UInt(k string) (v uint, ok bool) {
	_v, ok := m.UInt64(k)
	return uint(_v), ok
}

// UInt8 does the best to convert the value whose key is k to uint8.
func (m SMap) UInt8(k string) (v uint8, ok bool) {
	_v, ok := m.UInt64(k)
	return uint8(_v), ok
}

// UInt16 does the best to convert the value whose key is k to uint16.
func (m SMap) UInt16(k string) (v uint16, ok bool) {
	_v, ok := m.UInt64(k)
	return uint16(_v), ok
}

// UInt32 does the best to convert the value whose key is k to uint32.
func (m SMap) UInt32(k string) (v uint32, ok bool) {
	_v, ok := m.UInt64(k)
	return uint32(_v), ok
}
