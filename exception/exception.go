// Exception handler like "parent.child.sub-child...sub-child".
package exception

import (
	"fmt"
	"strings"
)

type Exception struct {
	_type string
	msg   string
}

func New(_type string, msg string) Exception {
	return Exception{_type: fixType(_type), msg: msg}
}

func fixType(_type string) string {
	_type = strings.TrimSpace(_type)
	return strings.Trim(_type, ".")
}

func (e Exception) Error() string {
	return e.String()
}

func (e Exception) String() string {
	return fmt.Sprintf("Exception<%s>: %s", e._type, e.msg)
}

func (e Exception) GetType() string {
	return e._type
}

func (e Exception) GetMsg() string {
	return e.msg
}

func (e Exception) IsChildByString(parent string) bool {
	return IsChildByString(e._type, parent)
}

func (e Exception) IsChild(parent Exception) bool {
	return IsChildByString(e._type, parent._type)
}

func (e Exception) IsSame(other Exception) bool {
	return IsSame(e._type, other._type)
}

func (e Exception) GetTypeList() (list []string) {
	return getListBySep(e._type, ".")
}

func getListBySep(s string, sep string) (list []string) {
	orig_list := strings.Split(s, sep)
	for _, t := range orig_list {
		if t = strings.TrimSpace(t); t != "" {
			list = append(list, t)
		}
	}
	return
}

func IsChildByString(child, parent string) bool {
	type1 := fixType(child)
	type2 := fixType(parent)

	// "aa.bb.cc" is not the child of "aa.bb.cc" or "aa.bb"
	if len(type1) <= len(type2) {
		return false
	}

	// "aa.bb.cc" is the child of "aa.bb" and "aa"
	return strings.Contains(type1, type2)
}

func IsSame(type1, type2 string) bool {
	type1 = fixType(type1)
	type2 = fixType(type2)

	if type1 == type2 {
		return true
	}
	return false
}

func IsChildByExc(child, parent Exception) bool {
	return IsChildByString(child._type, parent._type)
}
