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

package withframe_test

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/internal/benchmark"
	_ "github.com/JavierZunzunegui/zerrors/zmain" // makes zerrors capture frames
	pkg_errors "github.com/pkg/errors"
	"golang.org/x/xerrors"
)

const (
	fileName        = `withframe_test\.go`
	testPackagePath = `github\.com\/JavierZunzunegui\/zerrors\/internal\/benchmark\/withframe_test`
	repoPath        = `zerrors\/internal\/benchmark\/withframe\/withframe_test\.go`
)

func TestErrorsContainsFrames(t *testing.T) {
	for _, depth := range benchmark.Scenarios() {

		wErr := createWrap(depth)
		if reg, msg := zerrorsRegexp(depth), wErr.Error(); !reg.MatchString(msg) {
			t.Errorf("createWrap(%d): expected regexp %s got %s", depth, reg.String(), msg)
		}

		swErr := createSWrap(depth)
		if reg, msg := zerrorsRegexp(depth), swErr.Error(); !reg.MatchString(msg) {
			t.Errorf("createSWrap(%d): expected regexp %s got %s", depth, reg.String(), msg)
		}

		if reg, msg := pkgErrorsWithStackAndMessageRegexp(depth), printfWithPlusV(createPkgErrorsWithStackAndMessage(depth)); !reg.MatchString(msg) {
			t.Errorf("createPkgErrorsWithStackAndMessage(%d): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), msg)
		} else if zWrapMsg := zerrorsLikePkgErrors(wErr); !reg.MatchString(zWrapMsg) {
			t.Fatalf("zerrorsLikePkgErrors(createWrap(%d)): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), zWrapMsg)
		} else if zSWrapMsg := zerrorsLikePkgErrors(swErr); !reg.MatchString(zSWrapMsg) {
			t.Fatalf("zerrorsLikePkgErrors(createSWrap(%d)): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), zSWrapMsg)
		}

		if reg, msg := xerrorsRegexp(depth), printfWithPlusV(createXerrorsErrorf(depth)); !reg.MatchString(msg) {
			t.Errorf("createXerrorsErrorf(%d): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), msg)
		} else if zWrapMsg := zerrorsLikeXerrors(wErr); !reg.MatchString(zWrapMsg) {
			t.Fatalf("zerrorsLikeXerrors(createWrap(%d)): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), zWrapMsg)
		} else if zSWrapMsg := zerrorsLikeXerrors(swErr); !reg.MatchString(zSWrapMsg) {
			t.Fatalf("zerrorsLikeXerrors(createSWrap(%d)): expected regexp\n%s\ngot\n%s\n", depth, reg.String(), zSWrapMsg)
		}
	}
}

func zerrorsRegexp(depth int) *regexp.Regexp {
	const frame = `\(` + fileName + `\:[1-9][0-9]*\)`
	return regexp.MustCompile(fmt.Sprintf(`(wrapper %s\: ){%d}base %s`, frame, depth, frame))
}

func pkgErrorsWithStackAndMessageRegexp(depth int) *regexp.Regexp {
	const funcName = testPackagePath + `\.(createPkgErrorsWithStackAndMessage|createWrap|createSWrap)`
	const lineNumber = `[1-9][0-9]*`
	const linePath = `.*` + repoPath + `:` + lineNumber
	return regexp.MustCompile(fmt.Sprintf(`base(\n%s\n\t%s){%d}(\n.*\n *\t.*:%s)*(\nwrapper){%d}`, funcName, linePath, depth+1, lineNumber, depth))
}

func xerrorsRegexp(depth int) *regexp.Regexp {
	const funcName = testPackagePath + `\.(createXerrorsErrorf|createWrap|createSWrap)`
	const linePath = `.*` + repoPath + `:[1-9][0-9]*`
	const detailMsg = `wrapper:\n *` + funcName + `\n *` + linePath
	const baseMsg = `base:\n *` + funcName + `\n *` + linePath
	return regexp.MustCompile(fmt.Sprintf(`%s(\n *- %s){%d}\n *- %s`, detailMsg, detailMsg, depth-1, baseMsg))
}

func printfWithPlusV(err error) string {
	return fmt.Sprintf("%+v", err)
}

func zerrorsLikePkgErrors(err error) string {
	var frames []runtime.Frame
	var errs []error
	var lastFrameIndex int
	for i := 0; err != nil; err, i = errors.Unwrap(err), i+1 {
		errs = append(errs, zerrors.Value(err))
		frame, ok := zerrors.Frame(err)
		if ok {
			frames = append(frames, frame)
			lastFrameIndex = i
		}
	}

	var buf bytes.Buffer
	for firstError, i := true, len(errs)-1; i >= lastFrameIndex; firstError, i = false, i-1 {
		if !firstError {
			buf.WriteString("\n")
		}
		buf.WriteString(errs[i].Error())
	}

	for i := len(frames) - 1; i >= 0; i-- {
		frame := frames[i]
		buf.WriteString("\n")
		buf.WriteString(frame.Function)
		buf.WriteString("\n\t")
		buf.WriteString(frame.File)
		buf.WriteString(":")
		buf.WriteString(strconv.Itoa(frame.Line))
	}

	for i := lastFrameIndex - 1; i >= 0; i-- {
		buf.WriteString("\n")
		buf.WriteString(errs[i].Error())
	}

	return buf.String()
}

func zerrorsLikeXerrors(err error) string {
	var buf bytes.Buffer

	for firstError := true; err != nil; err, firstError = errors.Unwrap(err), false {
		frame, _ := zerrors.Frame(err)

		if !firstError {
			buf.WriteString("\n  - ")
		}
		buf.WriteString(zerrors.Value(err).Error())
		buf.WriteString(":\n    ")
		buf.WriteString(frame.Function)
		buf.WriteString("\n        ")
		buf.WriteString(frame.File)
		buf.WriteString(":")

		// alternative to buf.WriteString(strconv.Itoa(frame.Line)), without allocating.
		b := buf.Bytes()
		l := len(b)
		buf.Write(strconv.AppendInt(b, int64(frame.Line), 10)[l:])
	}

	return buf.String()
}

func createWrap(depth int) error {
	if depth == 0 {
		return zerrors.New(errors.New("base"))
	}
	return zerrors.Wrap(createWrap(depth-1), errors.New("wrapper"))
}

func BenchmarkZerrors_WrapWithFrame(b *testing.B) {
	benchmark.CreateAndError(b, createWrap, nil)
}

func createSWrap(depth int) error {
	if depth == 0 {
		return zerrors.SNew("base")
	}
	return zerrors.SWrap(createSWrap(depth-1), "wrapper")
}

func BenchmarkZerrors_SWrapWithFrame(b *testing.B) {
	benchmark.CreateAndError(b, createSWrap, nil)
}

func createPkgErrorsWithStackAndMessage(depth int) error {
	if depth == 0 {
		return pkg_errors.WithStack(errors.New("base"))
	}
	return pkg_errors.WithMessage(createPkgErrorsWithStackAndMessage(depth-1), "wrapper")
}

func BenchmarkPkgErrors_WithStackAndMessage(b *testing.B) {
	benchmark.CreateAndError(b, createPkgErrorsWithStackAndMessage, printfWithPlusV)
}

func createXerrorsErrorf(depth int) error {
	if depth == 0 {
		return xerrors.New("base")
	}
	return xerrors.Errorf("wrapper: %w", createXerrorsErrorf(depth-1))
}

func BenchmarkXerrors_Errorf(b *testing.B) {
	benchmark.CreateAndError(b, createXerrorsErrorf, printfWithPlusV)
}

func BenchmarkZerrors_WrapLikePkgErrors(b *testing.B) {
	benchmark.CreateAndError(b, createWrap, zerrorsLikePkgErrors)
}

func BenchmarkZerrors_SWrapLikePkgErrors(b *testing.B) {
	benchmark.CreateAndError(b, createSWrap, zerrorsLikePkgErrors)
}

func BenchmarkZerrors_WrapLikeXerrors(b *testing.B) {
	benchmark.CreateAndError(b, createWrap, zerrorsLikeXerrors)
}

func BenchmarkZerrors_SWrapLikeXerrors(b *testing.B) {
	benchmark.CreateAndError(b, createSWrap, zerrorsLikeXerrors)
}
