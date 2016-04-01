package extremum

// Return the minimum of both the slices.
// Return -1 if the type is not supported.
//
// The type of both the slices may be []int, []uint, []int64, []uint64.

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

// Return the minimum of both the slice of []int.
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

// Return the minimum of both the slice of []uint.
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

// Return the minimum of both the slice of []int64.
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

// Return the minimum of both the slice of []uint64.
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
