// Copyright 2022 xgfone
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

package set

import (
	"fmt"
)

func ExampleSet() {
	intset1 := NewSet(1, 2, 3)
	intset2 := NewSet(7, 8, 9)
	strset := NewSet("a", "b", "c")

	intset1.Add(3, 4, 5)
	strset.Add("c", "d", "e")
	fmt.Println(intset1.Size(), strset.Size())
	fmt.Println(intset1.Contains(4), intset1.Contains(5), intset1.Contains(6))
	fmt.Println(strset.Contains("d"), strset.Contains("e"), strset.Contains("z"))
	// 5 5
	// true true false
	// true true false

	intset1.Remove(1, 2, 9)
	strset.Remove("a", "b", "z")
	fmt.Println(intset1.Size(), strset.Size())
	fmt.Println(intset1.Contains(1), intset1.Contains(2))
	fmt.Println(strset.Contains("a"), strset.Contains("b"))
	// 3 3
	// false false
	// false false

	list1 := intset1.Slice()
	fmt.Println(list1[0])
	fmt.Println(list1[1])
	fmt.Println(list1[2])
	// 3
	// 4
	// 5

	list2 := strset.Slice()
	fmt.Println(list2[0])
	fmt.Println(list2[1])
	fmt.Println(list2[2])
	// c
	// d
	// e

	union := intset1.Union(intset2)
	diff := intset1.Difference(intset2)
	inter := intset1.Intersection(intset2)
	sdiff := intset1.SymmetricDifference(intset2)
	fmt.Println(union.Size(), diff.Size(), inter.Size(), sdiff.Size())
	// 6 3 0 6

	fmt.Println(union.Difference(intset1).Equal(intset2))
	fmt.Println(union.Equal(sdiff))
	// true
	// true

	union.Range(func(element int) {
		fmt.Println(element)
	})
	for _, v := range union.Slice() {
		fmt.Println(v)
	}
	// 3
	// 4
	// 5
	// 7
	// 8
	// 9

	// Unordered output:
	// 5 5
	// true true false
	// true true false
	// 3 3
	// false false
	// false false
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
	// 7
	// 8
	// 9
	// 3
	// 4
	// 5
	// 7
	// 8
	// 9
}
