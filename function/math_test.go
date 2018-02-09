package function

import (
	"fmt"
)

func ExampleAbsInt64() {
	fmt.Printf("AbsInt64(123) = %d\n", AbsInt64(123))
	fmt.Printf("AbsInt64(-123) = %d\n", AbsInt64(-123))

	// Output:
	// AbsInt64(123) = 123
	// AbsInt64(-123) = 123
}
