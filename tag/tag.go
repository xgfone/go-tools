package tag

import (
	"reflect"
	"strconv"
	"strings"
)

// GetFieldTags returns all tags in a fields of a struct.
//
// If the type of v is not string or reflect.StructTag, return nil.
//
// For each element in the returned slice, it's [2]string{TagName, TagValue}.
func GetFieldTags(t interface{}) [][2]string {
	var tag string
	switch v := t.(type) {
	case reflect.StructTag:
		tag = string(v)
	case string:
		tag = v
	default:
		return nil
	}

	tags := make([][2]string, 0)
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
				tags = append(tags, [2]string{name, value})
			}
		}
	}
	return tags
}

// GetStructTags returns all tags in all the fields of a struct.
//
// If v is not a struct or a pointer to struct, return nil.
//
// For each element in the returned slice,
// it's [3]string{FieldName, TagName, TagValue}.
func GetStructTags(v interface{}) [][3]string {
	_type := reflect.TypeOf(v)
	if _type.Kind() == reflect.Ptr {
		_type = _type.Elem()
	}
	if _type.Kind() != reflect.Struct {
		return nil
	}

	fieldNum := _type.NumField()
	tags := make([][3]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		field := _type.Field(i)
		fieldName := field.Name
		for _, tag := range GetFieldTags(field.Tag) {
			tags = append(tags, [3]string{fieldName, tag[0], tag[1]})
		}
	}
	return tags
}
