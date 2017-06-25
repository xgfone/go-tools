package tags

import (
	"reflect"
	"strconv"
	"strings"
)

// GetAllTags copys from the method Get() from reflect.StructTag.
// But it returns a map of Tag-Value.
//
// The type of the argument tag must be eithor string or reflect.StructTag.
// Or panic.
func GetAllTags(t interface{}) map[string]string {
	var tag string
	if _tag, ok := t.(reflect.StructTag); ok {
		tag = string(_tag)
	} else if _tag, ok := t.(string); ok {
		tag = _tag
	} else {
		panic("The type of the argument must be eithor string or reflect.StructTag")
	}

	tags := make(map[string]string)
	for tag != "" {
		// Strip the two-side whitespaces.
		tag = strings.Trim(tag, " \t\n")

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i := 0
		_len := len(tag)
		for i < _len && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] < 0x7f {
			i++
		}
		if i == 0 || i+1 >= _len || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		_len = len(tag)
		for i < _len && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= _len {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		if value, err := strconv.Unquote(qvalue); err == nil {
			if strings.TrimSpace(value) != "" {
				tags[name] = value
			}
		}
	}
	return tags
}
