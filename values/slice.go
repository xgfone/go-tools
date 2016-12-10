package values

type Slice []interface{}

func (s Slice) Byte(i int) (byte, bool) {
	if v, ok := s[i].(byte); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustByte(i int) byte {
	if v, ok := s.Byte(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Rune(i int) (rune, bool) {
	if v, ok := s[i].(rune); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustRune(i int) rune {
	if v, ok := s.Rune(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Int8(i int) (int8, bool) {
	if v, ok := s[i].(int8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt8(i int) int8 {
	if v, ok := s.Int8(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Int16(i int) (int16, bool) {
	if v, ok := s[i].(int16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt16(i int) int16 {
	if v, ok := s.Int16(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Int32(i int) (int32, bool) {
	if v, ok := s[i].(int32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt32(i int) int32 {
	if v, ok := s.Int32(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Int64(i int) (int64, bool) {
	if v, ok := s[i].(int64); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt64(i int) int64 {
	if v, ok := s.Int64(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Int(i int) (int, bool) {
	if v, ok := s[i].(int); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustInt(i int) int {
	if v, ok := s.Int(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) UInt8(i int) (uint8, bool) {
	if v, ok := s[i].(uint8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt8(i int) uint8 {
	if v, ok := s.UInt8(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) UInt16(i int) (uint16, bool) {
	if v, ok := s[i].(uint16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt16(i int) uint16 {
	if v, ok := s.UInt16(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) UInt32(i int) (uint32, bool) {
	if v, ok := s[i].(uint32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt32(i int) uint32 {
	if v, ok := s.UInt32(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) UInt64(i int) (uint64, bool) {
	if v, ok := s[i].(uint64); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt64(i int) uint64 {
	if v, ok := s.UInt64(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) UInt(i int) (uint, bool) {
	if v, ok := s[i].(uint); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUInt(i int) uint {
	if v, ok := s.UInt(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) String(i int) (string, bool) {
	if v, ok := s[i].(string); ok {
		return v, true
	} else {
		return "", false
	}
}

func (s Slice) MustString(i int) string {
	if v, ok := s.String(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Bool(i int) (bool, bool) {
	if v, ok := s[i].(bool); ok {
		return v, true
	} else {
		return false, false
	}
}

func (s Slice) MustBool(i int) bool {
	if v, ok := s.Bool(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Float32(i int) (float32, bool) {
	if v, ok := s[i].(float32); ok {
		return v, true
	} else {
		return FZERO32, false
	}
}

func (s Slice) MustFloat32(i int) float32 {
	if v, ok := s.Float32(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Float64(i int) (float64, bool) {
	if v, ok := s[i].(float64); ok {
		return v, true
	} else {
		return FZERO64, false
	}
}

func (s Slice) MustFloat64(i int) float64 {
	if v, ok := s.Float64(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Complex64(i int) (complex64, bool) {
	if v, ok := s[i].(complex64); ok {
		return v, true
	} else {
		return complex(FZERO32, FZERO32), false
	}
}

func (s Slice) MustComplex64(i int) complex64 {
	if v, ok := s.Complex64(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Complex128(i int) (complex128, bool) {
	if v, ok := s[i].(complex128); ok {
		return v, true
	} else {
		return complex(FZERO64, FZERO64), false
	}
}

func (s Slice) MustComplex128(i int) complex128 {
	if v, ok := s.Complex128(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Uintptr(i int) (uintptr, bool) {
	if v, ok := s[i].(uintptr); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) MustUintptr(i int) uintptr {
	if v, ok := s.Uintptr(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Interface(i int) (interface{}, bool) {
	return s[i], true
}

func (s Slice) MustInterface(i int) interface{} {
	if v, ok := s.Interface(i); !ok {
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) Slice(i int) (Slice, bool) {
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
		panic(ErrType)
	} else {
		return v
	}
}

func (s Slice) SMap(i int) (SMap, bool) {
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
		panic(ErrType)
	} else {
		return v
	}
}
