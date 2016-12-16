// Package str is the supplement of the standard library of strings.
package str

import (
	"bytes"
	"unicode"
)

// SplitSpace splits the string of s by the whitespace, which is equal to
// str.split() in Python.
//
// Notice: SplitSpace(s) == Split(s, unicode.IsSpace).
func SplitSpace(s string) []string {
	return Split(s, unicode.IsSpace)
}

// SplitString splits the string of s by sep, but is not the same as
// strings.Split(), which the rune in sep arbitrary combination. For example,
// SplitString("abcdefg-12345", "3-edc") == []string{"ab", "fg", "12", "45"}.
func SplitString(s string, sep string) []string {
	return Split(s, func(c rune) bool {
		for _, r := range sep {
			if r == c {
				return true
			}
		}
		return false
	})
}

// Split splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator.
func Split(s string, filter func(c rune) bool) []string {
	for i, c := range s {
		if !filter(c) {
			s = s[i:]
			break
		}
	}

	if len(s) == 0 {
		return nil
	}

	results := make([]string, 0)
	buf := bytes.NewBuffer(nil)
	is_new := false
	for _, c := range s {
		if filter(c) {
			is_new = true
			continue
		}

		if is_new {
			results = append(results, buf.String())
			buf = bytes.NewBuffer(nil)
			is_new = false
		}

		buf.WriteRune(c)
	}

	last := buf.String()
	if len(last) > 0 {
		results = append(results, last)
	}

	return results
}
