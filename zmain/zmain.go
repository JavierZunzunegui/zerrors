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

// Package zmain allows setting of some global state afeecting how all zerrors are wrapped or serialised.
// It is not meant to be imported by any normal library, those should rely on zerrors exclusively.
// It is recommended to imported in at most one place and for it's methods to be called at most once and exclusively
// during the package initialisation phase.
package zmain

import (
	// zerrors set some state in zerrors/internal that must be set before zmain is used.
	_ "github.com/JavierZunzunegui/zerrors"

	"github.com/JavierZunzunegui/zerrors/internal"
)

// UnsetFrameCapture disables all frame capturing. This may be used primarily as a performance optimisation if frame
// information is not desired at all.
func UnsetFrameCapture() {
	internal.UnsetFrameCapture()
}

// SetBasic modifies how the `Error() string` method serialises zerrors. Note this should not rely on the `Error()`
// method of zerrors or it may enter an infinite recursion.
func SetBasic(f func(error) string) {
	internal.SetBasic(f)
}

// SetDetail modifies how the `zerrors.Detail(error) string` method serialises zerrors. Note this should not rely on the
// `Detail(error)` function of zerrors or it may enter an infinite recursion.
func SetDetail(f func(error) string) {
	internal.SetDetail(f)
}
