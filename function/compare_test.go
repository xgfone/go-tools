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

func ExampleEQ() {
	fmt.Println(EQ(1, 1))
	fmt.Println(EQ(1, 2))

	// Output:
	// true
	// false
}

func ExampleLT() {
	fmt.Println(LT(1, 1))
	fmt.Println(LT(1, 2))
	fmt.Println(LT(2, 1))

	// Output:
	// false
	// true
	// false
}

func ExampleGT() {
	fmt.Println(GT(1, 1))
	fmt.Println(GT(1, 2))
	fmt.Println(GT(2, 1))

	// Output:
	// false
	// false
	// true
}

func ExampleLE() {
	fmt.Println(LE(1, 1))
	fmt.Println(LE(1, 2))
	fmt.Println(LE(2, 1))

	// Output:
	// true
	// true
	// false
}

func ExampleGE() {
	fmt.Println(GE(1, 1))
	fmt.Println(GE(1, 2))
	fmt.Println(GE(2, 1))

	// Output:
	// true
	// false
	// true
}
