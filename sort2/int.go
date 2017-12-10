package sort2

// Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int64Slice) Sort() { Sort(p) }

// Uint64Slice attaches the methods of Interface to []uint64, sorting in increasing order.
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint64Slice) Sort() { Sort(p) }

// UintSlice attaches the methods of Interface to []uint, sorting in increasing order.
type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p UintSlice) Sort() { Sort(p) }

// Int64s sorts a slice of []int64 in increasing order.
func Int64s(a []int64) { Sort(Int64Slice(a)) }

// Uint64s sorts a slice of []uint64 in increasing order.
func Uint64s(a []uint64) { Sort(Uint64Slice(a)) }

// Uints sorts a slice of []uint in increasing order.
func Uints(a []uint) { Sort(UintSlice(a)) }
