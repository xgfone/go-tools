package values

// UInt64 does the best to convert the value whose index is i to uint64.
func (s Slice) UInt64(i int) (uint64, error) {
	if len(s) <= i {
		return 0, ErrOutOfLen
	}
	return ToUInt64(s[i])
}

// IsUInt64 returns true when the type of the ith value is uint64; or false.
func (s Slice) IsUInt64(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uint64)
	return ok
}

// UInt does the best to convert the value whose index is i to uint.
func (s Slice) UInt(i int) (uint, error) {
	v, err := s.UInt64(i)
	return uint(v), err
}

// IsUInt returns true when the type of the ith value is uint; or false.
func (s Slice) IsUInt(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uint)
	return ok
}

// Byte does the best to convert the value whose index is i to bool.
func (s Slice) Byte(i int) (v byte, err error) {
	return s.UInt8(i)
}

// IsByte returns true when the type of the ith value is byte; or false.
func (s Slice) IsByte(i int) bool {
	return s.IsUInt8(i)
}

// UInt8 does the best to convert the value whose index is i to uint8.
func (s Slice) UInt8(i int) (uint8, error) {
	v, err := s.UInt64(i)
	return uint8(v), err
}

// IsUInt8 returns true when the type of the ith value is uint8; or false.
func (s Slice) IsUInt8(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uint8)
	return ok
}

// UInt16 does the best to convert the value whose index is i to uint16.
func (s Slice) UInt16(i int) (uint16, error) {
	v, err := s.UInt64(i)
	return uint16(v), err
}

// IsUInt16 returns true when the type of the ith value is uint16; or false.
func (s Slice) IsUInt16(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uint16)
	return ok
}

// UInt32 does the best to convert the value whose index is i to uint32.
func (s Slice) UInt32(i int) (uint32, error) {
	v, err := s.UInt64(i)
	return uint32(v), err
}

// IsUInt32 returns true when the type of the ith value is uint32; or false.
func (s Slice) IsUInt32(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uint32)
	return ok
}

// Uintptr does the best to convert the value whose index is i to uintptr.
func (s Slice) Uintptr(i int) (uintptr, error) {
	v, err := s.UInt(i)
	return uintptr(v), err
}

// IsUintptr returns true when the type of the ith value is uintptr; or false.
func (s Slice) IsUintptr(i int) bool {
	if len(s) <= i {
		return false
	}

	_, ok := s[i].(uintptr)
	return ok
}
