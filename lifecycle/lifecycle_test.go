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

package lifecycle

import (
	"fmt"
)

func ExampleManager() {
	lcm := NewManager()

	in := make(chan interface{})
	out := make(chan interface{})

	fmt.Println("Register the clean functions of apps")
	lcm.Register(func() {
		fmt.Println("Clean and wait app1 to end")
	}).RegisterChannel(in, out).Register(func() {
		fmt.Println("Clean and wait app3 to end")
	})

	fmt.Println("Apps do something ...")
	go func() {
		<-in // Block until the program exits, that's, calling the method lcm.Stop()
		fmt.Println("Clean and wait app2 to end")
		out <- true // Inform the main goruntine that the app has cleaned and ended.
	}()

	fmt.Println("The program is ready to exit ...")
	go lcm.Stop() // Apps clean
	lcm.Wait()
	fmt.Println("The program exited")

	// Output:
	// Register the clean functions of apps
	// Apps do something ...
	// The program is ready to exit ...
	// Clean and wait app3 to end
	// Clean and wait app2 to end
	// Clean and wait app1 to end
	// The program exited
}
