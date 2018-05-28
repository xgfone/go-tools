// Package tags manages the tags in a struct.
package tags

import (
	"bytes"
	"fmt"
	"reflect"
)

// Tag represents a tag of a field in a struct.
type Tag struct {
	// The field name which the tag belongs to.
	Field string

	// The name of the tag.
	Name string

	// The value of the tag which defined in Field.
	Value string
}

// Tags is a struct to manage the tags of a struct.
type Tags struct {
	debug  bool
	name   string
	caches []Tag

	f2t map[string]map[string]string
	t2f map[string]map[string]string
}

// New returns a new Tag to manage the tags in a certain struct.
//
// v is a struct variable or a pointer to a struct.
//
// If passing the second argument and it's true, it will enable the debug mode,
// which will output the procedure building the tags.
func New(v interface{}, debug ...bool) *Tags {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}

	sname := typ.Name()
	tags := newTag(sname)
	nf := typ.NumField()
	var tag Tag
	for i := 0; i < nf; i++ {
		field := typ.Field(i)
		fname := field.Name

		tags.f2t[fname] = make(map[string]string)
		for k, v := range GetAllTags(field.Tag) {
			tags.debugf("Struct=%s, Field=%s, TagName=%s, TagValue=%s",
				sname, fname, k, v)

			tag = Tag{Field: fname, Name: k, Value: v}
			tags.caches = append(tags.caches, tag)
			tags.f2t[fname][k] = v

			if tf, ok := tags.t2f[k]; ok {
				tf[fname] = v
			} else {
				tags.t2f[k] = map[string]string{
					fname: v,
				}
			}
		}
	}

	return tags
}

func newTag(name string) *Tags {
	return &Tags{
		name:   name,
		caches: make([]Tag, 0),

		f2t: make(map[string]map[string]string),
		t2f: make(map[string]map[string]string),
	}
}

func (t Tags) debugf(format string, args ...interface{}) {
	if t.debug {
		fmt.Printf(format, args...)
	}
}

// Name returns the name of Tag, which is the name of the struct.
func (t Tags) Name() string {
	return t.name
}

// GetAllValuesByTag returns the values of the corresponding tag.
//
// Notice: There are more than one fields that has this tag.
func (t Tags) GetAllValuesByTag(tag string) []string {
	if fv := t.t2f[tag]; fv != nil {
		_len := len(fv)
		i := 0
		vs := make([]string, _len)
		for _, v := range fv {
			vs[i] = v
			i++
		}
		return vs
	}
	return nil
}

// GetAllFieldsAndValuesByTag is the same as GetValuesByTags(),
// but also return the field name as the key.
func (t Tags) GetAllFieldsAndValuesByTag(tag string) map[string]string {
	if v, ok := t.t2f[tag]; ok {
		return v
	}
	return nil
}

// GetValueByFieldAndTag returns the value of the tag in a specified field.
// Return an empty string if the field doesn't have the tag.
func (t Tags) GetValueByFieldAndTag(field, tag string) string {
	if tv, ok := t.f2t[field]; ok {
		for t, v := range tv {
			if tag == t {
				return v
			}
		}
	}

	return ""
}

// GetAllFieldsByTag returns the names of all the field where defined this tag.
// Return nil if no field defines the tag.
func (t Tags) GetAllFieldsByTag(tag string) []string {
	if fv, ok := t.t2f[tag]; ok {
		_len := len(fv)
		i := 0
		fs := make([]string, _len)
		for f := range fv {
			fs[i] = f
			i++
		}
		return fs
	}

	return nil
}

// GetAllTagsAndValuesByField returns all the tags of the field.
// Return nil if the field has no tags.
func (t Tags) GetAllTagsAndValuesByField(field string) map[string]string {
	if v, ok := t.f2t[field]; ok {
		return v
	}
	return nil
}

// GetAll returns all the information parsed by the tag manager.
// Return nil if no tag is parsed. It's almost used to debug or traverse the
// tags in all the fields.
//
// The returned list is sorted on the basis of the order of the field which is
// defined in the struct.
func (t Tags) GetAll() []Tag {
	return t.caches
}

// Audit the result that the manager parses the tags upon the struct.
// It's almost used to debug. If having a question about the built result, you
// can use it and print the returned value to check.
//
// For the example above, the returned format is like:
// 	Name: TagTest
// 	FieldName=field1, TagName=tag1, TagValue=value1
// 	FieldName=field2, TagName=tag2, TagValue=value2
// 	FieldName=field3, TagName=tag3, TagValue=value3
// 	......
func (t Tags) Audit() string {
	buf := bytes.NewBufferString(fmt.Sprintf("Name: %v\n", t.name))

	for _, tag := range t.caches {
		buf.WriteString(fmt.Sprintf("FieldName=%s, TagName=%s, TagValue=%s\n",
			tag.Field, tag.Name, tag.Value))
	}

	return buf.String()
}

// TravelByTag travels the information of the tag.
//
// The type of the trvaeling function is func(string, string), which needs two
// arguments and no return value. The first argument is the name of the field
// where defined the tag, and the second is the value of the tag.
func (t Tags) TravelByTag(tag string, f func(string, string)) {
	if fs, ok := t.t2f[tag]; ok {
		for field, value := range fs {
			f(field, value)
		}
	}
}

// TravelByField travels the information of the field.
//
// The type of the trvaeling function is func(string, string), which needs two
// arguments and no return value. The first argument is the name of the tag,
// and the second is the value of the tag.
func (t Tags) TravelByField(field string, f func(string, string)) {
	if ts, ok := t.f2t[field]; ok {
		for tag, value := range ts {
			f(tag, value)
		}
	}
}
