// Package queue supplies the queue function.
//
// You only have to implement the interface Queue. The package has two
// implementations based on Channel and List.
package queue

import (
	"container/list"
	"fmt"
	"sync"
)

var (
	// ErrFull represents that the queue is full.
	ErrFull = fmt.Errorf("The queue is full")

	// ErrEmpty represents that the queue is empty.
	ErrEmpty = fmt.Errorf("The queue is empty")

	// ErrTimeout represents that the request is timeout.
	ErrTimeout = fmt.Errorf("The request is timeout")
)

// Queue is an queue interface.
type Queue interface {
	// Get returns an element from the queue.
	//
	// If the queue is empty, it maybe wait until an element is visiable,
	// or return ErrEmpty. If the implementation supports the timeout, it maybe
	// return ErrTimeout when the timeout arrives.
	//
	// Notice: Whether the method returns ErrEmpty or ErrTimeout or not, it is
	// decided by the implementation. And the implementation should indicate
	// which kind it supports.
	Get() (interface{}, error)

	// Put places a value as an element into the queue.
	//
	// If the queue is full, it maybe wait until the value can be put,
	// or return ErrFull to represent that the queue is full. If the
	// implementation supports the timeout, it maybe return ErrTimeout
	// when the timeout arrives.
	//
	// Notice: Whether the method returns ErrFull or ErrTimeout or not, it is
	// decided by the implementation. And the implementation should indicate
	// which kind it supports.
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
// the size is the size of the queue, which must be greater than 0; Or it will
// panic.
//
// When the size of the queue is very large, suggest to use NewListQueue,
// which is the queue based on the list, not pre-allocate the memory for the
// elements.
//
// Notice: the memory queue is the blocking model, and doesn't return an error
// forever.
func NewMemoryQueue(size int) Queue {
	if size < 1 {
		panic(fmt.Errorf("the queue size must be greater than 0"))
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

type listQueue struct {
	sync.Mutex
	block bool

	getWaiter int
	putWaiter int

	getChan chan struct{}
	putChan chan struct{}

	lists *list.List
	cap   int
}

// NewListQueue returns a Queue based on list.
//
// If the size is equal to or less than 0, the queue has no limit.
//
// If giving the second argument ant it's true, the queue doesn't return any
// error forerver. Or Get returns ErrEmpty when the queue is empty, and Put
// returns ErrFull when the queue is full. This implementation doesn't support
// the timeout, so both Put and Get don't return ErrTimeout.
//
// The queue does not pre-allocate the memory for the elements.
func NewListQueue(size int, block ...bool) Queue {
	var b bool
	if len(block) > 0 && block[0] {
		b = true
	}

	return &listQueue{
		cap:   size,
		lists: list.New(),
		block: b,

		getChan: make(chan struct{}, 1),
		putChan: make(chan struct{}, 1),
	}
}

// Get implements the method Get of the interface Queue.
func (l *listQueue) Get() (interface{}, error) {
	first := true
	for {
		if v, ok, err := l.get(first); err != nil {
			return nil, err
		} else if ok {
			return v, nil
		}
		first = false
	}
}

func (l *listQueue) get(first bool) (v interface{}, ok bool, err error) {
	l.Lock()
	if l.lists.Len() == 0 { // Empty
		if !l.block {
			err = ErrEmpty
			l.Unlock()
			return
		}

		if first {
			l.getWaiter++
		}
		l.Unlock()
		<-l.getChan
		return
	}

	v = l.lists.Remove(l.lists.Front())

	if !first {
		l.getWaiter--
	}

	waiter := l.putWaiter > 0
	l.Unlock()

	if waiter { // There are some goroutines to wait to get the value and wake it up.
		l.putChan <- struct{}{}
	}
	ok = true
	return
}

// Put implements the method Put of the interface Queue.
func (l *listQueue) Put(v interface{}) error {
	first := true
	for {
		if ok, err := l.put(v, first); err != nil {
			return err
		} else if ok {
			return nil
		}
		first = false
	}
}

func (l *listQueue) put(v interface{}, first bool) (ok bool, err error) {
	l.Lock()
	if l.full() {
		if !l.block {
			err = ErrFull
			l.Unlock()
			return
		}

		if first {
			l.putWaiter++ // Represent that there is a goroutine to wait to put a value.
		}
		l.Unlock()
		<-l.putChan
		return
	}

	l.lists.PushBack(v)

	if !first {
		l.putWaiter-- // Represent that the goroutine waiting to put a value has done.
	}

	waiter := l.getWaiter > 0
	l.Unlock()

	if waiter { // There are some goroutines to wait to get the value and wake it up.
		l.getChan <- struct{}{}
	}
	return true, nil
}

// Size implements the method Size of the interface Queue.
func (l *listQueue) Size() (v int, err error) {
	l.Lock()
	v = l.lists.Len()
	l.Unlock()
	return
}

// Full implements the method Full of the interface Queue.
func (l *listQueue) Full() (v bool, err error) {
	l.Lock()
	v = l.full()
	l.Unlock()
	return
}

func (l *listQueue) full() bool {
	if l.cap < 1 {
		return false
	}
	return l.lists.Len() == l.cap
}

// Empty implements the method Empty of the interface Queue.
func (l *listQueue) Empty() (v bool, err error) {
	l.Lock()
	v = l.lists.Len() == 0
	l.Unlock()
	return
}
