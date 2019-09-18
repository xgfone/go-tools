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

package execution

import (
	"context"
	"fmt"
)

func ExampleExecute() {
	if err := Execute(context.TODO(), "ls", "."); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleExecutes() {
	if err := Executes(context.TODO(), []string{"ls"}); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleOutput() {
	if _, err := Output(context.TODO(), "ls", "."); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleOutputs() {
	if _, err := Outputs(context.TODO(), []string{"ls"}); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleHook() {
	filterCmd := func(name string, args ...string) bool {
		// Disable run: rm -rf /
		if name == "rm" && len(args) >= 2 && args[0] == "-rf" && args[1] == "/" {
			return false
		}
		return true
	}
	printCmd := func(name string, args ...string) bool {
		// Print the cmd
		fmt.Printf("Run: %v\n", append([]string{name}, args...))
		return true
	}
	AppendHooks(filterCmd, printCmd)

	RunCmd(context.TODO(), "ls")
	if _, _, err := RunCmd(context.TODO(), "rm", "-rf", "/"); err.(CmdError).Err == ErrDeny {
		fmt.Println(`deny to run "rm -rf /"`)
	}

	// Output:
	// Run: [ls]
	// deny to run "rm -rf /"
}
