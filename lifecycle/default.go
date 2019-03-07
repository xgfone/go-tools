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

package lifecycle

var defaultManager *Manager

func init() {
	defaultManager = NewManager()
}

// Register registers the argument to the global default one.
func Register(functions ...func()) *Manager {
	return defaultManager.Register(functions...)
}

// RegisterChannel registers the argument to the global default one.
func RegisterChannel(in chan<- interface{}, out <-chan interface{}) *Manager {
	return defaultManager.RegisterChannel(in, out)
}

// Stop stops the global default one.
func Stop() {
	defaultManager.Stop()
}

// GetDefaultManager returns the default global Manager.
func GetDefaultManager() *Manager {
	return defaultManager
}

// ResetDefaultManager resets the default global manager.
func ResetDefaultManager(m *Manager) {
	if m == nil {
		panic("The argument is nil")
	}
	defaultManager = m
}

// IsStop returns true if the default global manager, or false.
func IsStop() bool {
	return defaultManager.IsStop()
}

// RunForever calls the method RunForever of the default global manager.
func RunForever() {
	defaultManager.RunForever()
}

// Wait calls the method Wait of the default global manager.
func Wait() {
	defaultManager.Wait()
}

// Exit calls the method Exit of the default global manager.
func Exit(code int) {
	defaultManager.Exit(code)
}
