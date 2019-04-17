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

// GetFieldTagsMap is the same as GetFieldTags, but returns a map, the key
// of which is the tag name, and the value of which is the tag value.
func GetFieldTagsMap(t interface{}) map[string]string {
	tags := GetFieldTags(t)
	if tags == nil {
		return nil
	}

	ms := make(map[string]string, len(tags))
	for _, tag := range tags {
		ms[tag[0]] = tag[1]
	}
	return ms
}

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

// GetStructTagsMap is the same as GetStructTags, but returns a map,
// which is `map[FieldName]map[TagName]TagValue`.
func GetStructTagsMap(t interface{}) map[string]map[string]string {
	tags := GetStructTags(t)
	if tags == nil {
		return nil
	}

	maps := make(map[string]map[string]string, len(tags))
	for _, tag := range tags {
		ms := maps[tag[0]]
		if ms == nil {
			ms = make(map[string]string, 4)
			maps[tag[0]] = ms
		}
		ms[tag[1]] = tag[2]
	}

	return maps
}
