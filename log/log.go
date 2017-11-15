package log

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
