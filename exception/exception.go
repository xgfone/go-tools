// Package exception is a exception handler like "parent.child.sub-child...sub-child".
//
// "A.B" is the parent exception is "A.B.C", "A.B.D" and "A.B.C.D", not "A.C".
//
package exception

import (
	"fmt"
	"strings"
)

var (
	// SEPARATOR is the separator between the parent and the child exceptions.
	SEPARATOR = "."
)

// Exception is the exception handler.
type Exception struct {
	_type string
	msg   string
}

// New returns a new Exception.
func New(_type string, msg string) Exception {
	return Exception{_type: fixType(_type), msg: msg}
}

func fixType(_type string) string {
	_type = strings.TrimSpace(_type)
	return strings.Trim(_type, SEPARATOR)
}

// Error implements the error interface.
func (e Exception) Error() string {
	return e.String()
}

// String implements the interface fmt.Stringer.
func (e Exception) String() string {
	return fmt.Sprintf("Exception<%s>: %s", e._type, e.msg)
}

// GetType returns the exception type.
func (e Exception) GetType() string {
	return e._type
}

// GetMsg returns the exception message.
func (e Exception) GetMsg() string {
	return e.msg
}

// IsChildByString is equal to IsChildByString(e.GetType(), parent).
func (e Exception) IsChildByString(parent string) bool {
	return IsChildByString(e._type, parent)
}

// IsChild is equal to IsChildByExc(e, parent).
func (e Exception) IsChild(parent Exception) bool {
	return IsChildByString(e._type, parent._type)
}

// IsParentByString is equal to IsChildByString(child, e.GetType()).
func (e Exception) IsParentByString(child string) bool {
	return IsChildByString(child, e._type)
}

// IsParent is equal to IsChildByExc(child, e).
func (e Exception) IsParent(child Exception) bool {
	return IsChildByString(child._type, e._type)
}

// IsSame is equal to IsSame(e.GetType(), other.GetType()).
func (e Exception) IsSame(other Exception) bool {
	return IsSame(e._type, other._type)
}

// GetTypeList returns a list of the exception types.
//
// For example, GetTypeList() on the exception type "e1.e2.e3" returns ["e1", "e2", "e3"].
func (e Exception) GetTypeList() (list []string) {
	return getListBySep(e._type, SEPARATOR)
}

func getListBySep(s string, sep string) (list []string) {
	origList := strings.Split(s, sep)
	for _, t := range origList {
		if t = strings.TrimSpace(t); t != "" {
			list = append(list, t)
		}
	}
	return
}

// IsChildByString check whether child is sub-exception of parent or not.
//
// Return true if yes, or false.
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

// IsSame check whether two exception types is the same.
//
// Return ture if yes, or false.
func IsSame(type1, type2 string) bool {
	type1 = fixType(type1)
	type2 = fixType(type2)

	if type1 == type2 {
		return true
	}
	return false
}

// IsChildByExc is same as IsChildByString, but the type is Exception, not string.
func IsChildByExc(child, parent Exception) bool {
	return IsChildByString(child._type, parent._type)
}
