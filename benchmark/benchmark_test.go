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

package benchmark_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	pkg_errors "github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func scenarios() []int { return []int{1, 2, 3, 5, 10, 20} }

func createWrap(depth int) error {
	if depth == 0 {
		return zerrors.New(errors.New("base"))
	}
	return zerrors.Wrap(createWrap(depth-1), errors.New("wrapper"))
}

func createSWrap(depth int) error {
	if depth == 0 {
		return zerrors.SNew("base")
	}
	return zerrors.SWrap(createSWrap(depth-1), "wrapper")
}

func createFmtErrorF(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return fmt.Errorf("wrapper: %w", createFmtErrorF(depth-1))
}

func createErrorsNew(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return errors.New("wrapper: " + createErrorsNew(depth-1).Error())
}

func createPkgErrorsWithMessage(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return pkg_errors.WithMessage(createPkgErrorsWithMessage(depth-1), "wrapper")
}

func createXerrorsErrorF(depth int) error {
	if depth == 0 {
		return xerrors.New("base")
	}
	return xerrors.Errorf("wrapper: %w", createXerrorsErrorF(depth-1))
}

func runBenchmarkCreateAndError(b *testing.B, f func(int) error) {
	for _, depth := range scenarios() {
		depth := depth

		b.Run(fmt.Sprintf("depth_%d", depth), func(b *testing.B) {
			var msg string
			for i := 0; i < b.N; i++ {
				err := f(depth)
				msg = err.Error()
			}

			b.StopTimer()
			if testing.Verbose() {
				b.Log(msg)
			}
		})
	}
}
