package values

type Slice []interface{}

func (s Slice) Byte(i int) (byte, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(byte); ok {
		return v, true
	} else {
		return 0, false
	}
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

func (s Slice) Rune(i int) (rune, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(rune); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustRune(i int) rune {
	if v, ok := s.Rune(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) RuneWithDefault(i int, _default rune) rune {
	if v, ok := s.Rune(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int8(i int) (int8, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt8(i int) int8 {
	if v, ok := s.Int8(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int8WithDefault(i int, _default int8) int8 {
	if v, ok := s.Int8(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int16(i int) (int16, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt16(i int) int16 {
	if v, ok := s.Int16(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int16WithDefault(i int, _default int16) int16 {
	if v, ok := s.Int16(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int32(i int) (int32, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt32(i int) int32 {
	if v, ok := s.Int32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int32WithDefault(i int, _default int32) int32 {
	if v, ok := s.Int32(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int64(i int) (int64, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int64); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt64(i int) int64 {
	if v, ok := s.Int64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Int64WithDefault(i int, _default int64) int64 {
	if v, ok := s.Int64(i); ok {
		return v
	}
	return _default
}

func (s Slice) Int(i int) (int, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(int); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt(i int) int {
	if v, ok := s.Int(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) IntWithDefault(i int, _default int) int {
	if v, ok := s.Int(i); ok {
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

func (s Slice) UInt64(i int) (uint64, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(uint64); ok {
		return v, true
	} else {
		return 0, false
	}
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

func (s Slice) String(i int) (string, bool) {
	if len(s) <= i {
		return "", false
	}

	if v, ok := s[i].(string); ok {
		return v, true
	} else {
		return "", false
	}
}

func (s Slice) MustString(i int) string {
	if v, ok := s.String(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) StringWithDefault(i int, _default string) string {
	if v, ok := s.String(i); ok {
		return v
	}
	return _default
}

func (s Slice) Bool(i int) (bool, bool) {
	if len(s) <= i {
		return false, false
	}

	if v, ok := s[i].(bool); ok {
		return v, true
	} else {
		return false, false
	}
}

func (s Slice) MustBool(i int) bool {
	if v, ok := s.Bool(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) BoolWithDefault(i int, _default bool) bool {
	if v, ok := s.Bool(i); ok {
		return v
	}
	return _default
}

func (s Slice) Float32(i int) (float32, bool) {
	if len(s) <= i {
		return FZERO32, false
	}

	if v, ok := s[i].(float32); ok {
		return v, true
	} else {
		return FZERO32, false
	}
}

func (s Slice) MustFloat32(i int) float32 {
	if v, ok := s.Float32(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Float32WithDefault(i int, _default float32) float64 {
	if v, ok := s.Float32(i); ok {
		return v
	}
	return _default
}

func (s Slice) Float64(i int) (float64, bool) {
	if len(s) <= i {
		return FZERO64, false
	}

	if v, ok := s[i].(float64); ok {
		return v, true
	} else {
		return FZERO64, false
	}
}

func (s Slice) MustFloat64(i int) float64 {
	if v, ok := s.Float64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Float64WithDefault(i int, _default float64) float64 {
	if v, ok := s.Float64(i); ok {
		return v
	}
	return _default
}

func (s Slice) Complex64(i int) (complex64, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(complex64); ok {
		return v, true
	} else {
		return complex(FZERO32, FZERO32), false
	}
}

func (s Slice) MustComplex64(i int) complex64 {
	if v, ok := s.Complex64(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Complex64WithDefault(i int, _default complex64) complex64 {
	if v, ok := s.Complex64(i); ok {
		return v
	}
	return _default
}

func (s Slice) Complex128(i int) (complex128, bool) {
	if len(s) <= i {
		return 0, false
	}

	if v, ok := s[i].(complex128); ok {
		return v, true
	} else {
		return complex(FZERO64, FZERO64), false
	}
}

func (s Slice) MustComplex128(i int) complex128 {
	if v, ok := s.Complex128(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) Complex128WithDefault(i int, _default complex128) complex128 {
	if v, ok := s.Complex128(i); ok {
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

func (s Slice) Interface(i int) (interface{}, bool) {
	if len(s) <= i {
		return nil, false
	}

	return s[i], true
}

func (s Slice) MustInterface(i int) interface{} {
	if v, ok := s.Interface(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) InterfaceWithDefault(i int, _default interface{}) interface{} {
	if v, ok := s.Interface(i); ok {
		return v
	}
	return _default
}

func (s Slice) Slice(i int) (Slice, bool) {
	if len(s) <= i {
		return nil, false
	}

	if v, ok := s[i].(Slice); ok {
		return v, true
	} else if v, ok := s[i].([]interface{}); ok {
		return Slice(v), true
	} else {
		return nil, false
	}
}

func (s Slice) MustSlice(i int) Slice {
	if v, ok := s.Slice(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) SliceWithDefault(i int, _default Slice) Slice {
	if v, ok := s.Slice(i); ok {
		return v
	}
	return _default
}

func (s Slice) SMap(i int) (SMap, bool) {
	if len(s) <= i {
		return nil, false
	}

	if v, ok := s[i].(SMap); ok {
		return v, true
	} else if v, ok := s[i].(map[string]interface{}); ok {
		return SMap(v), true
	} else {
		return nil, false
	}
}

func (s Slice) MustSMap(i int) SMap {
	if v, ok := s.SMap(i); !ok {
		panic(ErrTypeOrIndex)
	} else {
		return v
	}
}

func (s Slice) SMapWithDefault(i int, _default SMap) SMap {
	if v, ok := s.SMap(i); ok {
		return v
	}
	return _default
}
