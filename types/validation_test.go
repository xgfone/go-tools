package types

import (
	"fmt"
)

func ExampleNameToType() {
	fmt.Println(NameToType("string"))
	fmt.Println(NameToType("strings"))
	fmt.Println(NameToType("string2string"))

	// Output:
	// string
	// []string
	// map[string]string
}

func ExampleIsZero() {
	// bool
	fmt.Println("bool:", IsZero(true), IsZero(false))

	// int, int8, int16, int32, int64, rune
	fmt.Println("int:", IsZero(int(1)), IsZero(int(0)))

	// uint, uint8, uint16, uint32, uint64, byte
	fmt.Println("uint:", IsZero(uint(1)), IsZero(uint(0)))

	// complex64, complex128
	fmt.Println("complex:", IsZero(complex64(1.1+1.1i)),
		IsZero(complex64(0.0+0.0i)))

	// chan, func, map, slice
	var c chan string
	var f func()
	var m map[string]string
	var s []string
	fmt.Println("chan func map slice:", IsZero(c), IsZero(f), IsZero(m),
		IsZero(s))

	// ptr
	ii := 11
	fmt.Println("ptr:", IsZero(&ii), IsZero((*int)(nil)))

	// array
	fmt.Println("array:", IsZero([3]int{1, 2, 3}), IsZero([0]int{}))

	// string
	fmt.Println("string:", IsZero("123"), IsZero(""))

	// struct
	type S struct{}
	fmt.Println("struct:", IsZero(S{}))

	// interface
	fmt.Println("interface:", IsZero(interface{}(nil)))

	// Output:
	// bool: false true
	// int: false true
	// uint: false true
	// complex: false true
	// chan func map slice: true true true true
	// ptr: false true
	// array: false true
	// string: false true
	// struct: false
	// interface: true
}

func ExampleVerifyType() {
	fmt.Println(VerifyType(int64(1), "int64"))
	fmt.Println(VerifyType(int64(1), "int32"))
	fmt.Println(VerifyType(byte(1), "uint8"))
	fmt.Println(VerifyType(byte(1), "int8"))
	fmt.Println(VerifyType("123", "string"))
	fmt.Println(VerifyType("123", "bytes"))

	fmt.Println(VerifyType([]byte("123"), "bytes"))
	fmt.Println(VerifyType([]byte("123"), "string"))
	fmt.Println(VerifyType(map[string]string{"123": "abc"}, "string2string"))
	fmt.Println(VerifyType(map[string]string{"123": "abc"}, "string2interface"))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleVerifyMapValueType() {
	v1 := map[string]interface{}{
		"a": 11,
		"b": "22",
	}
	v2 := map[int]interface{}{
		1: "a",
		2: "b",
	}

	fmt.Println(VerifyMapValueType(v1, "a", "int"))
	fmt.Println(VerifyMapValueType(v1, "b", "string"))
	fmt.Println(VerifyMapValueType(v1, "b", "int"))
	fmt.Println(VerifyMapValueType(v1, "c", "string"))
	fmt.Println(VerifyMapValueType(v2, "1", "string"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}
