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
	return SplitSpaceN(s, -1)
}

// SplitString splits the string of s by sep, but is not the same as
// strings.Split(), which the rune in sep arbitrary combination. For example,
// SplitString("abcdefg-12345", "3-edc") == []string{"ab", "fg", "12", "45"}.
func SplitString(s string, sep string) []string {
	return SplitStringN(s, sep, -1)
}

// Split splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator.
func Split(s string, filter func(c rune) bool) []string {
	return SplitN(s, filter, -1)
}

// SplitSpaceN is the same as SplitStringN, but the whitespace.
func SplitSpaceN(s string, number int) []string {
	return SplitN(s, unicode.IsSpace, number)
}

// SplitStringN is the same as SplitN, but the separator is the string of sep.
func SplitStringN(s string, sep string, number int) []string {
	return SplitN(s, func(c rune) bool {
		for _, r := range sep {
			if r == c {
				return true
			}
		}
		return false
	}, number)
}

// SplitN splits the string of s by the filter. Split will pass each rune to the
// filter to determine whether it is the separator, but only number times.
//
// If number is equal to 0, don't split; greater than 0, only split number times;
// less than 0, don't limit. If the leading rune is the separator, it doesn't
// consume the split number.
//
// Notice: The result does not have the element of nil.
func SplitN(s string, filter func(c rune) bool, number int) []string {
	if number == 0 {
		return []string{s}
	}

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
	isNew := false
	for i, c := range s {
		if filter(c) {
			isNew = true
			continue
		}

		if isNew {
			results = append(results, buf.String())
			buf = bytes.NewBuffer(nil)
			isNew = false
			number--
			if number == 0 {
				buf.WriteString(s[i:])
				break
			}
		}

		buf.WriteRune(c)
	}

	last := buf.String()
	if len(last) > 0 {
		results = append(results, last)
	}

	return results
}
