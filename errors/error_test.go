package errors

import (
	"fmt"
)

func ExampleError() {
	ErrorType1 := NewType("e1")
	ErrorType2 := NewType("e2")
	ErrorType11 := ErrorType1.SubType("e11")

	err1 := ErrorType1.New("error1")
	err2 := ErrorType2.New("error2")
	err11 := ErrorType11.New("error11")

	fmt.Println(err1.IsType(ErrorType1))
	fmt.Println(err1.IsType(ErrorType2))
	fmt.Println(err1.IsType(ErrorType11))
	fmt.Println(err2.IsType(ErrorType1))
	fmt.Println(err2.IsType(ErrorType2))
	fmt.Println(err2.IsType(ErrorType11))
	fmt.Println(err11.IsType(ErrorType1))
	fmt.Println(err11.IsType(ErrorType2))
	fmt.Println(err11.IsType(ErrorType11))

	fmt.Println("---")

	fmt.Println(ErrorType1.IsChildOf(ErrorType2))
	fmt.Println(ErrorType1.IsChildOf(ErrorType11))
	fmt.Println(ErrorType2.IsChildOf(ErrorType1))
	fmt.Println(ErrorType2.IsChildOf(ErrorType11))
	fmt.Println(ErrorType11.IsChildOf(ErrorType1))
	fmt.Println(ErrorType11.IsChildOf(ErrorType2))

	fmt.Println("---")

	fmt.Println(ErrorType1.IsParentOf(ErrorType2))
	fmt.Println(ErrorType1.IsParentOf(ErrorType11))
	fmt.Println(ErrorType2.IsParentOf(ErrorType1))
	fmt.Println(ErrorType2.IsParentOf(ErrorType11))
	fmt.Println(ErrorType11.IsParentOf(ErrorType1))
	fmt.Println(ErrorType11.IsParentOf(ErrorType2))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// false
	// true
	// ---
	// false
	// false
	// false
	// false
	// true
	// false
	// ---
	// false
	// true
	// false
	// false
	// false
	// false
	//
}
