package zmain_test

import (
	"errors"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/JavierZunzunegui/zerrors"
	"github.com/JavierZunzunegui/zerrors/zmain"
)

func TestZMain(t *testing.T) {
	errWithFrames := zerrors.SWrap(zerrors.SNew("base"), "wrapper")
	var errWithoutFrames error

	t.Run("UnsetFrameCapture", func(t *testing.T) {
		const fileRegex = `\(zmain_test\.go:[1-9][0-9]*\)`
		if expected, got := regexp.MustCompile("wrapper "+fileRegex+": base "+fileRegex), zerrors.Detail(errWithFrames); !expected.MatchString(got) {
			t.Errorf("Detail(errWithFrames): expected matching regex %s, got %q", expected.String(), got)
		}
		if _, ok := zerrors.Frame(errWithFrames); !ok {
			t.Error("Frame(errWithFrames): expected a frame, got none")
		}

		// This should normally only be done in an init function, but done here for testing purposes.
		zmain.UnsetFrameCapture()

		// Frames not printed by Detail if available after zmain.UnsetFrameCapture(), for consistency.
		if expected, got := "wrapper: base", zerrors.Detail(errWithFrames); got != expected {
			t.Errorf("UnsetFrameCapture(); Detail(errWithFrames): expected %q. got %q", expected, got)
		}
		// The frame remains accessible.
		if _, ok := zerrors.Frame(errWithFrames); !ok {
			t.Error("UnsetFrameCapture(); Frame(errWithFrames): expected a frame, got none")
		}

		errWithoutFrames = zerrors.SWrap(zerrors.SNew("base"), "wrapper")
		if expected, got := "wrapper: base", errWithoutFrames.Error(); got != expected {
			t.Errorf("UnsetFrameCapture(); errWithoutFrames.Error(): expected %q. got %q", expected, got)
		}
		if frame, ok := zerrors.Frame(errWithoutFrames); ok {
			t.Errorf("UnsetFrameCapture(); Frame(errWithoutFrames): expected no frame, got %v", frame)
		}
	})

	t.Run("SetBasic", func(t *testing.T) {
		// This should normally only be done in an init function, but done here for testing purposes.
		zmain.SetBasic(func(err error) string {
			var ss []string
			for ; err != nil; err = errors.Unwrap(err) {
				s := zerrors.Value(err).Error()
				if f, ok := zerrors.Frame(err); ok {
					s += " (" + path.Base(f.Function) + ")"
				}
				ss = append(ss, s)
			}

			return strings.Join(ss, " - ")
		})

		if expected, got := "wrapper - base", errWithoutFrames.Error(); expected != got {
			t.Errorf("SetBasic(.); errWithoutFrames.Errror(): expected %q. got %q", expected, got)
		}
		const funcName = "(zmain_test.TestZMain)"
		if expected, got := "wrapper "+funcName+" - base "+funcName, errWithFrames.Error(); expected != got {
			t.Errorf("SetBasic(.); errWithFrames.Errror(): expected %q. got %q", expected, got)
		}
	})

	t.Run("SetDetail", func(t *testing.T) {
		// This should normally only be done in an init function, but done here for testing purposes.
		zmain.SetDetail(func(err error) string {
			var ss []string
			for ; err != nil; err = errors.Unwrap(err) {
				s := zerrors.Value(err).Error()
				if f, ok := zerrors.Frame(err); ok {
					s += " (" + f.Function + ")"
				}
				ss = append(ss, s)
			}

			return strings.Join(ss, " - ")
		})

		if expected, got := "wrapper - base", zerrors.Detail(errWithoutFrames); expected != got {
			t.Errorf("SetDetail(.); Detail(errWithoutFrames): expected %q. got %q", expected, got)
		}
		const funcName = "(github.com/JavierZunzunegui/zerrors/zmain_test.TestZMain)"
		if expected, got := "wrapper "+funcName+" - base "+funcName, zerrors.Detail(errWithFrames); expected != got {
			t.Errorf("SetDetail(.); Detail(errWithFrames): expected %q. got %q", expected, got)
		}
	})
}
