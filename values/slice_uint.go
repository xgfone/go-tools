package values

// UInt64 does the best to convert the value whose index is i to uint64.
func (s Slice) UInt64(i int) (uint64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToUInt64(s[i])
}

// UInt does the best to convert the value whose index is i to uint.
func (s Slice) UInt(i int) (uint, bool) {
	v, ok := s.UInt64(i)
	return uint(v), ok
}

// Byte does the best to convert the value whose index is i to bool.
func (s Slice) Byte(i int) (v byte, ok bool) {
	return s.UInt8(i)
}

// UInt8 does the best to convert the value whose index is i to uint8.
func (s Slice) UInt8(i int) (uint8, bool) {
	v, ok := s.UInt64(i)
	return uint8(v), ok
}

// UInt16 does the best to convert the value whose index is i to uint16.
func (s Slice) UInt16(i int) (uint16, bool) {
	v, ok := s.UInt64(i)
	return uint16(v), ok
}

// UInt32 does the best to convert the value whose index is i to uint32.
func (s Slice) UInt32(i int) (uint32, bool) {
	v, ok := s.UInt64(i)
	return uint32(v), ok
}

// Uintptr does the best to convert the value whose index is i to uintptr.
func (s Slice) Uintptr(i int) (uintptr, bool) {
	v, ok := s.UInt(i)
	return uintptr(v), ok
}
