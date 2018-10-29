// Package sort2 is the supplement of the standard library of sort, such as
// key-value slice sort and interface slice sort.
package sort2

import (
	"sort"

	"github.com/xgfone/go-tools/function"
)

// Less is the global default less function, which is `function.LT` by default.
// You can change it to a customized function, but notice that it's global.
//
// For the basic types, such as string, integer and float, the default is OK.
var Less func(first, second interface{}) bool

func init() {
	Less = function.LT
}

// Kv represents a key-value pair.
//
// DEPRECATED!!!
type Kv struct {
	Key   interface{}
	Value interface{}
}

// KvSliceByKey is a slice type of Kv, which is sorted by the key.
//
// Notice: The method Less will use the global function Less to compare the two
// keys.
//
// DEPRECATED!!!
type KvSliceByKey []Kv

func (kv KvSliceByKey) Len() int {
	return len(kv)
}

func (kv KvSliceByKey) Less(i, j int) bool {
	return Less(kv[i].Key, kv[j].Key)
}

func (kv KvSliceByKey) Swap(i, j int) {
	tmp := kv[i]
	kv[i] = kv[j]
	kv[j] = tmp
}

// KvSliceByValue is a slice type of Kv, which is sorted by the value.
//
// Notice: The method Less will use the global function Less to compare the two
// values.
//
// DEPRECATED!!!
type KvSliceByValue []Kv

func (kv KvSliceByValue) Len() int {
	return len(kv)
}

func (kv KvSliceByValue) Less(i, j int) bool {
	return Less(kv[i].Value, kv[j].Value)
}

func (kv KvSliceByValue) Swap(i, j int) {
	tmp := kv[i]
	kv[i] = kv[j]
	kv[j] = tmp
}

// Sort is the union of sort.Sort and sort.Reverse.
//
// if giving the second argument and it's true, it will sort in reverse order.
//
// DEPRECATED!!!
func Sort(kv sort.Interface, reverse ...bool) {
	if len(reverse) > 0 && reverse[0] {
		kv = sort.Reverse(kv)
	}
	sort.Sort(kv)
}

// MapToKvSliceByKey converts map[string]interface{} to KvSliceByKey.
//
// DEPRECATED!!!
func MapToKvSliceByKey(m map[string]interface{}) KvSliceByKey {
	if len(m) == 0 {
		return nil
	}

	i := 0
	results := make(KvSliceByKey, len(m))
	for key, value := range m {
		results[i] = Kv{Key: key, Value: value}
		i++
	}
	return results
}

// MapStringToKvSliceByKey converts map[string]string to KvSliceByKey.
//
// DEPRECATED!!!
func MapStringToKvSliceByKey(m map[string]string) KvSliceByKey {
	if len(m) == 0 {
		return nil
	}

	i := 0
	results := make(KvSliceByKey, len(m))
	for key, value := range m {
		results[i] = Kv{Key: key, Value: value}
		i++
	}
	return results
}

// MapToKvSliceByValue converts map[string]interface{} to KvSliceByValue.
//
// DEPRECATED!!!
func MapToKvSliceByValue(m map[string]interface{}) KvSliceByValue {
	if len(m) == 0 {
		return nil
	}

	i := 0
	results := make(KvSliceByValue, len(m))
	for key, value := range m {
		results[i] = Kv{Key: key, Value: value}
		i++
	}
	return results
}

// MapStringToKvSliceByValue converts map[string]string to KvSliceByValue.
//
// DEPRECATED!!!
func MapStringToKvSliceByValue(m map[string]string) KvSliceByValue {
	if len(m) == 0 {
		return nil
	}

	i := 0
	results := make(KvSliceByValue, len(m))
	for key, value := range m {
		results[i] = Kv{Key: key, Value: value}
		i++
	}
	return results
}
