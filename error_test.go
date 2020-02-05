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
	"regexp"
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/internal"
)

// filename needs updating if this file is renamed.
const fileName = "error_test.go"

const separator = `\:\s`

func TestZErrors(t *testing.T) {
	t.Run("without frames", func(t *testing.T) {
		testZErrors(t, false)
	})

	t.Run("with frames", func(t *testing.T) {
		testZErrors(t, true)
	})
}

func testZErrors(t *testing.T, withFrame bool) {
	t.Run("wrap", func(t *testing.T) {
		testWrap(t, withFrame)
	})

	t.Run("Is", testIs)

	t.Run("As", testAs)
}

func testWrap(t *testing.T, withFrame bool) {
	var s string
	if withFrame {
		frameReg := regexp.MustCompile(`\(` + regexp.QuoteMeta(fileName) + `\:[1-9][0-9]*\)`)
		if exampleFrame := "(" + fileName + ":100)"; !frameReg.MatchString(exampleFrame) {
			t.Fatalf("example frame %q not matched the frame format '%s'", exampleFrame, frameReg.String())
		}
		internal.SetFrameCapture()
		defer internal.UnsetFrameCapture()
		s = `\s` + frameReg.String()
	}

	t.Run("nil cases", func(t *testing.T) {
		if err := zerrors.Value(nil); err != nil {
			t.Errorf("Value(nil): expecting nil, got %q", err)
		}

		if frame, ok := zerrors.Frame(nil); ok {
			t.Errorf("Frame(nil): expecting false, got true with frame: %v", frame)
		}

		// Anti-pattern.
		if err := zerrors.New(nil); err != nil {
			t.Errorf("New(nil): expecting nil, got %q", err)
		}

		// Anti-pattern.
		if err := zerrors.Wrap(nil, nil); err != nil {
			t.Errorf("Wrap(nil, nil): expecting nil, got %q", err)
		}
	})

	const baseMsg = "some error"
	baseErr := errors.New(baseMsg)

	t.Run("non-wrapped cases", func(t *testing.T) {
		if err := zerrors.Value(baseErr); err != baseErr {
			t.Errorf("Value(baseErr): expecting baseErr=%q, got %q", baseErr, err)
		}

		if frame, ok := zerrors.Frame(baseErr); ok {
			t.Errorf("Frame(baseErr): expecting false, got true with frame: %v", frame)
		}

		if err := zerrors.Wrap(nil, baseErr); err != nil {
			t.Errorf("Wrap(nil, baseErr): expecting nil, got %q", err)
		}
		// Anti-pattern.
		if err := zerrors.Wrap(baseErr, nil); err != baseErr {
			t.Errorf("Wrap(baseErr, nil): expecting baseErr=%q, got %q", baseErr, err)
		}

		if err := zerrors.SWrap(nil, baseMsg); err != nil {
			t.Errorf("SWrap(nil, baseMsg): expecting nil, got %q", err)
		}
	})

	reg1 := regexp.MustCompile(regexp.QuoteMeta(baseMsg) + s)
	w1Err := zerrors.New(baseErr)
	s1Err := zerrors.SNew(baseMsg)

	t.Run("first level wrapping", func(t *testing.T) {
		if msg := w1Err.Error(); !reg1.MatchString(msg) {
			t.Errorf("w1Err.Error(): expecting regex '%s', got %q", reg1.String(), msg)
		}
		if err := zerrors.Value(w1Err); err != baseErr {
			t.Errorf("Value(w1Err): expecting baseErr=%q, got %q", baseErr, err)
		}
		if !withFrame {
			if frame, ok := zerrors.Frame(w1Err); ok {
				t.Errorf("Frame(w1Err): expecting false, got true with frame: %v", frame)
			}
		} else {
			if _, ok := zerrors.Frame(w1Err); !ok {
				t.Error("Frame(w1Err): expecting true, got false")
			}
		}
		if err := errors.Unwrap(w1Err); err != nil {
			t.Errorf("Unwrap(w1Err): expecting nil, got %q", err)
		}

		if msg := s1Err.Error(); !reg1.MatchString(msg) {
			t.Errorf("s1Err.Error(): expecting regex '%s', got %q", reg1.String(), msg)
		}
		if msg := zerrors.Value(s1Err).Error(); msg != baseMsg {
			t.Errorf("Value(SNew(baseMsg)).Error(): expecting err=%q, got %q", baseMsg, msg)
		}
		if !withFrame {
			if frame, ok := zerrors.Frame(s1Err); ok {
				t.Errorf("Frame(s1Err): expecting false, got true with frame: %v", frame)
			}
		} else {
			if _, ok := zerrors.Frame(s1Err); !ok {
				t.Error("Frame(s1Err): expecting true, got false")
			}
		}
		if err := errors.Unwrap(s1Err); err != nil {
			t.Errorf("Unwrap(s1Err): expecting nil, got %q", err)
		}

		if err := zerrors.New(w1Err); err != w1Err {
			t.Errorf("New(w1Err): expecting w1Err=%q, got %q", w1Err, err)
		}

		if err := zerrors.Wrap(nil, w1Err); err != nil {
			t.Errorf("Wrap(nil, w1Err): expecting nil, got %q", err)
		}
		// Anti-pattern.
		if err := zerrors.Wrap(w1Err, nil); err != w1Err {
			t.Errorf("Wrap(w1Err, nil): expecting w1Err=%q, got %q", w1Err, err)
		}
	})

	const secondMsg = "second error"
	secondErr := errors.New(secondMsg)
	reg2 := regexp.MustCompile(regexp.QuoteMeta(secondMsg) + s + separator + reg1.String())
	w2Err := zerrors.Wrap(w1Err, secondErr)
	s2Err := zerrors.SWrap(s1Err, secondMsg)

	t.Run("second level wrapping", func(t *testing.T) {
		if msg := w2Err.Error(); !reg2.MatchString(msg) {
			t.Errorf("w2Err.Error(): expecting regex '%s', got %q", reg2, msg)
		}
		if err := zerrors.Value(w2Err); err != secondErr {
			t.Errorf("Value(w2Err): expecting secondErr=%q, got %q", secondErr, err)
		}
		if !withFrame {
			if frame, ok := zerrors.Frame(w2Err); ok {
				t.Errorf("Frame(w2Err): expecting false, got true with frame: %v", frame)
			}
		} else {
			if _, ok := zerrors.Frame(w2Err); !ok {
				t.Error("Frame(w2Err): expecting true, got false")
			}
		}
		if err := errors.Unwrap(w2Err); err != w1Err {
			t.Errorf("Unwrap(w2Err): expecting w1Err=%q, got %q", w1Err, err)
		}
		if err := errors.Unwrap(errors.Unwrap(w2Err)); err != nil {
			t.Errorf("Unwrap(Unwrap(w2Err)): expecting nil, got %q", err)
		}

		if msg := s2Err.Error(); !reg2.MatchString(msg) {
			t.Errorf("s2Err.Error(): expecting regex '%s', got %q", reg2, msg)
		}
		if msg := zerrors.Value(s2Err).Error(); msg != secondMsg {
			t.Errorf("Value(s2Err).Error(): expecting secondMsg=%q, got %q", secondMsg, msg)
		}
		if !withFrame {
			if frame, ok := zerrors.Frame(s2Err); ok {
				t.Errorf("Frame(s2Err): expecting false, got true with frame: %v", frame)
			}
		} else {
			if _, ok := zerrors.Frame(s2Err); !ok {
				t.Error("Frame(s2Err): expecting true, got false")
			}
		}
		if msg := zerrors.Value(errors.Unwrap(s2Err)).Error(); msg != baseMsg {
			t.Errorf("Value(Unwrap(s2Err)).Error(): expecting baseMsg=%q, got %q", baseMsg, msg)
		}
		if err := errors.Unwrap(errors.Unwrap(s2Err)); err != nil {
			t.Errorf("Unwrap(Unwrap(s2Err)): expecting nil, got %q", err)
		}

		// Anti-pattern.
		if msg := zerrors.Wrap(w1Err, zerrors.New(secondErr)).Error(); !reg2.MatchString(msg) {
			t.Errorf("Wrap(w1Err, New(secondErr)): expecting regex '%s', got %q", reg2, msg)
		}
		// Anti-pattern.
		if msg, reg := zerrors.Wrap(baseErr, secondErr).Error(), regexp.MustCompile(regexp.QuoteMeta(secondMsg)+s+separator+regexp.QuoteMeta(baseMsg)); !reg.MatchString(msg) {
			t.Errorf("Wrap(baseErr, secondErr): expecting regex '%s', got %q", reg, msg)
		}
	})
}

