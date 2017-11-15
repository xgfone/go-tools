// Package sync2 is the supplement of the standard library `sync`.
package sync2

import (
	"container/list"
	"fmt"
	"sync"
)

// ResourceLock is an interface about the resource lock.
type ResourceLock interface {
	// Lock locks the resource, which id is the resource id.
	//
	// When the resource named id has been locked, it should wait to be unlocked
	// by the locker.
	//
	// If id is empty, it should panic.
	Lock(id string)

	// Unlock unlocks the resource, which id is the resource id.
	//
	// If the lock is not locked or id is empty, it should panic.
	Unlock(id string)
}

type hashResourceLock map[uint8]resourceLock

func (h hashResourceLock) getResource(id string) resourceLock {
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
// Notice: this implementation use the quadratic hash. The first hash does not
// use the lock, so it maybe have a better performance when the resources are
// not the same.
func NewHashResourceLock() ResourceLock {
	r := make(hashResourceLock, 256)
	for i := 0; i < 256; i++ {
		r[uint8(i)] = newResourceLock()
	}
	return r
}

// NewResourceLock is the alias of NewHashResourceLock.
//
// This is used to be compatible with the old.
func NewResourceLock() ResourceLock {
	return NewHashResourceLock()
}

// resourceLock is a lock to lock a certain resource by the resource id.
type resourceLock struct {
	locker    *sync.Mutex
	resources map[string]struct{}
	waiters   map[string]*list.List
}

// newResourceLock returns a new resourceLock.
func newResourceLock() resourceLock {
	return resourceLock{
		locker:    new(sync.Mutex),
		resources: make(map[string]struct{}),
		waiters:   make(map[string]*list.List),
	}
}

// Lock locks a certain resource by its id.
//
// Once Lock is called, Unlock must be called later, that's, Lock and Unlock
// must appear in pairs.
func (r resourceLock) Lock(id string) {
	for !r.lock(id) {
	}
}

func (r resourceLock) lock(id string) bool {
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
		waiter = list.New()
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

// Unlock unlocks a certain resource by its id.
//
// If the lock is not locked, it will panic.
func (r resourceLock) Unlock(id string) {
	r.locker.Lock()
	if _, ok := r.resources[id]; ok {
		delete(r.resources, id)
		if waiter, ok := r.waiters[id]; ok {
			e := waiter.Front()
			e.Value.(func())()
			waiter.Remove(e)

			if waiter.Len() == 0 {
				delete(r.waiters, id)
			}
		}
		r.locker.Unlock()
		return
	}
	r.locker.Unlock()
	panic(fmt.Errorf("The lock[%s] is not locked", id))
}
