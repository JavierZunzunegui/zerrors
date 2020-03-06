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

// Package internal contains stateful global variables internal to zerrors.
package internal

func init() {
	if linkerFrameCapture == "true" {
		// Set by the linker.
		SetFrameCapture()
	}
}

var (
	// frameCapture may be set to true during init phase, but never after (except in internal tests).
	frameCapture = false

	// linkerFrameCapture is a linker alternative to zmain to manipulate frameCapture.
	// If 'true', frameCapture is enabled regardless of zmain import.
	// If 'false', frameCapture is disabled regardless of zmain import.
	// Any other value is ignored.
	linkerFrameCapture string
)

// SetFrameCapture is only used in zmain's init func and in testing.
func SetFrameCapture() {
	if linkerFrameCapture == "false" {
		return
	}
	once.Do(initCache)
	frameCapture = true
}

// UnsetFrameCapture is only used in testing.
func UnsetFrameCapture() {
	if linkerFrameCapture != "true" {
		frameCapture = false
	}
}

// GetFrameCapture is a getter for frameCapture.
func GetFrameCapture() bool {
	return frameCapture
}
