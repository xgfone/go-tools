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

package option

import (
	"testing"
)

func TestOption(t *testing.T) {
	if NONE.IsSome() || NONE.SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if Some(123).IsNone() || Some(123).SomeOr(456).(int) != 123 {
		t.Fail()
	}
}

func TestNamedOption(t *testing.T) {
	if NamedNone("").IsSome() || NamedNone("").SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if NamedSome("", 123).IsNone() || NamedSome("", 123).SomeOr(456).(int) != 123 {
		t.Fail()
	}
}
