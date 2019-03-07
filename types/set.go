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

package types

// Set is a set type.
//
// The element of the set must be hashable, such as int, string, array, etc.
// Notice: slice and map is not hashable.
//
// The set supports the mixed types, but suggest to use the consistent type
// in a set.
type Set struct {
	cache map[interface{}]struct{}
}

// NewSet returns a new Set.
//
// If the element is string, it will ignore the empty string.
func NewSet(elements ...interface{}) Set {
	s := Set{cache: make(map[interface{}]struct{}, len(elements))}
	s.Add(elements...)
	return s
}

// NewSetFromStrings returns a new Set.
//
// It will ignore the empty string.
func NewSetFromStrings(elements ...string) Set {
	s := Set{cache: make(map[interface{}]struct{}, len(elements))}
	s.AddStrings(elements...)
	return s
}

// NewSetFromInts returns a new Set.
func NewSetFromInts(elements ...int) Set {
	s := Set{cache: make(map[interface{}]struct{}, len(elements))}
	s.AddInts(elements...)
	return s
}

// NewSetFromSet returns a new Set.
func NewSetFromSet(sets ...Set) Set {
	s := Set{cache: make(map[interface{}]struct{}, len(sets))}
	s.UnionUpdate(sets...)
	return s
}

// Add adds some elements into the set.
//
// If the element is string, it will ignore the empty string.
func (s Set) Add(elements ...interface{}) {
	for _, v := range elements {
		if v != nil {
			_v, ok := v.(string)
			if ok && _v == "" {
				continue
			}
			s.cache[v] = struct{}{}
		}
	}
}

// AddStrings adds some string elements into the set.
//
// It will ignore the empty string.
func (s Set) AddStrings(elements ...string) {
	for _, v := range elements {
		if v != "" {
			s.cache[v] = struct{}{}
		}
	}
}

// AddInts adds some int elements into the set.
func (s Set) AddInts(elements ...int) {
	for _, v := range elements {
		s.cache[v] = struct{}{}
	}
}

// Remove removes the elements from the set.
func (s Set) Remove(elements ...interface{}) {
	for _, v := range elements {
		delete(s.cache, v)
	}
}

// RemoveStrings removes the string elements from the set.
func (s Set) RemoveStrings(elements ...string) {
	for _, v := range elements {
		delete(s.cache, v)
	}
}

// RemoveInts removes the int elements from the set.
func (s Set) RemoveInts(elements ...int) {
	for _, v := range elements {
		delete(s.cache, v)
	}
}

// Pop removes and returns an arbitrary element from the set.
// But return nil if the set is empty.
func (s Set) Pop() interface{} {
	for e := range s.cache {
		delete(s.cache, e)
		return e
	}
	return nil
}

// Clear removes all the elements from the set.
func (s Set) Clear() {
	s.cache = make(map[interface{}]struct{})
}

// Has returns true if the element is in the set. Or return false.
func (s Set) Has(element interface{}) bool {
	_, ok := s.cache[element]
	return ok
}

// Equal returns true if s == other.
func (s Set) Equal(other Set) bool {
	for e := range s.cache {
		if !other.Has(e) {
			return false
		}
	}

	for e := range other.cache {
		if !s.Has(e) {
			return false
		}
	}

	return true
}

// Size returns the number of the elements in the set.
func (s Set) Size() int {
	return len(s.cache)
}

// List converts the set to a list type.
func (s Set) List() []interface{} {
	list := make([]interface{}, 0, len(s.cache))
	for e := range s.cache {
		list = append(list, e)
	}
	return list
}

// Copy returns a copy of the current set.
func (s Set) Copy() Set {
	cs := Set{cache: make(map[interface{}]struct{}, len(s.cache))}
	for e := range s.cache {
		cs.cache[e] = struct{}{}
	}
	return cs
}

// Walk travels the elements of the set.
func (s Set) Walk(f func(interface{})) {
	for e := range s.cache {
		f(e)
	}
}

//////////////////////////////////////////////////////////////////////////////

// UnionUpdate updates the set, adding the elements from all others.
func (s Set) UnionUpdate(others ...Set) {
	for _, set := range others {
		for e := range set.cache {
			s.cache[e] = struct{}{}
		}
	}
}

// DifferenceUpdate updates the set, removing the elements found in others.
func (s Set) DifferenceUpdate(others ...Set) {
	for _, set := range others {
		for e := range set.cache {
			delete(s.cache, e)
		}
	}
}

// IntersectionUpdate updates the set, keeping only elements found in it and all others.
func (s Set) IntersectionUpdate(others ...Set) {
	cache := make(map[interface{}]struct{})
	for e := range s.cache {
		var no bool
		for _, set := range others {
			if !set.Has(e) {
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
func (s Set) SymmetricDifferenceUpdate(other Set) {
	cache := make(map[interface{}]struct{})

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
func (s Set) Union(others ...Set) Set {
	r := s.Copy()
	for _, set := range others {
		for e := range set.cache {
			r.cache[e] = struct{}{}
		}
	}
	return r
}

// Difference returns a new set with elements in the set that are not in the others.
func (s Set) Difference(others ...Set) Set {
	r := s.Copy()
	for _, set := range others {
		for e := range set.cache {
			delete(r.cache, e)
		}
	}
	return r
}

// Intersection returns a new set with elements common to the set and all others.
func (s Set) Intersection(others ...Set) Set {
	r := NewSet()
	for e := range s.cache {
		var no bool
		for _, set := range others {
			if !set.Has(e) {
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
func (s Set) SymmetricDifference(other Set) Set {
	r := NewSet()

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
