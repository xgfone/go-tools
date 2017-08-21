package validation

import (
	"fmt"
	"reflect"
)

// VerifyFunc is a function to verifty whether the type of a value is the given
// type, the first argument of which is the value, and the second of which is
// the given type.
type VerifyFunc func(interface{}, string) bool

var (
	typeMap = make(map[string]VerifyFunc)
)

func init() {
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
	}
	return
}

// VerifyType verifies whether the type of v is t.
//
// t may be a basic builtin type, that's, "bool", "string", "byte", "rune",
// "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32",
// "uint64", "float32", "float64", "complex64", "complex128".
//
// t may be the list type of a basic builtin type as above. Notice: the type name
// is the name of the basic builtin type, which ends with "s".
//
// t may be a map type:
//   string2string for map[string]string,
//   string2interface for map[string]interface{}
//   int642interface for map[int64]interface{}
//   int642string for map[int64]string
//   int2string for map[int]string
//   int2interface for map[int]interface{}
func VerifyType(v interface{}, t string) bool {
	if f, ok := typeMap[t]; ok {
		return f(v, t)
	}
	panic(fmt.Errorf("Not support to verify the type of %s", t))
}

//VerifyMapValueType verifies whether the type of the value of the key in the map
// is the given type.
//
// Return false if m is not the map type or the map does not have the key.
//
// Notice: the type of the key of the map type must be string.
func VerifyMapValueType(m interface{}, k, t string) (ok bool) {
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
