// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package lifecycle offers a manager of the lifecycle of some apps in a program.
package lifecycle

import (
	"context"
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
	done       chan struct{}
	ctx        context.Context
	cancel     func()
}

// NewManager returns a new LifeCycleManager.
func NewManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	return &Manager{
		callbacks:  make([]func(), 0, 8),
		shouldStop: make(chan struct{}),
		done:       make(chan struct{}),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// RegisterChannel is same as Register, but using the channel, not the callback.
//
// The parameter out is used to notice the app to end. And in is used to notice
// the manager that the app has cleaned and ended successfully. They may be nil.
//
// NOTICE: the two parameters must not be a same channel.
//
// Exmaple: See the wait method.
func (m *Manager) RegisterChannel(out chan<- interface{}, in <-chan interface{}) *Manager {
	if reflect.ValueOf(in).Pointer() == reflect.ValueOf(out).Pointer() {
		panic(ErrSameArgs)
	}

	outf := func() {}
	if out != nil {
		outf = func() { out <- struct{}{} }
	}

	inf := func() {}
	if in != nil {
		inf = func() { <-in }
	}

	return m.Register(func() {
		outf()
		inf()
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

	m.lock.Lock()
	m.callbacks = append(m.callbacks, functions...)
	m.lock.Unlock()
	return m
}

// PrefixRegister is the same as Register, but adding the callback function
// before others.
func (m *Manager) PrefixRegister(functions ...func()) *Manager {
	if m.IsStop() {
		panic(ErrStopped)
	}

	callbacks := append([]func(){}, functions...)

	m.lock.Lock()
	callbacks = append(callbacks, m.callbacks...)
	m.callbacks = callbacks
	m.lock.Unlock()
	return m
}

// Stop terminates and cleans all the apps.
//
// This method will be blocked until all the apps finish the clean.
// If the cleaning function of a certain app panics, ignore it and continue to
// call the cleaning function of the next app.
func (m *Manager) Stop() {
	if atomic.CompareAndSwapInt32(&m.stoped, 0, 1) {
		m.lock.RLock()
		defer m.lock.RUnlock()

		for _len := len(m.callbacks) - 1; _len >= 0; _len-- {
			callFuncAndIgnorePanic(m.callbacks[_len])
		}
		m.cancel()
		close(m.done)
		close(m.shouldStop)
	}
}

func callFuncAndIgnorePanic(f func()) {
	defer recover()
	if f != nil {
		f()
	}
}

// IsStop returns true if the manager has been stoped, or false.
func (m *Manager) IsStop() bool {
	return atomic.LoadInt32(&m.stoped) != 0
}

// Done returns a channel to report whether the manager is stopped, that,s,
// the channel will be closed when the manager is stopped.
func (m *Manager) Done() <-chan struct{} {
	return m.done
}

// Context returns a Context, which will be canceled when the manager is stopped.
func (m *Manager) Context() context.Context {
	return m.ctx
}

// RunForever is the same as m.Wait(), but it should be called in main goroutine
// to wait to exit the program.
func (m *Manager) RunForever() {
	<-m.shouldStop
}

// Wait will wait that the manager stops.
func (m *Manager) Wait() {
	m.lock.Lock()
	if m.IsStop() {
		m.lock.Unlock()
		return
	}

	exit := make(chan struct{}, 1)
	finished := make(chan struct{}, 1)
	m.callbacks = append([]func(){func() { exit <- struct{}{}; <-finished }}, m.callbacks...)
	m.lock.Unlock()

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
