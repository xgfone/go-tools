package endian_test

import (
	"bytes"
	"testing"

	"github.com/xgfone/go-tools/net/endian"
)

func TestBigEndian(t *testing.T) {
	b := endian.Big.From64(0x1122334455667788)
	if bytes.Compare(b, []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}) != 0 {
		t.Fail()
	}

	if v := endian.Big.To16([]byte{0x11, 0x22}); v != 0x1122 {
		t.Fail()
	}

	if v := endian.Big.To32([]byte{0x11, 0x22, 0x33, 0x44}); v != 0x11223344 {
		t.Fail()
	}

	if v := endian.Big.To64([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}); v != 0x1122334455667788 {
		t.Fail()
	}
}
