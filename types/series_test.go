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
	fmt.Printf("key1=%s\n", child.GetParent().MustGetString("key1"))
	fmt.Printf("key2=%d\n", child.MustGetInt("key2"))
	fmt.Printf("key3=%v\n", child.GetBoolD("key3", false))

	// Output:
	// key1=xyz
	// key1=abc
	// key2=123
	// key3=false
}

func ExampleMergeSeriesKeyP2C() {
	parent := NewSeries()
	parent.Set("key1", "abc")
	parent.Set("key2", 123)

	child := NewSeries(parent)
	child.Set("key1", "xyz")

	vs := MergeSeriesKeyP2C(child, "key1")
	fmt.Printf("len=%d\n", len(vs))
	fmt.Printf("values=%v", vs)

	// Output:
	// len=2
	// values=[abc xyz]
}

func ExampleMergeSeriesKeyC2P() {
	parent := NewSeries()
	parent.Set("key1", "abc")
	parent.Set("key2", 123)

	child := NewSeries(parent)
	child.Set("key1", "xyz")

	vs := MergeSeriesKeyC2P(child, "key1")
	fmt.Printf("len=%d\n", len(vs))
	fmt.Printf("values=%v", vs)

	// Output:
	// len=2
	// values=[xyz abc]
}

func ExampleNamedSeries() {
	parent := NewSeries()
	parent.Set("key1", "abc")
	parent.Set("key2", 123)

	child := NewNamedSeries("child", parent)
	child.Set("key1", "xyz")

	fmt.Printf("name=%s\n", child.GetName())
	fmt.Printf("key1=%s\n", child.MustGetString("key1"))
	fmt.Printf("key1=%s\n", child.GetParent().MustGetString("key1"))
	fmt.Printf("key2=%d\n", child.MustGetInt("key2"))
	fmt.Printf("key3=%v\n", child.GetBoolD("key3", false))

	// Output:
	// name=child
	// key1=xyz
	// key1=abc
	// key2=123
	// key3=false
}
