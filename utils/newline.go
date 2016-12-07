// Some utility functions, which are classified to a certain package.
package utils

import "runtime"

var (
	Mac     = "\r"
	Unix    = "\n"
	Windows = "\r\n"
)

var newlines map[string]string

func init() {
	newlines = map[string]string{
		"windows": Windows,
		"darwin":  Mac,
		"linux":   Unix,
		"freebsd": Unix,
	}
}

// Return the newline of the current os. For example, windows' is "\r\n",
// linux's is "\n", Mac's is "\r", etc.
func NewLine() string {
	if v, ok := newlines[runtime.GOOS]; ok {
		return v
	} else {
		return Unix
	}
}
