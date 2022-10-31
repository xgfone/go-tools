// Copyright 2022 xgfone
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

package funcs

import "testing"

func TestMin(t *testing.T) {
	if v := Min(1, 2); v != 1 {
		t.Errorf("Min: expect %d, but got %d", 1, v)
	}
	if v := Min(2, 1); v != 1 {
		t.Errorf("Min: expect %d, but got %d", 1, v)
	}
}

func TestMax(t *testing.T) {
	if v := Max(1, 2); v != 2 {
		t.Errorf("Min: expect %d, but got %d", 2, v)
	}
	if v := Max(2, 1); v != 2 {
		t.Errorf("Min: expect %d, but got %d", 2, v)
	}
}

func TestCompare(t *testing.T) {
	if v1, v2 := 1, 2; GT(v1, v2) {
		t.Errorf("notexpect %d > %d", v1, v2)
	}
	if v1, v2 := 2, 1; !GT(v1, v2) {
		t.Errorf("expect %d > %d", v1, v2)
	}

	if v1, v2 := 1, 2; GE(v1, v2) {
		t.Errorf("notexpect %d >= %d", v1, v2)
	}
	if v1, v2 := 2, 1; !GE(v1, v2) {
		t.Errorf("expect %d >= %d", v1, v2)
	}
	if v1, v2 := 2, 2; !GE(v1, v2) {
		t.Errorf("expect %d >= %d", v1, v2)
	}

	if v1, v2 := 2, 1; LT(v1, v2) {
		t.Errorf("notexpect %d < %d", v1, v2)
	}
	if v1, v2 := 1, 2; !LT(v1, v2) {
		t.Errorf("expect %d < %d", v1, v2)
	}

	if v1, v2 := 2, 1; LE(v1, v2) {
		t.Errorf("notexpect %d <= %d", v1, v2)
	}
	if v1, v2 := 1, 2; !LE(v1, v2) {
		t.Errorf("expect %d <= %d", v1, v2)
	}
	if v1, v2 := 2, 2; !LE(v1, v2) {
		t.Errorf("expect %d <= %d", v1, v2)
	}
}
