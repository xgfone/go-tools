package types

import (
	"fmt"
)

func ExampleSeries() {
	parent := NewSeries()
	parent.Set("key1", "abc")
	parent.Set("key2", 123)

	child := NewSeries(parent)
	child.Set("key1", "xyz")

	fmt.Printf("key1=%s\n", child.MustGetString("key1"))
	fmt.Printf("key2=%d\n", child.MustGetInt("key2"))
	fmt.Printf("key3=%v\n", child.GetBoolD("key3", false))

	// Output:
	// key1=xyz
	// key2=123
	// key3=false
}
