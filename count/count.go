// Atomic count.
package count

import "sync/atomic"

// The count type, whose operations are atomic.
type Count uint64

// NewCount create a new Count. The default is 0.
func NewCount() *Count {
	var c Count = 0
	return &c
}

// Add 1 to the current value.
func (c *Count) Add() *Count {
	return c.AddWith(1)
}

// Subtract 1 from the current value.
func (c *Count) Sub() *Count {
	return c.SubWith(1)
}

// Add v to the current value.
func (c *Count) AddWith(v uint64) *Count {
	atomic.AddUint64((*uint64)(c), v)
	return c
}

// Subtract v from the current value.
func (c *Count) SubWith(v uint64) *Count {
	atomic.AddUint64((*uint64)(c), ^uint64(v-1))
	return c
}

// Set the current value to v.
func (c *Count) Set(v uint64) *Count {
	atomic.StoreUint64((*uint64)(c), v)
	return c
}

// Get the current value.
func (c *Count) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
