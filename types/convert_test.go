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

package types

import (
	"testing"
	"time"
)

func TestConverter(t *testing.T) {
	c := NewConverter(Int)
	if err := c.Scan("123"); err != nil {
		t.Error(err)
	} else if v := c.Value().(int); v != 123 {
		t.Error(v)
	}

	if v, err := Convert(Time, "0000-00-00 00:00:00"); err != nil {
		t.Error(err)
	} else if v != (time.Time{}) {
		t.Error(v)
	}
}
