package values

func (m SMap) UInt64(k string) (v uint64, ok bool) {
	_v, ok := m[k]
	if !ok {
		return
	}
	return ToUInt64(_v)
}

func (m SMap) UInt64WithDefault(k string, _default uint64) uint64 {
	if v, ok := m.UInt64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUInt64(k string) uint64 {
	if v, ok := m.UInt64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Byte(k string) (v byte, ok bool) {
	return m.UInt8(k)
}

func (m SMap) ByteWithDefault(k string, _default byte) byte {
	if v, ok := m.Byte(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustByte(k string) byte {
	if v, ok := m.Byte(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Uintptr(k string) (v uintptr, ok bool) {
	_v, ok := m.UInt(k)
	return uintptr(_v), ok
}

func (m SMap) UintptrWithDefault(k string, _default uintptr) uintptr {
	if v, ok := m.Uintptr(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUintptr(k string) uintptr {
	if v, ok := m.Uintptr(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt(k string) (v uint, ok bool) {
	_v, ok := m.UInt64(k)
	return uint(_v), ok
}

func (m SMap) UIntWithDefault(k string, _default uint) uint {
	if v, ok := m.UInt(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUInt(k string) uint {
	if v, ok := m.UInt(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt8(k string) (v uint8, ok bool) {
	_v, ok := m.UInt64(k)
	return uint8(_v), ok
}

func (m SMap) UInt8WithDefault(k string, _default uint8) uint8 {
	if v, ok := m.UInt8(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUInt8(k string) uint8 {
	if v, ok := m.UInt8(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt16(k string) (v uint16, ok bool) {
	_v, ok := m.UInt64(k)
	return uint16(_v), ok
}

func (m SMap) UInt16WithDefault(k string, _default uint16) uint16 {
	if v, ok := m.UInt16(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUInt16(k string) uint16 {
	if v, ok := m.UInt16(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt32(k string) (v uint32, ok bool) {
	_v, ok := m.UInt64(k)
	return uint32(_v), ok
}

func (m SMap) UInt32WithDefault(k string, _default uint32) uint32 {
	if v, ok := m.UInt32(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustUInt32(k string) uint32 {
	if v, ok := m.UInt32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}
