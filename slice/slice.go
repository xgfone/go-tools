// Copyright 2020 xgfone
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

// Package slice supplies some assistant functions about slice.
package slice

import "reflect"

// InSlice reports whether v is in vs.
func InSlice(v interface{}, vs []interface{}) bool {
	for _, _v := range vs {
		if reflect.DeepEqual(v, _v) {
			return true
		}
	}
	return false
}

// InStrings reports whether s is in ss.
func InStrings(s string, ss []string) bool {
	for _, _s := range ss {
		if _s == s {
			return true
		}
	}
	return false
}

// InInts reports whether v is in vs.
func InInts(v int, vs []int) bool {
	for _, i := range vs {
		if i == v {
			return true
		}
	}
	return false
}

// InUints reports whether v is in vs.
func InUints(v uint, vs []uint) bool {
	for _, i := range vs {
		if i == v {
			return true
		}
	}
	return false
}

// StringsEqual reports whether the two strings is equal.
func StringsEqual(s1, s2 []string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// ReverseStrings reverses the elements in the strings slice.
func ReverseStrings(ss []string) {
	slen := len(ss)
	for i, j := 0, slen/2; i < j; i++ {
		k := slen - i - 1
		ss[i], ss[k] = ss[k], ss[i]
	}
}
