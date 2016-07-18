package atomics_test

import (
	"fmt"

	"github.com/xgfone/go-tools/atomics"
)

func ExampleCount() {
	c := atomics.NewCount()
	c.Set(100)
	fmt.Println(c.Get())
	c.Add()
	c.Add()
	fmt.Println(c.Get())
	c.Sub()
	fmt.Println(c.Get())
	c.SubWith(100)
	fmt.Println(c.Get())

	// Output:
	// 100
	// 102
	// 101
	// 1
}
