// Package lifecycle offers a manager of the lifecycle of some apps in a program.
package lifecycle

import (
	"errors"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
)

var (
	// ErrStopped is a stop error.
	ErrStopped = errors.New("The lifecycle manager has been stopped")

	// ErrSameArgs is a arguments error.
	ErrSameArgs = errors.New("The arguments is the same")
)

// Manager manage the lifecycle of some apps in a program.
type Manager struct {
	lock       sync.RWMutex
	stoped     int32
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

func (m *Manager) getCallbacks() []func() {
	m.lock.RLock()
	cb := make([]func(), 0, len(m.callbacks))
	copy(cb, m.callbacks)
	m.lock.RUnlock()
	return cb
}

func (m *Manager) addCallbacks(cb ...func()) *Manager {
	m.lock.Lock()
	m.callbacks = append(m.callbacks, cb...)
	m.lock.Unlock()
	return m
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
func (m *Manager) Register(functions ...func()) *Manager {
	if m.IsStop() {
		panic(ErrStopped)
	}
	return m.addCallbacks(functions...)
}

// Stop terminates and cleans all the apps.
//
// This method will be blocked until all the apps finish the clean.
// If the cleaning function of a certain app panics, ignore it and continue to
// call the cleaning function of the next app.
func (m *Manager) Stop() {
	if atomic.CompareAndSwapInt32(&m.stoped, 0, 1) {
		functions := m.getCallbacks()
		for _len := len(functions) - 1; _len >= 0; _len-- {
			callFuncAndIgnorePanic(functions[_len])
		}
		m.shouldStop <- struct{}{}
	}
}

func callFuncAndIgnorePanic(f func()) {
	defer func() {
		recover()
	}()

	if f != nil {
		f()
	}
}

// IsStop returns true if the manager has been stoped, or false.
func (m *Manager) IsStop() bool {
	return atomic.LoadInt32(&m.stoped) != 0
}

// RunForever is the same as m.Wait(), but it should be called in main goroutine
// to wait to exit the program.
func (m *Manager) RunForever() {
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
	exit := make(chan struct{}, 1)
	finished := make(chan struct{}, 1)
	m.addCallbacks(func() { exit <- struct{}{}; <-finished })

	<-exit // Wait that the manager stops.
	// Here can do some cleanup works.
	finished <- struct{}{} // Notify the manager that the task finished.
}

// Exit executes the stop functions and exit the program with the code
// by calling os.Exit(code).
func (m *Manager) Exit(code int) {
	m.Stop()
	os.Exit(code)
}
