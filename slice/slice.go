// Copyright 2022 xgfone
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

// Package slice provides some assistant functions about slice.
package slice

// Reverse reverses the elements in the slice.
func Reverse[E any](vs []E) {
	_len := len(vs) - 1
	if _len <= 0 {
		return
	}

	for i, j := 0, _len/2; i <= j; i++ {
		k := _len - i
		vs[i], vs[k] = vs[k], vs[i]
	}
}

// Contains reports whether vs contains v.
func Contains[E comparable](vs []E, v E) bool {
	for i, _len := 0, len(vs); i < _len; i++ {
		if vs[i] == v {
			return true
		}
	}
	return false
}

// ToInterfaces converts []any to []interface{}.
func ToInterfaces[T any](vs []T) []interface{} {
	is := make([]interface{}, len(vs))
	for i, _len := 0, len(vs); i < _len; i++ {
		is[i] = vs[i]
	}
	return is
}
