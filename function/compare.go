package function

import (
	"fmt"
)

// Comparer is used to compare two values.
type Comparer interface {
	// Compare two values.
	//
	// The returned value is
	//   an positive integer when it is greater than value,
	//   0  when they are equal.
	//   an negative integer when it is less than value.
	Compare(value interface{}) int
}

// Compare compares two values.
//
// The returned value is
//   an positive integer when first is greater than second,
//   0  when they are equal.
//   an negative integer when first is less than second.
//
// It supports these types as follow:
//   int  int8  int16  int32  int64
//   uint uint8 uint16 uint32 uint64
//   float32 float64
//   Comparer
//
// Notice: the two values must have the same type and not be nil, or panic.
func Compare(first, second interface{}) int {
	if first == nil || second == nil {
		panic(fmt.Errorf("the value is nil"))
	}

	switch v1 := first.(type) {
	case int:
		return v1 - second.(int)
	case int8:
		return int(v1 - second.(int8))
	case int16:
		return int(v1 - second.(int16))
	case int32:
		return int(v1 - second.(int32))
	case int64:
		return int(v1 - second.(int64))
	case uint:
		return int(v1 - second.(uint))
	case uint8:
		return int(v1 - second.(uint8))
	case uint16:
		return int(v1 - second.(uint16))
	case uint32:
		return int(v1 - second.(uint32))
	case uint64:
		return int(v1 - second.(uint64))
	case float32:
		v2 := second.(float32)
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
		return 0
	case float64:
		v2 := second.(float64)
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
		return 0
	case Comparer:
		return v1.Compare(second)
	default:
		if v2, ok := second.(Comparer); ok {
			v := v2.Compare(first)
			if v == 0 {
				return 0
			}
			return -v
		}
		panic(fmt.Errorf("unsupported type '%T'", first))
	}
}
