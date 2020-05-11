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

	"github.com/xgfone/go-tools/v6/lifecycle"
)

// DefaultSignals is the default signals to be handled.
var DefaultSignals = []os.Signal{
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGINT,
	os.Interrupt,
}

// HandleSignal is the same as HandleSignalWithCallback, but using lifecycle.Stop
// as the callback function.
//
// Notice: If the signals are empty, it is DefaultSignals by default.
func HandleSignal(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultSignals
	}
	HandleSignalWithCallback(lifecycle.Stop, signals...)
}

// HandleSignalWithCallback calls the callback function when the signals are received.
func HandleSignalWithCallback(cb func(), signals ...os.Signal) {
	if len(signals) == 0 {
		panic("no singals")
	}

	ss := make(chan os.Signal, 1)
	signal.Notify(ss, signals...)
	for {
		<-ss
		cb()
	}
}
