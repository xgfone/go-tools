// Package queue supplies the queue function.
//
// You only have to implement the interface Queue.
//
// The package has implemented a memory queue based on Go channel.
package queue

import (
	"fmt"
)

// Queue is an queue interface.
type Queue interface {
	// Get returns an element from the queue.
	//
	// It will wait until an element is visible.
	Get() (interface{}, error)

	// Put places a value as an element into the queue.
	//
	// If the queue has the size limit, it will be blocked when the queue is full.
	Put(interface{}) error

	// Size returns the number of the elements in the queue.
	Size() (int, error)

	// Full returns true if the queue is full, or false.
	Full() (bool, error)

	// Empty returns true if the queue has no elements, or false.
	Empty() (bool, error)
}

type memoryQueue struct {
	cap    int
	caches chan interface{}
}

// NewMemoryQueue returns a new Queue based on the memory.
//
// the size is the size of the queue. if it's 0, the queue is a synchronized
// queue, which you can equate it with channel. If the size is a negative,
// it will panic.
//
// Notice: the memory queue doesn't return an error forever.
func NewMemoryQueue(size int) Queue {
	if size < 0 {
		panic(fmt.Errorf("the queue size must not be negative"))
	}
	return memoryQueue{cap: size, caches: make(chan interface{}, size)}
}

// Get implements the method Get of the interface Queue.
func (m memoryQueue) Get() (interface{}, error) {
	return <-m.caches, nil
}

// Put implements the method Put of the interface Queue.
func (m memoryQueue) Put(v interface{}) error {
	m.caches <- v
	return nil
}

// Size implements the method Size of the interface Queue.
func (m memoryQueue) Size() (int, error) {
	return len(m.caches), nil
}

// Full implements the method Full of the interface Queue.
func (m memoryQueue) Full() (bool, error) {
	return len(m.caches) == m.cap, nil
}

// Empty implements the method Empty of the interface Queue.
func (m memoryQueue) Empty() (bool, error) {
	return len(m.caches) == 0, nil
}
