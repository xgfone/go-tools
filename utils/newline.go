package utils

import "runtime"

var (
	default_nls = "\n"

	newlines map[string]string
)

func init() {
	newlines = map[string]string{
		"windows": "\r\n",
		"darwin":  "\r",
		"linux":   "\n",
		"freebsd": "\n",
	}
}

// Return the newline of the current os. For example, windows' is "\r\n",
// linux's is "\n", Mac's is "\r", etc.
func NewLine() string {
	if v, ok := newlines[runtime.GOOS]; ok {
		return v
	} else {
		return default_nls
	}
}
