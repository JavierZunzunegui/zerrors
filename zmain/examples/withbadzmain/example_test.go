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

// Package withbadzmain_test imports zmain but not as its first import.
// Non-global wrapErrors have frames, but global ones are initialised before the frameCapture is set.
package withbadzmain_test

import (
	"fmt"

	"github.com/JavierZunzunegui/zerrors/zmain/examples/internal"

	_ "github.com/JavierZunzunegui/zerrors/zmain"
)

func ExampleWithBadZMain() {
	fmt.Println(internal.ExampleFunc())
	// Output:
	// wrapper (example.go:27): base
}
