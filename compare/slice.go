package compare

import (
	"fmt"
	"strings"

	"github.com/xgfone/go-tools/extremum"
)

func compareSlice(v1, v2 interface{}) int {
	switch _v1 := v1.(type) {
	case []int:
		if _v2, ok := v2.([]int); !ok {
			panic("Type is not the same")
		} else {
			return compareIntSlice(_v1, _v2)
		}
	case []uint:
		if _v2, ok := v2.([]uint); !ok {
			panic("Type is not the same")
		} else {
			return compareUintSlice(_v1, _v2)
		}
	case []int8:
		if _v2, ok := v2.([]int8); !ok {
			panic("Type is not the same")
		} else {
			return compareInt8Slice(_v1, _v2)
		}
	case []uint8:
		if _v2, ok := v2.([]uint8); !ok {
			panic("Type is not the same")
		} else {
			return compareUint8Slice(_v1, _v2)
		}
	case []int16:
		if _v2, ok := v2.([]int16); !ok {
			panic("Type is not the same")
		} else {
			return compareInt16Slice(_v1, _v2)
		}
	case []uint16:
		if _v2, ok := v2.([]uint16); !ok {
			panic("Type is not the same")
		} else {
			return compareUint16Slice(_v1, _v2)
		}
	case []int32:
		if _v2, ok := v2.([]int32); !ok {
			panic("Type is not the same")
		} else {
			return compareInt32Slice(_v1, _v2)
		}
	case []uint32:
		if _v2, ok := v2.([]uint32); !ok {
			panic("Type is not the same")
		} else {
			return compareUint32Slice(_v1, _v2)
		}
	case []int64:
		if _v2, ok := v2.([]int64); !ok {
			panic("Type is not the same")
		} else {
			return compareInt64Slice(_v1, _v2)
		}
	case []uint64:
		if _v2, ok := v2.([]uint64); !ok {
			panic("Type is not the same")
		} else {
			return compareUint64Slice(_v1, _v2)
		}
	case []string:
		if _v2, ok := v2.([]string); !ok {
			panic("Type is not the same")
		} else {
			return compareStringSlice(_v1, _v2)
		}
	case []float32:
		if _v2, ok := v2.([]float32); !ok {
			panic("Type is not the same")
		} else {
			return compareFloat32Slice(_v1, _v2)
		}
	case []float64:
		if _v2, ok := v2.([]float64); !ok {
			panic("Type is not the same")
		} else {
			return compareFloat64Slice(_v1, _v2)
		}
	default:
		panic(fmt.Sprintf("Type is not supported: %v\n", _v1))
	}
}

func compareIntSlice(v1, v2 []int) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareUintSlice(v1, v2 []uint) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareInt8Slice(v1, v2 []int8) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareUint8Slice(v1, v2 []uint8) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareInt16Slice(v1, v2 []int16) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareUint16Slice(v1, v2 []uint16) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareInt32Slice(v1, v2 []int32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareUint32Slice(v1, v2 []uint32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareInt64Slice(v1, v2 []int64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareUint64Slice(v1, v2 []uint64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareFloat32Slice(v1, v2 []float32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareFloat64Slice(v1, v2 []float64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

func compareStringSlice(v1, v2 []string) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.Min(len1, len2).(int)
	for i := 0; i < _len; i++ {
		diff := strings.Compare(v1[i], v2[i])
		if diff != 0 {
			return diff
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}
