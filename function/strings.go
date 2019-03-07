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
	"bytes"
	"unicode"
)

// SplitSpace splits the string of s by the whitespace, which is equal to
// str.split() in Python.
//
// Notice: SplitSpace(s) == Split(s, unicode.IsSpace).
func SplitSpace(s string) []string {
	return SplitSpaceN(s, -1)
}

// SplitSpaceN is the same as SplitStringN, but the whitespace.
func SplitSpaceN(s string, maxsplit int) []string {
	return SplitN(s, unicode.IsSpace, maxsplit)
}

// SplitString splits the string of s by sep, but is not the same as
// strings.Split(), which the rune in sep arbitrary combination. For example,
// SplitString("abcdefg-12345", "3-edc") == []string{"ab", "fg", "12", "45"}.
func SplitString(s string, sep string) []string {
	return SplitStringN(s, sep, -1)
}

// SplitStringN is the same as SplitN, but the separator is the string of sep.
func SplitStringN(s string, sep string, maxsplit int) []string {
	return SplitN(s, func(c rune) bool {
		for _, r := range sep {
			if r == c {
				return true
			}
		}
		return false
	}, maxsplit)
}

// Split splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator.
func Split(s string, filter func(c rune) bool) []string {
	return SplitN(s, filter, -1)
}

// SplitN splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator, but only maxsplit times.
//
// If maxsplit is equal to 0, don't split; greater than 0, only split maxsplit times;
// less than 0, don't limit. If the leading rune is the separator, it doesn't
// consume the split maxsplit.
//
// Notice: The result does not have the element of nil.
func SplitN(s string, filter func(c rune) bool, maxsplit int) []string {
	if maxsplit == 0 {
		return []string{s}
	}

	j := 0
	for i, c := range s {
		if filter(c) {
			j = i
		} else {
			break
		}
	}
	if j != 0 {
		s = s[j+1:]
	}

	if len(s) == 0 {
		return nil
	}

	maxlen := maxsplit
	if maxlen < 1 {
		maxlen = 4
	}
	results := make([]string, 0, maxlen)
	buf := bytes.NewBuffer(nil)
	isNew := false
	for i, c := range s {
		if filter(c) {
			isNew = true
			continue
		}

		if isNew {
			results = append(results, buf.String())
			buf = bytes.NewBuffer(nil)
			isNew = false
			maxsplit--
			if maxsplit == 0 {
				buf.WriteString(s[i:])
				break
			}
		}

		buf.WriteRune(c)
	}

	last := buf.String()
	if len(last) > 0 {
		results = append(results, last)
	}

	return results
}
