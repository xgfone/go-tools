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
//
// DEPRECATED!!! Please use the sub-package atexit.
package lifecycle

import (
	"context"
	"os"
	"reflect"

	"github.com/xgfone/go-tools/v7/atexit"
)

// Manager manage the lifecycle of some apps in a program.
type Manager struct {
	m *atexit.Manager
}

// NewManager returns a new lifecycle manager.
func NewManager() *Manager { return &Manager{m: atexit.NewManager()} }

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
		panic("out must not be equal to in")
	}

	return m.Register(func() {
		if out != nil {
			out <- struct{}{}
		}
		if in != nil {
			<-in
		}
	})
}

// Register registers a callback function..
//
// When calling Stop(), the callback function will be called in turn
// by the order that they are registered.
func (m *Manager) Register(functions ...func()) *Manager {
	m.m.PushBack(functions...)
	return m
}

// PrefixRegister is the same as Register, but adding the callback function
// before others.
func (m *Manager) PrefixRegister(functions ...func()) *Manager {
	m.m.PushFront(functions...)
	return m
}

// Stop stops the manager and calls the registered functions.
//
// This method will be blocked until all the apps finish the clean.
// If the cleaning function of a certain app panics, ignore it and continue to
// call the cleaning function of the next app.
func (m *Manager) Stop() { m.m.Stop() }

// IsStop returns true if the manager has been stoped, or false.
func (m *Manager) IsStop() bool { return m.m.IsStopped() }

// Done returns a channel to report whether the manager is stopped, that,s,
// the channel will be closed when the manager is stopped.
func (m *Manager) Done() <-chan struct{} { return m.m.Done() }

// Context returns a Context, which will be canceled when the manager is stopped.
func (m *Manager) Context() context.Context { return m.m.Context() }

// RunForever is equal to Wait().
func (m *Manager) RunForever() { m.Wait() }

// Wait will wait that the manager stops.
func (m *Manager) Wait() { m.m.Wait() }

// Exit stops the manager and exit the program by calling os.Exit(code).
func (m *Manager) Exit(code int) { m.Stop(); os.Exit(code) }
