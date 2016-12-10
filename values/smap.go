package values

type SMap map[string]interface{}

func (m SMap) Bool(k string) (bool, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(bool); ok {
			return v2, true
		}
	}
	return false, false
}

func (m SMap) BoolWithDefault(k string, _default bool) bool {
	if v, ok := m.Bool(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustBool(k string) bool {
	if v, ok := m.Bool(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Byte(k string) (byte, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(byte); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) Complex64(k string) (complex64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(complex64); ok {
			return v2, true
		}
	}
	return complex(FZERO32, FZERO32), false
}

func (m SMap) Complex64WithDefault(k string, _default complex64) complex64 {
	if v, ok := m.Complex64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustComplex64(k string) complex64 {
	if v, ok := m.Complex64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Complex128(k string) (complex128, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(complex128); ok {
			return v2, true
		}
	}
	return complex(FZERO64, FZERO64), false
}

func (m SMap) Complex128WithDefault(k string, _default complex128) complex128 {
	if v, ok := m.Complex128(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustComplex128(k string) complex128 {
	if v, ok := m.Complex128(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Float32(k string) (float32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(float32); ok {
			return v2, true
		}
	}
	return FZERO32, false
}

func (m SMap) Float32WithDefault(k string, _default float32) float32 {
	if v, ok := m.Float32(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustFloat32(k string) float32 {
	if v, ok := m.Float32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Float64(k string) (float64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(float64); ok {
			return v2, true
		}
	}
	return FZERO64, false
}
func (m SMap) Float64WithDefault(k string, _default float64) float64 {
	if v, ok := m.Float64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustFloat64(k string) float64 {
	if v, ok := m.Float64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Rune(k string) (rune, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(rune); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) RuneWithDefault(k string, _default rune) rune {
	if v, ok := m.Rune(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustRune(k string) rune {
	if v, ok := m.Rune(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) String(k string) (string, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(string); ok {
			return v2, true
		}
	}
	return "", false
}
func (m SMap) StringWithDefault(k string, _default string) string {
	if v, ok := m.String(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustString(k string) string {
	if v, ok := m.String(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Uintptr(k string) (uintptr, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uintptr); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) Int(k string) (int, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) IntWithDefault(k string, _default int) int {
	if v, ok := m.Int(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt(k string) int {
	if v, ok := m.Int(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt(k string) (uint, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) Int8(k string) (int8, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int8); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int8WithDefault(k string, _default int8) int8 {
	if v, ok := m.Int8(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt8(k string) int8 {
	if v, ok := m.Int8(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int16(k string) (int16, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int16); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int16WithDefault(k string, _default int16) int16 {
	if v, ok := m.Int16(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt16(k string) int16 {
	if v, ok := m.Int16(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int32(k string) (int32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int32); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int32WithDefault(k string, _default int32) int32 {
	if v, ok := m.Int32(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt32(k string) int32 {
	if v, ok := m.Int32(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Int64(k string) (int64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int64); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int64WithDefault(k string, _default int64) int64 {
	if v, ok := m.Int64(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInt64(k string) int64 {
	if v, ok := m.Int64(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) UInt8(k string) (uint8, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint8); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) UInt16(k string) (uint16, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint16); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) UInt32(k string) (uint32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint32); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) UInt64(k string) (uint64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint64); ok {
			return v2, true
		}
	}
	return 0, false
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

func (m SMap) Interface(k string) (interface{}, bool) {
	if v1, ok := m[k]; ok {
		return v1, true
	}
	return nil, false
}

func (m SMap) InterfaceWithDefault(k string, _default interface{}) interface{} {
	if v, ok := m.Interface(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustInterface(k string) interface{} {
	if v, ok := m.Interface(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Slice(k string) (Slice, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(Slice); ok {
			return v2, true
		} else if v3, ok := v1.([]interface{}); ok {
			return Slice(v3), true
		}
	}
	return nil, false
}

func (m SMap) SliceWithDefault(k string, _default Slice) Slice {
	if v, ok := m.Slice(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustSlice(k string) Slice {
	if v, ok := m.Slice(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) SMap(k string) (SMap, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(SMap); ok {
			return v2, true
		} else if v3, ok := v1.(map[string]interface{}); ok {
			return SMap(v3), true
		}
	}
	return nil, false
}

func (m SMap) SMapWithDefault(k string, _default SMap) SMap {
	if v, ok := m.SMap(k); ok {
		return v
	}
	return _default
}

func (m SMap) MustSMap(k string) SMap {
	if v, ok := m.SMap(k); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (m SMap) Keys() []string {
	r := make([]string, 0)
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

func (m SMap) Values() []interface{} {
	r := make([]interface{}, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
