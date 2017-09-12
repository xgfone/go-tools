// Package sort2 is the supplement of the standard library of sort.
package sort2

import (
	"sort"

	"github.com/xgfone/go-tools/compare"
)

// KeyValuePair is a slice type of key-value pair. The first of the key-value
// pair is the key, and the second of that is the value.
type KeyValuePair [][2]interface{}

func (p KeyValuePair) Len() int {
	return len(p)
}

func (p KeyValuePair) Less(i, j int) bool {
	return compare.LT(p[i][1], p[j][1])
}

func (p KeyValuePair) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}

// SortMap returns the sorted key-value pairs.
//
// The result is the sorted key-value pair slice. the first of the pair is key,
// and the second of that is value.
func SortMap(sm map[interface{}]interface{}) [][2]interface{} {
	if len(sm) == 0 {
		return nil
	}

	results := make([][2]interface{}, 0, len(sm))
	for key, value := range sm {
		results = append(results, [2]interface{}{key, value})
	}

	sort.Sort(KeyValuePair(results))
	return results
}
