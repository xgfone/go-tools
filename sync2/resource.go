package sync2

import (
	"container/list"
	"fmt"
	"sync"
)

var (
	// ErrEmpty represents the error that the resource id is empty.
	ErrEmpty = fmt.Errorf("The resource id is empty")
)

// ResourceLock is an interface about the resource lock.
//
// DEPRECATED!!!
type ResourceLock interface {
	// Lock locks the resource, which id is the resource id.
	//
	// When the resource named id has been locked, it should wait to be unlocked
	// by the locker.
	//
	// If id is empty, it should panic, the value of which is ErrEmpty.
	Lock(id string)

	// Unlock unlocks the resource, which id is the resource id.
	//
	// If the lock is not locked or id is empty, it should panic, the value
	// of which is ErrEmpty.
	Unlock(id string)
}

type hashResourceLock map[uint8]BaseResourceLock

func (h hashResourceLock) getResource(id string) BaseResourceLock {
	if id == "" {
		panic(ErrEmpty)
	}
	return h[id[len(id)-1]]
}

func (h hashResourceLock) Lock(id string) {
	h.getResource(id).Lock(id)
}

func (h hashResourceLock) Unlock(id string) {
	h.getResource(id).Unlock(id)
}

// NewHashResourceLock returns a new ResourceLock based on hash.
//
// The implementation uses the once hash based on BaseResourceLock, that's,
// the implementation uses the quadratic hash. The first hash uses 256 buckets,
// which uses the last byte of the resource id as the hash key and has no lock.
// So it maybe have a better performance.
//
// DEPRECATED!!!
func NewHashResourceLock() ResourceLock {
	r := make(hashResourceLock, 256)
	for i := 0; i < 256; i++ {
		r[uint8(i)] = NewBaseResourceLock()
	}
	return r
}

// NewResourceLock is the alias of NewHashResourceLock.
//
// This is used to be compatible with the old.
//
// DEPRECATED!!!
func NewResourceLock() ResourceLock {
	return NewHashResourceLock()
}

// BaseResourceLock is a lock to lock a certain resource by the resource id.
//
// DEPRECATED!!!
type BaseResourceLock struct {
	locker     *sync.Mutex
	resources  map[string]struct{}
	waiters    map[string]*list.List
	waiterPool *sync.Pool
}

// NewBaseResourceLock returns a new BaseResourceLock.
//
// DEPRECATED!!!
func NewBaseResourceLock() BaseResourceLock {
	return BaseResourceLock{
		locker:    new(sync.Mutex),
		resources: make(map[string]struct{}),
		waiters:   make(map[string]*list.List),
		waiterPool: &sync.Pool{New: func() interface{} {
			return list.New()
		}},
	}
}

// Lock implements the method Lock of ResourceLock.
//
// Once Lock is called, Unlock should be called later, that's, Lock and Unlock
// should appear in pairs.
func (r BaseResourceLock) Lock(id string) {
	if id == "" {
		panic(ErrEmpty)
	}

	for !r.lock(id) {
	}
}

func (r BaseResourceLock) lock(id string) bool {
	r.locker.Lock()
	// Lock successfully
	if _, ok := r.resources[id]; !ok {
		r.resources[id] = struct{}{}
		r.locker.Unlock()
		return true
	}

	// Fail to lock, add the wait to wake up.
	waiter, ok := r.waiters[id]
	if !ok {
		waiter = r.waiterPool.Get().(*list.List)
		r.waiters[id] = waiter
	}
	wake := make(chan struct{})
	waiter.PushBack(func() {
		wake <- struct{}{}
	})
	r.locker.Unlock()
	<-wake
	return false
}

// Unlock implements the method Unlock of ResourceLock.
//
// If the lock is not locked, it will panic.
func (r BaseResourceLock) Unlock(id string) {
	if id == "" {
		panic(ErrEmpty)
	}

	r.locker.Lock()
	if _, ok := r.resources[id]; ok {
		delete(r.resources, id)
		if waiter, ok := r.waiters[id]; ok {
			e := waiter.Front()
			e.Value.(func())()
			waiter.Remove(e)

			if waiter.Len() == 0 {
				delete(r.waiters, id)
				r.waiterPool.Put(waiter)
			}
		}
		r.locker.Unlock()
		return
	}
	r.locker.Unlock()
	panic(fmt.Errorf("The lock[%s] is not locked", id))
}
