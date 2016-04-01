package extremum

// Return the maximal of both the slices.
// Return -1 if the type is not supported.
//
// The type of both the slices may be []int, []uint, []int64, []uint64.
func MaxInSlice(v interface{}) int {
	switch _v := v.(type) {
	case []int:
		return MaxIntInSlice(_v)
	case []uint:
		return MaxUintInSlice(_v)
	case []int64:
		return MaxInt64InSlice(_v)
	case []uint64:
		return MaxUint64InSlice(_v)
	default:
		return -1
	}
}

// Return the maximal of both the slice of []int.
func MaxIntInSlice(v []int) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] > v[result] {
			result = i
		}
	}

	return result
}

// Return the maximal of both the slice of []uint.
func MaxUintInSlice(v []uint) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] > v[result] {
			result = i
		}
	}

	return result
}

// Return the maximal of both the slice of []int64.
func MaxInt64InSlice(v []int64) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] > v[result] {
			result = i
		}
	}

	return result
}

// Return the maximal of both the slice of []uint64.
func MaxUint64InSlice(v []uint64) int {
	_len := len(v)
	if _len == 0 {
		return -1
	}
	result := 0
	for i := 1; i < _len; i++ {
		if v[i] > v[result] {
			result = i
		}
	}

	return result
}
