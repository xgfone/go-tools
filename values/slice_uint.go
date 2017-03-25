package values

// UInt64 is the same as Int64, but uint64.
func (s Slice) UInt64(i int) (uint64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToUInt64(s[i])
}

// MustUInt64 is the same as MustInt64, but uint64.
func (s Slice) MustUInt64(i int) uint64 {
	if v, ok := s.UInt64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt64WithDefault is the same as Int64WithDefault, but uint64.
func (s Slice) UInt64WithDefault(i int, _default uint64) uint64 {
	if v, ok := s.UInt64(i); ok {
		return v
	}
	return _default
}

// UInt is the same as Int64, but uint.
func (s Slice) UInt(i int) (uint, bool) {
	v, ok := s.UInt64(i)
	return uint(v), ok
}

// MustUInt is the same as MustInt64, but uint.
func (s Slice) MustUInt(i int) uint {
	if v, ok := s.UInt(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UIntWithDefault is the same as Int64WithDefault, but uint.
func (s Slice) UIntWithDefault(i int, _default uint) uint {
	if v, ok := s.UInt(i); ok {
		return v
	}
	return _default
}

// Byte is the same as Int64, but bool.
func (s Slice) Byte(i int) (v byte, ok bool) {
	return s.UInt8(i)
}

// MustByte is the same as MustInt64, but byte.
func (s Slice) MustByte(i int) byte {
	if v, ok := s.Byte(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// ByteWithDefault is the same as Int64WithDefault, but byte.
func (s Slice) ByteWithDefault(i int, _default byte) byte {
	if v, ok := s.Byte(i); ok {
		return v
	}
	return _default
}

// UInt8 is same as Int64, but uint8.
func (s Slice) UInt8(i int) (uint8, bool) {
	v, ok := s.UInt64(i)
	return uint8(v), ok
}

// MustUInt8 is the same as MustInt64, but uint8.
func (s Slice) MustUInt8(i int) uint8 {
	if v, ok := s.UInt8(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt8WithDefault is the same as Int64WithDefault, but uint8.
func (s Slice) UInt8WithDefault(i int, _default uint8) uint8 {
	if v, ok := s.UInt8(i); ok {
		return v
	}
	return _default
}

// UInt16 is the same as Int64, but uint16.
func (s Slice) UInt16(i int) (uint16, bool) {
	v, ok := s.UInt64(i)
	return uint16(v), ok
}

// MustUInt16 is the same as MustInt64, but uint16.
func (s Slice) MustUInt16(i int) uint16 {
	if v, ok := s.UInt16(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt16WithDefault is the same as Int64WithDefault, but uint16.
func (s Slice) UInt16WithDefault(i int, _default uint16) uint16 {
	if v, ok := s.UInt16(i); ok {
		return v
	}
	return _default
}

// UInt32 is the same as Int64, but uint32.
func (s Slice) UInt32(i int) (uint32, bool) {
	v, ok := s.UInt64(i)
	return uint32(v), ok
}

// MustUInt32 is the same as MustInt64, but uint32.
func (s Slice) MustUInt32(i int) uint32 {
	if v, ok := s.UInt32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt32WithDefault is the same as Int64WithDefault, but uint32.
func (s Slice) UInt32WithDefault(i int, _default uint32) uint32 {
	if v, ok := s.UInt32(i); ok {
		return v
	}
	return _default
}

// Uintptr is the same as Int64, but uintptr.
func (s Slice) Uintptr(i int) (uintptr, bool) {
	v, ok := s.UInt(i)
	return uintptr(v), ok
}

// MustUintptr is the same as MustInt64, but uintptr.
func (s Slice) MustUintptr(i int) uintptr {
	if v, ok := s.Uintptr(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UintptrWithDefault is the same as Int64WithDefault, but uintptr.
func (s Slice) UintptrWithDefault(i int, _default uintptr) uintptr {
	if v, ok := s.Uintptr(i); ok {
		return v
	}
	return _default
}
