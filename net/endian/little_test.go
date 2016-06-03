package endian_test

import (
	"bytes"
	"testing"

	"github.com/xgfone/go-tools/net/endian"
)

func TestLittleEndian(t *testing.T) {
	b := endian.Little.From64(0x1122334455667788)
	if bytes.Compare(b, []byte{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}) != 0 {
		t.Fail()
	}

	if v := endian.Little.To16([]byte{0x22, 0x11}); v != 0x1122 {
		t.Fail()
	}

	if v := endian.Little.To32([]byte{0x44, 0x33, 0x22, 0x11}); v != 0x11223344 {
		t.Fail()
	}

	if v := endian.Little.To64([]byte{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}); v != 0x1122334455667788 {
		t.Fail()
	}
}
