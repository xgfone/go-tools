// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package io2

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func ExampleReadN() {
	s := "1234567890"
	rbuf := bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 9); err != nil || string(v) != s[:9] {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	rbuf = bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 0); err != nil || string(v) != s {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	rbuf = bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 11); err != io.EOF || string(v) != s {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
	// OK
	// OK
}

func TestReadNWriter(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	reader := bytes.NewBufferString("test")
	if ReadNWriter(writer, reader, 4) != nil || writer.String() != "test" {
		t.Errorf("writer: %s", writer.String())
	}

	writer = bytes.NewBuffer(nil)
	reader = bytes.NewBufferString("test")
	if ReadNWriter(writer, reader, 2) != nil || writer.String() != "te" {
		t.Errorf("writer: %s", writer.String())
	} else if ReadNWriter(writer, reader, 2) != nil || writer.String() != "test" {
		t.Errorf("writer: %s", writer.String())
	}

	writer = bytes.NewBuffer(nil)
	reader = bytes.NewBufferString("test")
	if err := ReadNWriter(writer, reader, 5); err == nil {
		t.Error("non-nil")
	} else if err != io.EOF {
		t.Error(err)
	}
}
