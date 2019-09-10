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

package function

import (
	"../strings2"
)

// Some string splitting functions.
//
// For backward compatibility, we redefine them to those in strings2.
// See the sub-package strings2.
var (
	Split        = strings2.Split
	SplitN       = strings2.SplitN
	SplitSpace   = strings2.SplitSpace
	SplitSpaceN  = strings2.SplitSpaceN
	SplitString  = strings2.SplitSpace
	SplitStringN = strings2.SplitSpaceN
)
