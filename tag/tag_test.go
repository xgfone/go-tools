package tag

import (
	"reflect"
	"fmt"
)

func ExampleGetFieldTags() {
	type S struct {
		Field string `name1:"value1" name2:"value2"`
	}

	for _, tag := range GetFieldTags(reflect.TypeOf(S{}).Field(0).Tag) {
		fmt.Printf("Tag=%s, Value=%s\n", tag[0], tag[1])
	}

	// Output:
	// Tag=name1, Value=value1
	// Tag=name2, Value=value2
}

func ExampleGetStructTags() {
	type S struct {
		Field1 string `name1:"value1" name2:"value2"`
		Field2 string `name1:"value1" name2:"value2"`
	}

	for _, tag := range GetStructTags(S{}) {
		fmt.Printf("Field=%s, Tag=%s, Value=%s\n", tag[0], tag[1], tag[2])
	}

	// Output:
	// Field=Field1, Tag=name1, Value=value1
	// Field=Field1, Tag=name2, Value=value2
	// Field=Field2, Tag=name1, Value=value1
	// Field=Field2, Tag=name2, Value=value2
}
