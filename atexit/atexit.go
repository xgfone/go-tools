// Copyright 2021 xgfone
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

// Package atexit is used to manage the exit of a certain lifecycle.
package atexit

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
)

// Manager is used to manage the exit of a certain lifecycle.
type Manager struct {
	lock  sync.RWMutex
	funcs []func()

	stoped int32
	stop   chan struct{}
	done   chan struct{}
	ctx    context.Context
	cancel func()
}

// NewManager returns a new LifeCycleManager.
func NewManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	return &Manager{
		funcs:  make([]func(), 0, 8),
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

// PushBack appends a callback function to the end, which is executed
// when it is stopped.
func (m *Manager) PushBack(functions ...func()) {
	if m.IsStopped() {
		panic("the atexit manager has been stopped")
	}

	if len(functions) == 0 {
		return
	}

	m.lock.Lock()
	m.funcs = append(m.funcs, functions...)
	m.lock.Unlock()
}

// PushFront inserts a callback function to the front, which is executed
// when it is stopped.
func (m *Manager) PushFront(functions ...func()) {
	if m.IsStopped() {
		panic("the atexit manager has been stopped")
	}

	_len := len(functions)
	if _len == 0 {
		return
	}

	m.lock.Lock()
	funcs := make([]func(), _len+len(m.funcs))
	copy(funcs, functions)
	copy(funcs[_len:], m.funcs)
	m.funcs = funcs
	m.lock.Unlock()
}

// Stop stops the manager, which will call all the registered functions.
func (m *Manager) Stop() {
	if atomic.CompareAndSwapInt32(&m.stoped, 0, 1) {
		m.lock.RLock()
		defer m.lock.RUnlock()

		for _len := len(m.funcs) - 1; _len >= 0; _len-- {
			callFuncAndIgnorePanic(m.funcs[_len])
		}

		m.cancel()
		close(m.done)
		close(m.stop)
	}
}

func callFuncAndIgnorePanic(f func()) { defer recover(); f() }

// Exit stops the manager and exit the program with the code.
func (m *Manager) Exit(code int) { m.Stop(); os.Exit(code) }

// Wait will wait that the manager stops.
func (m *Manager) Wait() { <-m.stop }

// IsStopped reports whether the manager has been stoped.
func (m *Manager) IsStopped() bool { return atomic.LoadInt32(&m.stoped) != 0 }

// Done returns a channel, which will be closed when the manager is stopped.
func (m *Manager) Done() <-chan struct{} { return m.done }

// Context returns a Context, which will be canceled when the manager is stopped.
func (m *Manager) Context() context.Context { return m.ctx }
