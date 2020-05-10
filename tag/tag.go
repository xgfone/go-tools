// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tag is used to find and get the tags in a struct.
package tag

import (
	"reflect"
	"strconv"
	"strings"
)

// GetFieldTags returns all tags in a fields of a struct.
// The key is the tag name, and the value is the tag value.
//
// If the type of v is not string or reflect.StructTag, return nil.
func GetFieldTags(t interface{}) map[string]string {
	var tag string
	switch v := t.(type) {
	case reflect.StructTag:
		tag = string(v)
	case string:
		tag = v
	default:
		return nil
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

// GetStructTags returns all tags in all the fields of a struct.
//
// The returned type is map[FieldName]map[TagName]TagValue.
//
// If v is not a struct or a pointer to struct, return nil.
func GetStructTags(v interface{}) map[string]map[string]string {
	_type := reflect.TypeOf(v)
	if _type.Kind() == reflect.Ptr {
		_type = _type.Elem()
	}
	if _type.Kind() != reflect.Struct {
		return nil
	}

	fieldNum := _type.NumField()
	tags := make(map[string]map[string]string, fieldNum)
	for i := 0; i < fieldNum; i++ {
		field := _type.Field(i)
		if _tags := GetFieldTags(field.Tag); _tags != nil {
			tags[field.Name] = _tags
		}
	}
	return tags
}
