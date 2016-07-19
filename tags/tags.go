// Manage the tags in a struct.
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
	fields []ft
	f2t    map[string][]TV
	t2f    map[string][]FV
}

type ft struct {
	field string
	tag   reflect.StructTag
}

// It is used for the method GetToField().
type FV struct {
	// The field name which the tag belongs to.
	Field string

	// The value of the tag.
	Value string
}

// It is used for the method GetAllByField().
type TV struct {
	// The name of the tag.
	Tag string

	// The value of the tag.
	Value string
}

// It is used for the method GetAll().
type FT struct {
	// The name of the field.
	Field string

	// The name of the tag which defined in Field.
	Tag string

	// The value of the tag which defined in Field.
	Value string
}

func debugf(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", args...)
	}
}

// Create a new Tag to manage the tags in a certain struct.
//
// v is a struct variable or a pointer to a struct.
func NewTag(v interface{}) *Tag {
	typ := reflect.TypeOf(v)
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
		t.fields = append(t.fields, ft{field: field.Name, tag: field.Tag})
		t.f2t[field.Name] = make([]TV, 0)
	}

	return t
}

func newTag(name string) *Tag {
	return &Tag{
		name:   name,
		fields: make([]ft, 0),
		f2t:    make(map[string][]TV),
		t2f:    make(map[string][]FV),
	}
}

// Build the tag upon the struct. That's, analyze and get the values of the
// tag in all the fields of the struct. If the tag has been built, ignore it.
//
// Notice: Building the tag is in turn according to the order of the field. You
// can set Debug to true to see the building order. If a field defines the tag
// and its value is empty or only whitespaces, it's treated that it doesn't exist.
func (t *Tag) BuildTag(tag string) *Tag {
	tag = strings.TrimSpace(tag)
	if _, ok := t.t2f[tag]; ok {
		return t
	}
	t.t2f[tag] = make([]FV, 0)
	for _, fields := range t.fields {
		_tag := fields.tag
		field := fields.field
		if v := strings.TrimSpace(_tag.Get(tag)); v != "" {
			debugf("Building: Field:[%v] Tag:[%v] Value:[%v]", field, tag, v)
			t.t2f[tag] = append(t.t2f[tag], FV{Field: field, Value: v})
			t.f2t[field] = append(t.f2t[field], TV{Tag: tag, Value: v})
		}
	}
	return t
}

// Build a set of the tags, which is equivalent to calling BuildTag() more than
// once. See BuildTag().
func (t *Tag) BuildTags(tags []string) *Tag {
	for _, tag := range tags {
		t.BuildTag(tag)
	}
	return t
}

// Return the name of Tag, which is the name of the struct.
func (t Tag) Name() string {
	return t.name
}

// Get the value of the corresponding tag.
//
// If more than one field has the tag, return the value of the tag of the first
// field according to the order defined in the struct. Return an empty string
// if no field defines the tag.
func (t Tag) Get(tag string) string {
	_, v := t.GetWithField(tag)
	return v
}

// Same as Get(), but also return the field name except its value.
// Return ("", "") if no field defines the tag.
func (t Tag) GetWithField(tag string) (field, value string) {
	tag = strings.TrimSpace(tag)
	if v, ok := t.t2f[tag]; !ok {
		return "", ""
	} else if len(v) == 0 {
		return "", ""
	} else {
		return v[0].Field, v[0].Value
	}
}

// Return the value of the tag in a specified field. Return an empty string if
// the field doesn't have the tag.
func (t Tag) GetByField(tag, field string) string {
	if v, ok := t.f2t[field]; !ok {
		return ""
	} else if len(v) == 0 {
		return ""
	} else {
		for _, value := range v {
			if tag == value.Tag {
				return value.Value
			}
		}
		return ""
	}
}

// Get the information of all the tags defined in all the fields.
// Return nil if no field defines the tag.
func (t Tag) GetToField(tag string) (fv []FV) {
	if v, ok := t.t2f[tag]; !ok {
		return nil
	} else if len(v) == 0 {
		return nil
	} else {
		fv = make([]FV, len(v))
		copy(fv, v)
		return
	}
}

// Get all the tags of the field. Return nil if the field has no tags.
func (t Tag) GetAllByField(field string) (tv []TV) {
	if v, ok := t.f2t[field]; !ok {
		return nil
	} else if len(v) == 0 {
		return nil
	} else {
		tv = make([]TV, len(v))
		copy(tv, v)
		return
	}
}

// Get all the information parsed by the tag manager. Return nil if no tag is
// parsed. It's almost used to debug or traverse the tags in all the fields.
//
// The returned list is sorted on the basis of the order of the field which is
// defined in the struct. But the tags defined in the same field is unordered.
func (t Tag) GetAll() []FT {
	ft := make([]FT, 0)
	for field, tvs := range t.f2t {
		for _, tv := range tvs {
			ft = append(ft, FT{Field: field, Tag: tv.Tag, Value: tv.Value})
		}
	}
	return ft
}
