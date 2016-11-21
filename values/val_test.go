package values_test

import (
	"fmt"

	"github.com/xgfone/go-tools/values"
)

func ExampleIsZero() {
	// bool
	fmt.Println("bool:", values.IsZero(true), values.IsZero(false))

	// int, int8, int16, int32, int64, rune
	fmt.Println("int:", values.IsZero(int(1)), values.IsZero(int(0)))

	// uint, uint8, uint16, uint32, uint64, byte
	fmt.Println("uint:", values.IsZero(uint(1)), values.IsZero(uint(0)))

	// complex64, complex128
	fmt.Println("complex:", values.IsZero(complex64(1.1+1.1i)), values.IsZero(complex64(0.0+0.0i)))

	// chan, func, map, slice
	var c chan string
	var f func()
	var m map[string]string
	var s []string
	fmt.Println("chan func map slice:", values.IsZero(c), values.IsZero(f), values.IsZero(m), values.IsZero(s))

	// ptr
	ii := 11
	fmt.Println("ptr:", values.IsZero(&ii), values.IsZero((*int)(nil)))

	// array
	fmt.Println("array:", values.IsZero([3]int{1, 2, 3}), values.IsZero([0]int{}))

	// string
	fmt.Println("string:", values.IsZero("123"), values.IsZero(""))

	// struct
	type S struct{}
	fmt.Println("struct:", values.IsZero(S{}))

	// interface
	fmt.Println("interface:", values.IsZero(interface{}(nil)))

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
