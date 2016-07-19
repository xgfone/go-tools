package tags_test

import (
	"fmt"

	"github.com/xgfone/go-tools/tags"
)

func ExampleTag() {
	type TagTest struct {
		F1 string `tag1:"123" tag2:"456" tag3:"789" tag4:"000"`
		F2 string `tag1:"aaa" tag2:"bbb" tag3:"ccc" tag5:"zzz"`
		F3 string `tag1:"ddd" tag2:"eee" tag3:"fff" tag6:"yyy"`
	}
	//tags.Debug = true
	tag := tags.NewTag(TagTest{})
	tag.BuildTags([]string{"tag1", "tag2"}).BuildTag("tag5").BuildTag("tag6")

	// Get
	fmt.Println(tag.Get("tag1"))
	fmt.Println(tag.Get("tag2"))
	fmt.Println(tag.Get("tag4"))

	// GetWithField
	fmt.Println(tag.GetWithField("tag1"))
	field, value := tag.GetWithField("tag4")
	fmt.Println(field, value, "End") // End is for saving the output whitespaces.
	fmt.Println(tag.GetWithField("tag5"))
	fmt.Println(tag.GetWithField("tag6"))

	// GetByField
	fmt.Println(tag.GetByField("tag1", "F1"))
	fmt.Println(tag.GetByField("tag1", "F2"))
	fmt.Println(tag.GetByField("tag4", "F2"))

	// GetToField
	fmt.Println(tag.GetToField("tag1"))
	fmt.Println(tag.GetToField("tag2"))
	fmt.Println(tag.GetToField("tag3"))
	fmt.Println(tag.GetToField("tag4"))
	fmt.Println(tag.GetToField("tag5"))
	fmt.Println(tag.GetToField("tag6"))

	// Output:
	// 123
	// 456
	//
	// F1 123
	//   End
	// F2 zzz
	// F3 yyy
	// 123
	// aaa
	//
	// [F1 F2 F3]
	// [F1 F2 F3]
	// []
	// []
	// [F2]
	// [F3]
}
