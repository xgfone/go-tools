package tags

import (
	"reflect"
	"strconv"
	"strings"
)

// GetAllTags copys from the method Get() from reflect.StructTag.
// But it returns all the tags defined in the fields by the order
// that they appear, not the value of the specific tag.
//
// The type of the argument tag must be eithor string or reflect.StructTag.
// Or panic.
func GetAllTags(t interface{}) []TV {
	var tag reflect.StructTag
	if _tag, ok := t.(reflect.StructTag); ok {
		tag = _tag
	} else if _tag, ok := t.(string); ok {
		tag = reflect.StructTag(_tag)
	} else {
		panic("The type of the argument must be eithor string or reflect.StructTag")
	}

	_tags := make([]TV, 0)
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		if value, err := strconv.Unquote(qvalue); err == nil {
			if strings.TrimSpace(value) != "" {
				_tags = append(_tags, TV{Tag: name, Value: value})
			}
		}
	}
	return _tags
}
