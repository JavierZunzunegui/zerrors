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
	"runtime"
	"strings"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/internal"
)

func ExampleSWrap() {
	fmt.Println(zerrors.SWrap(nil, "irrelevant"))

	err := zerrors.SNew("base")

	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)

	err = zerrors.SWrap(err, "second wrapper")
	fmt.Println(err)
	// Output:
	// <nil>
	// first wrapper: base
	// second wrapper: first wrapper: base
}

func ExampleSWrap_customFormat() {
	err := zerrors.SNew("base")
	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err.Error())            // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// first wrapper: base
	// first wrapper - base
}

func ExampleSWrap_withFrame() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	err := zerrors.SNew("base")

	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)

	err = zerrors.SWrap(err, "second wrapper")
	fmt.Println(err)
	// Output:
	// first wrapper (error_example_test.go:61): base (error_example_test.go:59)
	// second wrapper (error_example_test.go:64): first wrapper (error_example_test.go:61): base (error_example_test.go:59)
}

func ExampleSWrap_withFrameCustomFormat() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	err := zerrors.SNew("base")
	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err.Error())            // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// first wrapper (error_example_test.go:77): base (error_example_test.go:76)
	// first wrapper (zerrors_test.ExampleSWrap_withFrameCustomFormat) - base (zerrors_test.ExampleSWrap_withFrameCustomFormat)
}

type codeError struct {
	code int
}

func (e codeError) Error() string { return fmt.Sprintf("code=%d", e.code) }

func ExampleWrap() {
	fmt.Println(zerrors.Wrap(nil, codeError{1} /* irrelevant */))

	err := zerrors.SNew("base")

	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)

	err = zerrors.Wrap(err, codeError{2})
	fmt.Println(err)
	// Output:
	// <nil>
	// code=1: base
	// code=2: code=1: base
}

func ExampleWrap_customFormat() {
	err := zerrors.SNew("base")
	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)                    // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// code=1: base
	// code=1 - base
}

func ExampleWrap_withFrame() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	err := zerrors.SNew("base")

	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)

	err = zerrors.Wrap(err, codeError{2})
	fmt.Println(err)
	// Output:
	// code=1 (error_example_test.go:124): base (error_example_test.go:122)
	// code=2 (error_example_test.go:127): code=1 (error_example_test.go:124): base (error_example_test.go:122)
}

func ExampleWrap_withFrameCustomFormat() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	err := zerrors.SNew("base")
	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err.Error())            // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// code=1 (error_example_test.go:140): base (error_example_test.go:139)
	// code=1 (zerrors_test.ExampleWrap_withFrameCustomFormat) - base (zerrors_test.ExampleWrap_withFrameCustomFormat)
}

// An example custom alternative to err.Error().
func customErrorFormat(err error) string {
	var ss []string
	for ; err != nil; err = errors.Unwrap(err) {
		var framePtr *runtime.Frame
		if frame, ok := zerrors.Frame(err); ok {
			framePtr = &frame
		}

		s := zerrors.Value(err).Error()
		if framePtr != nil {
			s += " (" + path.Base(framePtr.Function) + ")"
		}
		ss = append(ss, s)
	}
	return strings.Join(ss, " - ")
}
