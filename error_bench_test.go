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
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/internal"
)

func scenarios() []struct {
	name  string
	depth int
} {
	return []struct {
		name  string
		depth int
	}{
		{
			"depth-1",
			1,
		},
		{
			"depth-2",
			2,
		},
		{
			"depth-3",
			3,
		},
		{
			"depth-5",
			5,
		},
		{
			"depth-10",
			10,
		},
		{
			"depth-20",
			20,
		},
	}
}

func checkErrorsAreIdentical(b *testing.B) {
	for _, scenario := range scenarios() {
		if m := map[string]struct{}{
			createWrap(scenario.depth).Error():      {},
			createSWrap(scenario.depth).Error():     {},
			createFmtErrorF(scenario.depth).Error(): {},
			createErrorsNew(scenario.depth).Error(): {},
		}; len(m) != 1 {
			b.Fatalf("all createX methods should be identical, got %d values for depth %d: %v", len(m), scenario.depth, m)
		}
	}
	b.ResetTimer()
}

func createWrap(depth int) error {
	err := zerrors.New(errors.New("base"))
	for i := 0; i < depth; i++ {
		err = zerrors.Wrap(err, errors.New("wrapper"))
	}
	return err
}

func createSWrap(depth int) error {
	err := zerrors.SNew("base")
	for i := 0; i < depth; i++ {
		err = zerrors.SWrap(err, "wrapper")
	}
	return err
}

func createFmtErrorF(depth int) error {
	err := errors.New("base")
	for i := 0; i < depth; i++ {
		err = fmt.Errorf("wrapper: %w", err)
	}
	return err
}

func createErrorsNew(depth int) error {
	err := errors.New("base")
	for i := 0; i < depth; i++ {
		err = errors.New("wrapper: " + err.Error())
	}
	return err
}

func BenchmarkWrappingError_Error(b *testing.B) {
	checkErrorsAreIdentical(b)

	runBenchmarkWrappingError("Wrap", b, createWrap)

	internal.SetFrameCapture()
	runBenchmarkWrappingError("Wrap_WithFrame", b, createWrap)
	internal.UnsetFrameCapture()

	runBenchmarkWrappingError("SWrap", b, createSWrap)

	internal.SetFrameCapture()
	runBenchmarkWrappingError("SWrap_WithFrame", b, createSWrap)
	internal.UnsetFrameCapture()
}

func runBenchmarkWrappingError(name string, b *testing.B, f func(int) error) {
	b.Run(name, func(b *testing.B) {
		for _, scenario := range scenarios() {
			scenario := scenario

			err := f(scenario.depth)

			b.Run(scenario.name, func(b *testing.B) {
				b.ResetTimer()
				var msg string
				for i := 0; i < b.N; i++ {
					msg = err.Error()
				}

				b.StopTimer()
				if testing.Verbose() {
					b.Log(msg)
				}
			})
		}
	})
}

func BenchmarkCreateAndError(b *testing.B) {
	checkErrorsAreIdentical(b)

	runBenchmarkCreateAndError("Wrap", b, createWrap)

	internal.SetFrameCapture()
	runBenchmarkCreateAndError("Wrap_WithFrame", b, createWrap)
	internal.UnsetFrameCapture()

	runBenchmarkCreateAndError("SWrap", b, createSWrap)

	internal.SetFrameCapture()
	runBenchmarkCreateAndError("SWrap_WithFrame", b, createSWrap)
	internal.UnsetFrameCapture()

	runBenchmarkCreateAndError("fmt_Errorf", b, createFmtErrorF)

	runBenchmarkCreateAndError("errors_New", b, createErrorsNew)
}

func runBenchmarkCreateAndError(name string, b *testing.B, f func(int) error) {
	b.Run(name, func(b *testing.B) {
		for _, scenario := range scenarios() {
			scenario := scenario

			b.Run(scenario.name, func(b *testing.B) {
				var msg string
				for i := 0; i < b.N; i++ {
					err := f(scenario.depth)
					msg = err.Error()
				}

				b.StopTimer()
				if testing.Verbose() {
					b.Log(msg)
				}
			})
		}
	})
}
