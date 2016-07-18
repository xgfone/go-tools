package atomics_test

import (
	"fmt"

	"github.com/xgfone/go-tools/atomics"
)

func ExampleBool() {
	b := atomics.NewBool()
	fmt.Println(b.Get())
	b.SetTrue()
	fmt.Println(b.Get())
	b.SetFalse()
	fmt.Println(b.Get())

	// Output:
	// false
	// true
	// false
}
