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
	"testing"
)

type T struct {
	ID   int
	Name string
}

func (t T) Get() (int, string) {
	return t.ID, t.Name
}

func TestMethod(t *testing.T) {
	v := T{ID: 123, Name: "Aaron"}

	if !HasMethod(v, "Get") {
		t.Fail()
	}

	if HasMethod(v, "Method") {
		t.Fail()
	}

	if GetMethod(v, "ID") != nil {
		t.Fail()
	}

	if _, err := CallMethod(v, "ID"); err == nil {
		t.Fail()
	}

	if _, err := CallMethod(v, "Get", 11); err == nil {
		t.Fail()
	}

	if _, ok := GetMethod(v, "Get").(func(T) (int, string)); !ok {
		t.Fail()
	}

	vv, _ := CallMethod(v, "Get")
	if _vv, ok := vv[0].(int); !ok || _vv != 123 {
		t.Fail()
	}

	if _vv, ok := vv[1].(string); !ok || _vv != "Aaron" {
		t.Fail()
	}
}
