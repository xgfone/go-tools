package function

import "strings"

// Compare whether v1 is greater than v2.
// Return a positive integer if greater, 0 if equal, a negative if less.
//
// v1 and v2 may be a byte, rune, int, uint, int8, int16, int32, int64,
// uint8, uint16, uint32, uint64, float32, float64, string, or their slice,
// or a struct implementing the interface of Comparer.
//
// Notice: if the types of v1 and v2 are not identical, it will panic.
func Compare(v1, v2 interface{}) int {
	if _v1, ok := v1.(Comparer); ok {
		return _v1.Compare(v2)
	}

	var first, second float64
	switch _v1 := v1.(type) {
	case int:
		first, second = float64(_v1), float64(v2.(int))
	case uint:
		first, second = float64(_v1), float64(v2.(uint))
	case int8:
		first, second = float64(_v1), float64(v2.(int8))
	case uint8:
		first, second = float64(_v1), float64(v2.(uint8))
	case int32:
		first, second = float64(_v1), float64(v2.(int32))
	case uint32:
		first, second = float64(_v1), float64(v2.(uint32))
	case int16:
		first, second = float64(_v1), float64(v2.(int16))
	case uint16:
		first, second = float64(_v1), float64(v2.(uint16))
	case int64:
		first, second = float64(_v1), float64(v2.(int64))
	case uint64:
		first, second = float64(_v1), float64(v2.(uint64))
	case float32:
		first, second = float64(_v1), float64(v2.(float32))
	case float64:
		first, second = _v1, v2.(float64)
	case string:
		return strings.Compare(_v1, v2.(string))
	default:
		return compareSlice(v1, v2)
	}

	if first > second {
		return 1
	} else if first < second {
		return -1
	} else {
		return 0
	}
}

// LT is the same as Compare, but return true if v1 is less than v2, or return false.
func LT(v1, v2 interface{}) bool {
	return Compare(v1, v2) < 0
}

// GT is the same as LT, but greater than.
func GT(v1, v2 interface{}) bool {
	return Compare(v1, v2) > 0
}

// EQ is the same as LT, but equal to.
func EQ(v1, v2 interface{}) bool {
	return Compare(v1, v2) == 0
}

// NE is the same as LT, but not equal to.
func NE(v1, v2 interface{}) bool {
	return !EQ(v1, v2)
}

// GE is the same as LT, but greater than or equal to.
func GE(v1, v2 interface{}) bool {
	return GT(v1, v2) || EQ(v1, v2)
}

// LE is the same as LT, but less than or equal to.
func LE(v1, v2 interface{}) bool {
	return LT(v1, v2) || EQ(v1, v2)
}
