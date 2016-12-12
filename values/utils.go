package values

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt64(_v interface{}) (v int64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = int64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = int64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Int()
	case uint, uint8, uint16, uint32, uint64:
		v = int64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = int64(reflect.ValueOf(_v).Float())
	case string:
		if vv, err := strconv.ParseInt(_v.(string), 10, 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

func ToUInt64(_v interface{}) (v uint64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = uint64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = uint64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(_v).Uint()
	case uint, uint8, uint16, uint32, uint64:
		v = uint64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = uint64(reflect.ValueOf(_v).Float())
	case string:
		if vv, err := strconv.ParseUint(_v.(string), 10, 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

func ToString(_v interface{}) (v string, ok bool) {
	ok = true
	switch _v.(type) {
	case string:
		v = _v.(string)
	case []byte:
		v = string(_v.([]byte))
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		v = fmt.Sprintf("%v", _v)
	default:
		ok = false
	}
	return
}

func ToFloat64(_v interface{}) (v float64, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = float64(real(reflect.ValueOf(_v).Complex()))
	case bool:
		v = float64(Bool2Int(_v.(bool)))
	case int, int8, int16, int32, int64:
		v = float64(reflect.ValueOf(_v).Int())
	case uint, uint8, uint16, uint32, uint64:
		v = float64(reflect.ValueOf(_v).Uint())
	case float32, float64:
		v = reflect.ValueOf(_v).Float()
	case string:
		if vv, err := strconv.ParseFloat(_v.(string), 64); err == nil {
			v = vv
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

func ToComplex128(_v interface{}) (v complex128, ok bool) {
	ok = true
	switch _v.(type) {
	case complex64, complex128:
		v = complex128(reflect.ValueOf(_v).Complex())
	case bool:
		v = complex(float64(Bool2Int(_v.(bool))), 0)
	case int, int8, int16, int32, int64:
		v = complex(float64(reflect.ValueOf(_v).Int()), 0)
	case uint, uint8, uint16, uint32, uint64:
		v = complex(float64(reflect.ValueOf(_v).Uint()), 0)
	case float32, float64:
		v = complex(reflect.ValueOf(_v).Float(), 0)
	default:
		ok = false
	}
	return
}
