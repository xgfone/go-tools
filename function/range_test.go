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

package function

import (
	"fmt"
)

func ExampleRange() {
	fmt.Println(Ranges(1, 10, 2))
	fmt.Println(Ranges(10, 1, -2))

	// Output:
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}

func ExampleRanges() {
	fmt.Println(Range(10))
	fmt.Println(Range(1, 10))
	fmt.Println(Range(1, 10, 1))
	fmt.Println(Range(1, 10, 2))
	fmt.Println(Range(10, 0))

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 3 5 7 9]
	// []
}
