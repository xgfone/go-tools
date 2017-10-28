// Package signal2 is the supplement of the standard library of `signal`,
// such as `HandleSignal`.
package signal2

import (
	"os"
	"os/signal"

	"github.com/xgfone/go-tools/lifecycle"
)

// HandleSignal is the same as HandleSignalWithLifecycle, but using the global
// default lifecycle manager.
//
// It's equal to
//   HandleSignalWithLifecycle(lifecycle.GetDefaultManager(), signals...)
func HandleSignal(signals ...os.Signal) {
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
	ss := make(chan os.Signal, 1)
	signals = append(signals, os.Interrupt)
	signal.Notify(ss, signals...)
	<-ss
	m.Stop()
}
