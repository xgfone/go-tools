// Some atomic types, such as Bool.
package atomics

import "sync"

// The bool type, on which the operationsis atomic.
type Bool struct {
	sync.Mutex
	v bool
}

// New a new Bool variable. You can create it by &Bool{} directly.
func NewBool() *Bool {
	return &Bool{}
}

// Set it to true.
func (b *Bool) SetTrue() {
	b.Lock()
	defer b.Unlock()
	b.v = true
}

// Set it to false.
func (b *Bool) SetFalse() {
	b.Lock()
	defer b.Unlock()
	b.v = false
}

// Get its value.
func (b Bool) Get() bool {
	b.Lock()
	defer b.Unlock()
	return b.v
}
