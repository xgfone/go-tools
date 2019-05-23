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

import (
	"bytes"
	"testing"
)

func TestBuilder_TruncateBefore(t *testing.T) {
	b := NewBuilder(32)
	b.WriteString("abcdefg")
	b.TruncateBefore(3)
	if b.String() != "defg" {
		t.Error(b.String())
	} else if b.TruncateBefore(10); b.String() != "" {
		t.Error(b.String())
	}
}

func TestBuilder_TruncateAfter(t *testing.T) {
	b := NewBuilder(32)
	b.WriteString("abcdefg")
	b.TruncateAfter(3)
	if b.String() != "abcd" {
		t.Error(b.String())
	} else if b.TruncateAfter(10); b.String() != "" {
		t.Error(b.String())
	}
}

func TestBuilder_WriteRune(t *testing.T) {
	b := NewBuilderString("abc")
	b.WriteRune(rune('d'))
	b.WriteRune(rune('中'))
	if b.String() != "abcd中" {
		t.Error(b.String())
	}
}

func TestBuilder_TrimNewline(t *testing.T) {
	b := NewBuilderString("abcd   \n\n\n")
	b.TrimNewline()
	if b.String() != "abcd   " {
		t.Error(b.String())
	}
}

func TestBuilder_WriteTo(t *testing.T) {
	b := NewBuilder(32)
	b.AppendInt(123)

	buf := bytes.NewBuffer(nil)
	b.WriteTo(buf)

	if buf.String() != "123" {
		t.Error(buf.String())
	} else if b.String() != "123" {
		t.Error(b.String())
	}
}

func TestBuilder_AppendJSON(t *testing.T) {
	b := NewBuilder(32)
	b.AppendJSON(`a"b`)

	if b.String() != `"a\"b"` {
		t.Error(b.String())
	}

	b.Reset()
	b.AppendJSON("ab")
	if b.String() != `"ab"` {
		t.Error(b.String())
	}
}