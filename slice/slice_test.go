// Copyright 2020 xgfone
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

package slice

import "testing"

func TestInInts(t *testing.T) {
	if !InInts(1, []int{1, 2, 3}) {
		t.Fail()
	}
	if InInts(0, []int{1, 2, 3}) {
		t.Fail()
	}
}

func TestInUints(t *testing.T) {
	if !InUints(1, []uint{1, 2, 3}) {
		t.Fail()
	}
	if InUints(0, []uint{1, 2, 3}) {
		t.Fail()
	}
}

func TestInStrings(t *testing.T) {
	if !InStrings("a", []string{"a", "b", "c"}) {
		t.Fail()
	}
	if InStrings("z", []string{"a", "b", "c"}) {
		t.Fail()
	}
}

func TestInSlice(t *testing.T) {
	if !InSlice(1, []interface{}{1, 2, 3}) {
		t.Fail()
	}
	if InSlice(0, []interface{}{1, 2, 3}) {
		t.Fail()
	}
}

func TestReverseStrings(t *testing.T) {
	ss1 := []string{"a", "b", "c", "d"}
	ss2 := []string{"d", "c", "b", "a"}
	ReverseStrings(ss1)
	if !StringsEqual(ss1, ss2) {
		t.Error(ss1)
	}

	ss1 = []string{"a", "b", "c", "d", "e"}
	ss2 = []string{"e", "d", "c", "b", "a"}
	ReverseStrings(ss1)
	if !StringsEqual(ss1, ss2) {
		t.Error(ss1)
	}
}
