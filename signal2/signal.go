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

// Package signal2 is the supplement of the standard library of `signal`,
// such as `HandleSignal`.
package signal2

import (
	"os"
	"os/signal"
	"syscall"

	"../lifecycle"
)

// DefaultSignals is the default signals to be handled.
var DefaultSignals = []os.Signal{
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGINT,
}

// HandleSignal is the same as HandleSignalWithLifecycle, but using the global
// default lifecycle manager.
//
// It's equal to
//   HandleSignalWithLifecycle(lifecycle.GetDefaultManager(), signals...)
//
// Notice: If the signals are empty, it will be equal to
//   HandleSignal(DefaultSignals...)
func HandleSignal(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultSignals
	}
	HandleSignalWithLifecycle(lifecycle.GetDefaultManager(), signals...)
}

// HandleSignalWithLifecycle wraps and handles the signals.
//
// The default wraps os.Interrupt. And you can pass the extra signals,
// syscall.SIGTERM, syscall.SIGQUIT, etc, such as
//   m := lifecycle.GetDefaultManager()
//   HandleSignalWithLifecycle(m, syscall.SIGTERM, syscall.SIGQUIT)
//
// For running it in a goroutine, use
//   go HandleSignalWithLifecycle(m, syscall.SIGTERM, syscall.SIGQUIT)
func HandleSignalWithLifecycle(m *lifecycle.Manager, signals ...os.Signal) {
	HandleSignalWithFunc(func() { m.Stop() }, os.Interrupt, signals...)
}

// HandleSignalWithFunc calls the function f when the signals are received.
func HandleSignalWithFunc(f func(), sig os.Signal, signals ...os.Signal) {
	ss := make(chan os.Signal, 1)
	signals = append(signals, sig)
	signal.Notify(ss, signals...)
	for {
		<-ss
		f()
	}
}
