package values

// UInt64 converts the value whose key is k to uint64.
func (m SMap) UInt64(k string) (v uint64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToUInt64(_v)
}

// UInt64WithDefault is the same as UInt64, but return the default value,
// not ZERO, when failed.
func (m SMap) UInt64WithDefault(k string, _default uint64) uint64 {
	if v, ok := m.UInt64(k); ok {
		return v
	}
	return _default
}

// MustUInt64 is the same as UInt64, but panic when failed.
func (m SMap) MustUInt64(k string) uint64 {
	if v, ok := m.UInt64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Byte is the same as uint64, but byte.
func (m SMap) Byte(k string) (v byte, ok bool) {
	return m.UInt8(k)
}

// ByteWithDefault is the same as UInt64WithDefault, but byte.
func (m SMap) ByteWithDefault(k string, _default byte) byte {
	if v, ok := m.Byte(k); ok {
		return v
	}
	return _default
}

// MustByte is the same as MustUInt64, but byte.
func (m SMap) MustByte(k string) byte {
	if v, ok := m.Byte(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// Uintptr is the same as UInt64, but uintptr.
func (m SMap) Uintptr(k string) (v uintptr, ok bool) {
	_v, ok := m.UInt(k)
	return uintptr(_v), ok
}

// UintptrWithDefault is the same as UInt64WithDefault, but uintptr.
func (m SMap) UintptrWithDefault(k string, _default uintptr) uintptr {
	if v, ok := m.Uintptr(k); ok {
		return v
	}
	return _default
}

// MustUintptr is the same as MustUInt64, but uintptr.
func (m SMap) MustUintptr(k string) uintptr {
	if v, ok := m.Uintptr(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt is the same as UInt64, but uint.
func (m SMap) UInt(k string) (v uint, ok bool) {
	_v, ok := m.UInt64(k)
	return uint(_v), ok
}

// UIntWithDefault is the same as UInt64WithDefault, but uint.
func (m SMap) UIntWithDefault(k string, _default uint) uint {
	if v, ok := m.UInt(k); ok {
		return v
	}
	return _default
}

// MustUInt is the same as MustUInt64, but uint.
func (m SMap) MustUInt(k string) uint {
	if v, ok := m.UInt(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt8 is the same as UInt64, but uint8.
func (m SMap) UInt8(k string) (v uint8, ok bool) {
	_v, ok := m.UInt64(k)
	return uint8(_v), ok
}

// UInt8WithDefault is the same as UInt64WithDefault, but uint8.
func (m SMap) UInt8WithDefault(k string, _default uint8) uint8 {
	if v, ok := m.UInt8(k); ok {
		return v
	}
	return _default
}

// MustUInt8 is the same as MustUInt64, but uint8.
func (m SMap) MustUInt8(k string) uint8 {
	if v, ok := m.UInt8(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt16 is the same UInt64, but uint16.
func (m SMap) UInt16(k string) (v uint16, ok bool) {
	_v, ok := m.UInt64(k)
	return uint16(_v), ok
}

// UInt16WithDefault is the same as UInt64WithDefault, but uint16.
func (m SMap) UInt16WithDefault(k string, _default uint16) uint16 {
	if v, ok := m.UInt16(k); ok {
		return v
	}
	return _default
}

// MustUInt16 is the same MustUInt64, but uint16.
func (m SMap) MustUInt16(k string) uint16 {
	if v, ok := m.UInt16(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

// UInt32 is the same as UInt64, but uint32.
func (m SMap) UInt32(k string) (v uint32, ok bool) {
	_v, ok := m.UInt64(k)
	return uint32(_v), ok
}

// UInt32WithDefault is the UInt64WithDefault, but uint32.
func (m SMap) UInt32WithDefault(k string, _default uint32) uint32 {
	if v, ok := m.UInt32(k); ok {
		return v
	}
	return _default
}

// MustUInt32 is the same as MustUInt64, but uint32.
func (m SMap) MustUInt32(k string) uint32 {
	if v, ok := m.UInt32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
