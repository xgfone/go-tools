package sort2

import (
	"fmt"

	"github.com/xgfone/go-tools/function"
)

func ExampleInterfaceSlice() {
	data1 := []interface{}{3, 2, 4, 1, 5}
	InterfaceSlice(data1, function.LT)
	fmt.Println(data1)

	data2 := []interface{}{3, 2, 4, 1, 5}
	InterfaceSlice(data2, function.LT, true)
	fmt.Println(data2)

	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
}
