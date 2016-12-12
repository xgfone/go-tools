package values

func (s Slice) UInt64(i int) (uint64, bool) {
	if len(s) <= i {
		return 0, false
	}
	return ToUInt64(s[i])
}

func (s Slice) MustUInt64(i int) uint64 {
	if v, ok := s.UInt64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UInt64WithDefault(i int, _default uint64) uint64 {
	if v, ok := s.UInt64(i); ok {
		return v
	}
	return _default
}

func (s Slice) UInt(i int) (uint, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uint); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt(i int) uint {
	if v, ok := s.UInt(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UIntWithDefault(i int, _default uint) uint {
	if v, ok := s.UInt(i); ok {
		return v
	}
	return _default
}

func (s Slice) Byte(i int) (v byte, ok bool) {
	return
}

func (s Slice) MustByte(i int) byte {
	if v, ok := s.Byte(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) ByteWithDefault(i int, _default byte) byte {
	if v, ok := s.Byte(i); ok {
		return v
	}
	return _default
}

func (s Slice) UInt8(i int) (uint8, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uint8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt8(i int) uint8 {
	if v, ok := s.UInt8(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UInt8WithDefault(i int, _default uint8) uint8 {
	if v, ok := s.UInt8(i); ok {
		return v
	}
	return _default
}

func (s Slice) UInt16(i int) (uint16, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uint16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt16(i int) uint16 {
	if v, ok := s.UInt16(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UInt16WithDefault(i int, _default uint16) uint16 {
	if v, ok := s.UInt16(i); ok {
		return v
	}
	return _default
}

func (s Slice) UInt32(i int) (uint32, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uint32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt32(i int) uint32 {
	if v, ok := s.UInt32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UInt32WithDefault(i int, _default uint32) uint32 {
	if v, ok := s.UInt32(i); ok {
		return v
	}
	return _default
}

func (s Slice) Uintptr(i int) (uintptr, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uintptr); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUintptr(i int) uintptr {
	if v, ok := s.Uintptr(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) UintptrWithDefault(i int, _default uintptr) uintptr {
	if v, ok := s.Uintptr(i); ok {
		return v
	}
	return _default
}
