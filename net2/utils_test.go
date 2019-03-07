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

package net2

import "testing"

func TestJoinHostPort(t *testing.T) {
	result := "127.0.0.1:8000"

	if JoinHostPort("127.0.0.1", "8000") != result {
		t.Fail()
	}

	if JoinHostPort("127.0.0.1", 8000) != result {
		t.Fail()
	}

	if JoinHostPort([]byte("127.0.0.1"), 8000) != result {
		t.Fail()
	}

	if JoinHostPort([]byte("127.0.0.1"), "8000") != result {
		t.Fail()
	}
}

func TestGetAllIPs(t *testing.T) {
	if ips, err := GetAllIPs(); err != nil || len(ips) == 0 {
		t.Fail()
	}
}

func TestIPIsOnHost(t *testing.T) {
	if !IPIsOnHost("127.0.0.1") {
		t.Fail()
	}
}

func TestIsIP(t *testing.T) {
	if !IsIP("1.2.3.4") || !IsIP("fe80::acf4:ffff:feb7:bb24") {
		t.Fail()
	}

	if IsIP("1.2.3.4.5") || IsIP("fe80::acf4::ffff:feb7:bb24") {
		t.Fail()
	}
}
