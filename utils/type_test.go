package utils

import (
	"fmt"
)

func ExmapleVerifyType() {
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
