package function

import (
	"fmt"
)

func ExampleIsNil() {
	var i *int
	var v interface{} = i

	fmt.Printf("i == %v\n", i)
	fmt.Printf("v == %v\n", v)

	fmt.Printf("i == nil  ==>  %v\n", i == nil)
	fmt.Printf("v == nil  ==>  %v\n", v == nil)

	fmt.Printf("IsNil(i)  ==>  %v\n", IsNil(i))
	fmt.Printf("IsNil(v)  ==>  %v\n", IsNil(v))

	// Output:
	// i == <nil>
	// v == <nil>
	// i == nil  ==>  true
	// v == nil  ==>  false
	// IsNil(i)  ==>  true
	// IsNil(v)  ==>  true
}
