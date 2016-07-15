// Get the maximal or the minimal of both the values.
package extremum

// Return the maximal of both int64.
func Max(v1, v2 int64) int64 {
	return MaxInt64(v1, v2)
}

// Return the maximal of both int`.
func MaxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Return the maximal of both uint.
func MaxUint(v1, v2 uint) uint {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Return the maximal of both int64.
func MaxInt64(v1, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Return the maximal of both uint64.
func MaxUint64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}
	return v2
}
