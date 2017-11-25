package tags

import (
	"testing"

	"github.com/xgfone/go-tools/function"
)

type TagTest struct {
	F1 string `tag1:"123" tag2:"456" tag3:"789" tag4:"000"`
	F2 string `tag1:"aaa" tag2:"bbb" tag3:"ccc" tag5:"zzz"`
	F3 string `tag1:"ddd" tag2:"eee" tag3:"fff" tag6:"yyy"`
}

var _tags = New(TagTest{})

func TestGetAllValuesByTag(t *testing.T) {
	vs := _tags.GetAllValuesByTag("tag1")
	if len(vs) != 3 || !function.InSlice("123", vs) ||
		!function.InSlice("aaa", vs) || !function.InSlice("ddd", vs) {
		t.Fail()
	}
}

func TestGetValueByFieldAndTag(t *testing.T) {
	if _tags.GetValueByFieldAndTag("F1", "tag1") != "123" {
		t.Fail()
	}
}

func TestGetAllFieldsByTag(t *testing.T) {
	fs := _tags.GetAllFieldsByTag("tag2")
	if len(fs) != 3 || !function.InSlice("F1", fs) ||
		!function.InSlice("F2", fs) || !function.InSlice("F3", fs) {
		t.Fail()
	}
}
