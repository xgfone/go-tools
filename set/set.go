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

// Package set provides a common set based on generics.
package set

import (
	"bytes"
	"fmt"
)

// Set is a set type.
type Set[T comparable] struct {
	cache map[T]struct{}
}

// NewSet returns a new Set from a slice.
func NewSet[T comparable](elements ...T) Set[T] {
	s := Set[T]{cache: make(map[T]struct{}, len(elements))}
	s.Add(elements...)
	return s
}

// NewSetWithCap returns a new Set with the initialization capacity.
func NewSetWithCap[T comparable](cap int) Set[T] {
	return Set[T]{cache: make(map[T]struct{}, cap)}
}

// NewSetFromSet returns a new Set.
func NewSetFromSet[T comparable](sets ...Set[T]) Set[T] {
	s := Set[T]{cache: make(map[T]struct{}, len(sets))}
	s.UnionUpdate(sets...)
	return s
}

func (s Set[T]) String() string {
	var i int
	buf := bytes.NewBuffer(nil)
	buf.Grow(128)
	buf.WriteByte('{')
	for key := range s.cache {
		if i == 0 {
			fmt.Fprintf(buf, "%v", key)
		} else {
			fmt.Fprintf(buf, " %v", key)
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

// Add adds some elements into the set.
func (s Set[T]) Add(elements ...T) {
	for _, v := range elements {
		s.cache[v] = struct{}{}
	}
}

// Remove removes the elements from the set.
func (s Set[T]) Remove(elements ...T) {
	for _, v := range elements {
		delete(s.cache, v)
	}
}

// Pop removes and returns an arbitrary element from the set.
func (s Set[T]) Pop() (element T, ok bool) {
	for e := range s.cache {
		delete(s.cache, e)
		return e, true
	}
	return
}

// Clear removes all the elements from the set.
func (s Set[T]) Clear() {
	for key := range s.cache {
		delete(s.cache, key)
	}
}

// Contains returns true if the element is in the set. Or return false.
func (s Set[T]) Contains(element T) bool {
	_, ok := s.cache[element]
	return ok
}

// Equal returns true if s == other.
func (s Set[T]) Equal(other Set[T]) bool {
	for e := range s.cache {
		if !other.Contains(e) {
			return false
		}
	}

	for e := range other.cache {
		if !s.Contains(e) {
			return false
		}
	}

	return true
}

// Size returns the number of the elements in the set.
func (s Set[T]) Size() int {
	return len(s.cache)
}

// Slice converts the set to a slice.
func (s Set[T]) Slice() []T {
	list := make([]T, 0, len(s.cache))
	for e := range s.cache {
		list = append(list, e)
	}
	return list
}

// Clone returns a copy of the current set.
func (s Set[T]) Clone() Set[T] {
	cs := Set[T]{cache: make(map[T]struct{}, len(s.cache))}
	for e := range s.cache {
		cs.cache[e] = struct{}{}
	}
	return cs
}

// Range travels all the elements of the set.
func (s Set[T]) Range(f func(element T)) {
	for e := range s.cache {
		f(e)
	}
}

//////////////////////////////////////////////////////////////////////////////

// UnionUpdate updates the set, adding the elements from all others.
func (s Set[T]) UnionUpdate(others ...Set[T]) {
	for _, set := range others {
		for e := range set.cache {
			s.cache[e] = struct{}{}
		}
	}
}

// DifferenceUpdate updates the set, removing the elements found in others.
func (s Set[T]) DifferenceUpdate(others ...Set[T]) {
	for _, set := range others {
		for e := range set.cache {
			delete(s.cache, e)
		}
	}
}

// IntersectionUpdate updates the set, keeping only elements found in it and all others.
func (s Set[T]) IntersectionUpdate(others ...Set[T]) {
	cache := make(map[T]struct{})
	for e := range s.cache {
		var no bool
		for _, set := range others {
			if !set.Contains(e) {
				no = true
				break
			}
		}
		if !no {
			cache[e] = struct{}{}
		}
	}
	s.cache = cache
}

// SymmetricDifferenceUpdate updates the set, keeping only elements
// found in either set, but not in both.
func (s Set[T]) SymmetricDifferenceUpdate(other Set[T]) {
	cache := make(map[T]struct{})

	for e := range other.cache {
		if _, ok := s.cache[e]; !ok {
			cache[e] = struct{}{}
		}
	}

	for e := range s.cache {
		if _, ok := other.cache[e]; !ok {
			cache[e] = struct{}{}
		}
	}

	s.cache = cache
}

//////////////////////////////////////////////////////////////////////////////

// Union returns a new set with elements from the set and all others.
func (s Set[T]) Union(others ...Set[T]) Set[T] {
	r := s.Clone()
	for _, set := range others {
		for e := range set.cache {
			r.cache[e] = struct{}{}
		}
	}
	return r
}

// Difference returns a new set with elements in the set that are not in the others.
func (s Set[T]) Difference(others ...Set[T]) Set[T] {
	r := s.Clone()
	for _, set := range others {
		for e := range set.cache {
			delete(r.cache, e)
		}
	}
	return r
}

// Intersection returns a new set with elements common to the set and all others.
func (s Set[T]) Intersection(others ...Set[T]) Set[T] {
	r := NewSet[T]()
	for e := range s.cache {
		var no bool
		for _, set := range others {
			if !set.Contains(e) {
				no = true
				break
			}
		}
		if !no {
			r.cache[e] = struct{}{}
		}
	}
	return r
}

// SymmetricDifference returns a new set with elements in either the set
// or other but not both.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	r := NewSet[T]()

	for e := range other.cache {
		if _, ok := s.cache[e]; !ok {
			r.cache[e] = struct{}{}
		}
	}

	for e := range s.cache {
		if _, ok := other.cache[e]; !ok {
			r.cache[e] = struct{}{}
		}
	}

	return r
}
