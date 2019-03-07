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

package sort2

import (
	"fmt"
)

func ExampleInterfaces() {
	data1 := []interface{}{3, 2, 4, 1, 5}
	Interfaces(data1, func(v1, v2 interface{}) bool { return v1.(int) < v2.(int) })
	fmt.Println(data1)

	// Output:
	// [1 2 3 4 5]
}
