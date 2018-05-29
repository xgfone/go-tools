package types

import "fmt"

func ExampleSet() {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet("a", "b", "c")

	s1.Add(3, 4, 5)
	s2.Add("c", "d", "e")
	fmt.Println(s1.Size(), s1.Size())
	fmt.Println(s1.Has(4), s1.Has(5), s1.Has(6), s2.Has("d"), s2.Has("e"), s2.Has("z"))
	// Output:
	// 5 5
	// true true false true true false

	s1.RemoveInts(1, 2, 9)
	s2.RemoveStrings("a", "b", "z")
	fmt.Println(s1.Size(), s2.Size())
	fmt.Println(s1.Has(1), s1.Has(2), s2.Has("a"), s2.Has("b"))
	// Output:
	// 3 3
	// false false false false

	list1 := s1.List()
	fmt.Println(list1[0])
	fmt.Println(list1[1])
	fmt.Println(list1[2])
	// Unordered output:
	// 3
	// 4
	// 5

	list2 := s2.List()
	fmt.Println(list2[0])
	fmt.Println(list2[1])
	fmt.Println(list2[2])
	// Unordered output:
	// c
	// d
	// e

	union := s1.Union(s2)
	diff := s1.Difference(s2)
	inter := s1.Intersection(s2)
	sdiff := s1.SymmetricDifference(s2)
	fmt.Println(union.Size(), diff.Size(), inter.Size(), sdiff.Size())
	// Output:
	// 6 3 0 6

	fmt.Println(union.Difference(s1).Equal(s2))
	fmt.Println(union.Equal(sdiff))
	// Output:
	// true
	// true

	list := union.List()
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])
	fmt.Println(list[3])
	fmt.Println(list[4])
	fmt.Println(list[5])
	// Unordered output:
	// 3
	// 4
	// 5
	// c
	// d
	// e

	// Unordered output:
	// 5 5
	// true true false true true false
	// 3 3
	// false false false false
	// 3
	// 4
	// 5
	// c
	// d
	// e
	// 6 3 0 6
	// true
	// true
	// 3
	// 4
	// 5
	// c
	// d
	// e
}
