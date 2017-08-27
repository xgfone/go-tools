// Package sync2 is the supplement of the standard library `sync`.
package sync2

import (
	"container/list"
	"sync"
)

// ResourceLock is a lock to lock a certain resource by the resource id.
type ResourceLock struct {
	locker    *sync.Mutex
	resources map[string]struct{}
	waiters   map[string]*list.List
}

// NewResourceLock returns a new ResourceLock.
func NewResourceLock() *ResourceLock {
	return &ResourceLock{
		locker:    new(sync.Mutex),
		resources: make(map[string]struct{}),
		waiters:   make(map[string]*list.List),
	}
}

// Lock locks a certain resource by its id.
//
// Once Lock is called, Unlock must be called later, that's, Lock and Unlock
// must appear in pairs.
func (r *ResourceLock) Lock(id string) {
	for !r.lock(id) {
	}
}

func (r *ResourceLock) lock(id string) bool {
	r.locker.Lock()
	// Lock successfully
	if _, ok := r.resources[id]; !ok {
		r.resources[id] = struct{}{}
		r.locker.Unlock()
		return true
	}

	// Fail to lock, add the wait to wake up.
	if _, ok := r.waiters[id]; !ok {
		r.waiters[id] = list.New()
	}
	wake := make(chan struct{})
	r.waiters[id].PushBack(func() {
		wake <- struct{}{}
	})
	r.locker.Unlock()
	<-wake
	return false
}

// Unlock unlocks a certain resource by its id.
//
// If the lock is not locked, it's successful.
func (r *ResourceLock) Unlock(id string) {
	r.locker.Lock()
	if _, ok := r.resources[id]; ok {
		delete(r.resources, id)
		if _, ok := r.waiters[id]; ok {
			e := r.waiters[id].Front()
			e.Value.(func())()
			r.waiters[id].Remove(e)

			if r.waiters[id].Len() == 0 {
				delete(r.waiters, id)
			}
		}
	}
	r.locker.Unlock()
}
