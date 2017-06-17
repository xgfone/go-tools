// Package lifecycle offers a manager of the lifecycle of some apps in a program.
package lifecycle

import (
	"errors"
	"sync"
)

var (
	// ErrStopped is a stop error.
	ErrStopped = errors.New("The manager has been stopped")
)

// Manager manage the lifecycle of some apps in a program.
type Manager struct {
	sync.Mutex
	stoped     bool
	callbacks  []func()
	shouldStop chan struct{}
}

// NewManager returns a new LifeCycleManager.
func NewManager() *Manager {
	return &Manager{
		callbacks:  make([]func(), 0, 8),
		shouldStop: make(chan struct{}, 1),
	}
}

// RegisterChannel is same as Register, but using the channel, not the callback.
//
// The parameter in is used to notice the app to end. And out is used to notice
// the manager that the app has cleaned and ended successfully.
//
// NOTICE: The capacity of in and out must all be ZERO, that's, the two channels
// must be synchronized.
func (m *Manager) RegisterChannel(in chan<- interface{}, out <-chan interface{}) *Manager {
	if cap(in) != 0 || cap(out) != 0 {
		panic(errors.New("The capacity of the channel is not 0"))
	}

	return m.Register(func() {
		in <- true
		<-out
	})
}

// Register registers a callback function for the app.
//
// When calling Stop(), the callback function will be called in turn
// by the order that they are registered.
func (m *Manager) Register(f func()) *Manager {
	m.Lock()
	defer m.Unlock()

	if m.stoped {
		panic(ErrStopped)
	}

	m.callbacks = append(m.callbacks, f)
	return m
}

// Stop terminates and cleans all the apps.
//
// This method will be blocked until all the apps finish the clean.
// If the cleaning function of a certain app panics, ignore it and continue to
// call the cleaning function of the next app.
func (m *Manager) Stop() {
	m.Lock()
	defer m.Unlock()

	if m.stoped {
		return
	}

	for _, f := range m.callbacks {
		// f()
		callFuncAndIgnorePanic(f)
	}

	m.shouldStop <- struct{}{}
}

func callFuncAndIgnorePanic(f func()) {
	defer func() {
		recover()
	}()
	f()
}

// IsStop returns true if the manager has been stoped, or false.
func (m *Manager) IsStop() (yes bool) {
	m.Lock()
	yes = m.stoped
	m.Unlock()
	return
}

// RunForever is the same as m.Wait(), but it should be called in main goroutine
// to wait to exit the program.
func (m *Manager) RunForever() {
	if m.IsStop() {
		panic(ErrStopped)
	}

	<-m.shouldStop
}

// Wait will wait that the manager stops.
func (m *Manager) Wait() {
	if IsStop() {
		return
	}

	m.wait()
}

func (m *Manager) wait() {
	in := make(chan interface{})
	out := make(chan interface{})
	m.RegisterChannel(in, out)
	<-in
	out <- struct{}{}
}
