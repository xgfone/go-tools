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

package errors

import (
	"fmt"
)

func ExampleError() {
	ErrorType1 := NewType("e1")
	ErrorType2 := NewType("e2")
	ErrorType11 := ErrorType1.SubType("e11")

	err1 := ErrorType1.New("error1")
	err2 := ErrorType2.New("error2")
	err11 := ErrorType11.New("error11")

	fmt.Println(err1.IsType(ErrorType1))
	fmt.Println(err1.IsType(ErrorType2))
	fmt.Println(err1.IsType(ErrorType11))
	fmt.Println(err2.IsType(ErrorType1))
	fmt.Println(err2.IsType(ErrorType2))
	fmt.Println(err2.IsType(ErrorType11))
	fmt.Println(err11.IsType(ErrorType1))
	fmt.Println(err11.IsType(ErrorType2))
	fmt.Println(err11.IsType(ErrorType11))

	fmt.Println("---")

	fmt.Println(ErrorType1.IsChildOf(ErrorType2))
	fmt.Println(ErrorType1.IsChildOf(ErrorType11))
	fmt.Println(ErrorType2.IsChildOf(ErrorType1))
	fmt.Println(ErrorType2.IsChildOf(ErrorType11))
	fmt.Println(ErrorType11.IsChildOf(ErrorType1))
	fmt.Println(ErrorType11.IsChildOf(ErrorType2))

	fmt.Println("---")

	fmt.Println(ErrorType1.IsParentOf(ErrorType2))
	fmt.Println(ErrorType1.IsParentOf(ErrorType11))
	fmt.Println(ErrorType2.IsParentOf(ErrorType1))
	fmt.Println(ErrorType2.IsParentOf(ErrorType11))
	fmt.Println(ErrorType11.IsParentOf(ErrorType1))
	fmt.Println(ErrorType11.IsParentOf(ErrorType2))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// false
	// true
	// ---
	// false
	// false
	// false
	// false
	// true
	// false
	// ---
	// false
	// true
	// false
	// false
	// false
	// false
	//
}
