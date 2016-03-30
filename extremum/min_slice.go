package extremum

func MinInSlice(v interface{}) int {
	switch _v := v.(type) {
	case []int:
		return MinIntInSlice(_v)
	case []uint:
		return MinUintInSlice(_v)
	case []int64:
		return MinInt64InSlice(_v)
	case []uint64:
		return MinUint64InSlice(_v)
	default:
		return -1
	}
}

func MinIntInSlice(v []int) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] < v[result] {
			result = i
		}
	}

	return result
}

func MinUintInSlice(v []uint) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] < v[result] {
			result = i
		}
	}

	return result
}

func MinInt64InSlice(v []int64) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] < v[result] {
			result = i
		}
	}

	return result
}

func MinUint64InSlice(v []uint64) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] < v[result] {
			result = i
		}
	}

	return result
}
