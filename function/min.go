package function

import (
	"fmt"
	"reflect"
)

// Min returns the minimum of both values.
//
// Support the types: int, int8, int16, int32, int64, uint, uint8, uint16,
// uint32, uint64.
//
// If the types of both is not the same, it will panic.
//
// DEPRECATED!!!
func Min(v1, v2 interface{}) interface{} {
	_v1 := reflect.ValueOf(v1)
	_v2 := reflect.ValueOf(v2)
	if _v1.Kind() != _v2.Kind() {
		panic(errType)
	}

	switch v1.(type) {
	case int, int8, int16, int32, int64:
		if _v1.Int() > _v2.Int() {
			return v2
		}
		return v1
	case uint8, uint16, uint32, uint64:
		if _v1.Uint() > _v2.Uint() {
			return v2
		}
		return v1
	default:
		panic(errType)
	}
}

// MinInSlice returns the minimum in slice v. Return nil if the slice is ZERO.
//
// The type of the element in slice must be int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64. Or panic.
//
// DEPRECATED!!!
func MinInSlice(v interface{}) (min interface{}) {
	if v == nil {
		return
	}

	_v := reflect.ValueOf(v)
	if _v.Kind() != reflect.Slice && _v.Kind() != reflect.Array {
		panic(fmt.Errorf("the type is not slice or array"))
	}

	vlen := _v.Len()
	if vlen == 0 {
		return
	}

	switch v.(type) {
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		min = _v.Index(0).Interface()
		if vlen == 1 {
			return
		}

		for i := 1; i < vlen; i++ {
			min = Min(min, _v.Index(i).Interface())
		}
		return
	default:
		panic(errType)
	}
}
