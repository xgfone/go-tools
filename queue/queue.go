// Package queue supplies the queue function.
//
// You only have to implement the interface Queue.
//
// The package has implemented a memory queue based on Go channel.
package queue

import (
	"container/list"
	"fmt"
	"sync"
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

type listQueue struct {
	sync.Mutex

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
// Notice: the memory queue doesn't return an error forever.
func NewListQueue(size int) Queue {
	return &listQueue{
		cap:   size,
		lists: list.New(),

		getChan: make(chan struct{}, 1),
		putChan: make(chan struct{}, 1),
	}
}

// Get implements the method Get of the interface Queue.
func (l *listQueue) Get() (v interface{}, err error) {
	first := true
	for {
		if v, ok := l.get(first); ok {
			return v, nil
		}
		first = false
	}
}

func (l *listQueue) get(first bool) (v interface{}, ok bool) {
	l.Lock()
	if l.lists.Len() == 0 { // Empty
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
func (l *listQueue) Put(v interface{}) (err error) {
	first := true
	for !l.put(v, first) {
		first = false
	}
	return
}

func (l *listQueue) put(v interface{}, first bool) (ok bool) {
	l.Lock()
	if l.full() {
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
	return true
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
