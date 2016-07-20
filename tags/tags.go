// Manage the tags in a struct.
package tags

import (
	"bytes"
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
	fields []*ft
	f2t    map[string][]*FT
	t2f    map[string][]*FT
}

type ft struct {
	Field string
	Tag   reflect.StructTag
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
		t.fields = append(t.fields, &ft{Field: field.Name, Tag: field.Tag})
		t.f2t[field.Name] = make([]*FT, 0)
	}

	return t
}

func newTag(name string) *Tag {
	return &Tag{
		name:   name,
		fields: make([]*ft, 0),
		f2t:    make(map[string][]*FT),
		t2f:    make(map[string][]*FT),
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
	t.t2f[tag] = make([]*FT, 0)
	for _, field := range t.fields {
		stag := (*field).Tag
		_field := (*field).Field
		if v := strings.TrimSpace(stag.Get(tag)); v != "" {
			debugf("Building: Field:[%v] Tag:[%v] Value:[%v]", _field, tag, v)
			value := &FT{Field: _field, Tag: tag, Value: v}
			t.t2f[tag] = append(t.t2f[tag], value)
			t.f2t[_field] = append(t.f2t[_field], value)
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
		value := *v[0]
		return value.Field, value.Value
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
			if tag == (*value).Tag {
				return (*value).Value
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
		i := 0
		for _, value := range v {
			fv[i] = FV{Field: (*value).Field, Value: (*value).Value}
			i++
		}
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
		i := 0
		for _, value := range v {
			_v := *value
			tv[i] = TV{Tag: _v.Tag, Value: _v.Value}
			i++
		}
		return
	}
}

// Get all the information parsed by the tag manager. Return nil if no tag is
// parsed. It's almost used to debug or traverse the tags in all the fields.
//
// The returned list is sorted on the basis of the order of the field which is
// defined in the struct. And the tags is based on the order which they are
// building.
func (t Tag) GetAll() []FT {
	ft := make([]FT, 0)
	for field, tvs := range t.f2t {
		for _, tv := range tvs {
			ft = append(ft, FT{Field: field, Tag: (*tv).Tag, Value: (*tv).Value})
		}
	}
	return ft
}

// Audit the result that the manager parses the tags upon the struct.
// It's almost used to debug.
//
// For the example above, the returned format is like:
// 	Name: TagTest
//
// 	Fields:
// 	{Field:F1 Tag:tag1:"123" tag2:"456" tag3:"789" tag4:"000"}
// 	{Field:F2 Tag:tag1:"aaa" tag2:"bbb" tag3:"ccc" tag5:"zzz"}
// 	{Field:F3 Tag:tag1:"ddd" tag2:"eee" tag3:"fff" tag6:"yyy"}
//
// 	Field To Tag:
// 	Field: F1, Value: [{F1 tag1 123}, {F1 tag2 456}]
// 	Field: F2, Value: [{F2 tag1 aaa}, {F2 tag2 bbb}, {F2 tag5 zzz}]
// 	Field: F3, Value: [{F3 tag1 ddd}, {F3 tag2 eee}, {F3 tag6 yyy}]
//
// 	Tag To Field:
// 	Tag: tag1, Value: [{F1 tag1 123}, {F2 tag1 aaa}, {F3 tag1 ddd}]
// 	Tag: tag2, Value: [{F1 tag2 456}, {F2 tag2 bbb}, {F3 tag2 eee}]
// 	Tag: tag5, Value: [{F2 tag5 zzz}]
// 	Tag: tag6, Value: [{F3 tag6 yyy}]
func (t Tag) Audit() string {
	buf := bytes.NewBufferString("")

	buf.WriteString(fmt.Sprintf("Name: %v\n", t.name))

	buf.WriteString("\nFields:\n")
	for _, ft := range t.fields {
		buf.WriteString(fmt.Sprintf("%+v\n", *ft))
	}

	buf.WriteString("\nField To Tag:\n")
	for field, values := range t.f2t {
		buf.WriteString(fmt.Sprintf("Field: %v, Value: [", field))
		first := true
		for _, ft := range values {
			if !first {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%v", *ft))
			first = false
		}
		buf.WriteString("]\n")
	}

	buf.WriteString("\nTag To Field:\n")
	for tag, values := range t.t2f {
		buf.WriteString(fmt.Sprintf("Tag: %v, Value: [", tag))
		first := true
		for _, ft := range values {
			if !first {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%v", *ft))
			first = false
		}
		buf.WriteString("]\n")
	}

	return buf.String()
}

// Travel the information of the tag, which is the funcational programming of
// GetToField.
//
// The type of the trvaeling function is func(string, string), which needs two
// arguments and no return value. The first argument is the name of the tag,
// and the second is the value of the tag.
func (t Tag) TravelByTag(tag string, f func(string, string)) {
	if fts, ok := t.t2f[tag]; ok {
		for _, ft := range fts {
			f((*ft).Field, (*ft).Value)
		}
	}
}

// Travel the information of the field, which is the funcational programming of
// GetAllByField.
//
// The type of the trvaeling function is func(string, string), which needs two
// arguments and no return value. The first argument is the name of the field,
// and the second is the value of the tag.
func (t Tag) TravelByField(field string, f func(string, string)) {
	if fts, ok := t.f2t[field]; ok {
		for _, ft := range fts {
			f((*ft).Tag, (*ft).Value)
		}
	}
}
