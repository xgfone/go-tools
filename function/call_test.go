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
	"errors"
	"fmt"
	"testing"
)

func get(i int, j int) (int, error) {
	return i + j, nil
}

func TestCall(t *testing.T) {
	if ret, err := Call(get, 1, 2); err != nil {
		t.Fail()
	} else {
		if ret[0].(int) != 3 || ret[1] != nil {
			t.Fail()
		}
	}
}

func TestCallWithPointer(t *testing.T) {
	f := func(v *int) (old int) {
		old = *v
		*v++
		return
	}

	v := 1
	ret, _ := Call(f, &v)
	// The returned value is the old, which is 1, and v became 2.
	if ret[0].(int) != 1 || v != 2 {
		t.Fail()
	}
}

func ExampleCall() {
	f := func(i int, j int) (int, error) {
		return i + j, errors.New("This is not an error")
	}

	ret, _ := Call(f, 1, 2)

	// Since the first result is an integer, and it's not necessary to check
	// whether it is nil, so you may omit it, and infer this type directly.
	if ret[0] != nil {
		fmt.Println(ret[0].(int))
	}

	// Since the second result may be nil, so you MUST check whether it is nil firstly.
	if ret[1] != nil {
		fmt.Println(ret[1].(error))
	}
	// Output:
	// 3
	// This is not an error
}
