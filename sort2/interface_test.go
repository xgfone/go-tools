package sort2

import (
	"fmt"
)

func ExampleInterfaces() {
	data1 := []interface{}{3, 2, 4, 1, 5}
	Interfaces(data1, func(v1, v2 interface{}) bool { return v1.(int) < v2.(int) })
	fmt.Println(data1)

	// Output:
	// [1 2 3 4 5]
}
