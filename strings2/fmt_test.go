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

package strings2

import "fmt"

func ExampleFormat() {
	format := NewFormat("{{", "}}")

	s1 := format.Format("hello {{name}}. {{name}}.", "name", "world")
	fmt.Println(s1)

	s2 := format.Format("hello {{name}}. {{name}}.", "name", "world", "age", 123)
	fmt.Println(s2)

	s3 := format.Format("hello {{name}}. {{name}}.", "age", 123)
	fmt.Println(s3)

	s4 := format.FormatByMap("hello {{name}}. {{name}}.", map[string]interface{}{"name": "world"})
	fmt.Println(s4)

	s5 := format.FormatByMap("hello {{name}}. {{name}}.", map[string]interface{}{"name": "world", "age": 123})
	fmt.Println(s5)

	s6 := format.FormatByMap("hello {{name}}. {{name}}.", map[string]interface{}{"age": 123})
	fmt.Println(s6)

	s7 := format.FormatByMap("hello {{name}}. You are [{{age:6d}}].", map[string]interface{}{"name": "world", "age": 123})
	fmt.Println(s7)

	s8 := format.Format("hello {{name}}. You are [{{age:6d}}].", "name", "world", "age", 123)
	fmt.Println(s8)

	s9 := format.Format("hello {{name}}.", "name", func() interface{} { return "world" })
	fmt.Println(s9)

	// Output:
	// hello world. world.
	// hello world. world.
	// hello {{name}}. {{name}}.
	// hello world. world.
	// hello world. world.
	// hello {{name}}. {{name}}.
	// hello world. You are [   123].
	// hello world. You are [   123].
	// hello world.
}

func ExampleFmtString() {
	s1 := FmtString("hello {name}. {name}.", "name", "world")
	fmt.Println(s1)

	s2 := FmtString("hello {name}. {name}.", "name", "world", "age", 123)
	fmt.Println(s2)

	s3 := FmtString("hello {name}. {name}.", "age", 123)
	fmt.Println(s3)

	s4 := FmtString("hello {name}. You are [{age:6d}].", "name", "world", "age", 123)
	fmt.Println(s4)

	// Output:
	// hello world. world.
	// hello world. world.
	// hello {name}. {name}.
	// hello world. You are [   123].
}

func ExampleFmtStringByMap() {
	s1 := FmtStringByMap("hello {name}. {name}.", map[string]interface{}{"name": "world"})
	fmt.Println(s1)

	s2 := FmtStringByMap("hello {name}. {name}.", map[string]interface{}{"name": "world", "age": 123})
	fmt.Println(s2)

	s3 := FmtStringByMap("hello {name}. {name}.", map[string]interface{}{"age": 123})
	fmt.Println(s3)

	s4 := FmtStringByMap("hello {name}. You are [{age:6d}].", map[string]interface{}{"name": "world", "age": 123})
	fmt.Println(s4)

	// Output:
	// hello world. world.
	// hello world. world.
	// hello {name}. {name}.
	// hello world. You are [   123].
}
