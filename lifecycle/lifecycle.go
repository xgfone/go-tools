// Package lifecycle offers a manager of the lifecycle of some apps in a program.
package lifecycle

import (
	"errors"
	"reflect"
	"sync"
)

var (
	// ErrStopped is a stop error.
	ErrStopped = errors.New("The manager has been stopped")

	// ErrSameArgs is a arguments error.
	ErrSameArgs = errors.New("The arguments is the same")
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
// The parameter out is used to notice the app to end. And in is used to notice
// the manager that the app has cleaned and ended successfully.
//
// NOTICE: the tow parameters must not be a same channel.
//
// Exmaple: See the wait method.
func (m *Manager) RegisterChannel(out chan<- interface{}, in <-chan interface{}) *Manager {
	if reflect.ValueOf(in).Pointer() == reflect.ValueOf(out).Pointer() {
		panic(ErrSameArgs)
	}

	return m.Register(func() {
		out <- struct{}{}
		<-in
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
	exit := make(chan interface{}, 1)
	finished := make(chan interface{}, 1)
	m.RegisterChannel(exit, finished)

	<-exit // Wait that the manager stops.
	// Here can do some cleanup works.
	finished <- struct{}{} // Notify the manager that the task finished.
}
