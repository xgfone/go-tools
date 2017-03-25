package extremum

// Min returns the minimum of both int64.
func Min(v1, v2 int64) int64 {
	return MinInt64(v1, v2)
}

// MinInt returns the minimum of both int.
func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

// MinUint returns the minimum of both uint.
func MinUint(v1, v2 uint) uint {
	if v1 < v2 {
		return v1
	}
	return v2
}

// MinInt64 returns the minimum of both int64.
func MinInt64(v1, v2 int64) int64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

// MinUint64 returns the minimum of both uint64.
func MinUint64(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v1
	}
	return v2
}
