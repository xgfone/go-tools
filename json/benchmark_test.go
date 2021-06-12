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
	"testing"
)

func BenchmarkMarshalJSON(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	buf.Grow(64)

	ms := map[string]interface{}{"number": 123, "name": "abc"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MarshalJSON(buf, ms)
	}
}

func BenchmarkMarshal(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	buf.Grow(64)

	ms := map[string]interface{}{"number": 123, "name": "abc"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		json.NewEncoder(buf).Encode(ms)
	}
}
