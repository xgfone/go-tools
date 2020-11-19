// Copyright 2020 xgfone
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

package namespace

import (
	"strings"
	"testing"
)

func TestNameSpace(t *testing.T) {
	ns := NewNameSpace("_test_ns_")

	if exist, err := ns.IsExist(); err != nil {
		t.Error(err)
	} else if exist {
		t.Fail()
	}

	if err := ns.Create(); err != nil {
		if strings.Contain(err.Error(), "Permission denied") {
			return
		}
		t.Fatal(err)
	}

	if exist, err := ns.IsExist(); err != nil {
		t.Error(err)
	} else if !exist {
		t.Fail()
	}

	if _, err := ns.Exec("ls"); err != nil {
		t.Error(err)
	}

	if nss, err := GetAllNameSpace(); err != nil {
		t.Error(err)
	} else {
		var exist bool
		for _, _ns := range nss {
			if _ns.Name == ns.Name {
				exist = true
			}
		}

		if !exist {
			t.Errorf("NS '%s' does not exist", ns.Name)
		}
	}

	if err := ns.Delete(); err != nil {
		t.Fatal(err)
	}

	if nss, err := GetAllNameSpace(); err != nil {
		t.Error(err)
	} else {
		var exist bool
		for _, _ns := range nss {
			if _ns.Name == ns.Name {
				exist = true
			}
		}

		if exist {
			t.Errorf("NS '%s' exist", ns.Name)
		}
	}
}
