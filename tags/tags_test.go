package tags

import "testing"

type TagTest struct {
	F1 string `tag1:"123" tag2:"456" tag3:"789" tag4:"000"`
	F2 string `tag1:"aaa" tag2:"bbb" tag3:"ccc" tag5:"zzz"`
	F3 string `tag1:"ddd" tag2:"eee" tag3:"fff" tag6:"yyy"`
}

var _tags = New(TagTest{})

func TestGetAllValuesByTag(t *testing.T) {
	vs := _tags.GetAllValuesByTag("tag1")
	if len(vs) != 3 || vs[0] != "123" || vs[1] != "aaa" || vs[2] != "ddd" {
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
	if len(fs) != 3 || fs[0] != "F1" || fs[1] != "F2" || fs[2] != "F3" {
		t.Fail()
	}
}
