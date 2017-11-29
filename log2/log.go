// Package log2 supplies a global debug and error log function, and a log
// function adaptor.
package log2

import (
	"fmt"
	"log"
)

var (
	// DebugF is used to output the debug logging by other packages,
	// which supplies a common way. The logging is outputted to os.Stderr
	// by default.
	DebugF func(format string, args ...interface{})

	// ErrorF is used to output the error logging by other packages,
	// which supplies a common way. The logging is outputted to os.Stderr
	// by default.
	ErrorF func(format string, args ...interface{})
)

// Logger is a logger interface based on the level.
//
// For the level, Critical > Error > Warning > Info > Debug.
type Logger interface {
	// Debug records the log as the debug level.
	Debug(format string, args ...interface{})

	// Info records the log as the info level.
	Info(format string, args ...interface{})

	// Warning records the log as the warning level.
	Warning(format string, args ...interface{})

	// Error records the log as the error level.
	Error(format string, args ...interface{})

	// Critical records the log as the critical level.
	Critical(format string, args ...interface{})
}

func init() {
	DebugF = func(format string, args ...interface{}) {
		format = fmt.Sprintf("[DEBUG] %s\n", format)
		log.Printf(format, args...)
	}

	ErrorF = func(format string, args ...interface{}) {
		format = fmt.Sprintf("[ERROR] %s\n", format)
		log.Printf(format, args...)
	}
}

// LoggerFunc converts a function to a suitable logger interface.
type LoggerFunc func(format string, args ...interface{})

// Printf implements the logger interface including the method Printf.
func (l LoggerFunc) Printf(format string, args ...interface{}) {
	l(format, args...)
}

// Debug implements the logger interface including the method Debug.
func (l LoggerFunc) Debug(format string, args ...interface{}) {
	l(format, args...)
}

// Debugf implements the logger interface including the method Debugf.
func (l LoggerFunc) Debugf(format string, args ...interface{}) {
	l(format, args...)
}

// Info implements the logger interface including the method Info.
func (l LoggerFunc) Info(format string, args ...interface{}) {
	l(format, args...)
}

// Infof implements the logger interface including the method Infof.
func (l LoggerFunc) Infof(format string, args ...interface{}) {
	l(format, args...)
}

// Warningf implements the logger interface including the method Warningf.
func (l LoggerFunc) Warningf(format string, args ...interface{}) {
	l(format, args...)
}

// Warning implements the logger interface including the method Warning.
func (l LoggerFunc) Warning(format string, args ...interface{}) {
	l(format, args...)
}

// Error implements the logger interface including the method Error.
func (l LoggerFunc) Error(format string, args ...interface{}) {
	l(format, args...)
}

// Errorf implements the logger interface including the method Errorf.
func (l LoggerFunc) Errorf(format string, args ...interface{}) {
	l(format, args...)
}

// Warnf implements the logger interface including the method Warnf.
func (l LoggerFunc) Warnf(format string, args ...interface{}) {
	l(format, args...)
}

// Warn implements the logger interface including the method Warn.
func (l LoggerFunc) Warn(format string, args ...interface{}) {
	l(format, args...)
}
