// Package os2 is the supplement of the standard library of `os`, such as `Exit`.
package os2

import (
	"os"

	"github.com/xgfone/go-tools/lifecycle"
)

// Exit will call lifecycle.Stop() before calling os.Exit(code).
//
// DEPRECATED!!!
func Exit(code int) {
	lifecycle.Stop()
	os.Exit(code)
}
