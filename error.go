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

// Package zerrors provides error wrapping functionality.
package zerrors

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
	"sync"

	"github.com/JavierZunzunegui/zerrors/internal"
)

func init() {
	internal.SetBasic(func(err error) string {
		return err.(*wrapError).basic()
	})

	internal.SetDetail(func(err error) string {
		return err.(*wrapError).detail()
	})
}

var bufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// wrapError is a linked list of errors, providing the wrapping functionality.
type wrapError struct {
	value error
	next  *wrapError
	pc    [1]uintptr
}

// Error uses the 'basic' function to serialise the error.
// By default it is, and under custom options it is expected to be, the shorter and more efficient form of wrapError
// serialisation.
// By default this is in the form "{err1}: {err2}: ... : {errN}" without any frame information.
func (w *wrapError) Error() string {
	return internal.Basic(w)
}

// basic provides a default format to how wrapped errors will be serialised under %s, %v, %q format or with the Error
// method :
// "{err1}: {err2}: ... : {errN}"
func (w *wrapError) basic() string {
	if w.next == nil && w.pc[0] == 0 {
		// Taking a shortcut.
		return w.value.Error()
	}

	return w.basicViaBuf()
}

func (w *wrapError) basicViaBuf() string {
	buf := bufPool.Get().(*bytes.Buffer)

	buf.WriteString(w.value.Error())

	for current := w.next; current != nil; current = current.next {
		buf.WriteString(": ")
		buf.WriteString(current.value.Error())

	}

	s := buf.String()

	buf.Reset()
	bufPool.Put(buf)

	return s
}

// Detail is an alternative error encoding to err.Error(), equivalent to the %+v format form.
// It can be called on any error but for non-zerror errors is encodes it in the standard %+v fmt form.
// For zerror errors, it uses the 'detail' function to serialise and it and is expected to (and by default, is) more
// detailed and expensive than that err.Error() and contains frame information.
// By default this is in the form:
// "{err1} ({file}:{line}): {err2} ({file}:{line}): ... : {errN} ({file}:{line})".
func Detail(err error) string {
	if _, ok := err.(*wrapError); ok {
		return internal.Detail(err)
	}
	return fmt.Sprintf("%+v", err)
}

// detail provides a default format to how wrapped errors will be serialised under %v format or with the Detail
// function:
// "{err1} ({file}:{line}): {err2} ({file}:{line}): ... : {errN} ({file}:{line})"
func (w *wrapError) detail() string {
	if !internal.GetFrameCapture() {
		return w.basic()
	}

	if w.next == nil && w.pc[0] == 0 {
		// Taking a shortcut.
		return w.value.Error()
	}

	return w.detailViaBuf()
}

func (w *wrapError) detailViaBuf() string {
	buf := bufPool.Get().(*bytes.Buffer)

	buf.WriteString(w.value.Error())
	if w.pc[0] != 0 {
		w.frameToBuffer(buf)
	}

	for current := w.next; current != nil; current = current.next {
		buf.WriteString(": ")
		buf.WriteString(current.value.Error())
		if current.pc[0] != 0 {
			current.frameToBuffer(buf)
		}
	}

	s := buf.String()

	buf.Reset()
	bufPool.Put(buf)

	return s
}

func (w *wrapError) frameToBuffer(buf *bytes.Buffer) {
	if w.pc[0] == 0 {
		return
	}

	frame, _ := runtime.CallersFrames(w.pc[:]).Next()

	buf.WriteString(" (")
	buf.WriteString(path.Base(frame.File))
	buf.WriteString(":")

	// alternative to buf.WriteString(strconv.Itoa(frame.Line)), without allocating.
	b := buf.Bytes()
	l := len(b)
	buf.Write(strconv.AppendInt(b, int64(frame.Line), 10)[l:])

	buf.WriteString(")")
}

// Unwrap is required by the current error wrapping functionality in pkg/errors.
// It returns either a non-nil WrapError or nil.
func (w *wrapError) Unwrap() error {
	if w.next == nil {
		return nil
	}

	return w.next
}

// Is is the custom method required by pkg/errors.Is.
func (w *wrapError) Is(err error) bool {
	return errors.Is(w.value, err)
}

// As is the custom method required by pkg/errors.As.
func (w *wrapError) As(i interface{}) bool {
	return errors.As(w.value, i)
}

