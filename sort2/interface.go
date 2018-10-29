package sort2

import "sort"

type interfaceSlice struct {
	data []interface{}
	less func(interface{}, interface{}) bool
}

func (s interfaceSlice) Len() int {
	return len(s.data)
}

func (s interfaceSlice) Less(i, j int) bool {
	return s.less(s.data[i], s.data[j])
}

func (s interfaceSlice) Swap(i, j int) {
	tmp := s.data[i]
	s.data[i] = s.data[j]
	s.data[j] = tmp
}

// InterfaceSlice sorts the interface slice, then returns the sorted slice.
//
// The elements in data should have the same type, or it maybe panic, which
// depends on the less function.
//
// less is a function to compare the two elements of the slice data,
// which returns true if the first is less than the second, or returns false.
// For the type, byte, rune, int, uint, int8, int16, int32, int64,
// uint8, uint16, uint32, uint64, float32, float64, string, you can use the
// function LT in the sub-package function, that's, function.LT. See example.
//
// If giving the second argument and it's true, sort the data in reverse.
func InterfaceSlice(data []interface{}, less func(first, second interface{}) bool) {
	if len(data) > 1 {
		sort.Sort(interfaceSlice{data: data, less: less})
	}
}
