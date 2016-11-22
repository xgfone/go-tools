package io2_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/xgfone/go-tools/io2"
)

func TestReadN(t *testing.T) {
	buf := bytes.NewBufferString("abcdefghijklmnopqrstuvwxyz\n1234567890")
	if b, err := io2.ReadN(buf, 26); err != nil || string(b) != "abcdefghijklmnopqrstuvwxyz" {
		t.Fail()
	}

	if b, err := io2.ReadN(buf, 1); err != nil || b[0] != '\n' {
		t.Fail()
	}

	if b, err := io2.ReadN(buf, 11); err != io.EOF || string(b) != "1234567890" {
		t.Fail()
	}

}

func TestWriteN(t *testing.T) {
	rbuf := []byte("abcdefghijklmnopqrstuvwxyz\n1234567890\n")
	wbuf := bytes.NewBuffer(nil)

	if err := io2.WriteN(wbuf, rbuf); err != nil || wbuf.String() != string(rbuf) {
		t.Fail()
	}
}
