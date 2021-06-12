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

package sort

import "fmt"

func ExampleInt64s() {
	vs := []int64{1, 3, 4, 2}
	Int64s(vs)
	fmt.Println(vs)

	// Output:
	// [1 2 3 4]
}

func ExampleUint64s() {
	vs := []uint64{1, 3, 4, 2}
	Uint64s(vs)
	fmt.Println(vs)

	// Output:
	// [1 2 3 4]
}

func ExampleUints() {
	vs := []uint{1, 3, 4, 2}
	Uints(vs)
	fmt.Println(vs)

	// Output:
	// [1 2 3 4]
}
