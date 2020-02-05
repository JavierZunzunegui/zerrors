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

func func0(b bool) error {
	if b {
		return zerrors.SNew("some error")
	}
	return nil
}

func func1(b bool) error {
	return zerrors.SWrap(func0(b), fmt.Sprintf("with b=%t", b))
}

func func2(b bool) error {
	if err := func1(b); err != nil {
		return zerrors.SWrap(err, "with if err check")
	}
	return nil
}

func func3(b bool) error {
	return zerrors.SWrap(func2(b), "no if err check")
}

func func4(b bool) error {
	return zerrors.Wrap(func3(b), myError{})
}

type myError struct{}

func (myError) Error() string { return "my error type" }

func ExampleError() {
	fmt.Println("without frames:")
	fmt.Println(func4(true))
	fmt.Println(func4(false))

	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	fmt.Println("")
	fmt.Println("with frames:")
	fmt.Println(func4(true))
	fmt.Println(func4(false))

	// Output:
	// without frames:
	// my error type: no if err check: with if err check: with b=true: some error
	// <nil>
	//
	// with frames:
	// my error type (error_example_test.go:51): no if err check (error_example_test.go:47): with if err check (error_example_test.go:41): with b=true (error_example_test.go:36): some error (error_example_test.go:30)
	// <nil>
}

func ExampleUnwrapAndValue() {
	fmt.Println("without frames:")
	for err := func4(true); err != nil; err = errors.Unwrap(err) {
		fmt.Println(zerrors.Value(err))
	}

	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	fmt.Println("")
	fmt.Println("with frames:")
	for err := func4(true); err != nil; err = errors.Unwrap(err) {
		fmt.Println(zerrors.Value(err))
	}

	// Output:
	// without frames:
	// my error type
	// no if err check
	// with if err check
	// with b=true
	// some error
	//
	// with frames:
	// my error type
	// no if err check
	// with if err check
	// with b=true
	// some error
}

func ExampleUnwrapAndFrame() {
	fmt.Println("without frames:")
	for err := func4(true); err != nil; err = errors.Unwrap(err) {
		if f, ok := zerrors.Frame(err); ok {
			fmt.Printf("file=%q, line=%d, function=%q\n", path.Base(path.Dir(f.File))+"/"+path.Base(f.File), f.Line, f.Function)
		} else {
			fmt.Println("(no frame)")
		}
	}

	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	fmt.Println("")
	fmt.Println("with frames:")
	for err := func4(true); err != nil; err = errors.Unwrap(err) {
		if f, ok := zerrors.Frame(err); ok {
			fmt.Printf("file=%q, line=%d, function=%q\n", path.Base(path.Dir(f.File))+"/"+path.Base(f.File), f.Line, f.Function)
		} else {
			fmt.Println("(no frame)")
		}
	}

	// Output:
	// without frames:
	// (no frame)
	// (no frame)
	// (no frame)
	// (no frame)
	// (no frame)
	//
	// with frames:
	// file="zerrors/error_example_test.go", line=51, function="github.com/JavierZunzunegui/zerrors_test.func4"
	// file="zerrors/error_example_test.go", line=47, function="github.com/JavierZunzunegui/zerrors_test.func3"
	// file="zerrors/error_example_test.go", line=41, function="github.com/JavierZunzunegui/zerrors_test.func2"
	// file="zerrors/error_example_test.go", line=36, function="github.com/JavierZunzunegui/zerrors_test.func1"
	// file="zerrors/error_example_test.go", line=30, function="github.com/JavierZunzunegui/zerrors_test.func0"
}

func customError(err error, joiner func([]string) string, f func(error, *runtime.Frame) string) string {
	var ss []string
	for ; err != nil; err = errors.Unwrap(err) {
		var framePtr *runtime.Frame
		if frame, ok := zerrors.Frame(err); ok {
			framePtr = &frame
		}
		ss = append(ss, f(zerrors.Value(err), framePtr))
	}
	return joiner(ss)
}

func ExampleCustomFormat() {
	internal.SetFrameCapture()
	defer internal.UnsetFrameCapture()

	target := func4(true)

	fmt.Println("' - '")
	fmt.Println(customError(target, func(ss []string) string { return strings.Join(ss, " - ") }, func(err error, _ *runtime.Frame) string { return err.Error() }))

	fmt.Println("")
	fmt.Println(`'\n\t'`)
	fmt.Println(customError(
		target,
		func(ss []string) string {
			var sep string
			for i := range ss {
				ss[i] = sep + ss[i]
				sep += "\t"
			}
			return strings.Join(ss, "\n")
		},
		func(err error, _ *runtime.Frame) string { return err.Error() },
	))

	fmt.Println("")
	fmt.Println(`reversed`)
	fmt.Println(customError(
		target,
		func(ss []string) string {
			rev := make([]string, len(ss))
			for i := range ss {
				rev[i] = ss[len(ss)-i-1]
			}
			return strings.Join(ss, " - ")
		},
		func(err error, _ *runtime.Frame) string { return err.Error() },
	))

	fmt.Println("")
	fmt.Println(`frame function only`)
	fmt.Println(customError(
		target,
		func(ss []string) string { return strings.Join(ss, ": ") },
		func(err error, frame *runtime.Frame) string {
			s := err.Error()
			if frame != nil {
				s += " (" + path.Base(frame.Function) + ")"
			}
			return s
		},
	))

	// Output:
	// ' - '
	// my error type - no if err check - with if err check - with b=true - some error
	//
	// '\n\t'
	// my error type
	// 	no if err check
	//		with if err check
	//			with b=true
	//				some error
	//
	// reversed
	// my error type - no if err check - with if err check - with b=true - some error
	//
	// frame function only
	// my error type (zerrors_test.func4): no if err check (zerrors_test.func3): with if err check (zerrors_test.func2): with b=true (zerrors_test.func1): some error (zerrors_test.func0)
}
