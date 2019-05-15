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
	"fmt"
	"strings"

	"github.com/xgfone/go-tools/types"
)

// DefaultFormat is the default global formatter.
var DefaultFormat = NewFormat("{", "}")

// FmtString formats the string s by DefaultFormat, which is short for
//   DefaultFormat.Format(s, kwargs...)
func FmtString(s string, kwargs ...interface{}) string {
	return DefaultFormat.Format(s, kwargs...)
}

// FmtStringByMap formats the string s by DefaultFormat, which is short for
//   DefaultFormat.FormatByMap(s, kwargs)
func FmtStringByMap(s string, kwargs map[string]interface{}) string {
	return DefaultFormat.FormatByMap(s, kwargs)
}

// Format is used to format a string based on the key placeholder
// that is replaced by the value.
type Format struct {
	Left  string
	Right string
}

// NewFormat returns a new Format.
func NewFormat(left, right string) Format {
	return Format{Left: left, Right: right}
}

// FormatByMap formats the string s, which will replaces the placeholder key
// with the value in kwargs.
//
// Notice: if the placeholder key does not have a corresponding value, it will
// persist and not be replaced.
func (f Format) FormatByMap(s string, kwargs map[string]interface{}) string {
	if len(kwargs) == 0 {
		return s
	}

	buf := bytes.NewBuffer(nil)
	buf.Grow(len(s))

	for {
		leftIndex := strings.Index(s, f.Left)
		if leftIndex == -1 {
			break
		}
		buf.WriteString(s[:leftIndex])
		s = s[leftIndex+len(f.Left):]

		rightIndex := strings.Index(s, f.Right)
		if rightIndex == -1 {
			break
		}
		valueEndIndex := rightIndex + len(f.Right)

		if value, ok := kwargs[s[:rightIndex]]; ok {
			if v, err := types.ToString(value); err == nil {
				buf.WriteString(v)
			} else {
				panic(fmt.Errorf("cannot convert '%v' to string", value))
			}
		} else {
			buf.WriteString(f.Left)
			buf.WriteString(s[:valueEndIndex])
		}
		s = s[valueEndIndex:]
	}

	buf.WriteString(s)
	return buf.String()
}

// Format formats the string s, which will convert kwargs to map[string]interface{}
// and call the method FormatByMap().
//
// Notice: the number of kwargs must be even, and the odd argument must be
// string.
func (f Format) Format(s string, kwargs ...interface{}) string {
	_len := len(kwargs)
	if _len%2 == 1 {
		panic("the number of kwargs must be even")
	}

	ms := make(map[string]interface{}, _len)
	for i := 0; i < _len; i += 2 {
		switch s := kwargs[i].(type) {
		case string:
			ms[s] = kwargs[i+1]
		default:
			panic("the key must be string")
		}
	}

	return f.FormatByMap(s, ms)
}