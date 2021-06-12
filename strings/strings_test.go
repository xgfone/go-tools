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

package strings

import (
	"bytes"
	"fmt"
)

func ExampleSafeWriteString() {
	buf := bytes.NewBuffer(nil)

	SafeWriteString(buf, `abcefg`, true, true)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abcefg`, true, false)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abcefg`, false, true)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abcefg`, false, false)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abc"efg`, true, true)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abc"efg`, true, false)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abc"efg`, false, true)
	buf.WriteByte('\n')

	SafeWriteString(buf, `abc"efg`, false, false)
	buf.WriteByte('\n')

	fmt.Print(buf.String())

	// Output:
	// "abcefg"
	// abcefg
	// "abcefg"
	// abcefg
	// "abc\"efg"
	// abc\"efg
	// "abc"efg"
	// abc"efg
}
