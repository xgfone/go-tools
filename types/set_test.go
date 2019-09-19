// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"sort"
)

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

	set := NewSet([2]int{1, 2}, [2]int{3, 4})
	fmt.Println(set.Has([2]int{1, 2}))
	fmt.Println(set.Has([2]int{2, 3}))
	fmt.Println(set.Has([3]int{2, 3, 4}))
	fmt.Println(set.Has([2]string{"a", "b"}))
	// true
	// false
	// false
	// false

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
	// true
	// false
	// false
	// false
}

func ExampleSet_AddSlices() {
	set := NewSetFromSlices([]string{"a", "b"}, []string{"c", "d"})
	set.AddSlices([]string{"o", "p", "q"}, []string{"r", "s", "t"})
	set.AddSlices([]string{"x", "y", "z"})

	ss := make([]string, 0, set.Size())
	set.Walk(func(v interface{}) { ss = append(ss, v.(string)) })
	sort.Strings(ss)

	fmt.Println(set.Size())
	fmt.Println(len(ss))
	fmt.Println(ss)

	// Output:
	// 13
	// 13
	// [a b c d o p q r s t x y z]
}
