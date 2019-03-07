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

func ExampleIsNil() {
	var i *int
	var v interface{} = i

	fmt.Printf("i == %v\n", i)
	fmt.Printf("v == %v\n", v)

	fmt.Printf("i == nil  ==>  %v\n", i == nil)
	fmt.Printf("v == nil  ==>  %v\n", v == nil)

	fmt.Printf("IsNil(i)  ==>  %v\n", IsNil(i))
	fmt.Printf("IsNil(v)  ==>  %v\n", IsNil(v))

	// Output:
	// i == <nil>
	// v == <nil>
	// i == nil  ==>  true
	// v == nil  ==>  false
	// IsNil(i)  ==>  true
	// IsNil(v)  ==>  true
}
