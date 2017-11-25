package function

import (
	"fmt"
	"strings"
)

func compareLen(len1, len2 int) int {
	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareSlice(v1, v2 interface{}) int {
	switch _v1 := v1.(type) {
	case []int:
		return compareIntSlice(_v1, v2.([]int))
	case []uint:
		return compareUintSlice(_v1, v2.([]uint))
	case []int8:
		return compareInt8Slice(_v1, v2.([]int8))
	case []uint8:
		return compareUint8Slice(_v1, v2.([]uint8))
	case []int16:
		return compareInt16Slice(_v1, v2.([]int16))
	case []uint16:
		return compareUint16Slice(_v1, v2.([]uint16))
	case []int32:
		return compareInt32Slice(_v1, v2.([]int32))
	case []uint32:
		return compareUint32Slice(_v1, v2.([]uint32))
	case []int64:
		return compareInt64Slice(_v1, v2.([]int64))
	case []uint64:
		return compareUint64Slice(_v1, v2.([]uint64))
	case []string:
		return compareStringSlice(_v1, v2.([]string))
	case []float32:
		return compareFloat32Slice(_v1, v2.([]float32))
	case []float64:
		return compareFloat64Slice(_v1, v2.([]float64))
	default:
		panic(fmt.Sprintf("Type is not supported: %t\n", _v1))
	}
}

func compareIntSlice(v1, v2 []int) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareUintSlice(v1, v2 []uint) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareInt8Slice(v1, v2 []int8) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareUint8Slice(v1, v2 []uint8) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareInt16Slice(v1, v2 []int16) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareUint16Slice(v1, v2 []uint16) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareInt32Slice(v1, v2 []int32) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareUint32Slice(v1, v2 []uint32) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareInt64Slice(v1, v2 []int64) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareUint64Slice(v1, v2 []uint64) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareFloat32Slice(v1, v2 []float32) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareFloat64Slice(v1, v2 []float64) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	return compareLen(len1, len2)
}

func compareStringSlice(v1, v2 []string) int {
	len1, len2 := len(v1), len(v2)
	_len := Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		diff := strings.Compare(v1[i], v2[i])
		if diff != 0 {
			return diff
		}
	}

	return compareLen(len1, len2)
}
