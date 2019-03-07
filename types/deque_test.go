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