func testIs(t *testing.T) {
	inErr := errors.New("some error")
	wErr := zerrors.New(inErr)
	outErr := errors.New("wrapping error")
	wErr = zerrors.Wrap(wErr, outErr)
	if !errors.Is(wErr, inErr) {
		t.Error("Is(wErr, inErr): should be true")
	}
	if !errors.Is(wErr, outErr) {
		t.Error("Is(wErr, outErr): should be true")
	}
	if errors.Is(wErr, errors.New("unknown error")) {
		t.Error("Is(Wrap(err1, err2), err3): should be false")
	}
}

type custom1Error struct{ msg string }

func (err custom1Error) Error() string { return err.msg }

type custom2Error struct{ msg string }

func (err custom2Error) Error() string { return err.msg }

func testAs(t *testing.T) {
	inErr := custom1Error{"some error"}
	wErr := zerrors.New(inErr)
	outErr := errors.New("wrapping error")
	wErr = zerrors.Wrap(wErr, outErr)
	err1 := custom1Error{}
	if !errors.As(wErr, &err1) {
		t.Error("As(wErr, &err1): should be true")
	} else if msg := err1.Error(); msg != inErr.Error() {
		t.Errorf("As(wErr, &err1): expecting err1.Error()=%q, got %q", inErr.Error(), msg)
	}
	err2 := custom2Error{}
	if errors.As(wErr, &err2) {
		t.Error("As(wErr, &err2): should be false")
	}
}
