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

// Package withzmain_test imports zmain as its first import.
// All wrapErrors have frames, including the ones in global variables.
package withzmain_test

import (
	"fmt"

	"github.com/JavierZunzunegui/zerrors/zmain"
	"github.com/JavierZunzunegui/zerrors/zmain/examples/internal"
)

func init() {
	zmain.UnsetFrameCapture()
}

func ExampleWithZMain() {
	fmt.Println(internal.ExampleFunc())
	// Output:
	// wrapper (example.go:27): base (example.go:22)
}
