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

package atexit

import "context"

// DefaultManager is the default global atexit manager.
var DefaultManager = NewManager()

// PushBack is equal to DefaultManager.PushBack(functions...).
func PushBack(functions ...func()) { DefaultManager.PushBack(functions...) }

// PushFront is equal to DefaultManager.PushFront(functions...).
func PushFront(functions ...func()) { DefaultManager.PushFront(functions...) }

// Stop is equal to DefaultManager.Stop().
func Stop() { DefaultManager.Stop() }

// Exit is equal to DefaultManager.Exit(code).
func Exit(code int) { DefaultManager.Exit(code) }

// Wait is equal to DefaultManager.Wait().
func Wait() { DefaultManager.Wait() }

// IsStopped is equal to DefaultManager.IsStopped().
func IsStopped() bool { return DefaultManager.IsStopped() }

// Done is equal to DefaultManager.Done().
func Done() <-chan struct{} { return DefaultManager.Done() }

// Context is equal to DefaultManager.Context().
func Context() context.Context { return DefaultManager.Context() }
