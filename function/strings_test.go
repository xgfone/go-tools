package function

import (
	"fmt"
	"unicode"
)

func ExampleSplitSpace() {
	s := "   1   2   3   "
	ss := SplitSpace(s)
	fmt.Printf("[len=%v: %v-%v-%v]\n", len(ss), ss[0], ss[1], ss[2])

	// Output:
	// [len=3: 1-2-3]
}

func ExampleSplit() {
	s := "   1   2   3   "
	ss := Split(s, unicode.IsSpace)
	fmt.Printf("[len=%v: %v-%v-%v]\n", len(ss), ss[0], ss[1], ss[2])

	// Output:
	// [len=3: 1-2-3]
}

func ExampleSplitString() {
	s := "abcdefg-12345"
	ss := SplitString(s, "3-edc")
	fmt.Printf("[len=%v: %v-%v-%v-%v]\n", len(ss), ss[0], ss[1], ss[2], ss[3])

	// Output:
	// [len=4: ab-fg-12-45]
}

func ExampleSplitN() {
	s := "   1   2   3   "

	s1 := SplitN(s, unicode.IsSpace, -1)
	fmt.Printf("[len=%v: -%v-%v-%v-]\n", len(s1), s1[0], s1[1], s1[2])

	s2 := SplitN(s, unicode.IsSpace, 0)
	fmt.Printf("[len=%v: -%v-]\n", len(s2), s2[0])

	s3 := SplitN(s, unicode.IsSpace, 1)
	fmt.Printf("[len=%v: -%v-%v-]\n", len(s3), s3[0], s3[1])

	s4 := SplitN(s, unicode.IsSpace, 5)
	if len(s4) == 3 {
		fmt.Printf("[len=%v: -%v-%v-%v-]\n", len(s1), s1[0], s1[1], s1[2])
	}

	fmt.Println(len(SplitSpace("   ")))

	// Output:
	// [len=3: -1-2-3-]
	// [len=1: -   1   2   3   -]
	// [len=2: -1-2   3   -]
	// [len=3: -1-2-3-]
	// 0
}
