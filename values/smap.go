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

func (m SMap) Byte(k string) (byte, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(byte); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Complex64(k string) (complex64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(complex64); ok {
			return v2, true
		}
	}
	return complex(FZERO32, FZERO32), false
}

func (m SMap) Complex128(k string) (complex128, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(complex128); ok {
			return v2, true
		}
	}
	return complex(FZERO64, FZERO64), false
}

func (m SMap) Float32(k string) (float32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(float32); ok {
			return v2, true
		}
	}
	return FZERO32, false
}

func (m SMap) Float64(k string) (float64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(float64); ok {
			return v2, true
		}
	}
	return FZERO64, false
}

func (m SMap) Rune(k string) (rune, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(rune); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) String(k string) (string, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(string); ok {
			return v2, true
		}
	}
	return "", false
}

func (m SMap) Uintptr(k string) (uintptr, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uintptr); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int(k string) (int, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) UInt(k string) (uint, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int8(k string) (int8, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int8); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int16(k string) (int16, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int16); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int32(k string) (int32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int32); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Int64(k string) (int64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(int64); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) UInt8(k string) (uint8, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint8); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) UInt16(k string) (uint16, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint16); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) UInt32(k string) (uint32, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint32); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) UInt64(k string) (uint64, bool) {
	if v1, ok := m[k]; ok {
		if v2, ok := v1.(uint64); ok {
			return v2, true
		}
	}
	return 0, false
}

func (m SMap) Interface(k string) (interface{}, bool) {
	if v1, ok := m[k]; ok {
		return v1, true
	}
	return nil, false
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
