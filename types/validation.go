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

import (
	"fmt"
	"html/template"
	"reflect"
)

var stype2type = map[string]string{
	"zero":             "zero",
	"bool":             "bool",
	"string":           "string",
	"byte":             "byte",
	"rune":             "rune",
	"int":              "int",
	"int8":             "int8",
	"int16":            "int16",
	"int32":            "int32",
	"int64":            "int64",
	"uint":             "uint",
	"uint8":            "uint8",
	"uint16":           "uint16",
	"uint32":           "uint32",
	"uint64":           "uint64",
	"float32":          "float32",
	"float64":          "float64",
	"complex64":        "complex64",
	"complex128":       "complex128",
	"bools":            "[]bool",
	"strings":          "[]string",
	"bytes":            "[]byte",
	"runes":            "[]rune",
	"ints":             "[]int",
	"int8s":            "[]int8",
	"int16s":           "[]int16",
	"int32s":           "[]int32",
	"int64s":           "[]int64",
	"uints":            "[]uint",
	"uint8s":           "[]uint8",
	"uint16s":          "[]uint16",
	"uint32s":          "[]uint32",
	"uint64s":          "[]uint64",
	"float32s":         "[]float32",
	"float64s":         "[]float64",
	"complex64s":       "[]complex64",
	"complex128s":      "[]complex128",
	"string2string":    "map[string]string",
	"string2interface": "map[string]interface{}",
	"int642interface":  "map[int64]interface{}",
	"int642string":     "map[int64]string",
	"int2string":       "map[int]string",
	"int2interface":    "map[int]interface{}",
	"uint642interface": "map[uint64]interface{}",
	"uint642string":    "map[uint64]string",
	"uint2string":      "map[uint]string",
	"uint2interface":   "map[uint]interface{}",
}

// NameToType returns the type string by the name.
//
// The name is the name used by VerifyType, such as string for string, strings
// for []string, or string2string for map[string]string.
func NameToType(name string) string {
	return stype2type[name]
}

// IsZero judges whether a value is ZERO.
//
// For "", 0, 0.0, false, 0+0i, nil, and the slice, array or map that the length
// is 0, they are all ZERO. Others is not ZERO.
func IsZero(v interface{}) bool {
	ok, _ := template.IsTrue(v)
	return !ok
}

// VerifyFunc is a function to verifty whether the type of a value is the given
// type, the first argument of which is the value, and the second of which is
// the given type.
type VerifyFunc func(value interface{}, _type string) bool

var (
	typeMap = make(map[string]VerifyFunc)
)

