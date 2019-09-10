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

import "time"

// MustToTime is equal to ToTime, but panic if there is an error.
func MustToTime(v interface{}, layout ...string) time.Time {
	_v, err := ToTime(v, layout...)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToBool is equal to ToBool, but panic if there is an error.
func MustToBool(v interface{}) bool {
	_v, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToFloat32 is equal to ToFloat32, but panic if there is an error.
func MustToFloat32(v interface{}) float32 {
	_v, err := ToFloat32(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToFloat64 is equal to ToFloat64, but panic if there is an error.
func MustToFloat64(v interface{}) float64 {
	_v, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt is equal to ToInt, but panic if there is an error.
func MustToInt(v interface{}) int {
	_v, err := ToInt(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt8 is equal to ToInt8, but panic if there is an error.
func MustToInt8(v interface{}) int8 {
	_v, err := ToInt8(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt16 is equal to ToInt16, but panic if there is an error.
func MustToInt16(v interface{}) int16 {
	_v, err := ToInt16(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt32 is equal to ToInt32, but panic if there is an error.
func MustToInt32(v interface{}) int32 {
	_v, err := ToInt32(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToInt64 is equal to ToInt64, but panic if there is an error.
func MustToInt64(v interface{}) int64 {
	_v, err := ToInt64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint is equal to ToUint, but panic if there is an error.
func MustToUint(v interface{}) uint {
	_v, err := ToUint(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint8 is equal to ToUint8, but panic if there is an error.
func MustToUint8(v interface{}) uint8 {
	_v, err := ToUint8(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint16 is equal to ToUint16, but panic if there is an error.
func MustToUint16(v interface{}) uint16 {
	_v, err := ToUint16(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint32 is equal to ToUint32, but panic if there is an error.
func MustToUint32(v interface{}) uint32 {
	_v, err := ToUint32(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToUint64 is equal to ToUint64, but panic if there is an error.
func MustToUint64(v interface{}) uint64 {
	_v, err := ToUint64(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToString is equal to ToString, but panic if there is an error.
func MustToString(v interface{}) string {
	_v, err := ToString(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToStringMap is equal to ToStringMap, but panic if there is an error.
func MustToStringMap(v interface{}) map[string]interface{} {
	_v, err := ToStringMap(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToSlice is equal to ToSlice, but panic if there is an error.
func MustToSlice(v interface{}) []interface{} {
	_v, err := ToSlice(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToMapKeys is equal to ToMapKeys, but panic if there is an error.
func MustToMapKeys(v interface{}) []string {
	_v, err := ToMapKeys(v)
	if err != nil {
		panic(err)
	}
	return _v
}

// MustToMapValues is equal to ToMapValues, but panic if there is an error.
func MustToMapValues(v interface{}) []interface{} {
	_v, err := ToMapValues(v)
	if err != nil {
		panic(err)
	}
	return _v
}
