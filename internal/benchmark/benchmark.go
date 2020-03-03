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

// Package benchmark contains internal, common code for all benchmarks.
package benchmark

import (
	"fmt"
	"testing"
)

// Scenarios are the different depths at which wrapping is being tested in all benchmarks.
func Scenarios() []int { return []int{1, 2, 3, 5, 10, 20} }

func defaultPrinter(err error) string { return err.Error() }

// CreateAndError runs a benchmark consisting creating an error via wrapping it multiple times and serialising it.
func CreateAndError(b *testing.B, f func(int) error, printer func(error) string) {
	if printer == nil {
		printer = defaultPrinter
	}

	for _, depth := range Scenarios() {
		depth := depth

		b.Run(fmt.Sprintf("depth_%d", depth), func(b *testing.B) {
			var msg string
			for i := 0; i < b.N; i++ {
				err := f(depth)
				msg = printer(err)
			}

			b.StopTimer()
			if testing.Verbose() {
				b.Log(msg)
			}
		})
	}
}
