//Copyright 2020 Google LLC
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//https://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package internal is just used for some examples.
package internal

import (
	"github.com/JavierZunzunegui/zerrors"
)

var e = zerrors.SNew("base")

// ExampleFunc returns a 2-tier wrapper error.
// The base error is a global variable.
func ExampleFunc() error {
	return zerrors.SWrap(e, "wrapper")
}
