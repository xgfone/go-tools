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
	"reflect"
)

// InSlice reports whether the value is in the slice.
//
// Returns false if slice is not a Slice/Array type or the type is not compatible.
func InSlice(value interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	if kind := s.Kind(); kind != reflect.Slice && kind != reflect.Array {
		return false
	}

	for i, slen := 0, s.Len(); i < slen; i++ {
		if reflect.DeepEqual(value, s.Index(i).Interface()) {
			return true
		}
	}

	return false
}
