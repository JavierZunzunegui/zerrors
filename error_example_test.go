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
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.UnsetFrameCapture()

	fmt.Println(zerrors.SWrap(nil, "irrelevant"))

	err := zerrors.SNew("base")

	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))

	err = zerrors.SWrap(err, "second wrapper")
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))
	// Output:
	// <nil>
	// first wrapper: base
	// first wrapper: base
	// second wrapper: first wrapper: base
	// second wrapper: first wrapper: base
}

func ExampleSWrap_customFormat() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.UnsetFrameCapture()

	err := zerrors.SNew("base")
	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)                    // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// first wrapper: base
	// first wrapper - base
}

func ExampleSWrap_withFrame() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()

	err := zerrors.SNew("base")

	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))

	err = zerrors.SWrap(err, "second wrapper")
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))
	// Output:
	// first wrapper: base
	// first wrapper (error_example_test.go:70): base (error_example_test.go:68)
	// second wrapper: first wrapper: base
	// second wrapper (error_example_test.go:74): first wrapper (error_example_test.go:70): base (error_example_test.go:68)
}

func ExampleSWrap_withFrameCustomFormat() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()

	err := zerrors.SNew("base")
	err = zerrors.SWrap(err, "first wrapper")
	fmt.Println(err)                    // normal format
	fmt.Println(zerrors.Detail(err))    // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// first wrapper: base
	// first wrapper (error_example_test.go:89): base (error_example_test.go:88)
	// first wrapper (zerrors_test.ExampleSWrap_withFrameCustomFormat) - base (zerrors_test.ExampleSWrap_withFrameCustomFormat)
}

type codeError struct {
	code int
}

func (e codeError) Error() string { return fmt.Sprintf("code=%d", e.code) }

func ExampleWrap() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.UnsetFrameCapture()

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
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.UnsetFrameCapture()

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

	err := zerrors.SNew("base")

	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))

	err = zerrors.Wrap(err, codeError{2})
	fmt.Println(err)
	fmt.Println(zerrors.Detail(err))
	// Output:
	// code=1: base
	// code=1 (error_example_test.go:143): base (error_example_test.go:141)
	// code=2: code=1: base
	// code=2 (error_example_test.go:147): code=1 (error_example_test.go:143): base (error_example_test.go:141)
}

func ExampleWrap_withFrameCustomFormat() {
	// This is just in testing, elsewhere import "github.com/JavierZunzunegui/zerrors/zmain" in your main.go.
	internal.SetFrameCapture()

	err := zerrors.SNew("base")
	err = zerrors.Wrap(err, codeError{1})
	fmt.Println(err)                    // normal format
	fmt.Println(zerrors.Detail(err))    // normal format
	fmt.Println(customErrorFormat(err)) // custom format
	// Output:
	// code=1: base
	// code=1 (error_example_test.go:162): base (error_example_test.go:161)
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
