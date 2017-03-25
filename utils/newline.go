// Package utils supplys some utility functions, which are classified to a certain package.
package utils

import "runtime"

var (
	// MacNL is the newline in Mac.
	MacNL = "\r"

	// UnixNL is the newline in Unix/Linux.
	UnixNL = "\n"

	// WindowsNL is the newline in Windows.
	WindowsNL = "\r\n"
)

var newlines map[string]string

func init() {
	newlines = map[string]string{
		"windows": WindowsNL,
		"darwin":  MacNL,
		"linux":   UnixNL,
		"freebsd": UnixNL,
	}
}

// NewLine returns the newline of the current os. For example, windows' is "\r\n",
// linux's is "\n", Mac's is "\r", etc.
func NewLine() string {
	if v, ok := newlines[runtime.GOOS]; ok {
		return v
	}
	return UnixNL
}
