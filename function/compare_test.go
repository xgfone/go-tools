package function

import (
	"fmt"
)

func ExampleEQ() {
	fmt.Println(EQ(1, 1))
	fmt.Println(EQ(1, 2))

	// Output:
	// true
	// false
}

func ExampleLT() {
	fmt.Println(LT(1, 1))
	fmt.Println(LT(1, 2))
	fmt.Println(LT(2, 1))

	// Output:
	// false
	// true
	// false
}

func ExampleGT() {
	fmt.Println(GT(1, 1))
	fmt.Println(GT(1, 2))
	fmt.Println(GT(2, 1))

	// Output:
	// false
	// false
	// true
}

func ExampleLE() {
	fmt.Println(LE(1, 1))
	fmt.Println(LE(1, 2))
	fmt.Println(LE(2, 1))

	// Output:
	// true
	// true
	// false
}

func ExampleGE() {
	fmt.Println(GE(1, 1))
	fmt.Println(GE(1, 2))
	fmt.Println(GE(2, 1))

	// Output:
	// true
	// false
	// true
}
