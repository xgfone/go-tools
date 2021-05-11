// Copyright 2021 xgfone
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

package atexit

import "fmt"

func ExampleManager() {
	m := NewManager()

	fmt.Println("Register the exit functions of apps")
	m.PushBack(func() {
		fmt.Println("exit1")
	})
	m.PushFront(func() {
		fmt.Println("exit2")
	})

	fmt.Println("Apps do something ...")
	m.Stop()
	fmt.Println("The program exited")

	// Output:
	// Register the exit functions of apps
	// Apps do something ...
	// exit1
	// exit2
	// The program exited
}
