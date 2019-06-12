// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// ToComplex128 does the best to convert any certain value to complex128.
//
// DEPRECATED!!!
func ToComplex128(_v interface{}) (v complex128, err error) {
	switch t := _v.(type) {
	case nil:
	case complex64:
		v = complex128(t)
	case complex128:
		v = t
	case float32:
		v = complex(float64(t), 0)
	case float64:
		v = complex(t, 0)
	case bool:
		if t {
			v = complex(1, 0)
		} else {
			v = complex(0, 0)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		f, _ := ToFloat64(_v)
		v = complex(f, 0)
	default:
		err = ErrUnknownType
	}
	return
}

// MustToComplex128 is equal to ToComplex128, but panic if there is an error.
//
// DEPRECATED!!!
func MustToComplex128(v interface{}) complex128 {
	_v, err := ToComplex128(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// ToMap is equal to ToStringMap.
//
// DEPRECATED!!! Please use ToStringMap.
func ToMap(v interface{}) (map[string]interface{}, error) {
	return ToStringMap(v)
}

// MustToMap is equal to ToMap, but panic if there is an error.
func MustToMap(v interface{}) map[string]interface{} {
	_v, err := ToMap(v)
	if err != nil {
		panic(err)
	}
	return _v
}
