package types

import "fmt"

func ExampleDeque() {
	de := NewDeque()
	de.PushBack(1)
	de.PushBack(2)
	de.PushBack(3)
	de.PushFront("a")
	de.PushFront("b")
	de.PushFront("c")

	de.Each(func(v interface{}) {
		fmt.Println(v)
	})

	fmt.Println(de.PopBack())
	fmt.Println(de.PopFront())

	// Output:
	// c
	// b
	// a
	// 1
	// 2
	// 3
	// 3 true
	// c true
}