func init() {
	RegisterVerifyFunc("nil", verifyBasicBuiltinType)
	RegisterVerifyFunc("zero", verifyBasicBuiltinType)
	RegisterVerifyFunc("bool", verifyBasicBuiltinType)
	RegisterVerifyFunc("string", verifyBasicBuiltinType)
	RegisterVerifyFunc("byte", verifyBasicBuiltinType)
	RegisterVerifyFunc("rune", verifyBasicBuiltinType)
	RegisterVerifyFunc("int", verifyBasicBuiltinType)
	RegisterVerifyFunc("int8", verifyBasicBuiltinType)
	RegisterVerifyFunc("int16", verifyBasicBuiltinType)
	RegisterVerifyFunc("int32", verifyBasicBuiltinType)
	RegisterVerifyFunc("int64", verifyBasicBuiltinType)
	RegisterVerifyFunc("uint", verifyBasicBuiltinType)
	RegisterVerifyFunc("uint8", verifyBasicBuiltinType)
	RegisterVerifyFunc("uint16", verifyBasicBuiltinType)
	RegisterVerifyFunc("uint32", verifyBasicBuiltinType)
	RegisterVerifyFunc("uint64", verifyBasicBuiltinType)
	RegisterVerifyFunc("float32", verifyBasicBuiltinType)
	RegisterVerifyFunc("float64", verifyBasicBuiltinType)
	RegisterVerifyFunc("complex64", verifyBasicBuiltinType)
	RegisterVerifyFunc("complex128", verifyBasicBuiltinType)

	RegisterVerifyFunc("bools", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("strings", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("bytes", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("runes", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("ints", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("int8s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("int16s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("int32s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("int64s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("uints", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("uint8s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("uint16s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("uint32s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("uint64s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("float32s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("float64s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("complex64s", verifyBasicBuiltinTypeList)
	RegisterVerifyFunc("complex128s", verifyBasicBuiltinTypeList)

	RegisterVerifyFunc("string2string", verifyMapType)
	RegisterVerifyFunc("string2interface", verifyMapType)
	RegisterVerifyFunc("int642interface", verifyMapType)
	RegisterVerifyFunc("int642string", verifyMapType)
	RegisterVerifyFunc("int2string", verifyMapType)
	RegisterVerifyFunc("int2interface", verifyMapType)
}

// RegisterVerifyFunc registers a type verification function.
func RegisterVerifyFunc(t string, f VerifyFunc) error {
	if _, ok := typeMap[t]; ok {
		return fmt.Errorf("has registered the type of %s", t)
	}
	typeMap[t] = f
	return nil
}

func verifyBasicBuiltinType(v interface{}, t string) (ok bool) {
	switch t {
	case "zero":
		return IsZero(v)
	case "nil":
		ok = v == nil
	case "bool":
		_, ok = v.(bool)
	case "string":
		_, ok = v.(string)
	case "int":
		_, ok = v.(int)
	case "int8":
		_, ok = v.(int8)
	case "int16":
		_, ok = v.(int16)
	case "int32", "rune":
		_, ok = v.(int32)
	case "int64":
		_, ok = v.(int64)
	case "uint":
		_, ok = v.(uint)
	case "uint8", "byte":
		_, ok = v.(uint8)
	case "uint16":
		_, ok = v.(uint16)
	case "uint32":
		_, ok = v.(uint32)
	case "uint64":
		_, ok = v.(uint64)
	case "float32":
		_, ok = v.(float32)
	case "float64":
		_, ok = v.(float64)
	case "complex64":
		_, ok = v.(complex64)
	case "complex128":
		_, ok = v.(complex128)
	}
	return
}

func verifyBasicBuiltinTypeList(v interface{}, t string) (ok bool) {
	switch t {
	case "bools":
		_, ok = v.([]bool)
	case "strings":
		_, ok = v.([]string)
	case "ints":
		_, ok = v.([]int)
	case "int8s":
		_, ok = v.([]int8)
	case "int16s":
		_, ok = v.([]int16)
	case "int32s", "rune":
		_, ok = v.([]int32)
	case "int64s":
		_, ok = v.([]int64)
	case "uints":
		_, ok = v.([]uint)
	case "uint8s", "bytes":
		_, ok = v.([]uint8)
	case "uint16s":
		_, ok = v.([]uint16)
	case "uint32s":
		_, ok = v.([]uint32)
	case "uint64s":
		_, ok = v.([]uint64)
	case "float32s":
		_, ok = v.([]float32)
	case "float64s":
		_, ok = v.([]float64)
	case "complex64s":
		_, ok = v.([]complex64)
	case "complex128s":
		_, ok = v.([]complex128)
	}
	return
}

func verifyMapType(v interface{}, t string) (ok bool) {
	switch t {
	case "string2string":
		_, ok = v.(map[string]string)
	case "string2interface":
		_, ok = v.(map[string]interface{})
	case "int642interface":
		_, ok = v.(map[int64]interface{})
	case "int642string":
		_, ok = v.(map[int64]string)
	case "int2string":
		_, ok = v.(map[int]string)
	case "int2interface":
		_, ok = v.(map[int]interface{})
	case "uint642interface":
		_, ok = v.(map[uint64]interface{})
	case "uint642string":
		_, ok = v.(map[uint64]string)
	case "uint2string":
		_, ok = v.(map[uint]string)
	case "uint2interface":
		_, ok = v.(map[uint]interface{})
	}
	return
}

// VerifyType verifies whether the type of v is t.
//
// The supported types are below:
//
//     t(string)           Go Type / Function Call
//     -------------------------------------------
//     "zero"              IsZero(v)
//     "nil"               nil
//     "bool"              bool
//     "string"            string
//     "byte"              byte
//     "rune"              rune
//     "int"               int
//     "int8"              int8
//     "int16"             int16
//     "int32"             int32
//     "int64"             int64
//     "uint"              uint
//     "uint8"             uint8
//     "uint16"            uint16
//     "uint32"            uint32
//     "uint64"            uint64
//     "float32"           float32
//     "float64"           float64
//     "complex64"         complex64
//     "complex128"        complex128
//     "bools"             []bool
//     "strings"           []string
//     "bytes"             []byte
//     "runes"             []rune
//     "ints"              []int
//     "int8s"             []int8
//     "int16s"            []int16
//     "int32s"            []int32
//     "int64s"            []int64
//     "uints"             []uint
//     "uint8s"            []uint8
//     "uint16s"           []uint16
//     "uint32s"           []uint32
//     "uint64s"           []uint64
//     "float32s"          []float32
//     "float64s"          []float64
//     "complex64s"        []complex64
//     "complex128s"       []complex128
//     "string2string"     map[string]string,
//     "string2interface"  map[string]interface{}
//     "int642interface"   map[int64]interface{}
//     "int642string"      map[int64]string
//     "int2string"        map[int]string
//     "int2interface"     map[int]interface{}
//     "uint642interface"  map[uint64]interface{}
//     "uint642string"     map[uint64]string
//     "uint2string"       map[uint]string
//     "uint2interface"    map[uint]interface{}
//
// Notice: You can add the new type verification by RegisterVerifyFunc it.
func VerifyType(v interface{}, t string) bool {
	if f, ok := typeMap[t]; ok {
		return f(v, t)
	}
	panic(fmt.Errorf("Not support to verify the type of %s", t))
}

// VerifyMapValueType verifies whether the type of the value of the key
// in the map is the given type.
//
// Return false if m is not the map type or the map does not have the key.
//
// Notice: the type of the key of the map type must be string, or return false.
func VerifyMapValueType(m interface{}, k, t string) (ok bool) {
	// Optimize the types of map[string]interface{}, map[string]string,
	// and map[string]int.
	switch _m := m.(type) {
	case map[string]interface{}:
		if v, ok := _m[k]; ok {
			return VerifyType(v, t)
		}
		return
	case map[string]string:
		if v, ok := _m[k]; ok {
			return VerifyType(v, t)
		}
		return
	case map[string]int:
		if v, ok := _m[k]; ok {
			return VerifyType(v, t)
		}
		return
	}

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map || v.Type().Key().Kind() != reflect.String {
		return
	}

	for _, _v := range v.MapKeys() {
		if _v.String() == k {
			return VerifyType(v.MapIndex(_v).Interface(), t)
		}
	}
	return
}

// VerifySliceValueType verifies whether the type of the ith value of the slice
// is the given type.
//
// Return false if the ith value of slice does not exist, that's, i>=len(slice).
// Return false if slice is not a slice or array type, too.
func VerifySliceValueType(slice interface{}, i int, t string) (ok bool) {
	// Optimize the types of []interface, []string, and []int.
	switch s := slice.(type) {
	case []interface{}:
		if i < len(s) {
			return VerifyType(s[i], t)
		}
		return
	case []string:
		if i < len(s) {
			return VerifyType(s[i], t)
		}
		return
	case []int:
		if i < len(s) {
			return VerifyType(s[i], t)
		}
		return
	}

	v := reflect.ValueOf(slice)
	kind := v.Kind()
	if kind != reflect.Slice && kind != reflect.Array && i >= v.Len() {
		return
	}
	return VerifyType(v.Index(i), t)
}
