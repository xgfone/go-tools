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

package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func ExampleMarshalJSON() {
	buf := bytes.NewBuffer(nil)

	MarshalJSON(buf, 123)
	buf.WriteByte('\n')
	MarshalJSON(buf, 1.23)
	buf.WriteByte('\n')
	MarshalJSON(buf, "123")
	buf.WriteByte('\n')
	MarshalJSON(buf, `double"quotation`)
	buf.WriteByte('\n')
	MarshalJSON(buf, time.Time{})
	buf.WriteByte('\n')
	MarshalJSON(buf, []int{1, 2, 3})
	buf.WriteByte('\n')
	MarshalJSON(buf, []string{"a", "b", "c"})
	buf.WriteByte('\n')
	MarshalJSON(buf, []float64{1.2, 1.4, 1.6})
	buf.WriteByte('\n')
	// MarshalJSON(buf, map[string]interface{}{"number": 123, "name": "abc"}) // {"number":123,"name":"abc"}
	buf.WriteByte('\n')

	fmt.Printf("%s", buf.String())

	// Output:
	// 123
	// 1.23
	// "123"
	// "double\"quotation"
	// "0001-01-01T00:00:00Z"
	// [1,2,3]
	// ["a","b","c"]
	// [1.2,1.4,1.6]
}

func ExampleMarshalKvJSON() {
	buf := bytes.NewBuffer(nil)
	MarshalKvJSON(buf, "nil", nil, "bool", true, "string", "abc", "int", 123,
		"double_quotation", `a"b`, "float", 1.23, "slice", []interface{}{"abc", 123},
		"sslice", []string{"a", "b", "c"}, "map", map[string]interface{}{"key": "xyz"})

	fmt.Println(buf.String())

	// Check whether the json string is valid.
	data := make(map[string]interface{})
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		fmt.Println(err)
	}

	// Output:
	// {"nil":null,"bool":true,"string":"abc","int":123,"double_quotation":"a\"b","float":1.23,"slice":["abc",123],"sslice":["a","b","c"],"map":{"key":"xyz"}}
}
