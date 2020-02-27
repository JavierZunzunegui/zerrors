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
	"testing"
)

func BenchmarkFmt_Errorf(b *testing.B) {
	runBenchmarkCreateAndError(b, createFmtErrorF)
}

func BenchmarkErrors_New(b *testing.B) {
	runBenchmarkCreateAndError(b, createErrorsNew)
}

func BenchmarkPkgErrors_WithMessage(b *testing.B) {
	runBenchmarkCreateAndError(b, createPkgErrorsWithMessage)
}

func BenchmarkXerrors_Errorf(b *testing.B) {
	runBenchmarkCreateAndError(b, createXerrorsErrorF)
}
