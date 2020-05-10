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

package tag

import (
	"fmt"
	"reflect"
)

func ExampleGetFieldTags() {
	type S struct {
		Field string `name1:"value1" name2:"value2"`
	}

	for name, value := range GetFieldTags(reflect.TypeOf(S{}).Field(0).Tag) {
		fmt.Printf("Tag=%s, Value=%s\n", name, value)
	}

	// Unordered output:
	// Tag=name1, Value=value1
	// Tag=name2, Value=value2
}

func ExampleGetStructTags() {
	type S struct {
		Field1 string `name1:"value1" name2:"value2"`
		Field2 string `name1:"value1" name2:"value2"`
	}

	for field, tags := range GetStructTags(S{}) {
		for name, value := range tags {
			fmt.Printf("Field=%s, Tag=%s, Value=%s\n", field, name, value)
		}
	}

	// Unordered output:
	// Field=Field1, Tag=name1, Value=value1
	// Field=Field1, Tag=name2, Value=value2
	// Field=Field2, Tag=name1, Value=value1
	// Field=Field2, Tag=name2, Value=value2
}
