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

// Package funcs provides some useful generic functions.
package funcs

import (
	"golang.org/x/exp/constraints"
)

// Min returns the minimal one between v1 and v2.
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v2
	}

	return v1
}

// Max returns the maximum one between v1 and v2.
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v2
	}

	return v1
}

// Compare compares v1 and v2, and returns
//
//	-1 if v1 < v2
//	 0 if v1 == v2
//	 1 if v1 > v2
func Compare[T constraints.Ordered](v1, v2 T) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
}

// GT reports whether v1 is greater than v2.
func GT[T constraints.Ordered](v1, v2 T) bool {
	return Compare(v1, v2) == 1
}

// GE reports whether v1 is greater than or equal to v2.
func GE[T constraints.Ordered](v1, v2 T) bool {
	return Compare(v1, v2) != -1
}

// LT reports whether v1 is less than v2.
func LT[T constraints.Ordered](v1, v2 T) bool {
	return Compare(v1, v2) == -1
}

// LE reports whether v1 is less than or equal to v2.
func LE[T constraints.Ordered](v1, v2 T) bool {
	return Compare(v1, v2) != 1
}
