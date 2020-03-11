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

// Package zmain enables frame capturing as a side effect of importing it.
// If imported, it should be done within the main package itself, and ideally as the first import there.
package zmain

import (
	_ "github.com/JavierZunzunegui/zerrors"

	"github.com/JavierZunzunegui/zerrors/internal"
)

// UnsetFrameCapture is only used in testing.
func UnsetFrameCapture() {
	internal.UnsetFrameCapture()
}

func SetBasic(f func(error) string) {
	internal.SetBasic(f)
}

func SetDetail(f func(error) string) {
	internal.SetDetail(f)
}