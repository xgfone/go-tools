package values

type Slice []interface{}

func (s Slice) Byte(i int) (byte, bool) {
	if v, ok := s[i].(byte); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Rune(i int) (rune, bool) {
	if v, ok := s[i].(rune); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Int8(i int) (int8, bool) {
	if v, ok := s[i].(int8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Int16(i int) (int16, bool) {
	if v, ok := s[i].(int16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Int32(i int) (int32, bool) {
	if v, ok := s[i].(int32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Int64(i int) (int64, bool) {
	if v, ok := s[i].(int64); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Int(i int) (int, bool) {
	if v, ok := s[i].(int); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) UInt8(i int) (uint8, bool) {
	if v, ok := s[i].(uint8); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) UInt16(i int) (uint16, bool) {
	if v, ok := s[i].(uint16); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) UInt32(i int) (uint32, bool) {
	if v, ok := s[i].(uint32); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) UInt64(i int) (uint64, bool) {
	if v, ok := s[i].(uint64); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) UInt(i int) (uint, bool) {
	if v, ok := s[i].(uint); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) String(i int) (string, bool) {
	if v, ok := s[i].(string); ok {
		return v, true
	} else {
		return "", false
	}
}

func (s Slice) Bool(i int) (bool, bool) {
	if v, ok := s[i].(bool); ok {
		return v, true
	} else {
		return false, false
	}
}

func (s Slice) Float32(i int) (float32, bool) {
	if v, ok := s[i].(float32); ok {
		return v, true
	} else {
		return FZERO32, false
	}
}

func (s Slice) Float64(i int) (float64, bool) {
	if v, ok := s[i].(float64); ok {
		return v, true
	} else {
		return FZERO64, false
	}
}

func (s Slice) Complex64(i int) (complex64, bool) {
	if v, ok := s[i].(complex64); ok {
		return v, true
	} else {
		return complex(FZERO32, FZERO32), false
	}
}

func (s Slice) Complex128(i int) (complex128, bool) {
	if v, ok := s[i].(complex128); ok {
		return v, true
	} else {
		return complex(FZERO64, FZERO64), false
	}
}

func (s Slice) Uintptr(i int) (uintptr, bool) {
	if v, ok := s[i].(uintptr); ok {
		return v, true
	} else {
		return 0, false
	}
}

func (s Slice) Interface(i int) (interface{}, bool) {
	return s[i], true
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

func (s Slice) SMap(i int) (SMap, bool) {
	if v, ok := s[i].(SMap); ok {
		return v, true
	} else if v, ok := s[i].(map[string]interface{}); ok {
		return SMap(v), true
	} else {
		return nil, false
	}
}
