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

package net

import (
	"fmt"
	"net"
)

// JoinHostPort is same as net.JoinHostPort, but it receives the arguments of
// any type, not only string.
//
// Recommend: Only use string or []byte as the type of host, and string or
// integer as that of port.
func JoinHostPort(host, port interface{}) string {
	if _host, ok := host.([]byte); ok {
		host = string(_host)
	}
	return net.JoinHostPort(fmt.Sprintf("%v", host), fmt.Sprintf("%v", port))
}
