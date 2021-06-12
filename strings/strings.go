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

// Package strings is the supplement of the standard library of `strings`.
package strings

import (
	"io"
	"strconv"
)

var doubleQuotationByte = []byte{'"'}

// SafeWriteString writes s into w.
//
// If escape is true, it will convert '"' to '\"'.
//
// if quote is true, it will output a '"' on both sides of s.
func SafeWriteString(w io.Writer, s string, escape, quote bool) (n int, err error) {
	// Check whether it needs to be escaped.
	if escape {
		escape = false
		for _, c := range s {
			if c == '"' {
				escape = true
			}
		}
		if escape {
			s = strconv.Quote(s)
			s = s[1 : len(s)-1]
		}
	}

	if quote {
		if n, err = w.Write(doubleQuotationByte); err != nil {
			return
		}
	}

	if ws, ok := w.(interface{ WriteString(string) (int, error) }); ok {
		if n, err = ws.WriteString(s); err != nil {
			return
		}
	} else {
		if n, err = w.Write([]byte(s)); err != nil {
			return
		}
	}

	if quote {
		if n, err = w.Write(doubleQuotationByte); err != nil {
			return
		}
	}

	return len(s), nil
}

// WriteString writes s into w.
//
// Notice: it will escape the double-quotation.
func WriteString(w io.Writer, s string, quote ...bool) (n int, err error) {
	if len(quote) > 0 && quote[0] {
		return SafeWriteString(w, s, true, true)
	}
	return SafeWriteString(w, s, true, false)
}
