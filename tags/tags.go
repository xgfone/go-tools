package tags

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// If true, output the procedure building the tags.
	Debug = false
)

// A struct to manage the tags of a struct.
type Tag struct {
	name   string
	fields []fTag
	f2t    map[string][]tTag
	t2f    map[string][]tField
}

type fTag struct {
	field string
	tag   reflect.StructTag
}

type tField struct {
	field string
	value string
}

type tTag struct {
	tag   string
	value string
}

func debugf(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", args...)
	}
}

// Create a new Tag to manage the tags of a certain struct.
//
// s is a struct variable or a pointer to a struct.
func NewTag(s interface{}) *Tag {
	typ := reflect.TypeOf(s)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	t := newTag(typ.Name())
	nf := typ.NumField()
	for i := 0; i < nf; i++ {
		field := typ.Field(i)
		t.fields = append(t.fields, fTag{field: field.Name, tag: field.Tag})
		t.f2t[field.Name] = make([]tTag, 0)
	}

	return t
}

func newTag(name string) *Tag {
	return &Tag{
		name:   name,
		fields: make([]fTag, 0),
		f2t:    make(map[string][]tTag),
		t2f:    make(map[string][]tField),
	}
}

// Build the tag on the struct. That's, analyze and get the information of all
// the tags in the struct.
func (t *Tag) BuildTag(tag string) *Tag {
	tag = strings.TrimSpace(tag)
	if _, ok := t.t2f[tag]; ok {
		return t
	}
	t.t2f[tag] = make([]tField, 0)
	for _, fields := range t.fields {
		_tag := fields.tag
		field := fields.field
		if v := strings.TrimSpace(_tag.Get(tag)); v != "" {
			debugf("Build: Field:[%v] Tag:[%v] Value:[%v]", field, tag, v)
			t.t2f[tag] = append(t.t2f[tag], tField{field: field, value: v})
			t.f2t[field] = append(t.f2t[field], tTag{tag: tag, value: v})
		}
	}
	return t
}

// Build a set of the tags. See BuildTag().
func (t *Tag) BuildTags(tags []string) *Tag {
	for _, tag := range tags {
		t.BuildTag(tag)
	}
	return t
}

// Return the name of Tag, that's, the name of the struct.
func (t Tag) Name() string {
	return t.name
}

// Get the value of the corresponding tag.
//
// If more than one field has the tag, return the value of the tag of the first
// field according to the order defined in the struct.
func (t Tag) Get(tag string) string {
	_, v := t.GetWithField(tag)
	return v
}

// Same as Get(), but also return the field name except its value.
func (t Tag) GetWithField(tag string) (field, value string) {
	tag = strings.TrimSpace(tag)
	if v, ok := t.t2f[tag]; !ok {
		return "", ""
	} else if len(v) == 0 {
		return "", ""
	} else {
		return v[0].field, v[0].value
	}
}

// Return the value of the tag in a specified field.
func (t Tag) GetByField(tag, field string) string {
	if v, ok := t.f2t[field]; !ok {
		return ""
	} else if len(v) == 0 {
		return ""
	} else {
		for _, value := range v {
			if tag == value.tag {
				return value.value
			}
		}
		return ""
	}
}

// Return the list of the fields which defines the tag.
func (t Tag) GetToField(tag string) []string {
	if v, ok := t.t2f[tag]; !ok {
		return nil
	} else if len(v) == 0 {
		return nil
	} else {
		fields := make([]string, 0)
		for _, value := range v {
			fields = append(fields, value.field)
		}
		return fields
	}
}
