package server

import (
	"fmt"
	"log"
	"os"
)

var (
	// NotOutputLog is on/off. If true, don't output the any log.
	NotOutputLog bool

	// NotOutputNewLine is on/off. If true, don't append a new line to the log.
	NotOutputNewLine bool
)

var _logger logger

func init() {
	_logger = logger{_logger: log.New(os.Stderr, "[TCP Server] ", log.LstdFlags)}
}

// Logger is a interface to implement a logger.
//
// DEPRECATED!!! Please ErrorF and DebugF in the sub-package log instead.
// It is only reserved, and not effect.
type Logger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
	_logger Logger
}

func (l logger) Printf(format string, v ...interface{}) {
	if !NotOutputLog {
		if !NotOutputNewLine {
			format = fmt.Sprintln(format)
		}
		l._logger.Printf(format, v...)
	}
}

// SetLogger replaces the default logger.
//
// DEPRECATED!!! Please ErrorF and DebugF in the sub-package log instead.
// It is only reserved, and not effect.
func SetLogger(logger Logger) {
	_logger._logger = logger
}

// GetLogger returns the current logger.
//
// DEPRECATED!!! Please ErrorF and DebugF in the sub-package log instead.
// It is only reserved, and not effect.
func GetLogger() Logger {
	return _logger._logger
}
