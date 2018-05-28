// Package strings2 is the supplement of the standard library of strings.
package strings2

import (
	"fmt"
	"strings"
)

var (
	leftDelimiter  = "{"
	rightDelimiter = "}"
)

// SetFmtDelimiter sets the delimiters which are used by KvFmt.
//
// The left delimiter is "{", and the right delimiter is "}".
func SetFmtDelimiter(left, right string) {
	if left == "" || right == "" {
		panic("The arguments cannot be empty")
	}
	leftDelimiter = left
	rightDelimiter = right
}

// KvFmt formats the string like the key-value method format of str in Python,
// which the placeholder is appointed by the key name of the values.
//
// Notice: the formatter will use %v to convert the value of the key to string.
// The delimiters are "{" and "}" by default, and you can reset them by the
// function SetFmtDelimiter.
func KvFmt(s string, values map[string]interface{}) string {
	for key, value := range values {
		key = fmt.Sprintf("%s%s%s", leftDelimiter, key, rightDelimiter)
		s = strings.Replace(s, key, fmt.Sprintf("%v", value), -1)
	}
	return s
}
