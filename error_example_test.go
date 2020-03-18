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

package zerrors_test

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/JavierZunzunegui/zerrors"
)

func ExampleSWrap() {
	err := zerrors.SNew("base")
	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)                          // basic standard format
	fmt.Println(zerrors.Detail(err))          // detail standard format
	fmt.Println(basicCustomErrorFormat(err))  // basic custom format
	fmt.Println(detailCustomErrorFormat(err)) // detail custom format
	// Output:
	// first wrapper: base
	// first wrapper (error_example_test.go:28): base (error_example_test.go:27)
	// first wrapper - base
	// first wrapper (zerrors_test.ExampleSWrap) - base (zerrors_test.ExampleSWrap)
}

type codeError struct {
	code int
}

func (e codeError) Error() string { return fmt.Sprintf("code=%d", e.code) }

func ExampleWrap() {
	err := zerrors.SNew("base")
	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)                          // basic standard format
	fmt.Println(zerrors.Detail(err))          // detail standard format
	fmt.Println(basicCustomErrorFormat(err))  // basic custom format
	fmt.Println(detailCustomErrorFormat(err)) // detail custom format
	// Output:
	// code=1: base
	// code=1 (error_example_test.go:48): base (error_example_test.go:47)
	// code=1 - base
	// code=1 (zerrors_test.ExampleWrap) - base (zerrors_test.ExampleWrap)
}

// An example custom alternative to err.Error().
func basicCustomErrorFormat(err error) string {
	var ss []string
	for ; err != nil; err = errors.Unwrap(err) {
		ss = append(ss, zerrors.Value(err).Error())
	}
	return strings.Join(ss, " - ")
}

// An example custom alternative to zerrors.Detail(err).
func detailCustomErrorFormat(err error) string {
	var ss []string
	for ; err != nil; err = errors.Unwrap(err) {
		s := zerrors.Value(err).Error()

		if frame, ok := zerrors.Frame(err); ok {
			s += " (" + path.Base(frame.Function) + ")"
		}

		ss = append(ss, s)
	}
	return strings.Join(ss, " - ")
}