func (w *wrapError) deepCopy() *wrapError {
	if w == nil {
		return nil
	}
	return &wrapError{value: w.value, next: w.next.deepCopy(), pc: w.pc}
}

// Format supports the fmt Printf (and variants) for wrapError.
// The 'basic' encoder (that of .Error()) is used in all encodings except for %+v, which uses the 'detail' one.
func (w *wrapError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, internal.Detail(w))
			return
		}
		_, _ = io.WriteString(s, w.Error())
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}

// New produces a wrapError from a non-wrapError.
// If frame capture is enabled, it will get one at the caller of New.
// If err is a wrapError it returns itself, with no frame changes. This should be avoided.
func New(err error) error {
	if err == nil {
		return err
	}

	if _, ok := err.(*wrapError); ok {
		// Anti-pattern.
		return err
	}

	wErr := wrapError{value: err}

	if internal.GetFrameCapture() {
		_ = runtime.Callers(2, wErr.pc[:])
	}

	return &wErr
}

// Wrap produces a wrapError using any error type.
// The first argument is the error being wrapped, the second is the error being added to it.
// If the first error is nil, Wrap does nothing and returns nil.
// If the second error is nil, it returns the first error. This should be avoided.
// Otherwise it returns a wrapError.
// If frame capture is enabled, it will get one at the caller of Wrap.
func Wrap(inErr, outErr error) error {
	if inErr == nil {
		return nil
	}

	if outErr == nil {
		// Anti-pattern.
		return inErr
	}

	wInErr, ok := inErr.(*wrapError)
	if !ok {
		wInErr = &wrapError{value: inErr}
	}

	if wOutErr, ok := outErr.(*wrapError); ok {
		// Anti-pattern.
		wErr := wOutErr.deepCopy()
		current := wErr
		for ; current.next != nil; current = current.next {
		}
		current.next = wInErr
		return wErr
	}

	wErr := wrapError{
		value: outErr,
		next:  wInErr,
	}

	if internal.GetFrameCapture() {
		_ = runtime.Callers(2, wErr.pc[:])
	}

	return &wErr
}

// errorString is a trivial implementation of error, identical to errors.errorString.
// It is implemented separately as a performance optimisation, as it saves one allocation.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// SNew produces a wrapError with a simple string error.
// If frame capture is enabled, it will get one at the caller of SNew.
func SNew(s string) error {
	// a performance optimisation, makes wrapError and errorString in a single allocation.
	alloc := struct {
		sErr errorString
		wErr wrapError
	}{}

	alloc.sErr = errorString{s}
	alloc.wErr = wrapError{
		value: &alloc.sErr,
	}

	if internal.GetFrameCapture() {
		_ = runtime.Callers(2, alloc.wErr.pc[:])
	}

	return &alloc.wErr
}

// SWrap produces a wrapError using a simple string error.
// The first argument is the error being wrapped, the second is the string error being added to it.
// If the first error is nil, Wrap does nothing and returns nil.
// Otherwise it returns a wrapError.
// If frame capture is enabled, it will get one at the caller of SWrap.
func SWrap(inErr error, s string) error {
	if inErr == nil {
		return nil
	}

	wInErr, ok := inErr.(*wrapError)
	if !ok {
		wInErr = &wrapError{value: inErr}
	}

	// a performance optimisation, makes wrapError and errorString in a single allocation.
	alloc := struct {
		sErr errorString
		wErr wrapError
	}{}
	alloc.sErr = errorString{s}
	alloc.wErr = wrapError{
		value: &alloc.sErr,
		next:  wInErr,
	}

	if internal.GetFrameCapture() {
		_ = runtime.Callers(2, alloc.wErr.pc[:])
	}

	return &alloc.wErr
}

// Value returns the error contained at this particular node in the error chain.
// For non-wrapError errors, it returns itself.
// The return value is never a wrapError, and is only nil if the input is also nil.
func Value(err error) error {
	if wErr, ok := err.(*wrapError); ok {
		return wErr.value
	}

	return err
}

// Frame returns the frame contained at this particular node in the error chain.
// For non-wrapError errors or wrapError errors without a frame, it returns an empty frame and false.
// Otherwise it returns the populated frame and true.
func Frame(err error) (runtime.Frame, bool) {
	wErr, ok := err.(*wrapError)
	if !ok || wErr.pc[0] == 0 {
		return runtime.Frame{}, false
	}

	frame, _ := runtime.CallersFrames(wErr.pc[:]).Next()
	return frame, true
}
