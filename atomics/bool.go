// Package atomics supplies some atomic types, such as Bool.
package atomics

import "sync"

// Bool is a bool type, on which the operationsis atomic.
type Bool struct {
	sync.Mutex
	v bool
}

// NewBool returns a new Bool variable. You can create it by &Bool{} directly.
func NewBool() *Bool {
	return &Bool{}
}

// SetTrue sets it to true.
func (b *Bool) SetTrue() {
	b.Lock()
	defer b.Unlock()
	b.v = true
}

// SetFalse sets it to false.
func (b *Bool) SetFalse() {
	b.Lock()
	defer b.Unlock()
	b.v = false
}

// Get returns its value.
func (b *Bool) Get() bool {
	b.Lock()
	defer b.Unlock()
	return b.v
}
