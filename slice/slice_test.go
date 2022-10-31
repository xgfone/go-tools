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

package slice

import "fmt"

func ExampleContains() {
	fmt.Println(Contains([]int{1, 2, 3}, 0))
	fmt.Println(Contains([]int{1, 2, 3}, 1))
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 3))
	fmt.Println(Contains([]int{1, 2, 3}, 4))

	// Output:
	// false
	// true
	// true
	// true
	// false
}

func ExampleReverse() {
	vs1 := []string{"a", "b", "c", "d"}
	Reverse(vs1)
	fmt.Println(vs1)

	vs2 := []int{1, 2, 3, 4}
	Reverse(vs2)
	fmt.Println(vs2)

	// Output:
	// [d c b a]
	// [4 3 2 1]
}

func ExampleToInterfaces() {
	ss := []string{"a", "b", "c"}
	vs1 := ToInterfaces(ss)
	fmt.Printf("%T: %v\n", vs1, vs1)

	ints := []int{1, 2, 3}
	vs2 := ToInterfaces(ints)
	fmt.Printf("%T: %v\n", vs2, vs2)

	// Output:
	// []interface {}: [a b c]
	// []interface {}: [1 2 3]
}
