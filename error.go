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
	"path"
	"runtime"
	"strconv"
	"sync"

	"github.com/JavierZunzunegui/zerrors/internal"
)

var bufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// wrapError is a linked list of errors, providing the wrapping functionality.
type wrapError struct {
	value error
	next  *wrapError
	pc    [1]uintptr
}

// Error provides a default format to how wrapped errors will be serialised:
// "{err1}: {err2}: ... : {errN}"
func (w *wrapError) Error() string {
	if w.next == nil && w.pc[0] == 0 {
		// Taking a shortcut.
		return w.value.Error()
	}

	return w.errorViaBuf()
}

func (w *wrapError) errorViaBuf() string {
	buf := bufPool.Get().(*bytes.Buffer)

	var framer internal.Framer
	if internal.GetFrameCapture() {
		framer = internal.GetFramer()
	}

	w.errorToBuffer(buf, framer)

	for current := w.next; current != nil; current = current.next {
		buf.WriteString(": ")
		current.errorToBuffer(buf, framer)
	}

	s := buf.String()

	buf.Reset()
	bufPool.Put(buf)

	return s
}

func (w *wrapError) errorToBuffer(buf *bytes.Buffer, framer internal.Framer) {
	buf.WriteString(w.value.Error())

	if w.pc[0] == 0 {
		return
	}

	frame := framer.Get(w.pc[0])

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

	return internal.GetFramer().Get(wErr.pc[0]), true
}
