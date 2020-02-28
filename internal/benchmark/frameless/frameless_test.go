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

package frameless_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/internal/benchmark"
	pkg_errors "github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func TestErrorsAreIdentical(t *testing.T) {
	for _, depth := range benchmark.Scenarios() {
		expected := strings.Repeat("wrapper: ", depth) + "base"

		if got := createWrap(depth).Error(); got != expected {
			t.Errorf("createWrap(%d).Error(): expected %q got %q", depth, expected, got)
		}

		if got := createSWrap(depth).Error(); got != expected {
			t.Errorf("createSWrap(%d).Error(): expected %q got %q", depth, expected, got)
		}

		if got := createFmtErrorf(depth).Error(); got != expected {
			t.Errorf("createFmtErrorf(%d).Error(): expected %q got %q", depth, expected, got)
		}

		if got := createErrorsNew(depth).Error(); got != expected {
			t.Errorf("createErrorsNew(%d).Error(): expected %q got %q", depth, expected, got)
		}

		if got := createPkgErrorsWithMessage(depth).Error(); got != expected {
			t.Errorf("createPkgErrorsWithMessage(%d).Error(): expected %q got %q", depth, expected, got)
		}
	}
}

func createWrap(depth int) error {
	if depth == 0 {
		return zerrors.New(errors.New("base"))
	}
	return zerrors.Wrap(createWrap(depth-1), errors.New("wrapper"))
}

func BenchmarkZerrors_Wrap(b *testing.B) {
	benchmark.CreateAndError(b, createWrap, nil)
}

func createSWrap(depth int) error {
	if depth == 0 {
		return zerrors.SNew("base")
	}
	return zerrors.SWrap(createSWrap(depth-1), "wrapper")
}

func BenchmarkZerrors_SWrap(b *testing.B) {
	benchmark.CreateAndError(b, createSWrap, nil)
}

func createFmtErrorf(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return fmt.Errorf("wrapper: %w", createFmtErrorf(depth-1))
}

func BenchmarkFmt_Errorf(b *testing.B) {
	benchmark.CreateAndError(b, createFmtErrorf, nil)
}

func createErrorsNew(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return errors.New("wrapper: " + createErrorsNew(depth-1).Error())
}

func BenchmarkErrors_New(b *testing.B) {
	benchmark.CreateAndError(b, createErrorsNew, nil)
}

func createPkgErrorsWithMessage(depth int) error {
	if depth == 0 {
		return errors.New("base")
	}
	return pkg_errors.WithMessage(createPkgErrorsWithMessage(depth-1), "wrapper")
}

func BenchmarkPkgErrors_WithMessage(b *testing.B) {
	benchmark.CreateAndError(b, createPkgErrorsWithMessage, nil)
}

func createXerrorsErrorF(depth int) error {
	if depth == 0 {
		return xerrors.New("base")
	}
	return xerrors.Errorf("wrapper: %w", createXerrorsErrorF(depth-1))
}

func BenchmarkXerrors_Errorf(b *testing.B) {
	benchmark.CreateAndError(b, createXerrorsErrorF, nil)
}
