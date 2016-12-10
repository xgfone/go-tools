package str_test

import (
	"fmt"
	"unicode"

	"github.com/xgfone/go-tools/str"
)

func ExampleSplitSpace() {
	s := "   1   2   3   "
	ss := str.SplitSpace(s)
	fmt.Printf("[len=%v: %v-%v-%v]\n", len(ss), ss[0], ss[1], ss[2])

	// Output:
	// [len=3: 1-2-3]
}

func ExampleSplit() {
	s := "   1   2   3   "
	ss := str.Split(s, unicode.IsSpace)
	fmt.Printf("[len=%v: %v-%v-%v]\n", len(ss), ss[0], ss[1], ss[2])

	// Output:
	// [len=3: 1-2-3]
}
